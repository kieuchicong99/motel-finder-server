package service

import (
	"server/model"
)

type userService struct {
	userRepository model.UserRepositoryInterface
}

func (u userService) RemoveMotelFavourites(motelCode string, userCode string) (response *model.GetOneResponse, httpCode int) {
	return u.userRepository.RemoveMotelFavourites(motelCode, userCode)
}

func (u userService) AddMotelFavourites(motelCode string, userCode string) (response *model.GetOneResponse, httpCode int) {
	return u.userRepository.AddMotelFavourites(motelCode, userCode)
}

func (u userService) ChangePass(newPass string, userCode string) (response *model.GetOneResponse, httpCode int) {
	return u.userRepository.ChangePass(newPass, userCode)
}

func (u userService) MakeIndexes() {
	u.userRepository.MakeIndexes()
}

func (u userService) Insert(user *model.User) (insertResponse *model.InsertResponse, httpCode int) {
	insertResponse, httpCode = u.userRepository.Insert(user)
	return insertResponse, httpCode
}

func (u userService) Update(code string, user *model.User) (updateResponse *model.UpdateResponse, httpCode int) {
	return u.userRepository.Update(code, user)
}

func (u userService) GetAll(page, pageSize int) (GetManyResponse *model.GetManyResponse, totalResult int, httpCode int) {
	return u.userRepository.GetAll(1, 20)
}

func (u userService) GetByCode(code string) (getOne *model.GetOneResponse, httpCode int) {
	return u.userRepository.GetByCode(code)
}

func (u userService) Login(username string, password string) (updateResponse *model.GetOneResponse, httpCode int) {
	return u.userRepository.Login(username, password)
}

func NewUser(repository model.UserRepositoryInterface) model.UserRepositoryInterface {
	return &userService{
		userRepository: repository,
	}
}
