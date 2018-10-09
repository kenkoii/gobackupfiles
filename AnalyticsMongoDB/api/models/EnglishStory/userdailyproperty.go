package models

import (
	"encoding/json"
	"io"
	"time"

	"golang.org/x/net/context"

	"github.com/asaskevich/govalidator"
	"google.golang.org/appengine/datastore"
)

// UserDailyProperty is the information of the user on that day, received once a day/daily.
type UserDailyProperty struct {
	ID             int64       `json:"id" datastore:"-"`
	UserID         int64       `json:"userId"`
	LoginDate      time.Time   `json:"loginDate"`
	Gacha          gachaAmount `json:"gacha"`
	Quest          questAmount `json:"quest"`
	Event          eventAmount `json:"event"`
	CardAmount     int64       `json:"cardAmount"`
	PlayerLevel    int64       `json:"playerLevel"`
	ReachStage     int64       `json:"reachStage"`
	QuestionAmount int64       `json:"questionAmount"`
	FactoryAmount  int64       `json:"factoryAmount"`
	Strengthen     int64       `json:"strengthen"`
	Progress       int64       `json:"progress"`
	Sell           int64       `json:"sell"`
	ExceedLimit    int64       `json:"exceedLimit"`
}

type gachaAmount struct {
	GachaID int64 `json:"gachaId"`
	Amount  int64 `json:"amount"`
}

type questAmount struct {
	QuestID int64 `json:"questId"`
	Amount  int64 `json:"amount"`
}

type eventAmount struct {
	EventID int64 `json:"eventId"`
	Amount  int64 `json:"amount"`
}

func (userDailyProperty *UserDailyProperty) key(c context.Context) *datastore.Key {
	if userDailyProperty.ID == 0 {
		return datastore.NewIncompleteKey(c, "UserDailyProperty", nil)
	}
	return datastore.NewKey(c, "UserDailyProperty", "", userDailyProperty.ID, nil)
}

func (userDailyProperty *UserDailyProperty) save(c context.Context) error {
	_, err := govalidator.ValidateStruct(userDailyProperty)
	if err != nil {
		return err
	}

	k, err := datastore.Put(c, userDailyProperty.key(c), userDailyProperty)
	if err != nil {
		return err
	}

	userDailyProperty.ID = k.IntID()
	return nil
}

// NewUserDailyProperty inserts a new entry into the datastore
func NewUserDailyProperty(c context.Context, r io.ReadCloser) (*UserDailyProperty, error) {

	var userDailyProperty UserDailyProperty
	// userDailyProperty.Timestamp = time.Now()
	err := json.NewDecoder(r).Decode(&userDailyProperty)
	if err != nil {
		return nil, err
	}

	userDailyProperty.ID = 0

	err = userDailyProperty.save(c)
	if err != nil {
		return nil, err
	}

	return &userDailyProperty, nil
}

// GetAllUserDailyProperties fetches all user property entries from datastore
func GetAllUserDailyProperties(c context.Context) ([]UserDailyProperty, error) {
	q := datastore.NewQuery("UserDailyProperty").Order("UserID")

	var userDailyProperties []UserDailyProperty
	_, err := q.GetAll(c, &userDailyProperties)
	if err != nil {
		return nil, err
	}

	return userDailyProperties, nil
}
