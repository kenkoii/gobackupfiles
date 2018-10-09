package models

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"io"
	"time"
)

type Word struct {
	Id        int64     `json:"id" datastore:"-"`
	TopicId   int64     `json:"topic_id" valid:"required"`
	Name      string    `json:"name"`
	Posted_By string    `json:"posted_by"`
	Timestamp time.Time `json:"timestamp"`
}

func (word *Word) save(c context.Context) error {
	_, err := govalidator.ValidateStruct(word)
	if err != nil {
		return err
	}

	k, err := datastore.Put(c, word.key(c), word)
	if err != nil {
		return err
	}

	word.Id = k.IntID()
	return nil
}

func (word *Word) key(c context.Context) *datastore.Key {
	if word.Id == 0 {
		return datastore.NewIncompleteKey(c, "Word", nil)
	}
	return datastore.NewKey(c, "Word", "", word.Id, nil)
}

func GetWords(c context.Context) ([]Word, error) {
	q := datastore.NewQuery("Word").Order("Name")

	var words []Word
	keys, err := q.GetAll(c, &words)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(words); i++ {
		words[i].Id = keys[i].IntID()
	}

	return words, nil
}

func GetWord(c context.Context, id int64) (*Word, error) {
	var word Word
	word.Id = id

	k := word.key(c)
	err := datastore.Get(c, k, &word)
	if err != nil {
		return nil, err
	}

	word.Id = k.IntID()

	return &word, nil
}

func GetWordsByIds(c context.Context, ids []int64) ([]Word, error) {
	var keys []*datastore.Key

	for _, id := range ids {
		keys = append(keys, datastore.NewKey(c, "Word", "", id, nil))
	}

	words := make([]Word, len(keys))
	err := datastore.GetMulti(c, keys, words)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(words); i++ {
		words[i].Id = keys[i].IntID()
	}

	return words, nil
}

func NewWord(c context.Context, r io.ReadCloser) (*Word, error) {

	var word Word
	word.Timestamp = time.Now()
	err := json.NewDecoder(r).Decode(&word)
	if err != nil {
		return nil, err
	}

	word.Id = 0

	err = word.save(c)
	if err != nil {
		return nil, err
	}
	//change to AddWordId
	_, err = AddWordId(c, word.TopicId, word.Id)
	return &word, nil
}

func RemoveWord(c context.Context, id int64) (*Word, error) {

	word, err := GetWord(c, id)
	if err != nil {
		return nil, err
	}

	err = datastore.Delete(c, word.key(c))
	if err != nil {
		return nil, err
	}
	//change to RemoveWordId
	_, err = RemoveWordId(c, word.TopicId, word.Id)
	return word, nil

}

func UpdateWord(c context.Context, id int64, r io.ReadCloser) (*Word, error) {

	var word Word
	word.Id = id

	k := word.key(c)
	err := datastore.Get(c, k, &word)
	if err != nil {
		return nil, err
	}

	var w Word
	err = json.NewDecoder(r).Decode(&w)
	if err != nil {
		return nil, err
	}

	word.Name = w.Name
	err = word.save(c)
	if err != nil {
		return nil, err
	}
	return &word, nil
}
