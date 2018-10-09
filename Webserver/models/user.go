package models

import "time"

type User struct{
	Id		int
	Username	string		`json:"username"`
	Name		string		`json:"name"`
	Completed	bool		`json:"completed"`
	Due		time.Time	`json:"due"`
}

type Users []User