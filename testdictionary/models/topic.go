package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Topic struct{
	Id	bson.ObjectId	`json:"id" bson:"_id,omitempty"`
	Name	string		`json:"name" bson:"name,omitempty"`
	Words	[]Word		`json:"words" bson:"words,omitempty"`
	AddedAt	time.Time	`json:"added_at" bson:"added_at"`
}

type Topics []Topic
