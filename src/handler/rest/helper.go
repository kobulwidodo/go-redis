package rest

import (
	"medium-go-redis/src/business/entity"

	"github.com/gin-gonic/gin"
)

func (r *rest) httpRespSuccess(ctx *gin.Context, code int, message string, data interface{}) {
	resp := entity.Response{
		Meta: entity.Meta{
			Message: message,
			Code:    code,
			IsError: false,
		},
		Data: data,
	}
	ctx.JSON(code, resp)
}

func (r *rest) httpRespError(ctx *gin.Context, code int, err error) {
	resp := entity.Response{
		Meta: entity.Meta{
			Message: err.Error(),
			Code:    code,
			IsError: true,
		},
		Data: nil,
	}
	ctx.AbortWithStatusJSON(code, resp)
}
