package models

import (
	"encoding/json"
	"io"
	"time"

	"golang.org/x/net/context"

	"github.com/asaskevich/govalidator"

	"google.golang.org/appengine/datastore"
)

// Package is the model for a package struct/object
type Package struct {
	ID        int64      `json:"id" datastore:"-"`
	Name      string     `json:"name"`
	Questions []Question `json:"questions,omitempty" datastore:"-"`
	Created   time.Time  `json:"created"`
}

func (p *Package) key(c context.Context) *datastore.Key {
	if p.ID == 0 {
		return datastore.NewIncompleteKey(c, "Package", nil)
	}
	return datastore.NewKey(c, "Package", "", p.ID, nil)
}

func (p *Package) save(c context.Context) (*datastore.Key, error) {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return nil, err
	}
	k, err := datastore.Put(c, p.key(c), p)
	if err != nil {
		return nil, err
	}

	return k, nil
}

func GetPackages(c context.Context) ([]Package, error) {
	q := datastore.NewQuery("Package").Order("Created")

	var packages []Package
	keys, err := q.GetAll(c, &packages)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(packages); i++ {
		packages[i].ID = keys[i].IntID()
	}

	return packages, nil
}

func GetPackage(c context.Context, id int64) (*Package, error) {
	var pack Package
	var questions []Question
	pack.ID = id

	k := pack.key(c)
	err := datastore.Get(c, k, &pack)
	if err != nil {
		return nil, err
	}

	pack.ID = k.IntID()

	q := datastore.NewQuery("Question").Filter("PackageID = ", id)
	keys, err := q.GetAll(c, &questions)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(questions); i++ {
		questions[i].ID = keys[i].IntID()
	}
	pack.Questions = questions

	return &pack, nil
}

func NewPackage(c context.Context, r io.ReadCloser) (*Package, error) {

	var pack Package
	pack.Created = time.Now()
	err := json.NewDecoder(r).Decode(&pack)
	if err != nil {
		return nil, err
	}

	pack.ID = 0
	k, err := pack.save(c)
	if err != nil {
		return nil, err
	}

	pack.ID = k.IntID()
	for i := 0; i < len(pack.Questions); i++ {
		pack.Questions[i].PackageID = pack.ID
		pack.Questions[i].Created = time.Now()
	}

	keys, err := datastore.PutMulti(c, pack.keys(c), pack.Questions)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(pack.Questions); i++ {
		pack.Questions[i].ID = keys[i].IntID()
	}
	// _, err = AddTopicId(c, topic.CategoryId, topic.Id)
	return &pack, nil
}
