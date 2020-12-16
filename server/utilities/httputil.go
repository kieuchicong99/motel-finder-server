package utilities

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
	Code    int    `json:"Code" `
	Message string `json:"Message" `
}


// HTTPError BadRequest
type ErrBadRequest struct {
	Code    int    `json:"Code" example:"400"`
	Message string `json:"Message" example:"Bad Request"`
}

// HTTPError NotFound
type ErrNotFound struct {
	Code    int    `json:"Code" example:"404"`
	Message string `json:"Message" example:"Not Found"`
}

// HTTPError InternalServer
type ErrInternalServer struct {
	Code    int    `json:"Code" example:"500"`
	Message string `json:"Message" example:"InternalServer"`
}
