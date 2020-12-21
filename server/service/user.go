package service

import (
	"server/model"
)

type userService struct {
	userRepository model.UserRepositoryInterface
}

func (u userService) MakeIndexes() {
 u.userRepository.MakeIndexes()
}

func (u userService) Insert(user *model.User) (insertResponse *model.InsertResponse, httpCode int) {
	insertResponse, httpCode = u.userRepository.Insert(user)
	return insertResponse, httpCode
}

func (u userService) Update(code string, user *model.User) (updateResponse *model.UpdateResponse, httpCode int) {
	panic("implement me")
}

func (u userService) GetAll(page, pageSize int) (GetManyResponse *model.GetManyResponse, totalResult int, httpCode int) {
	panic("implement me")
}

func (u userService) GetByCode(code string) (getOne *model.GetOneResponse, httpCode int) {
	panic("implement me")
}

func (u userService) Login(username string, password string) (updateResponse *model.GetOneResponse, httpCode int) {
	updateResponse, httpCode = u.userRepository.Login(username, password)
	return updateResponse, httpCode
}

func NewUser(repository model.UserRepositoryInterface) model.UserRepositoryInterface {
	return &userService{
		userRepository: repository,
	}
}
