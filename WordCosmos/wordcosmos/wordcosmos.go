package wordcosmos

import "time"

type AppReviewInfo struct {
	Timestamp int64     `datastore:"timestamp" json:"timestamp"`
	Message   string    `datastore:"message" json:"message"`
	Rating    int       `datastore:"rating" json:"rating"`
	Created   time.Time `datastore:"created"`
}

type AppReviewInfoService interface {
	CreateAppReviewInfo(a *AppReviewInfo) (AppReviewInfo, error)
}
