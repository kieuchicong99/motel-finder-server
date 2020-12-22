package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BaseRepository interface {
	Create(object interface{}) error
	FindOne(id string, query interface{}, object interface{}) error
	UpdateOne(id string, query interface{}, object interface{}) error
	UpdateStatusOne(id string, query interface{}, status bool, object interface{}) error
	FindAll(query bson.D, page int, pageSize int, object interface{},opt*options.FindOptions) (total int64, err error)

}

