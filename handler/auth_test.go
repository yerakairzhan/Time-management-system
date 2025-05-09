package handler

import (
	"TimeManagementSystem/service"
	"TimeManagementSystem/service/mocks"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http/httptest"
	"testing"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *mocks.MockAuthorization)

	testTable := []struct {
		name                string
		inputBody           string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"email":"testt", "password":"test"}`,
			mockBehavior: func(s *mocks.MockAuthorization) {
				s.EXPECT().CreateUser(CreateUserMatcher{
					ExpectedEmail:     "testt",
					PlaintextPassword: "test",
				}).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:      "Empty Email",
			inputBody: `{"email":"", "password":"test"}`,
			mockBehavior: func(s *mocks.MockAuthorization) {
			},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Invalid JSON",
			inputBody: `{"email":"testt", "password":}`,
			mockBehavior: func(s *mocks.MockAuthorization) {
			},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"email":"testt", "password":"test"}`,
			mockBehavior: func(s *mocks.MockAuthorization) {
				s.EXPECT().CreateUser(CreateUserMatcher{
					ExpectedEmail:     "testt",
					PlaintextPassword: "test",
				}).Return(0, fmt.Errorf("failed to create user"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"failed to create user"}`,
		},
		{
			name:      "Missing Password",
			inputBody: `{"email":"testt", "password":""}`,
			mockBehavior: func(s *mocks.MockAuthorization) {
			},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			auth := mocks.NewMockAuthorization(ctrl)
			test.mockBehavior(auth)

			services := &service.Service{Authorization: auth}
			h := NewHandler(services.TaskService, services.Authorization, services.NotificationService)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(test.inputBody))
			req.Header.Set("Content-Type", "application/json")

			r := gin.New()
			r.POST("/sign-up", h.signUp)

			r.ServeHTTP(w, req)

			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.JSONEq(t, test.expectedRequestBody, w.Body.String())
		})
	}
}
