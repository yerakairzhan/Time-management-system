package handler

import (
	db "TimeManagementSystem/db/sqlc"
	"TimeManagementSystem/service"
	"TimeManagementSystem/service/mocks"
	"bytes"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler_createTask(t *testing.T) {
	type mockBehaviour func(s *mocks.MockTaskService, task db.Task)
	type authBehaviour func(a *mocks.MockAuthorization)

	tests := []struct {
		name               string
		inputBody          string
		mockBehaviour      mockBehaviour
		authBehaviour      authBehaviour
		expectedStatusCode int
		expectedBody       string
	}{
		{
			name: "OK",
			inputBody: `{
				"Name":"Write a test",
				"Description":"in GO",
				"Category":"Code",
				"Priority":"High",
				"Deadline":"2025-04-16T14:30:00Z"
			}`,
			authBehaviour: func(a *mocks.MockAuthorization) {
				a.EXPECT().ParseToken("mockToken").Return(0, nil)
			},
			mockBehaviour: func(s *mocks.MockTaskService, task db.Task) {
				s.EXPECT().CreateTask(0, task).Return(0, nil)
			},
			expectedStatusCode: 200,
			expectedBody:       `{"id": 0}`,
		},
		{
			name: "Invalid JSON",
			inputBody: `{
				"Name":"Write a test",,
				"Description":"in GO",
				"Category":"Code",
				"Priority":"high",
				"Deadline":"2025-04-16T14:30:00Z"
			}`,
			authBehaviour: nil,
			mockBehaviour: func(s *mocks.MockTaskService, task db.Task) {
				s.EXPECT().CreateTask(gomock.Any(), gomock.Any()).Times(0)
			},
			expectedStatusCode: 400,
			expectedBody:       `{"error":"invalid request body"}`,
		},
		{
			name: "Invalid Deadline Format",
			inputBody: `{
				"Name":"Test",
				"Description":"test",
				"Category":"General",
				"Priority":"Low",
				"Deadline":"April 16th"
			}`,
			authBehaviour: func(a *mocks.MockAuthorization) {
				a.EXPECT().ParseToken("mockToken").Return(1, nil)
			},
			mockBehaviour:      func(s *mocks.MockTaskService, task db.Task) {},
			expectedStatusCode: 400,
			expectedBody:       `{"message":"invalid deadline format"}`,
		},
		{
			name: "Unauthorized",
			inputBody: `{
				"Name":"Test",
				"Description":"test",
				"Category":"General",
				"Priority":"Medium",
				"Deadline":"2025-04-16 14:30:00"
			}`,
			authBehaviour: func(a *mocks.MockAuthorization) {
				a.EXPECT().ParseToken("mockToken").Return(0, errors.New("unauthorized"))
			},
			mockBehaviour:      func(s *mocks.MockTaskService, task db.Task) {},
			expectedStatusCode: 401,
			expectedBody:       `{"message":"unauthorized"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			taskService := mocks.NewMockTaskService(ctrl)
			auth := mocks.NewMockAuthorization(ctrl)

			// Mock the expected token parsing in every test
			if test.authBehaviour != nil {
				test.authBehaviour(auth)
			}

			// Mock the task service behaviour for each test
			if test.mockBehaviour != nil {
				parsedDeadline, _ := time.Parse(time.RFC3339, "2025-04-16T14:30:00Z")
				expectedTask := db.Task{
					Name:        "Write a test",
					Description: sql.NullString{String: "in GO", Valid: true},
					Category:    "Code",
					Priority:    "High",
					Deadline:    sql.NullTime{Time: parsedDeadline.UTC(), Valid: true},
				}
				test.mockBehaviour(taskService, expectedTask)
			}

			// Initialize the services and handler
			services := &service.Service{
				TaskService:         taskService,
				Authorization:       auth,
				NotificationService: nil,
			}

			h := NewHandler(services.TaskService, services.Authorization, services.NotificationService)

			// Simulate HTTP request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/tasks", bytes.NewBufferString(test.inputBody))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer mockToken")

			r := gin.New()
			r.POST("/tasks", h.createTask)

			// Execute the test
			r.ServeHTTP(w, req)

			// Assert the results
			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.JSONEq(t, test.expectedBody, w.Body.String())
		})
	}
}
