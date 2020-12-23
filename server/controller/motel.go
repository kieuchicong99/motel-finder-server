package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	jwt "server/middleware"
	"server/model"
	MongoDriver "server/repository"
	MotelService "server/service"
	"server/utilities"
	"time"
)

var mongoDriver = MongoDriver.NewMongoGoDriver()
var motelService = MotelService.New(mongoDriver)

//@Tags Motel
// @Summary Tạo bài đăng
// @Description CreateMotel
// @Produce  json
// @Security ApiKeyAuth
// @Param request body model.CreateMotelPayload true "request information"
// @Success 200 {object} model.GetOneResponse
// @Router /motel [post]
func (c *Controller) CreateMotel(ctx *gin.Context) {
	var payload model.CreateMotelPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		utilities.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	//fmt.Printf("BODY: %v\n", payload)
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
	if userInfo == nil || userInfo.RoleCode == "RENTER" {
		utilities.InfoLog.Printf("ERR: %v\n", err)
		ctx.JSON(401, model.GetOneResponse{
			Message: "Fail",
			Error:   "Unauthorized or Use invalid Token",
			Data:    nil,
		})
		return
	}
	utilities.InfoLog.Printf("UserInfor: %v", userInfo)
	// neu nguoi tao la ADMIN thi trang thai mac dinh la true
	// new nguoi tao la OWNER thi trang thai la false
	var status = false
	if userInfo.RoleCode == "ADMIN" {
		status = true
	}

	motel := &model.Motel{
		MotelCode:        primitive.NewObjectID(),
		Address:          payload.Address,
		Images:           payload.Images,
		Image:            payload.Image,
		Tags:             payload.Tags,
		Description:      payload.Description,
		Title:            payload.Title,
		Cost:             payload.Cost,
		Latitude:         payload.Latitude,
		Longitude:        payload.Longitude,
		CreatedAt:        time.Now(),
		FinishedAt:       time.Now().AddDate(0, 1, 0),
		Status:           status,
		Bathroom:         payload.Bathroom,
		Kitchen:          payload.Kitchen,
		HasAirCondition:  payload.HasAirCondition,
		HasBalcony:       payload.HasBalcony,
		ElectricityPrice: payload.ElectricityPrice,
		WaterPrice:       payload.WaterPrice,
		UserCode:         userInfo.UserCode,
	}
	result, httpCode := motelService.Insert(motel)

	ctx.JSON(httpCode, result)

}

//@Tags Motel
// @Summary Cập nhật thông tin bài đăng
// @Description UpdateMotelInfo
// @Produce  json
// @Security ApiKeyAuth
// @Param  code path string true "Motel code"
// @Param request body model.UpdateMotelPayload true "request information"
// @Success 200 {object} model.GetOneResponse
// @Router /motel/info/{code} [patch]
func (c *Controller) UpdateMotelInfo(ctx *gin.Context) {
	var payload model.UpdateMotelPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		utilities.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	utilities.InfoLog.Printf("BODY: %v\n", payload)
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
	if userInfo == nil || userInfo.RoleCode == "RENTER" {
		utilities.InfoLog.Printf("ERR: %v\n", err)
		ctx.JSON(401, model.GetOneResponse{
			Message: "Fail",
			Error:   "Unauthorized or Use invalid Token",
			Data:    nil,
		})
		return
	}
	utilities.InfoLog.Printf("UserInfor: %v", userInfo)
	// neu nguoi tao la ADMIN thi trang thai mac dinh la true
	// new nguoi tao la OWNER thi trang thai la false
	var status = false
	if userInfo.RoleCode == "ADMIN" {
		status = true
	}
	motel := &model.Motel{
		Address:          payload.Address,
		Images:           payload.Images,
		Image:            payload.Image,
		Tags:             payload.Tags,
		Description:      payload.Description,
		Title:            payload.Title,
		Cost:             payload.Cost,
		Latitude:         payload.Latitude,
		Longitude:        payload.Longitude,
		FinishedAt:       payload.FinishedAt,
		Status:           status,
		Bathroom:         payload.Bathroom,
		Kitchen:          payload.Kitchen,
		HasAirCondition:  payload.HasAirCondition,
		HasBalcony:       payload.HasBalcony,
		ElectricityPrice: payload.ElectricityPrice,
		WaterPrice:       payload.WaterPrice,
	}
	code := ctx.Param("code")
	result, httpCode := motelService.UpdateInfo(code, motel)
	ctx.JSON(httpCode, result)
}

