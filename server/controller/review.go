package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/model"
	"server/service"
	"server/utilities"
)

type IReviewController interface {
	CreateReview(ctx *gin.Context)
	GetListReview(ctx *gin.Context)
	UpdateReviewStatus(ctx *gin.Context)
}

func NewReviewController() IReviewController {
	return &reviewControllerIml{
		reviewService: service.NewReviewService(),
	}
}

type reviewControllerIml struct {
	reviewService service.IReviewService
}

func (r *reviewControllerIml) CreateReport(ctx *gin.Context) {
	panic("implement me")
}

func (r *reviewControllerIml) GetListReport(ctx *gin.Context) {
	panic("implement me")
}

//@Tags Review
// @Summary Tạo bài đăng
// @Description CreateMotel
// @Produce  json
// @Param  reviewID path string true "Motel code"
// @Param request body model.UpdateStatusReview true "request information"
// @Success 200 {object} model.InsertResponse
// @Router /reviews/{reviewID} [patch]
func (r *reviewControllerIml) UpdateReviewStatus(ctx *gin.Context) {
	id:=ctx.Param("reviewID")
	var payload model.UpdateStatusReview
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		utilities.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	result, httpCode := r.reviewService.UpdateStatusReview(id,payload.Status)

	ctx.JSON(httpCode, result)
}

//@Tags Review
// @Summary Tạo bài đăng
// @Description CreateMotel
// @Produce  json
// @Param request body model.CreateReviewPayload true "request information"
// @Success 200 {object} model.InsertResponse
// @Router /reviews [post]
func (r *reviewControllerIml) CreateReview(ctx *gin.Context) {
	var payload model.CreateReviewPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		utilities.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	review := new(model.Review)
	review.BaseReview = payload.BaseReview
	result, httpCode := r.reviewService.CreateReview(review)

	ctx.JSON(httpCode, result)
}

func (r *reviewControllerIml) GetListReview(ctx *gin.Context) {
	sta := ctx.Query("status")
	status:= false
	if sta!="true"{
		status= true
	}
	userID := ctx.Query("user_code")
	motelCode := ctx.Query("motel_code")
	result,_,httpCode := r.reviewService.GetAllReview(1,10000,userID,status,motelCode)

	ctx.JSON(httpCode, result)
}
