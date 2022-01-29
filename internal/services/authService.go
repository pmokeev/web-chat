package services

import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"os"
	"pmokeev/web-chat/internal/models"
	"pmokeev/web-chat/internal/storage"
	"strconv"
	"time"
)

type AuthService struct {
	authStorage *storage.AuthStorage
}

func NewAuthService(authStorage *storage.AuthStorage) *AuthService {
	return &AuthService{authStorage: authStorage}
}

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashPassword), err
}

func CompareHashPasswords(correctPassword, requestPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(correctPassword), []byte(requestPassword))
	return err
}

func (authService *AuthService) SignUP(registerForm models.RegisterForm) error {
	hashPassword, err := HashPassword(registerForm.PasswordHash)
	if err != nil {
		return err
	}

	registerForm.PasswordHash = hashPassword
	err = authService.authStorage.AddNewUser(registerForm)

	return err
}

func (authService *AuthService) SignIn(loginForm models.LoginForm) (string, error) {
	user, err := authService.authStorage.GetUserPassword(loginForm)
	if err != nil {
		return "", err
	}

	err = CompareHashPasswords(user.PasswordHash, loginForm.Password)
	if err != nil {
		return "", err
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(user.ID),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(os.Getenv("JWTSecretKey")))
	return token, err
}

func (authService *AuthService) JWTVerify(JWTTokenString string) (bool, error) {
	decodedToken, err := jwt.ParseWithClaims(JWTTokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWTSecretKey")), nil
	})

	return decodedToken.Valid, err
}
