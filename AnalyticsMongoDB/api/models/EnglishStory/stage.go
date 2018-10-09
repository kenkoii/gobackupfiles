package models

import (
	"encoding/json"
	"io"
	"time"

	"golang.org/x/net/context"

	"github.com/asaskevich/govalidator"
	"google.golang.org/appengine/datastore"
)

// Stage is data for us to determine if the stage is balance
type Stage struct {
	ID           int64     `json:"id"`
	UserID       int64     `json:"userId"`
	Date         time.Time `json:"date"`
	State        int       `json:"state"`
	IsFirst      bool      `json:"isFirst"`
	Continue     int       `json:"continue"`
	PlayerLevel  int64     `json:"playerLevel"`
	StageID      int64     `json:"stageId"`
	ProvinceID   int64     `json:"provinceId"`
	AreaID       int64     `json:"areaId"`
	Cookie       int64     `json:"cookie"`
	Chocolate    int64     `json:"chocolate"`
	Achievement1 bool      `json:"achievement1"`
	Achievement2 bool      `json:"achievement2"`
	Achievement3 bool      `json:"achievement3"`
	FriendID     int64     `json:"friendId"`
	Turn         int64     `json:"turn"`
	Correct      int64     `json:"correct"`
}

func (stage *Stage) key(c context.Context) *datastore.Key {
	if stage.ID == 0 {
		return datastore.NewIncompleteKey(c, "Stage", nil)
	}
	return datastore.NewKey(c, "Stage", "", stage.ID, nil)
}

func (stage *Stage) save(c context.Context) error {
	_, err := govalidator.ValidateStruct(stage)
	if err != nil {
		return err
	}

	k, err := datastore.Put(c, stage.key(c), stage)
	if err != nil {
		return err
	}

	stage.ID = k.IntID()
	return nil
}

// NewStage inserts a new stage into the datastore
func NewStage(c context.Context, r io.ReadCloser) (*Stage, error) {

	var stage Stage
	// stage.Timestamp = time.Now()
	err := json.NewDecoder(r).Decode(&stage)
	if err != nil {
		return nil, err
	}

	stage.ID = 0

	err = stage.save(c)
	if err != nil {
		return nil, err
	}

	return &stage, nil
}

// GetAllStages fetches all stage entries from datastore
func GetAllStages(c context.Context) ([]Stage, error) {
	q := datastore.NewQuery("Stage").Order("UserID")

	var stages []Stage
	_, err := q.GetAll(c, &stages)
	if err != nil {
		return nil, err
	}

	return stages, nil
}
