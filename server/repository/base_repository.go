package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "server/infrastructure"
)


type baseRepository struct {
	primaryCollection *mongo.Collection
}

func (b *baseRepository) Create(object interface{}) error {
	res, err := b.primaryCollection.InsertOne(context.TODO(), object)
	if err != nil {
		return err
	}
	var _id, _ = res.InsertedID.(primitive.ObjectID)



	filler := bson.D{
		{"_id", _id},
	}
	err = b.primaryCollection.FindOne(context.TODO(), filler).Decode(object)
	return err

}
func (b *baseRepository) FindOne(id string, query interface{}, object interface{}) error {
	filter := bson.D{
		{id, query},
	}

	err := b.primaryCollection.FindOne(context.TODO(), filter).Decode(object)
	return err
}

func (b *baseRepository) UpdateOne(id string, query interface{}, object interface{}) error {
	opt := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := bson.D{
		{id, query},
	}
	update := bson.D{
		{"$set", object},
	}
	return b.primaryCollection.FindOneAndUpdate(context.TODO(), filter, update, opt).Decode(object)
}

func (b *baseRepository) UpdateStatusOne(id string, query interface{}, status bool, object interface{}) error {
	opt := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := bson.D{
		{id, query},
	}
	update := bson.D{
		{"$set", bson.D{
			{"Status", status},
		}},
	}
	return b.primaryCollection.FindOneAndUpdate(context.TODO(), filter, update, opt).Decode(object)
}
func (b *baseRepository) FindAll(query bson.D, page int, pageSize int, object interface{},opt*options.FindOptions) (total int64, err error) {
	findOptions := opt
	var cur *mongo.Cursor

	if page < 0 || pageSize < 0 && opt != nil {
		if cur, err = b.primaryCollection.Find(context.TODO(), query); err != nil {

			return
		}

		if err = cur.All(context.TODO(), object); err != nil {

			return
		}

	} else {
		if cur, err = b.primaryCollection.Find(context.TODO(), query, findOptions); err != nil {

			return
		}

		if err = cur.All(context.TODO(), object); err != nil {

			return
		}

	}
	if err = cur.Close(context.TODO()); err != nil {
		return 0, err
	}

	if total, err = b.primaryCollection.CountDocuments(context.TODO(), query); err != nil {

		return
	}
	return
}

