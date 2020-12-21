package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"server/model"
	MongoDriver "server/repository"
	UserService "server/service"
	"server/utilities"
	"time"
)

var usermongoDriver = MongoDriver.UserNewMongoGoDriver()
var userService = UserService.NewUser(usermongoDriver)

func init() {
	userService.MakeIndexes()
}

//@Tags User
// @Summary Tạo tài khoản
// @Description CreateUser
// @Produce  json
// @Param request body model.CreateUserPayload true "request information"
// @Success 200 {object} model.GetOneResponse
// @Router /user [post]
func (c *Controller) CreateUser(ctx *gin.Context) {
	utilities.InfoLog.Printf("CreateUser is running")
	var payload model.CreateUserPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		utilities.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	//fmt.Printf("BODY: %v\n", payload)
	newUser := &model.User{
		UserCode: primitive.NewObjectID(),
		UserName:  payload.UserName,
		RoleCode:  payload.RoleCode,
		FullName:  payload.FullName,
		PassWord:  payload.PassWord,
		Email:     "",
		Status:    false,
		CreatedOn: time.Time{},
	}
	result, httpCode := userService.Insert(newUser)

	ctx.JSON(httpCode, result)

}

//@Tags User
// @Summary Đăng nhập tài khoản
// @Description LoginUser
// @Produce  json
// @Param request body model.LoginPayload true "request information"
// @Success 200 {object} model.GetOneResponse
// @Router /user/login [post]
func (c *Controller) LoginUser(ctx *gin.Context) {
	var payload model.LoginPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		utilities.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	//fmt.Printf("BODY: %v\n", payload)
	result, httpCode := userService.Login(payload.UserName, payload.PassWord)

	ctx.JSON(httpCode, result)

}


