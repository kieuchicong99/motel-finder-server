package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/model"
	"server/service"
	"server/utilities"
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
// @Summary Tạo bài đăng
// @Description CreateMotel
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

func (r *reportControllerIml) GetListReport(ctx *gin.Context) {
	userID := ctx.Query("user_code")
	motelCode := ctx.Query("motel_code")
	result, _, httpCode := r.reportService.GetAllReport(1, 10000, userID, motelCode)

	ctx.JSON(httpCode, result)
}
