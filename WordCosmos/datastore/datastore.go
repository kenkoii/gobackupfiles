package datastore

import (
	"github.com/asaskevich/govalidator"
	"github.com/kenkoii/WordCosmos/wordcosmos"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

type DatastoreObject interface {
	kind() string
	key(context.Context) *datastore.Key
	save(context.Context) error
	search(context.Context) error
}

type AppReviewInfoDS struct {
	appReviewInfo wordcosmos.AppReviewInfo
}

type AppReviewInfoService struct{}

func (a *AppReviewInfoService) CreateAppReviewInfo(ari wordcosmos.AppReviewInfo, c context.Context) (*wordcosmos.AppReviewInfo, error) {
	ads := &AppReviewInfoDS{
		appReviewInfo: ari,
	}
	err := ads.save(c)
	if err != nil {
		return nil, err
	}

	return &ads.appReviewInfo, nil
}

func (ads *AppReviewInfoDS) kind() string {
	return "AppReviewInfo"
}

func (ads *AppReviewInfoDS) key(c context.Context) *datastore.Key {
	return datastore.NewIncompleteKey(c, ads.kind(), nil)
}

func (ads *AppReviewInfoDS) save(c context.Context) error {
	_, err := govalidator.ValidateStruct(ads)
	if err != nil {
		return err
	}
	_, err = datastore.Put(c, ads.key(c), ads)
	if err != nil {
		return err
	}
	return nil
}

func (ads *AppReviewInfoDS) search(c context.Context) error {
	err := ads.search(c)
	if err != nil {
		e := datastore.Get(c, ads.key(c), ads.appReviewInfo)
		if e != nil {
			return e
		}
		ads.save(c)
	}
	return nil
}
