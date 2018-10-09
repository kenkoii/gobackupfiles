package models

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"io"
	"time"
)

type Topic struct {
	Id          int64     `json:"id" datastore:"-"`
	CategoryId  int64     `json:"category_id" valid:"required"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	WordIds     []int64   `json:"word_ids"`
	Words       []Topic   `json:"words" datastore:"-"`
	Posted_By   string    `json:"posted_by"`
	Timestamp   time.Time `json:"timestamp"`
}

func (topic *Topic) save(c context.Context) error {
	_, err := govalidator.ValidateStruct(topic)
	if err != nil {
		return err
	}

	k, err := datastore.Put(c, topic.key(c), topic)
	if err != nil {
		return err
	}

	topic.Id = k.IntID()
	return nil
}

func (topic *Topic) key(c context.Context) *datastore.Key {
	if topic.Id == 0 {
		return datastore.NewIncompleteKey(c, "Topic", nil)
	}
	return datastore.NewKey(c, "Topic", "", topic.Id, nil)
}

func GetTopics(c context.Context) ([]Topic, error) {
	q := datastore.NewQuery("Topic").Order("Name")

	var topics []Topic
	keys, err := q.GetAll(c, &topics)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(topics); i++ {
		topics[i].Id = keys[i].IntID()
	}

	return topics, nil
}

func GetTopic(c context.Context, id int64) (*Topic, error) {
	var topic Topic
	topic.Id = id

	k := topic.key(c)
	err := datastore.Get(c, k, &topic)
	if err != nil {
		return nil, err
	}

	topic.Id = k.IntID()

	return &topic, nil
}

func GetTopicsByIds(c context.Context, ids []int64) ([]Topic, error) {
	var keys []*datastore.Key

	for _, id := range ids {
		keys = append(keys, datastore.NewKey(c, "Topic", "", id, nil))
	}

	topics := make([]Topic, len(keys))
	err := datastore.GetMulti(c, keys, topics)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(topics); i++ {
		topics[i].Id = keys[i].IntID()
	}

	return topics, nil
}

func NewTopic(c context.Context, r io.ReadCloser) (*Topic, error) {

	var topic Topic
	topic.Timestamp = time.Now()
	err := json.NewDecoder(r).Decode(&topic)
	if err != nil {
		return nil, err
	}

	topic.Id = 0

	err = topic.save(c)
	if err != nil {
		return nil, err
	}
	_, err = AddTopicId(c, topic.CategoryId, topic.Id)
	return &topic, nil

}

func RemoveTopic(c context.Context, id int64) (*Topic, error) {

	topic, err := GetTopic(c, id)
	if err != nil {
		return nil, err
	}

	err = datastore.Delete(c, topic.key(c))
	if err != nil {
		return nil, err
	}
	_, err = RemoveTopicId(c, topic.CategoryId, topic.Id)
	return topic, nil

}

func UpdateTopic(c context.Context, id int64, r io.ReadCloser) (*Topic, error) {

	var topic Topic
	topic.Id = id

	k := topic.key(c)
	err := datastore.Get(c, k, &topic)
	if err != nil {
		return nil, err
	}

	var t Topic
	err = json.NewDecoder(r).Decode(&t)
	if err != nil {
		return nil, err
	}

	topic.Name = t.Name
	topic.Description = t.Description
	err = topic.save(c)
	if err != nil {
		return nil, err
	}
	return &topic, nil
}

func AddWordId(c context.Context, id int64, word_id int64) (*Topic, error) {

	var topic Topic
	topic.Id = id

	k := topic.key(c)
	err := datastore.Get(c, k, &topic)
	if err != nil {
		return nil, err
	}

	topic.WordIds = append(topic.WordIds, []int64{word_id}...)
	err = topic.save(c)
	if err != nil {
		return nil, err
	}
	return &topic, nil
}

func RemoveWordId(c context.Context, id int64, word_id int64) (*Topic, error) {

	var topic Topic
	topic.Id = id

	k := topic.key(c)
	err := datastore.Get(c, k, &topic)
	if err != nil {
		return nil, err
	}

	for key, value := range topic.WordIds {
		if value == word_id {
			topic.WordIds[key] = topic.WordIds[len(topic.WordIds)-1]
			topic.WordIds = topic.WordIds[:len(topic.WordIds)-1]
			break
		}
	}
	err = topic.save(c)
	if err != nil {
		return nil, err
	}
	return &topic, nil
}
