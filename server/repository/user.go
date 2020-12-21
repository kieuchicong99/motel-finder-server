package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"server/infrastructure"
	"server/model"
	"server/utilities"
	jwt "server/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

// mongoGoDriver
type UserMongoGoDriverRepository struct {
	collection *mongo.Collection
}

func (u UserMongoGoDriverRepository) MakeIndexes() {
	_, _ = u.collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "UserName", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
}

func (u UserMongoGoDriverRepository) Insert(user *model.User) (insertResponse *model.InsertResponse, httpCode int) {
	utilities.InfoLog.Printf("Insert in Repository\n")
	result, err := u.collection.InsertOne(context.TODO(), user)
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
	var userCreated *model.User
	err = u.collection.FindOne(context.TODO(), filler).Decode(&userCreated)
	//utilities.InfoLog.Printf("Result : ", motelCreated)
	return &model.InsertResponse{
		Message: "Success",
		Error:   "",
		Data:    userCreated,
	}, 200
}

func (u UserMongoGoDriverRepository) Update(code string, user *model.User) (updateResponse *model.UpdateResponse, httpCode int) {
	opt := options.FindOneAndUpdate().SetReturnDocument(options.After)
	c, _ := primitive.ObjectIDFromHex(code)
	filter := bson.M{
		"UserCode": c,
	}
	update := bson.D{
		{"$set", user},
	}
	var userUpdated model.User
	u.collection.FindOneAndUpdate(context.TODO(), filter, update, opt).Decode(&userUpdated)
	utilities.InfoLog.Printf("motelUpdated: %v", userUpdated)
	return &model.UpdateResponse{
		Message: "Success",
		Error:   "",
		Data:    userUpdated,
	}, 200
}

func (u UserMongoGoDriverRepository) GetAll(page, pageSize int) (GetManyResponse *model.GetManyResponse, totalResult int, httpCode int) {
	query := bson.M{"_id": bson.M{"$exists": true}}
	opts := options.Find().SetSort(bson.M{"code": 1}).SetSkip(int64((page - 1) * pageSize)).SetLimit(int64(pageSize))

	cursor, err := u.collection.Find(context.TODO(), query, opts)
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

	data := []model.User{}
	if err := cursor.All(context.TODO(), &data); err != nil {
		utilities.ErrLog.Printf("Mongo cursor.All, err:%v\n", err)
		GetManyResponse.Error = fmt.Sprintf("%s", err)
		return GetManyResponse, 0, 500
	}

	GetManyResponse.Message = "Success"
	dataConvert := make([]interface{}, len(data))
	for i, v := range data {
		dataConvert[i] = v
	}
	GetManyResponse.Data = dataConvert
	return GetManyResponse, len(data), 200
}

func (u UserMongoGoDriverRepository) GetByCode(code string) (getOne *model.GetOneResponse, httpCode int) {
	c, _ := primitive.ObjectIDFromHex(code)
	utilities.InfoLog.Printf("UserCode:%v\n", code)
	filter := bson.M{
		"MotelCode": c,
	}
	var user model.User
	u.collection.FindOne(context.TODO(), filter).Decode(&user)
	utilities.InfoLog.Printf("user: %v", user)
	return &model.GetOneResponse{
		Message: "Success",
		Error:   "",
		Data:    user,
	}, 200
}

func (u UserMongoGoDriverRepository) Login(username string, password string) (updateResponse *model.GetOneResponse, httpCode int) {
	filter := bson.M{
		"UserName": username,
	}
	var user model.User
	u.collection.FindOne(context.TODO(), filter).Decode(&user)
	utilities.InfoLog.Printf("user: %v", user)
	token, err:=jwt.CreateToken(&user)
	if err != nil {
		return &model.GetOneResponse{
			Message: "Error",
			Error:   fmt.Sprintf("%v", err),
			Data:    nil,
		}, 500
	}
	return &model.GetOneResponse{
		Message: "Success",
		Error:   "",
		Data:    token,
	}, 200
}

func UserNewMongoGoDriver() (repository model.UserRepositoryInterface) {
	var mongoGoDriver = &infrastructure.MongoGoDriverConnector{}
	client, err := mongoGoDriver.Connect()
	if err != nil {
		return nil
	}
	var r = &UserMongoGoDriverRepository{
		collection: client.Database("motelfinder").Collection("user"),
	}
	return r
}
