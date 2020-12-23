package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/model"
	"server/service"
	"server/utilities"
	"strconv"
)

type IReportController interface {
	CreateReport(ctx *gin.Context)
	GetListReport(ctx *gin.Context)
}

func NewReportController()IReportController  {
	return &reportControllerIml{
		reportService: service.NewReportService(),
	}
}
type  reportControllerIml struct {
	reportService service.IReportService
}
//@Tags Report
// @Summary Report bài đăng vi phạm
// @Description CreateReport
// @Produce  json
// @Param request body model.CreateReportPayload true "request information"
// @Success 200 {object} model.InsertResponse
// @Router /reports [post]
func (r *reportControllerIml) CreateReport(ctx *gin.Context) {
	var payload model.CreateReportPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		utilities.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	report := new(model.Report)
	report.BaseReport = payload.BaseReport
	result, httpCode := r.reportService.CreateReport(report)

	ctx.JSON(httpCode, result)
}

//@Tags Report
// @Summary Lấy danh dách Report của một bài đăng
// @Description GetListReport
// @Produce  json
// @Param user_code query integer false "user_code"
// @Param motel_code query integer false "motel_code"
// @Param page query integer false "page"
// @Param page_size query integer false "page_size"
// @Success 200 {object} model.GetManyResponse
// @Router /reports [get]
func (r *reportControllerIml) GetListReport(ctx *gin.Context) {
	userID := ctx.Query("user_code")
	motelCode := ctx.Query("motel_code")
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ :=  strconv.Atoi(ctx.Query("page_size"))
	result, _, httpCode := r.reportService.GetAllReport(page, pageSize, userID, motelCode)

	ctx.JSON(httpCode, result)
}
