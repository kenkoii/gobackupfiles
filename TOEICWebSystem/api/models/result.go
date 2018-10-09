package models

import (
	"encoding/json"
	"io"

	"golang.org/x/net/context"

	"time"

	"github.com/asaskevich/govalidator"
	"google.golang.org/appengine/datastore"
)

type TestResult struct {
	ID                int64      `json:"id" datastore:"-"`
	UserID            string     `json:"userId"`
	PackageID         int64      `json:"packageId"`
	IsFirst           bool       `json:"isFirst"`
	Answers           []Answer   `json:"answers"`
	Feedbacks         []Feedback `json:"feedbacks"`
	TotalScore        float64    `json:"totalScore"`
	DurationInSeconds int64      `json:"durationInSeconds"`
	Created           time.Time  `json:"created"`
}

type Answer struct {
	QuestionID int64  `json:"questionId"`
	Choice     Choice `json:"choice"`
}

type Feedback struct {
	QuestionID int64  `json:"questionId"`
	Comment    string `json:"comment"`
	Rating     int    `json:"rating"`
}

func (t *TestResult) key(c context.Context) *datastore.Key {
	if t.ID == 0 {
		return datastore.NewIncompleteKey(c, "TestResult", nil)
	}
	return datastore.NewKey(c, "TestResult", "", t.ID, nil)
}

func (t *TestResult) save(c context.Context) (*datastore.Key, error) {
	_, err := govalidator.ValidateStruct(t)
	if err != nil {
		return nil, err
	}
	k, err := datastore.Put(c, t.key(c), t)
	if err != nil {
		return nil, err
	}

	return k, nil
}

func NewTestResult(c context.Context, r io.ReadCloser) (*TestResult, error) {

	var testResult TestResult
	testResult.Created = time.Now()
	err := json.NewDecoder(r).Decode(&testResult)
	if err != nil {
		return nil, err
	}

	testResult.IsFirst = CheckIfFirstTake(c, testResult.UserID, testResult.PackageID)

	testResult.ID = 0
	k, err := testResult.save(c)
	if err != nil {
		return nil, err
	}
	testResult.ID = k.IntID()
	_, _ = AddUserScore(c, testResult.UserID, Score{PackageID: testResult.PackageID, Score: testResult.TotalScore, TestResultID: testResult.ID, Created: testResult.Created, DurationInSeconds: testResult.DurationInSeconds})
	return &testResult, nil
}

func NewFeedbackResult(c context.Context, id int64, r io.ReadCloser) (*TestResult, error) {

	var testResult TestResult
	testResult.ID = id
	testResult.Created = time.Now()
	k := testResult.key(c)
	err := datastore.Get(c, k, &testResult)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(r).Decode(&testResult.Feedbacks)
	if err != nil {
		return nil, err
	}

	k, err = testResult.save(c)
	if err != nil {
		return nil, err
	}
	testResult.ID = k.IntID()

	return &testResult, nil
}

func GetResult(c context.Context, id int64) (*TestResult, error) {
	var testResult TestResult
	testResult.ID = id

	k := testResult.key(c)
	err := datastore.Get(c, k, &testResult)
	if err != nil {
		return nil, err
	}

	return &testResult, nil
}

func GetResultsByPackage(c context.Context, id int64) (*[]TestResult, error) {
	var testResults []TestResult

	q := datastore.NewQuery("TestResult").Filter("PackageID = ", id)
	keys, err := q.GetAll(c, &testResults)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(testResults); i++ {
		testResults[i].ID = keys[i].IntID()
	}

	if len(testResults) == 0 {
		return &[]TestResult{}, nil
	}
	return &testResults, nil
}

func GetFirstResultsByPackage(c context.Context, id int64) (*[]TestResult, error) {
	var testResults []TestResult

	q := datastore.NewQuery("TestResult").Filter("PackageID = ", id).Filter("IsFirst = ", true)
	keys, err := q.GetAll(c, &testResults)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(testResults); i++ {
		testResults[i].ID = keys[i].IntID()
	}

	if len(testResults) == 0 {
		return &[]TestResult{}, nil
	}
	return &testResults, nil
}
