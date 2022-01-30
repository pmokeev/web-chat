package controllers

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"pmokeev/web-chat/internal/models"
	services "pmokeev/web-chat/internal/services"
	mock_services "pmokeev/web-chat/internal/services/mocks"
	"testing"
)

func TestAuthController_SignUp(t *testing.T) {
	type mockBehavior func(service *mock_services.MockAuthorizationService, registerForm models.RegisterForm)

	testTable := []struct {
		name                 string
		inputBody            string
		inputUser            models.RegisterForm
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"test","email":"test@test.test","password":"test"}`,
			inputUser: models.RegisterForm{
				Name:         "test",
				Email:        "test@test.test",
				PasswordHash: "test",
			},
			mockBehavior: func(service *mock_services.MockAuthorizationService, registerForm models.RegisterForm) {
				service.EXPECT().SignUp(registerForm).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"message":"correct"}`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            `{"name":"test","email":"test@test.test"}`,
			inputUser:            models.RegisterForm{},
			mockBehavior:         func(service *mock_services.MockAuthorizationService, registerForm models.RegisterForm) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"invalid input body"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"name":"test","email":"test@test.test","password":"test"}`,
			inputUser: models.RegisterForm{
				Name:         "test",
				Email:        "test@test.test",
				PasswordHash: "test",
			},
			mockBehavior: func(service *mock_services.MockAuthorizationService, registerForm models.RegisterForm) {
				service.EXPECT().SignUp(registerForm).Return(errors.New("internal error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"error":"internal error"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_services.NewMockAuthorizationService(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			service := &services.Service{AuthorizationService: auth}
			controller := AuthController{authService: service}

			// Test Server
			router := gin.New()
			router.POST("/sign-up", controller.SignUp)

			// Test request
			w := httptest.NewRecorder()
			request := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(testCase.inputBody))

			// Perform request
			router.ServeHTTP(w, request)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}

func TestAuthController_SignIn(t *testing.T) {
	type mockBehavior func(service *mock_services.MockAuthorizationService, registerForm models.LoginForm)

	testTable := []struct {
		name                 string
		inputBody            string
		inputUser            models.LoginForm
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"email":"test@test.test","password":"test"}`,
			inputUser: models.LoginForm{
				Email:    "test@test.test",
				Password: "test",
			},
			mockBehavior: func(service *mock_services.MockAuthorizationService, registerForm models.LoginForm) {
				service.EXPECT().SignIn(registerForm).Return("token", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"message":"correct"}`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            `{"name":"test","email":"test@test.test"}`,
			inputUser:            models.LoginForm{},
			mockBehavior:         func(service *mock_services.MockAuthorizationService, registerForm models.LoginForm) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"invalid input body"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_services.NewMockAuthorizationService(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			service := &services.Service{AuthorizationService: auth}
			controller := AuthController{authService: service}

			// Test Server
			router := gin.New()
			router.POST("/sign-in", controller.SignIn)

			// Test request
			w := httptest.NewRecorder()
			request := httptest.NewRequest("POST", "/sign-in", bytes.NewBufferString(testCase.inputBody))

			// Perform request
			router.ServeHTTP(w, request)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
