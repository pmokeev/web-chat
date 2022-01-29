package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"net/http"
	"pmokeev/web-chat/internal/models"
	"pmokeev/web-chat/internal/services"
	"regexp"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (authController *AuthController) JWTVerify(context *gin.Context) {
	cookie, err := context.Request.Cookie("jwt")
	if err != nil {
		if err == http.ErrNoCookie {
			context.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	JWTTokenString := cookie.Value
	isValid, err := authController.authService.JWTVerify(JWTTokenString)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			context.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	if !isValid {
		context.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]string{
		"verify": "ok",
	})
}

func (authController *AuthController) SignUp(context *gin.Context) {
	var registerForm models.RegisterForm
	if err := context.BindJSON(&registerForm); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := authController.authService.SignUP(registerForm); err != nil {
		duplicate := regexp.MustCompile(`\(SQLSTATE 23505\)$`)
		if duplicate.MatchString(err.Error()) {
			context.AbortWithStatusJSON(http.StatusConflict, err.Error())
			return
		}
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
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			context.AbortWithStatusJSON(http.StatusConflict, err.Error())
			return
		}
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.SetCookie("jwt", token, 60*60*24, "/", "localhost", false, true)
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

func (authController *AuthController) GetProfile(context *gin.Context) {
	cookie, err := context.Request.Cookie("jwt")
	if err != nil {
		if err == http.ErrNoCookie {
			context.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	JWTTokenString := cookie.Value
	userProfile, err, isValid := authController.authService.GetUserInformation(JWTTokenString)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			context.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	if !isValid {
		context.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}

	context.JSON(http.StatusOK, userProfile)
}

func (authController *AuthController) ChangePassword(context *gin.Context) {
	var changePasswordForm models.ChangePassword
	if err := context.BindJSON(&changePasswordForm); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	err := authController.authService.ChangeUserPassword(changePasswordForm)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]string{
		"changed": "correct",
	})
}
