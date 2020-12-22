package controller

import "github.com/gin-gonic/gin"

type IReportController interface {
	CreateReport(ctx *gin.Context)
	GetListReport(ctx *gin.Context)
}

func NewReportController()IReportController  {
	return &reportControllerIml{
		
	}
}
type  reportControllerIml struct {

}

func (r reportControllerIml) CreateReport(ctx *gin.Context) {
	panic("implement me")
}

func (r reportControllerIml) GetListReport(ctx *gin.Context) {
	panic("implement me")
}
