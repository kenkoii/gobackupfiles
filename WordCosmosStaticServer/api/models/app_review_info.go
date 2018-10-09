package models

import (
	"context"
	"encoding/json"
	"io"
	"time"

	"github.com/asaskevich/govalidator"
	"google.golang.org/appengine/datastore"
)

type AppReviewInfo struct {
	Timestamp int64     `datastore:"timestamp" json:"timestamp"`
	Message   string    `datastore:"message" json:"message"`
	Rating    int       `datastore:"rating" json:"rating"`
	Created   time.Time `datastore:"created"`
}

func (appReviewInfo *AppReviewInfo) key(c context.Context) *datastore.Key {
	return datastore.NewKey(c, "AppReviewInfo", "", appReviewInfo.Timestamp, nil)
}

func (appReviewInfo *AppReviewInfo) save(c context.Context) error {
	_, err := govalidator.ValidateStruct(appReviewInfo)
	if err != nil {
		return err
	}
	_, err = datastore.Put(c, appReviewInfo.key(c), appReviewInfo)
	if err != nil {
		return err
	}
	return nil
}

func (appReviewInfo *AppReviewInfo) search(c context.Context) error {
	err := appReviewInfo.search(c)
	if err != nil {
		e := datastore.Get(c, appReviewInfo.key(c), appReviewInfo)
		if e != nil {
			return e
		}
		appReviewInfo.save(c)
	}
	return nil
}

func NewAppReviewInfo(c context.Context, r io.ReadCloser) (*AppReviewInfo, error) {
	var appReviewInfo AppReviewInfo
	appReviewInfo.Created = time.Now()
	err := json.NewDecoder(r).Decode(&appReviewInfo)
	if err != nil {
		return nil, err
	}

	err = appReviewInfo.save(c)
	if err != nil {
		return nil, err
	}
	return &appReviewInfo, nil
}
