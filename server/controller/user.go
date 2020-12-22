package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	jwt "server/middleware"
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


//@Tags User
// @Summary Lấy danh sách Tài khoản
// @Description GetUsersByFilter
// @Produce  json
// @Security ApiKeyAuth
// @Success 200  {object} model.GetManyResponse
// @Router /user [get]
func (c *Controller) GetUsersByFilter(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.JSON(401, model.GetOneResponse{
			Message: "Fail",
			Error:   "Use invalid Token",
			Data:    nil,
		})
		return
	}
	utilities.InfoLog.Printf("Token: %s", token)
	userInfo, err := jwt.ExtractTokenMetadata(ctx.Request)
	if userInfo== nil ||userInfo.RoleCode != "ADMIN"{
		utilities.InfoLog.Printf("ERR: %v\n", err)
		ctx.JSON(401, model.GetOneResponse{
			Message: "Fail",
			Error:   "Unauthorized or Use invalid Token",
			Data:    nil,
		})
		return
	}
	utilities.InfoLog.Printf("UserInfor: %v", userInfo)
	result, _, httpCode := userService.GetAll(1,20)
	ctx.JSON(httpCode, result)

}

//@Tags User
// @Summary Lấy Thông tin User theo code
// @Description GetUserByCode
// @Produce  json
// @Param  code path string true "User code"
// @Success 200  {object} model.GetOneResponse
// @Router /user/{code} [get]
func (c *Controller) GetUserByCode(ctx *gin.Context) {
	code := ctx.Param("code")
	result, httpCode := userService.GetByCode(code)
	ctx.JSON(httpCode, result)
}

//@Tags User
// @Summary Cập nhật Thông tin tài khoản
// @Description UpdateUser
// @Produce  json
// @Security ApiKeyAuth
// @Param request body model.UpdateUserInfoPayload true "request information"
// @Success 200 {object} model.GetOneResponse
// @Router /user/info [patch]
func (c *Controller) UpdateUser(ctx *gin.Context) {
	var payload model.UpdateUserInfoPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		utilities.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	//utilities.InfoLog.Printf("BODY: %v\n", payload)

	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.JSON(401, model.GetOneResponse{
			Message: "Fail",
			Error:   "Use invalid Token",
			Data:    nil,
		})
		return
	}
	//utilities.InfoLog.Printf("Token: %s", token)
	userInfo, err := jwt.ExtractTokenMetadata(ctx.Request)
	if userInfo== nil ||userInfo.RoleCode != "ADMIN"{
		utilities.InfoLog.Printf("ERR: %v\n", err)
		ctx.JSON(401, model.GetOneResponse{
			Message: "Fail",
			Error:   "Unauthorized or Use invalid Token",
			Data:    nil,
		})
		return
	}
	user := &model.User{
		Phone: payload.Phone,
		FullName: payload.FullName,
		Email: payload.Email,
	}
	utilities.InfoLog.Printf("UserInfo: %s", userInfo)
	result, httpCode := userService.Update( userInfo.UserCode.Hex(), user)
	ctx.JSON(httpCode, result)
}


//@Tags User
// @Summary Đổi mật khẩu tài khoản
// @Description ChangePass
// @Produce  json
// @Security ApiKeyAuth
// @Param request body model.ChangePassPayload true "request information"
// @Success 200 {object} model.GetOneResponse
// @Router /user/change-pass [patch]
func (c *Controller) ChangePass(ctx *gin.Context) {
	var payload model.ChangePassPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		utilities.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	//utilities.InfoLog.Printf("BODY: %v\n", payload)

	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.JSON(401, model.GetOneResponse{
			Message: "Fail",
			Error:   "Use invalid Token",
			Data:    nil,
		})
		return
	}
	//utilities.InfoLog.Printf("Token: %s", token)
	userInfo, err := jwt.ExtractTokenMetadata(ctx.Request)
	if userInfo== nil ||userInfo.RoleCode != "ADMIN"{
		utilities.InfoLog.Printf("ERR: %v\n", err)
		ctx.JSON(401, model.GetOneResponse{
			Message: "Fail",
			Error:   "Unauthorized or Use invalid Token",
			Data:    nil,
		})
		return
	}
	if userInfo.PassWord != payload.OldPass{
		ctx.JSON(400, model.UpdateResponse{
			Message: "Fail",
			Error:   "Fail due to type PassWord incorrect ",
			Data:    nil,
		} )
		return
	}
	result, httpCode := userService.ChangePass(payload.NewPass, userInfo.UserCode.Hex())
	ctx.JSON(httpCode, result)
}


//@Tags User
// @Summary Thêm bài đăng yêu thích
// @Description AddMotelFavourites
// @Produce  json
// @Security ApiKeyAuth
// @Param request body model.AddFavouritePayload true "request information"
// @Success 200 {object} model.GetOneResponse
// @Router /user/add-favourite-motel [post]
func (c *Controller) AddMotelFavourites(ctx *gin.Context) {
	var payload model.AddFavouritePayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		utilities.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	//utilities.InfoLog.Printf("BODY: %v\n", payload)
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.JSON(401, model.GetOneResponse{
			Message: "Fail",
			Error:   "Use invalid Token",
			Data:    nil,
		})
		return
	}
	//utilities.InfoLog.Printf("Token: %s", token)
	userInfo, err := jwt.ExtractTokenMetadata(ctx.Request)
	if userInfo== nil ||userInfo.RoleCode != "ADMIN"{
		utilities.InfoLog.Printf("ERR: %v\n", err)
		ctx.JSON(401, model.GetOneResponse{
			Message: "Fail",
			Error:   "Unauthorized or Use invalid Token",
			Data:    nil,
		})
		return
	}
	result, httpCode := userService.AddMotelFavourites(payload.MotelCode, userInfo.UserCode.Hex() )
	ctx.JSON(httpCode, result)
}