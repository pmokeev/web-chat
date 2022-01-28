package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"pmokeev/web-chat/internal/controllers"
	"pmokeev/web-chat/internal/services"
	"time"
)

type AuthRouter struct {
	authController *controllers.AuthController
}

func NewAuthRouter(authService *services.AuthService) *AuthRouter {
	return &AuthRouter{authController: controllers.NewAuthController(authService)}
}

func (authRouter *AuthRouter) InitAuthRouter() *gin.Engine {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://localhost:5000"},
		AllowMethods:     []string{"POST", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://localhost:5000"
		},
		MaxAge: 24 * time.Hour,
	}))

	auth := router.Group("/api/auth")
	{
		auth.POST("/sign-up", authRouter.authController.SignUp)
		auth.POST("/sign-in", authRouter.authController.SignIn)
		auth.POST("/logout", authRouter.authController.Logout)
	}

	return router
}
