package service

import (
	"server/model"
)

type motelService struct {
	motelRepository model.MotelRepositoryInterface
}

func (s motelService) MakeIndexes() {
	panic("implement me")
}

func (s motelService) Insert(motel *model.Motel) (insertResponse *model.InsertResponse, httpCode int) {
	insertResponse, httpCode = s.motelRepository.Insert(motel)
	return insertResponse, httpCode
}

func (s motelService) Update(code string, motel *model.Motel) (updateResponse *model.UpdateResponse, httpCode int) {
	updateResponse, httpCode = s.motelRepository.Update(code, motel)
	return updateResponse, httpCode
}

func (s motelService) GetAll(page, pageSize int) (GetManyResponse *model.GetManyResponse, totalResult int, httpCode int) {
	GetManyResponse, totalResult, httpCode = s.motelRepository.GetAll(page, pageSize)
	return GetManyResponse, totalResult, httpCode
}

func (s motelService) GetByCode(code string) (getOne *model.GetOneResponse, httpCode int) {
	getOne, httpCode = s.motelRepository.GetByCode(code)
	return getOne, httpCode
}

func New(repository model.MotelRepositoryInterface) model.MotelRepositoryInterface {
	return &motelService{
		motelRepository: repository,
	}
}
