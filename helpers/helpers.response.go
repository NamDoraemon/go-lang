package helpers

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Method     string      `json:"method"`
	StatusCode int         `json:"status_code"`
	Status     bool        `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func APIResponse(ctx *gin.Context, Message string, StatusCode int, Method string, Data interface{}) {
	jsonResponse := &Response{
		StatusCode: StatusCode,
		Method:     Method,
		Message:    Message,
		Data:       Data,
	}

	if StatusCode >= 400 {
		ctx.AbortWithStatusJSON(StatusCode, jsonResponse)
	} else {
		ctx.JSON(StatusCode, jsonResponse)
	}
}
