package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	UserCode        primitive.ObjectID `json:"UserCode" bson:"UserCode"`
	UserName        string             `json:"UserName" bson:"UserName" example:"kieuchicong"`
	RoleCode        string             `json:"RoleCode" bson:"RoleCode" example:"admin" `
	FullName        string             `json:"FullName" bson:"FullName" example:"Kiều Chí Công"`
	Phone           string             `json:"Phone" bson:"Phone" example:"0795048768"`
	PassWord        string             `json:"PassWord" bson:"PassWord" example:"abc123"`
	Email           string             `json:"Email" bson:"Email" example:"example@gmail.com" `
	Status          bool               `json:"Status" bson:"Status" example:"true"`
	CreatedOn       time.Time          `json:"CreatedOn" bson:"CreatedOn" example:"2020-11-01 00:00:00" format:"2020-11-01 00:00:00"`
	MotelFavourites map[string]string  `json:"MotelFavourites" bson:"MotelFavourites"`
}

type UserInfo struct {
	UserCode primitive.ObjectID `json:"UserCode" bson:"UserCode"`
	UserName string             `json:"UserName" bson:"UserName" example:"kieuchicong"`
	RoleCode string             `json:"RoleCode" bson:"RoleCode" example:"admin" `
	FullName string             `json:"FullName" bson:"FullName" example:"Kiều Chí Công"`
	PassWord string             `json:"PassWord" bson:"PassWord" example:"abc123"`
	Email    string             `json:"Email" bson:"Email" example:"example@gmail.com"`
	Phone    string             `json:"Phone" bson:"Phone" example:"0795048768"`
}

type CreateUserPayload struct {
	UserName string `json:"UserName" bson:"UserName" example:"kieuchicong"`
	RoleCode string `json:"RoleCode" bson:"RoleCode" example:"ADMIN" `
	FullName string `json:"FullName" bson:"FullName" example:"Kiều Chí Công"`
	PassWord string `json:"PassWord" bson:"PassWord" example:"abc123"`
}
type LoginPayload struct {
	UserName string `json:"UserName" bson:"UserName" example:"kieuchicong"`
	PassWord string `json:"PassWord" bson:"PassWord" example:"abc123"`
}

type UpdateUserInfoPayload struct {
	FullName string `json:"FullName" example:"Kiều Chí Công"`
	Email    string `json:"Email" example:"example@gmail.com"`
	Phone    string `json:"Phone" example:"0795048768"`
}

type ChangePassPayload struct {
	OldPass string `json:"OldPass" example:"abc@123"`
	NewPass string `json:"NewPass" example:"abc@124"`
}

type AddFavouritePayload struct {
	MotelCode string `json:"MotelCode"`
}

type UserRepositoryInterface interface {
	MakeIndexes()
	Insert(user *User) (insertResponse *InsertResponse, httpCode int)
	Update(code string, user *User) (updateResponse *UpdateResponse, httpCode int)
	GetAll(page, pageSize int) (GetManyResponse *GetManyResponse, totalResult int, httpCode int)
	GetByCode(code string) (getOne *GetOneResponse, httpCode int)
	Login(username string, password string) (response *GetOneResponse, httpCode int)
	ChangePass(newPass string, userCode string) (response *GetOneResponse, httpCode int)
	AddMotelFavourites(motelCode string, userCode string) (response *GetOneResponse, httpCode int)
}
