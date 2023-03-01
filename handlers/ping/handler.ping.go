package ping

import (
	"github.com/gin-gonic/gin"
	"github.com/namth/go-examples/helpers"
	"net/http"
	"time"
)

type handlerPing struct {
}

func NewHandlerPing() *handlerPing {
	return &handlerPing{}
}

func (h *handlerPing) PingHandler(ctx *gin.Context) {
	time.Sleep(5 * time.Second)
	helpers.APIResponse(ctx, "pong", http.StatusOK, http.MethodGet, nil)
}
