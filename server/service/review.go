package service

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"server/model"
	"server/repository"
	"server/utilities"
)

type IReviewService interface {
	CreateReview(review *model.Review) (insertResponse *model.InsertResponse, httpCode int)
	UpdateReview(code string, review *model.Review) (updateResponse *model.UpdateResponse, httpCode int)
	GetAllReview(page, pageSize int, userID string, status bool, motelCode string) (GetManyResponse *model.GetManyResponse, totalResult, httpCode int)
	GetByCode(code string) (getOne *model.GetOneResponse, httpCode int)
	UpdateStatusReview(id string,status bool)(insertResponse *model.InsertResponse, httpCode int)
}

func NewReviewService() IReviewService {
	return &reviewServiceIml{
		reviewRepository: repository.NewReviewRepository(),
	}
}
type reviewServiceIml struct {
	reviewRepository model.ReviewRepositoryInterface
}

func (r *reviewServiceIml) UpdateStatusReview(id string, status bool) (insertResponse *model.InsertResponse, httpCode int) {
	idO,_:=primitive.ObjectIDFromHex(id)
	review:=new(model.Review)
	insertResponse= new(model.InsertResponse)
	utilities.InfoLog.Print(status)
	err:=r.reviewRepository.UpdateStatusOne("_id",idO,status,review)
	if err!= nil{
		insertResponse.Data= review
		insertResponse.Error=err.Error()
		return insertResponse,500
	}
	insertResponse.Data= review
	return insertResponse,200
}

func (r *reviewServiceIml) CreateReview(review *model.Review) (insertResponse *model.InsertResponse, httpCode int) {
	 insertResponse= new(model.InsertResponse)
      err:=r.reviewRepository.Create(review)
      if err!= nil{
        insertResponse.Data= review
        insertResponse.Error=err.Error()
       return insertResponse,500
	  }
	  insertResponse.Data= review
	  return insertResponse,200

}

func (r *reviewServiceIml) UpdateReview(code string, review *model.Review) (updateResponse *model.UpdateResponse, httpCode int) {
	panic("implement me")
}

func (r *reviewServiceIml) GetAllReview(page, pageSize int, userID string, status bool, motelCode string) (GetManyResponse *model.GetManyResponse, totalResult, httpCode int) {
	GetManyResponse= new(model.GetManyResponse)
	data:= []model.Review{}
	filter:= bson.D{
	{"Status",status},
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
	total,err:=r.reviewRepository.FindAll(filter,page,pageSize,&data,options.Find())
	if err!= nil{
		GetManyResponse.Data= data
		GetManyResponse.Error=err.Error()
		return GetManyResponse,int(total),500
	}
	GetManyResponse.Data= data
	return GetManyResponse,int(total),200
}

func (r *reviewServiceIml) GetByCode(code string) (getOne *model.GetOneResponse, httpCode int) {
	panic("implement me")
}

