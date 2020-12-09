package httputil

import "github.com/gin-gonic/gin"

// NewError example
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

// HTTPError
type HTTPError struct {
	Code    int    `json:"code" `
	Message string `json:"message" `
}


// HTTPError BadRequest
type ErrBadRequest struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Bad Request"`
}

// HTTPError NotFound
type ErrNotFound struct {
	Code    int    `json:"code" example:"404"`
	Message string `json:"message" example:"Not Found"`
}

// HTTPError InternalServer
type ErrInternalServer struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"InternalServer"`
}
