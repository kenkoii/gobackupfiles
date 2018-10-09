package models

import (
	"encoding/json"
	"io"
	"time"

	"github.com/asaskevich/govalidator"

	"golang.org/x/net/context"

	"google.golang.org/appengine/datastore"
)

// UserProperty is the latest information of the user, updated once a day/daily.
type UserProperty struct {
	UserID          int64           `json:"userId"`
	DownloadDate    time.Time       `json:"downloadDate"`
	LoginDate       time.Time       `json:"loginDate"`
	OS              operatingSystem `json:"operatingSystem"`
	Version         float32         `json:"version"`
	DownloadVersion float32         `json:"downloadVersion"`
	WordLevel       int64           `json:"wordLevel"`
	GrammarLevel    int64           `json:"grammarLevel"`
	PlayerLevel     int64           `json:"playerLevel"`
	CardAmount      int64           `json:"cardAmount"`
	CardLimit       int64           `json:"cardLimit"`
	Chocolate       int64           `json:"chocolate"`
	Gold            int64           `json:"Gold"`
	Cookie          int64           `json:"Cookie"`
	EventTicket     int64           `json:"eventTicket"`
	Lottery         int64           `json:"lottery"`
	DeckInfo        string          `json:"deckInfo"` //tentative
	LeaderCard      int64           `json:"leaderCard"`
	HelperCard      int64           `json:"helperCard"`
	Setting         string          `json:"setting"`
	Friend          int64           `json:"friend"`
	FriendLimit     int64           `json:"friendLimit"`
	QuestionTotal   int64           `json:"questionTotal"`
	FactoryTotal    int64           `json:"factoryTotal"`
	UraTotal        int64           `json:"uraTotal"`
	Fast            int64           `json:"fast"`
	Light           int64           `json:"light"`
}

type operatingSystem int

const (
	android operatingSystem = 1 + iota
	iOS
	amazon
)

// Datastore functions
func (userProperty *UserProperty) key(c context.Context) *datastore.Key {
	return datastore.NewKey(c, "UserProperty", "", userProperty.UserID, nil)
}

func (userProperty *UserProperty) save(c context.Context) error {
	_, err := govalidator.ValidateStruct(userProperty)
	if err != nil {
		return err
	}

	_, err = datastore.Put(c, userProperty.key(c), userProperty)
	if err != nil {
		return err
	}

	return nil
}

func (userProperty *UserProperty) search(c context.Context) error {
	err := datastore.Get(c, userProperty.key(c), userProperty)
	if err != nil {
		return err
	}
	return nil
}

// NewUserProperty is a method to create a new entry
func NewUserProperty(c context.Context, r io.ReadCloser) (*UserProperty, error) {

	var userProperty UserProperty
	err := json.NewDecoder(r).Decode(&userProperty)
	if err != nil {
		return nil, err
	}

	err = userProperty.save(c)
	if err != nil {
		return nil, err
	}

	return &userProperty, nil
}

// GetAllUserProperties fetches all user property entries from datastore
func GetAllUserProperties(c context.Context) ([]UserProperty, error) {
	q := datastore.NewQuery("UserProperty").Order("UserID")

	var userProperties []UserProperty
	_, err := q.GetAll(c, &userProperties)
	if err != nil {
		return nil, err
	}

	return userProperties, nil
}

// GetUserProperty fetches specific user property entry from datastore using ID
func GetUserProperty(c context.Context, id int64) (interface{}, error) {
	var userProperty UserProperty
	userProperty.UserID = id
	k := userProperty.key(c)
	err := datastore.Get(c, k, &userProperty)
	if err != nil {
		return nil, err
	}
	return &userProperty, nil
}
