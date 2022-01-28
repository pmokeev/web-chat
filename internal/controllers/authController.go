package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pmokeev/web-chat/internal/services"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (authController *AuthController) Hello(context *gin.Context) {
	fmt.Println("Hello world!")

	context.JSON(http.StatusOK, map[string]string{
		"ok": "ok",
	})
}