//@Tags Motel
// @Summary Phê duyệt trạng thái bài đăng
// @Description UpdateMotelStatus
// @Produce  json
// @Security ApiKeyAuth
// @Param  code path string true "Motel code"
// @Param request body model.UpdateMotelStatusPayload true "request information"
// @Success 200 {object} model.GetOneResponse
// @Router /motel/status/{code} [patch]
func (c *Controller) UpdateMotelStatus(ctx *gin.Context) {
	var payload model.UpdateMotelStatusPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		utilities.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	utilities.InfoLog.Printf("BODY: %v\n", payload)
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
	if userInfo == nil || userInfo.RoleCode != "ADMIN" {
		utilities.InfoLog.Printf("ERR: %v\n", err)
		ctx.JSON(401, model.GetOneResponse{
			Message: "Fail",
			Error:   "Unauthorized or Use invalid Token",
			Data:    nil,
		})
		return
	}
	utilities.InfoLog.Printf("UserInfor: %v", userInfo)
	motel := &model.Motel{
		Status: payload.Status,
	}
	code := ctx.Param("code")
	result, httpCode := motelService.UpdateStatus(code, motel)
	ctx.JSON(httpCode, result)
}

//@Tags Motel
// @Summary Cập nhật trạng thái đã/chưa có người thuê
// @Description UpdateMotelAvailable
// @Produce  json
// @Security ApiKeyAuth
// @Param  code path string true "Motel code"
// @Param request body model.UpdateMotelAvailablePayload true "request information"
// @Success 200 {object} model.GetOneResponse
// @Router /motel/available/{code} [patch]
func (c *Controller) UpdateMotelAvailable(ctx *gin.Context) {
	var payload model.UpdateMotelAvailablePayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		utilities.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	utilities.InfoLog.Printf("BODY: %v\n", payload)
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
	if userInfo == nil || userInfo.RoleCode == "RENTER" {
		utilities.InfoLog.Printf("ERR: %v\n", err)
		ctx.JSON(401, model.GetOneResponse{
			Message: "Fail",
			Error:   "Unauthorized or Use invalid Token",
			Data:    nil,
		})
		return
	}
	motel := &model.Motel{
		Available: payload.Available,
	}
	code := ctx.Param("code")
	result, httpCode := motelService.UpdateAvailable(code, motel)
	ctx.JSON(httpCode, result)
}

//@Tags Motel
// @Summary Lấy danh sách bài đăng
// @Description GetMotelsByFilter
// @Produce  json
// @Security ApiKeyAuth
// @Success 200  {object} model.GetManyResponse
// @Router /motel [get]
func (c *Controller) GetMotelsByFilter(ctx *gin.Context) {
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
	if userInfo == nil || userInfo.RoleCode != "ADMIN" {
		utilities.InfoLog.Printf("ERR: %v\n", err)
		ctx.JSON(401, model.GetOneResponse{
			Message: "Fail",
			Error:   "Unauthorized or Use invalid Token",
			Data:    nil,
		})
		return
	}
	utilities.InfoLog.Printf("UserInfor: %v", userInfo)
	result, _, httpCode := motelService.GetAll(1, 20)
	ctx.JSON(httpCode, result)

}

//@Tags Motel
// @Summary Lấy danh sách bài đăng theo OWNER
// @Description GetMotelsByOwner
// @Produce  json
// @Security ApiKeyAuth
// @Success 200  {object} model.GetManyResponse
// @Router /motel/by-owner/list-motel [get]
func (c *Controller) GetMotelsByOwner(ctx *gin.Context) {
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
	if userInfo == nil || userInfo.RoleCode != "ADMIN" {
		utilities.InfoLog.Printf("ERR: %v\n", err)
		ctx.JSON(401, model.GetOneResponse{
			Message: "Fail",
			Error:   "Unauthorized or Use invalid Token",
			Data:    nil,
		})
		return
	}
	utilities.InfoLog.Printf("UserInfor: %v", userInfo)
	result, _, httpCode := motelService.GetAllByOwnerCode(userInfo.UserCode.Hex())
	ctx.JSON(httpCode, result)

}

//@Tags Motel
// @Summary Lấy bài đăng theo code
// @Description GetMotelByCode
// @Produce  json
// @Param  code path string true "Motel code"
// @Success 200  {object} model.GetOneResponse
// @Router /motel/by-code/{code} [get]
func (c *Controller) GetMotelByCode(ctx *gin.Context) {
	code := ctx.Param("code")
	result, httpCode := motelService.GetByCode(code)
	ctx.JSON(httpCode, result)
}
