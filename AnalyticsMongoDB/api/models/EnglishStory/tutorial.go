package models

import (
	"encoding/json"
	"io"
	"time"

	"golang.org/x/net/context"

	"github.com/asaskevich/govalidator"
	"google.golang.org/appengine/datastore"
)

// Tutorial is data for us to know when to stop
type Tutorial struct {
	ID         int64     `json:"id"`
	UserID     int64     `json:"userId"`
	TutorialID int64     `json:"tutorialId"`
	Date       time.Time `json:"date"`
}

func (tutorial *Tutorial) key(c context.Context) *datastore.Key {
	if tutorial.ID == 0 {
		return datastore.NewIncompleteKey(c, "Tutorial", nil)
	}
	return datastore.NewKey(c, "Tutorial", "", tutorial.ID, nil)
}

func (tutorial *Tutorial) save(c context.Context) error {
	_, err := govalidator.ValidateStruct(tutorial)
	if err != nil {
		return err
	}

	k, err := datastore.Put(c, tutorial.key(c), tutorial)
	if err != nil {
		return err
	}

	tutorial.ID = k.IntID()
	return nil
}

// NewTutorial inserts a new tutorial into the datastore
func NewTutorial(c context.Context, r io.ReadCloser) (*Tutorial, error) {

	var tutorial Tutorial
	// tutorial.Timestamp = time.Now()
	err := json.NewDecoder(r).Decode(&tutorial)
	if err != nil {
		return nil, err
	}

	tutorial.ID = 0

	err = tutorial.save(c)
	if err != nil {
		return nil, err
	}

	return &tutorial, nil
}

// GetAllTutorials fetches all tutorial entries from datastore
func GetAllTutorials(c context.Context) ([]Tutorial, error) {
	q := datastore.NewQuery("Tutorial").Order("UserID")

	var tutorials []Tutorial
	_, err := q.GetAll(c, &tutorials)
	if err != nil {
		return nil, err
	}

	return tutorials, nil
}
