package handlers

import (
	"fm.auth/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handlerPing struct {
}

func NewHandlerPing() *handlerPing {
	return &handlerPing{}
}

func (h *handlerPing) PingHandler(ctx *gin.Context) {
	helpers.APIResponse(ctx, "pong", http.StatusOK, nil)
}
