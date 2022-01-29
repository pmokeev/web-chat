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
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "GET"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	auth := router.Group("/api/auth")
	{
		auth.POST("/sign-up", authRouter.authController.SignUp)
		auth.POST("/sign-in", authRouter.authController.SignIn)
		auth.POST("/logout", authRouter.authController.Logout)
		auth.GET("/jwtverify", authRouter.authController.JWTVerify)
		auth.GET("/profile", authRouter.authController.GetProfile)
		auth.POST("/change-password", authRouter.authController.ChangePassword)
	}

	return router
}
