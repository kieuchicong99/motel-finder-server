package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"server/infrastructure"
	"server/model"
	"server/utilities"
)

// mongoGoDriver
type mongoGoDriverRepository struct {
	collection *mongo.Collection
}

func (m mongoGoDriverRepository) Insert(motel *model.Motel) (insertResponse *model.InsertResponse, httpCode int) {
	result, err := m.collection.InsertOne(context.TODO(), motel)
	if err != nil {
		utilities.ErrLog.Printf("Insert err, err: %v", err)
		insertResponse = &model.InsertResponse{
			Message: "Repository => InsertOne Error",
			Error:   fmt.Sprintf("%v", err),
			Data:    nil,
		}
		return insertResponse, 500
	}
	_id, _ := result.InsertedID.(primitive.ObjectID)
	filler := bson.D{
		{"_id", _id},
	}
	var motelCreated *model.Motel
	err = m.collection.FindOne(context.TODO(), filler).Decode(&motelCreated)
	//utilities.InfoLog.Printf("Result : ", motelCreated)
	return &model.InsertResponse{
		Message: "Success",
		Error:   "",
		Data:    motelCreated,
	}, 200
}

func (m mongoGoDriverRepository) Update(code string, motel *model.Motel) (updateResponse *model.UpdateResponse, httpCode int) {
	opt := options.FindOneAndUpdate().SetReturnDocument(options.After)
	c, _ := primitive.ObjectIDFromHex(code)
	filter := bson.M{
		"MotelCode": c,
	}
	utilities.InfoLog.Printf("Address: %v", motel)
	update := bson.D{
		{"$set", motel},
	}
	var motelUpdated model.Motel
	m.collection.FindOneAndUpdate(context.TODO(), filter, update, opt).Decode(&motelUpdated)
	utilities.InfoLog.Printf("motelUpdated: %v", motelUpdated)
	return &model.UpdateResponse{
		Message: "Success",
		Error:   "",
		Data:    motelUpdated,
	}, 200
}

func (m mongoGoDriverRepository) GetAll(page, pageSize int) (GetManyResponse *model.GetManyResponse, totalResult int, httpCode int) {
	query := bson.M{"_id": bson.M{"$exists": true}}
	opts := options.Find().SetSort(bson.M{"code": 1}).SetSkip(int64((page - 1) * pageSize)).SetLimit(int64(pageSize))

	cursor, err := m.collection.Find(context.TODO(), query, opts)
	defer cursor.Close(context.TODO())

	GetManyResponse = &model.GetManyResponse{
		Message: "Repository => GetManyResponse Error",
		Error:   "",
		Data:    nil,
	}
	if err != nil {
		utilities.ErrLog.Printf("Mongo Find error, err:%v\n", err)
		GetManyResponse.Error = fmt.Sprintf("%s", err)
		return GetManyResponse, 0, 500
	}

	data := []model.Motel{}
	if err := cursor.All(context.TODO(), &data); err != nil {
		utilities.ErrLog.Printf("Mongo cursor.All, err:%v\n", err)
		GetManyResponse.Error = fmt.Sprintf("%s", err)
		return GetManyResponse, 0, 500
	}

	GetManyResponse.Message = "Success"
	dataConvert := make([]interface{}, len(data))
	for i, v := range data {
		dataConvert[i] = v
		//utilities.InfoLog.Printf("Parse Result:%v\n", v)
	}
	//utilities.InfoLog.Printf("=========:%v\n", dataConvert)
	GetManyResponse.Data = dataConvert
	return GetManyResponse, len(data), 200
}

func (m mongoGoDriverRepository) GetByCode(motelCode string) (getOne *model.GetOneResponse, httpCode int) {
	c, _ := primitive.ObjectIDFromHex(motelCode)
	utilities.InfoLog.Printf("modelCode:%v\n", motelCode)
	filter := bson.M{
		"MotelCode": c,
	}
	var motel model.Motel
	m.collection.FindOne(context.TODO(), filter).Decode(&motel)
	utilities.InfoLog.Printf("motel: %v", motel)
	return &model.GetOneResponse{
		Message: "Success",
		Error:   "",
		Data:    motel,
	}, 200
}

func (m mongoGoDriverRepository) MakeIndexes() {
	panic("implement me")
}

// mongoGoDriver

//func NewMongoMgo ()(repository model.MotelRepositoryInterface){
//	var r = & mongoMgoRepository{
//		collection: "",
//	}
//	return r
//}

func NewMongoGoDriver() (repository model.MotelRepositoryInterface) {
	var mongoGoDriver = &infrastructure.MongoGoDriverConnector{}
	client, err := mongoGoDriver.Connect()
	if err != nil {
		return nil
	}
	var r = &mongoGoDriverRepository{
		collection: client.Database("motelfinder").Collection("motel"),
	}
	return r
}
