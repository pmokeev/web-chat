package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pmokeev/web-chat/internal/models"
	"pmokeev/web-chat/internal/services"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (authController *AuthController) SignUp(context *gin.Context) {
	var registerForm models.RegisterForm
	if err := context.BindJSON(&registerForm); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := authController.authService.SignUP(registerForm); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]string{
		"register": "ok",
	})
}

func (authController *AuthController) SignIn(context *gin.Context) {
	var loginForm models.LoginForm
	if err := context.BindJSON(&loginForm); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	token, err := authController.authService.SignIn(loginForm)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.SetCookie("jwt", token, 3600, "/", "localhost", false, true)
	context.JSON(http.StatusOK, map[string]string{
		"login": "correct",
	})
}

func (authController *AuthController) Logout(context *gin.Context) {
	context.SetCookie("jwt", "", -1, "/", "localhost", false, true)
	context.JSON(http.StatusOK, map[string]string{
		"logout": "correct",
	})
}
