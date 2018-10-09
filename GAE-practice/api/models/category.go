package models

import (
	"time"

	"encoding/json"
	"github.com/asaskevich/govalidator"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"io"
)

type Category struct {
	Id          int64     `json:"id" datastore:"-"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Posted_By   string    `json:"posted_by"`
	Timestamp   time.Time `json:"timestamp"`
	TopicIds    []int64   `json:"topic_ids"`
	Topics      []Topic   `json:"topics" datastore:"-"`
}

func (category *Category) save(c context.Context) error {
	_, err := govalidator.ValidateStruct(category)
	if err != nil {
		return err
	}

	k, err := datastore.Put(c, category.key(c), category)
	if err != nil {
		return err
	}

	category.Id = k.IntID()
	return nil
}

func (category *Category) key(c context.Context) *datastore.Key {
	if category.Id == 0 {
		return datastore.NewIncompleteKey(c, "Category", nil)
	}
	return datastore.NewKey(c, "Category", "", category.Id, nil)
}

func GetCategories(c context.Context) ([]Category, error) {
	q := datastore.NewQuery("Category").Order("Name")

	var categories []Category
	keys, err := q.GetAll(c, &categories)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(categories); i++ {
		categories[i].Id = keys[i].IntID()
	}

	return categories, nil
}

func GetCategory(c context.Context, id int64) (*Category, error) {
	var category Category
	category.Id = id

	k := category.key(c)
	err := datastore.Get(c, k, &category)
	if err != nil {
		return nil, err
	}

	if category.TopicIds != nil {
		topics, err := GetTopicsByIds(c, category.TopicIds)
		if err != nil {
			return nil, err
		}
		category.Topics = topics
	}

	category.Id = k.IntID()

	return &category, nil
}

func NewCategory(c context.Context, r io.ReadCloser) (*Category, error) {

	var category Category
	category.Timestamp = time.Now()
	err := json.NewDecoder(r).Decode(&category)
	if err != nil {
		return nil, err
	}

	category.Id = 0

	err = category.save(c)
	if err != nil {
		return nil, err
	}

	return &category, nil

}

func RemoveCategory(c context.Context, id int64) (*Category, error) {

	category, err := GetCategory(c, id)
	if err != nil {
		return nil, err
	}

	err = datastore.Delete(c, category.key(c))
	if err != nil {
		return nil, err
	}

	return category, nil

}

func UpdateCategory(c context.Context, id int64, r io.ReadCloser) (*Category, error) {

	var category Category
	category.Id = id

	k := category.key(c)
	err := datastore.Get(c, k, &category)
	if err != nil {
		return nil, err
	}

	var cat Category
	err = json.NewDecoder(r).Decode(&cat)
	if err != nil {
		return nil, err
	}

	category.Name = cat.Name
	category.Description = cat.Description
	err = category.save(c)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func AddTopicId(c context.Context, id int64, topic_id int64) (*Category, error) {

	var category Category
	category.Id = id

	k := category.key(c)
	err := datastore.Get(c, k, &category)
	if err != nil {
		return nil, err
	}

	category.TopicIds = append(category.TopicIds, []int64{topic_id}...)
	err = category.save(c)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func RemoveTopicId(c context.Context, id int64, topic_id int64) (*Category, error) {

	var category Category
	category.Id = id

	k := category.key(c)
	err := datastore.Get(c, k, &category)
	if err != nil {
		return nil, err
	}

	for key, value := range category.TopicIds {
		if value == topic_id {
			category.TopicIds[key] = category.TopicIds[len(category.TopicIds)-1]
			category.TopicIds = category.TopicIds[:len(category.TopicIds)-1]
			break
		}
	}
	err = category.save(c)
	if err != nil {
		return nil, err
	}
	return &category, nil
}
