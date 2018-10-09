package models

import (
	"encoding/json"
	"io"
	"time"

	"golang.org/x/net/context"

	"github.com/asaskevich/govalidator"
	"google.golang.org/appengine/datastore"
)

// Event is data for us to determine if stage is balanced
type Event struct {
	ID      int64     `json:"id" datastore:"-"`
	UserID  int64     `json:"userId"`
	Date    time.Time `json:"date"`
	EventID int64     `json:"eventId"`
	IsFirst bool      `json:"isFirst"`
	Type    int       `json:"type"`
	State   int       `json:"state"`
}

func (event *Event) key(c context.Context) *datastore.Key {
	if event.ID == 0 {
		return datastore.NewIncompleteKey(c, "Event", nil)
	}
	return datastore.NewKey(c, "Event", "", event.ID, nil)
}

func (event *Event) save(c context.Context) error {
	_, err := govalidator.ValidateStruct(event)
	if err != nil {
		return err
	}

	k, err := datastore.Put(c, event.key(c), event)
	if err != nil {
		return err
	}

	event.ID = k.IntID()
	return nil
}

// NewEvent inserts a new entry into the datastore
func NewEvent(c context.Context, r io.ReadCloser) (*Event, error) {

	var event Event
	// event.Timestamp = time.Now()
	err := json.NewDecoder(r).Decode(&event)
	if err != nil {
		return nil, err
	}

	event.ID = 0

	err = event.save(c)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

// GetAllEvents fetches all event entries from datastore
func GetAllEvents(c context.Context) ([]Event, error) {
	q := datastore.NewQuery("Event").Order("UserID")

	var events []Event
	_, err := q.GetAll(c, &events)
	if err != nil {
		return nil, err
	}

	return events, nil
}
