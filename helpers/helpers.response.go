package helpers

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func APIResponse(ctx *gin.Context, Message string, StatusCode int, Data interface{}) {
	jsonResponse := &Response{
		Message: Message,
		Data:    Data,
	}

	if StatusCode >= 0 {
		jsonResponse.Status = false
		ctx.AbortWithStatusJSON(StatusCode, jsonResponse)
	} else {
		jsonResponse.Status = true
		ctx.JSON(StatusCode, jsonResponse)
	}
}
