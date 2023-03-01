package auth

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/namth/go-examples/entities"
	"github.com/namth/go-examples/helpers"
	"github.com/namth/go-examples/repositories"
	"io/ioutil"
	"net/http"
)

type handlerAuth struct {
}

func NewHandlerAuth() *handlerAuth {
	return &handlerAuth{}
}

func (h *handlerAuth) LoginHandler(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Chuyển đổi dữ liệu JSON thành đối tượng User
	var login entities.Login
	err = json.Unmarshal(body, &login)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err, token := repositories.RepoLogin(login.Email, login.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := entities.Token{Token: token}
	helpers.APIResponse(ctx, "Login successful", http.StatusOK, resp)
}

func (h *handlerAuth) ValidateTokenHandler(ctx *gin.Context) {
	helpers.APIResponse(ctx, "pong", http.StatusOK, nil)
}
