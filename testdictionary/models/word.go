package models

import(
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Word struct{
	Id	bson.ObjectId	`json:"id" bson:"_id,omitempty"`
	Topic	string		`json:"topic" bson:"topic`
	Name	string		`json:"name,omitempty"`
	AddedAt	time.Time	`json:"added_at" bson:"added_at"`

}

type Words []Word
