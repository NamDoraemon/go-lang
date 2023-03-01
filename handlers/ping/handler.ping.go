package ping

import (
	"github.com/gin-gonic/gin"
	"github.com/namth/go-examples/helpers"
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
