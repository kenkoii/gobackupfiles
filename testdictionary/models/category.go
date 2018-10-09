package models

type Category struct{
	ID	string		`json:"id,omitempty"`
	Name	string		`json:"name,omitempty"`
	Topics	Topics		`json:"topics"`
}

type Categories []Category