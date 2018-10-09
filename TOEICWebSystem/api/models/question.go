package models

import (
	"encoding/json"
	"io"
	"time"

	"golang.org/x/net/context"

	"github.com/asaskevich/govalidator"

	"google.golang.org/appengine/datastore"
)

// Question is the model for a question struct/object
type Question struct {
	ID          int64     `json:"id" datastore:"-"`
	TestID      int64     `json:"testId"`
	Question    string    `json:"question"`
	PackageID   int64     `json:"packageId"`
	Words       int64     `json:"words"`
	Level       int64     `json:"level"`
	Score       int64     `json:"score"`
	Explanation string    `json:"explanation"`
	Japanese    string    `json:"japanese"`
	Category    int64     `json:"category"`
	Choices     []Choice  `json:"choices"`
	Created     time.Time `json:"created"`
}

// Choice is the model for a choice struct/object
type Choice struct {
	Choice      string `json:"choice"`
	Translation string `json:"translation"`
	Correct     bool   `json:"correct"`
}

func (q *Question) key(c context.Context) *datastore.Key {
	if q.ID == 0 {
		return datastore.NewIncompleteKey(c, "Question", nil)
	}
	return datastore.NewKey(c, "Question", "", q.ID, nil)
}

func (p *Package) keys(c context.Context) []*datastore.Key {
	keys := []*datastore.Key{}
	for i := 0; i < len(p.Questions); i++ {
		keys = append(keys, datastore.NewIncompleteKey(c, "Question", nil))
	}
	return keys
}

func (q *Question) save(c context.Context) error {
	_, err := govalidator.ValidateStruct(q)
	if err != nil {
		return err
	}
	_, err = datastore.Put(c, q.key(c), q)
	if err != nil {
		return err
	}

	return nil
}

func NewQuestion(c context.Context, r io.ReadCloser) (*Question, error) {

	var question Question
	question.Created = time.Now()
	err := json.NewDecoder(r).Decode(&question)
	if err != nil {
		return nil, err
	}

	question.ID = 0
	err = question.save(c)
	if err != nil {
		return nil, err
	}

	return &question, nil
}
