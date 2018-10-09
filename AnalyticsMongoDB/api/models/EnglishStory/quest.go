package models

import (
	"encoding/json"
	"io"
	"time"

	"golang.org/x/net/context"

	"github.com/asaskevich/govalidator"
	"google.golang.org/appengine/datastore"
)

// Quest is data regarding ingame quests
type Quest struct {
	ID      int64     `json:"id"`
	UserID  int64     `json:"userId"`
	Date    time.Time `json:"date"`
	QuestID int64     `json:"questId"`
}

func (quest *Quest) key(c context.Context) *datastore.Key {
	if quest.ID == 0 {
		return datastore.NewIncompleteKey(c, "Quest", nil)
	}
	return datastore.NewKey(c, "Quest", "", quest.ID, nil)
}

func (quest *Quest) save(c context.Context) error {
	_, err := govalidator.ValidateStruct(quest)
	if err != nil {
		return err
	}

	k, err := datastore.Put(c, quest.key(c), quest)
	if err != nil {
		return err
	}

	quest.ID = k.IntID()
	return nil
}

// NewQuest inserts a new entry into the datastore
func NewQuest(c context.Context, r io.ReadCloser) (*Quest, error) {

	var quest Quest
	// quest.Timestamp = time.Now()
	err := json.NewDecoder(r).Decode(&quest)
	if err != nil {
		return nil, err
	}

	quest.ID = 0

	err = quest.save(c)
	if err != nil {
		return nil, err
	}

	return &quest, nil
}

// GetAllQuests fetches all quest entries from datastore
func GetAllQuests(c context.Context) ([]Quest, error) {
	q := datastore.NewQuery("Quest").Order("UserID")

	var quests []Quest
	_, err := q.GetAll(c, &quests)
	if err != nil {
		return nil, err
	}

	return quests, nil
}
