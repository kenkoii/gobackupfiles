package models

import "context"

type DatastoreObject interface {
	key(context.Context)
	save(context.Context)
	search(context.Context)
}
