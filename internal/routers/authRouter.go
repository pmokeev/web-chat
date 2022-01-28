package routers

import (
	"github.com/gin-gonic/gin"
	"pmokeev/web-chat/internal/controllers"
	"pmokeev/web-chat/internal/services"
)

type AuthRouter struct {
	authController *controllers.AuthController
}

func NewAuthRouter(authService *services.AuthService) *AuthRouter {
	return &AuthRouter{authController: controllers.NewAuthController(authService)}
}

func (authRouter *AuthRouter) InitAuthRouter() *gin.Engine {
	router := gin.New()

	auth := router.Group("/api/auth")
	{
		auth.GET("/hello", authRouter.authController.Hello)
	}

	return router
}
