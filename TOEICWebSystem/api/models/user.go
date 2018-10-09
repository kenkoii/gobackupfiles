package models

import (
	"encoding/json"
	"io"
	"time"

	"golang.org/x/net/context"

	"github.com/asaskevich/govalidator"
	"google.golang.org/appengine/datastore"
)

// User is the model for a user
type User struct {
	ID          string    `json:"userid"`
	DisplayName string    `json:"displayName"`
	PhotoURL    string    `json:"photoURL"`
	ToeicScore  int       `json:"toeicScore"`
	Scores      []Score   `json:"scores"`
	Role        string    `json:"role"`
	Created     time.Time `json:"created"`
}

type Score struct {
	PackageID         int64     `json:"packageId"`
	Score             float64   `json:"score"`
	TestResultID      int64     `json:"resultId"`
	DurationInSeconds int64     `json:"durationInSeconds"`
	Created           time.Time `json:"created"`
}

// ErrorMessage is the model for a user
type ErrorMessage struct {
	Message string `json:"error"`
}

func (user *User) key(c context.Context) *datastore.Key {
	return datastore.NewKey(c, "User", user.ID, 0, nil)
}

func (user *User) save(c context.Context) error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}

	_, err = datastore.Put(c, user.key(c), user)
	if err != nil {
		return err
	}

	return nil
}

func (user *User) search(c context.Context) error {
	err := datastore.Get(c, user.key(c), user)
	if err != nil {
		return err
	}
	return nil
}

func NewUser(c context.Context, r io.ReadCloser) (*User, error) {

	var user User
	user.Created = time.Now()
	err := json.NewDecoder(r).Decode(&user)
	if err != nil {
		return nil, err
	}

	err = user.save(c)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(c context.Context, id string, r io.ReadCloser) (*User, error) {

	var user User
	user.Created = time.Now()
	err := json.NewDecoder(r).Decode(&user)
	if err != nil {
		return nil, err
	}

	err = user.save(c)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(c context.Context, id string) (interface{}, error) {
	var user User
	user.ID = id
	k := user.key(c)
	err := datastore.Get(c, k, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func AddUserScore(c context.Context, id string, score Score) (interface{}, error) {
	var user User
	user.ID = id
	k := user.key(c)
	err := datastore.Get(c, k, &user)
	if err != nil {
		return nil, err
	}
	user.Scores = append(user.Scores, score)
	user.save(c)
	return &user, nil
}

func CheckIfFirstTake(c context.Context, userId string, packageId int64) bool {
	var isTaken = true
	var user User
	user.ID = userId
	k := user.key(c)
	err := datastore.Get(c, k, &user)
	if err != nil {
		return isTaken
	}
	for i := 0; i < len(user.Scores); i++ {
		if user.Scores[i].PackageID == packageId {
			isTaken = false
			return isTaken
		}
	}
	return isTaken
}
