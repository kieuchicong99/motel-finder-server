package service

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"server/model"
	"server/repository"
)

type IReportService interface {
	CreateReport(review *model.Report) (insertResponse *model.InsertResponse, httpCode int)
	GetAllReport(page, pageSize int, userID string, motelCode string) (GetManyResponse *model.GetManyResponse, totalResult, httpCode int)
}

func NewReportService() IReportService {
	return &reportServiceIml{
		reportRepository: repository.NewReportRepository(),
	}
}

type reportServiceIml struct {
	reportRepository model.ReportRepositoryInterface
}

func (r *reportServiceIml) CreateReport(review *model.Report) (insertResponse *model.InsertResponse, httpCode int) {
	insertResponse= new(model.InsertResponse)
	err:=r.reportRepository.Create(review)
	if err!= nil{
		insertResponse.Data= review
		insertResponse.Error=err.Error()
		return insertResponse,500
	}
	insertResponse.Data= review
	return insertResponse,200

}

func (r *reportServiceIml) GetAllReport(page, pageSize int, userID string, motelCode string) (GetManyResponse *model.GetManyResponse, totalResult, httpCode int) {
	GetManyResponse= new(model.GetManyResponse)
	data:= []model.Report{}
	filter:= bson.D{

	}
	if userID!=""{
		idO,_:=primitive.ObjectIDFromHex(userID)
		filter=append(filter, bson.E{
			Key:   "UserCode",
			Value: idO,
		})
	}
	if motelCode!=""{
		idO,_:=primitive.ObjectIDFromHex(motelCode)
		filter=append(filter, bson.E{
			Key:   "MotelCode",
			Value: idO,
		})
	}
	total,err:=r.reportRepository.FindAll(filter,page,pageSize,&data,options.Find())
	if err!= nil{
		GetManyResponse.Data= data
		GetManyResponse.Error=err.Error()
		return GetManyResponse,int(total),500
	}
	GetManyResponse.Data= data
	return GetManyResponse,int(total),200
}
