package services

import (
	"github.com/gorilla/websocket"
	"pmokeev/web-chat/internal/models"
	"pmokeev/web-chat/internal/storage"
)

type AuthorizationService interface {
	SignUp(registerForm models.RegisterForm) error
	SignIn(loginForm models.LoginForm) (string, error)
	JWTVerify(JWTTokenString string) (bool, error)
	GetUserInformation(JWTTokenString string) (map[string]string, error, bool)
	ChangeUserPassword(changePasswordForm models.ChangePassword) error
}

type ChattingService interface {
	RunChat()
	AddUser(username string, connection *websocket.Conn)
	connectUser(user *models.User)
	broadcastMessage(message *models.Message)
	disconnectUser(user *models.User)
}

type Service struct {
	AuthorizationService
	ChattingService
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		AuthorizationService: NewAuthService(storage.AuthorizationStorage),
		ChattingService:      NewChatSerivce()}
}
