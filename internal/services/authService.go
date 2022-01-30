package services

import (
	"fmt"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"os"
	"pmokeev/web-chat/internal/models"
	"pmokeev/web-chat/internal/storage"
	"strconv"
	"time"
)

type AuthService struct {
	authStorage storage.AuthorizationStorage
}

func NewAuthService(authStorage storage.AuthorizationStorage) *AuthService {
	return &AuthService{authStorage: authStorage}
}

func (authService *AuthService) SignUP(registerForm models.RegisterForm) error {
	hashPassword, err := HashPassword(registerForm.PasswordHash)
	if err != nil {
		return err
	}

	registerForm.PasswordHash = hashPassword
	err = authService.authStorage.CreateUser(registerForm)

	return err
}

func (authService *AuthService) SignIn(loginForm models.LoginForm) (string, error) {
	user, err := authService.authStorage.GetUser(loginForm)
	if err != nil {
		return "", err
	}

	err = CompareHashPasswords(user.PasswordHash, loginForm.Password)
	if err != nil {
		return "", err
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"iss": strconv.Itoa(user.ID),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"data": map[string]string{
			"id":    strconv.Itoa(user.ID),
			"name":  user.Name,
			"email": user.Email,
		},
	},
	)

	token, err := claims.SignedString([]byte(os.Getenv("JWTSecretKey")))
	return token, err
}

func (authService *AuthService) JWTVerify(JWTTokenString string) (bool, error) {
	decodedToken, err := jwt.ParseWithClaims(JWTTokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWTSecretKey")), nil
	})

	return decodedToken.Valid, err
}

func (authService *AuthService) GetUserInformation(JWTTokenString string) (map[string]string, error, bool) {
	token, err := jwt.Parse(JWTTokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWTSecretKey")), nil
	})
	if err != nil {
		return make(map[string]string, 0), err, false
	}

	claims := token.Claims.(jwt.MapClaims)
	data := claims["data"].(map[string]interface{})

	return map[string]string{
		"id":    fmt.Sprintf("%v", data["id"]),
		"name":  fmt.Sprintf("%v", data["name"]),
		"email": fmt.Sprintf("%v", data["email"]),
	}, nil, token.Valid
}

func (authService *AuthService) ChangeUserPassword(changePasswordForm models.ChangePassword) error {
	loginForm := models.LoginForm{Email: changePasswordForm.Email, Password: changePasswordForm.OldPassword}
	user, err := authService.authStorage.GetUser(loginForm)
	if err != nil {
		return err
	}

	err = CompareHashPasswords(user.PasswordHash, loginForm.Password)
	if err != nil {
		return err
	}

	changePasswordForm.NewPassword, err = HashPassword(changePasswordForm.NewPassword)
	if err != nil {
		return err
	}

	err = authService.authStorage.ChangeUserPassword(changePasswordForm)
	return err
}
