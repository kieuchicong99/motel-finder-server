package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/model"
	"server/service"
	"server/utilities"
	"strconv"
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
// @Summary Cập nhật Đánh giá bài đăng
// @Description UpdateReviewStatus
// @Produce  json
// @Param  reviewID path string true " Review ID"
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
// @Summary Đánh giá bài đăng
// @Description CreateReview
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
//@Tags Review
// @Summary Lấy danh sách Đánh giá bài đăng
// @Description GetListReview
// @Produce  json
// @Param status query string false "status"
// @Param user_code query integer false "user_code"
// @Param motel_code query integer false "motel_code"
// @Param page query integer false "page"
// @Param page_size query integer false "page_size"
// @Success 200 {object} model.InsertResponse
// @Router /reviews [post]
func (r *reviewControllerIml) GetListReview(ctx *gin.Context) {
	sta := ctx.Query("status")
	status:= false
	if sta!="true"{
		status= true
	}
	userID := ctx.Query("user_code")
	motelCode := ctx.Query("motel_code")
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ :=  strconv.Atoi(ctx.Query("page_size"))
	result,_,httpCode := r.reviewService.GetAllReview(page ,pageSize, userID, status, motelCode)

	ctx.JSON(httpCode, result)
}
