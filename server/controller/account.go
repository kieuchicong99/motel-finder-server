package controller

//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"server/model"
//	"server/utilities"
//	"strconv"
//	"time"
//)
//
////@Tags User
//// ShowAccount godoc
//// @Summary Show an account
//// @Description get string by ID
//// @Accept  json
//// @Produce  json
//// @Param id path int true "Account ID"
//// @Success 200 {object} model.User
//// @Failure 400 {object} utilities.HTTPError
//// @Failure 404 {object} utilities.HTTPError
//// @Failure 500 {object} utilities.HTTPError
//// @Router /accounts/{id} [get]
//func (c *Controller) ShowAccount(ctx *gin.Context) {
//	id := ctx.Param("id")
//	aid, err := strconv.Atoi(id)
//	if err != nil {
//		utilities.NewError(ctx, http.StatusBadRequest, err)
//		return
//	}
//	account, err := model.AccountOne(aid)
//	if err != nil {
//		utilities.NewError(ctx, http.StatusNotFound, err)
//		return
//	}
//	ctx.JSON(http.StatusOK, account)
//}
//
////@Tags User
//// ListAccounts godoc
//// @Summary List accounts
//// @Description get accounts
//// @Accept  json
//// @Produce  json
//// @Param q query string false "name search by q" Format(email)
//// @Success 200 {array} model.User
//// @Failure 400 {object} utilities.ErrBadRequest
//// @Failure 404 {object} utilities.ErrNotFound
//// @Failure 500 {object} utilities.ErrInternalServer
//// @Router /accounts [get]
//func (c *Controller) ListAccounts(ctx *gin.Context) {
//
//	//q := ctx.Request.URL.Query().Get("q")
//	//accounts, err := model.AccountsAll(q)
//	//rows, err := db.Query(`SELECT username, password, email, created_on  from accounts`)
//	//
//	//if err != nil {
//	//	utilities.NewError(ctx, http.StatusNotFound, err)
//	//	return
//	//}
//	//var users []*model.User // declare a slice of courses that will hold all of the Course instances scanned from the rows object
//	//for rows.Next() { // this stops when there are no more rows
//	//	c := new(model.User) // initialize a new instance
//	//	rows.Scan(&c.UserName,&c.PassWord, &c.Email, &c.CreatedOn) // scan contents of the current row into the instance
//	//	fmt.Printf("col1: %v, col2: %v \n", c.UserName, c.Email, c.CreatedOn)
//	//	users = append(users, c) // add each instance to the slice
//	//}
//	user:= &model.User{
//		UserName:  "",
//		PassWord:  "",
//		Email:     "",
//		CreatedOn: time.Time{},
//	}
//	ctx.JSON(http.StatusOK, user)
//}
//
////@Tags User
//// AddAccount godoc
//// @Summary Add an account
//// @Description add by json account
//// @Accept  json
//// @Produce  json
//// @Param account body model.AddAccount true "Add account"
//// @Success 200 {array} model.User
//// @Failure 400 {object} utilities.ErrBadRequest
//// @Failure 404 {object} utilities.ErrNotFound
//// @Failure 500 {object} utilities.ErrInternalServer
//// @Router /accounts [post]
//func (c *Controller) AddAccount(ctx *gin.Context) {
//	var addAccount model.AddAccount
//	if err := ctx.ShouldBindJSON(&addAccount); err != nil {
//		utilities.NewError(ctx, http.StatusBadRequest, err)
//		return
//	}
//	if err := addAccount.Validation(); err != nil {
//		utilities.NewError(ctx, http.StatusBadRequest, err)
//		return
//	}
//	account := model.Account{
//		Name: addAccount.Name,
//	}
//	lastID, err := account.Insert()
//	if err != nil {
//		utilities.NewError(ctx, http.StatusBadRequest, err)
//		return
//	}
//	account.ID = lastID
//	ctx.JSON(http.StatusOK, account)
//}
//
////@Tags User
//// UpdateAccount godoc
//// @Summary Update an account
//// @Description Update by json account
//// @Accept  json
//// @Produce  json
//// @Param  id path int true "Account ID"
//// @Param  account body model.UpdateAccount true "Update account"
//// @Success 200 {array} model.User
//// @Failure 400 {object} utilities.ErrBadRequest
//// @Failure 404 {object} utilities.ErrNotFound
//// @Failure 500 {object} utilities.ErrInternalServer
//// @Router /accounts/{id} [patch]
//func (c *Controller) UpdateAccount(ctx *gin.Context) {
//	id := ctx.Param("id")
//	aid, err := strconv.Atoi(id)
//	if err != nil {
//		utilities.NewError(ctx, http.StatusBadRequest, err)
//		return
//	}
//	var updateAccount model.UpdateAccount
//	if err := ctx.ShouldBindJSON(&updateAccount); err != nil {
//		utilities.NewError(ctx, http.StatusBadRequest, err)
//		return
//	}
//	account := model.Account{
//		ID:   aid,
//		Name: updateAccount.Name,
//	}
//	err = account.Update()
//	if err != nil {
//		utilities.NewError(ctx, http.StatusNotFound, err)
//		return
//	}
//	ctx.JSON(http.StatusOK, account)
//}
//
////@Tags User
//// DeleteAccount godoc
//// @Summary Delete an account
//// @Description Delete by account ID
//// @Accept  json
//// @Produce  json
//// @Param  id path int true "Account ID" Format(int64)
//// @Success 204 {object} model.User
//// @Failure 400 {object} utilities.ErrBadRequest
//// @Failure 404 {object} utilities.ErrNotFound
//// @Failure 500 {object} utilities.ErrInternalServer
//// @Router /accounts/{id} [delete]
//func (c *Controller) DeleteAccount(ctx *gin.Context) {
//	id := ctx.Param("id")
//	aid, err := strconv.Atoi(id)
//	if err != nil {
//		utilities.NewError(ctx, http.StatusBadRequest, err)
//		return
//	}
//	err = model.Delete(aid)
//	if err != nil {
//		utilities.NewError(ctx, http.StatusNotFound, err)
//		return
//	}
//	ctx.JSON(http.StatusNoContent, gin.H{})
//}
//
////@Tags User
//// UploadAccountImage godoc
//// @Summary Upload account image
//// @Description Upload file
//// @Accept  multipart/form-data
//// @Produce  json
//// @Param  id path int true "Account ID"
//// @Param file formData file true "account image"
//// @Success 200 {object} model.User
//// @Failure 400 {object} utilities.ErrBadRequest
//// @Failure 404 {object} utilities.ErrNotFound
//// @Failure 500 {object} utilities.ErrInternalServer
//// @Router /accounts/{id}/images [post]
//func (c *Controller) UploadAccountImage(ctx *gin.Context) {
//	id, err := strconv.Atoi(ctx.Param("id"))
//	if err != nil {
//		utilities.NewError(ctx, http.StatusBadRequest, err)
//		return
//	}
//	file, err := ctx.FormFile("file")
//	if err != nil {
//		utilities.NewError(ctx, http.StatusBadRequest, err)
//		return
//	}
//	ctx.JSON(http.StatusOK, Message{Message: fmt.Sprintf("upload complete userID=%d filename=%s", id, file.Filename)})
//}
