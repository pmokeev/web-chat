package controllers

import (
	"github.com/gin-gonic/gin"
	"pmokeev/web-chat/internal/services"
)

type AuthorizationController interface {
	SignUp(context *gin.Context)
	SignIn(context *gin.Context)
	Logout(context *gin.Context)
	GetProfile(context *gin.Context)
	JWTVerify(context *gin.Context)
	ChangePassword(context *gin.Context)
}

type Controller struct {
	AuthorizationController
}

func NewController(service *services.Service) *Controller {
	return &Controller{AuthorizationController: NewAuthController(service.AuthorizationService)}
}
