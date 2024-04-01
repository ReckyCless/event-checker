package handler

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	app "github.com/reckycless/event-checker"
	"github.com/reckycless/event-checker/pkg/service"
	service_mocks "github.com/reckycless/event-checker/pkg/service/mocks"
)

func TestHandler_signUp(t *testing.T) {
	// Init Test Table
	type mockBehaivor func(s *service_mocks.MockAuthorization, user app.User)

	var image_path_test = "image_path"
	var phone_test = "+71234567890"

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            app.User
		mockBehavior         mockBehaivor
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "OK",
			inputBody: `{
				"name": "Test Name", 
				"surname": "Test Surname", 
				"patronymic": "Test Patronymic", 
				"password": "qwerty",
				"birth_date": "2012-03-20T00:00:00Z",
				"email": "testemail@mail.ru",
				"phone": "+71234567890",
				"image_path": "image_path",
				"role_id": 1,
				"password": "qwerty"
			}`,
			inputUser: app.User{
				Name:       "Test Name",
				Surname:    "Test Surname",
				Patronymic: "Test Patronymic",
				BirthDate:  time.Date(2012, 03, 20, 0, 0, 0, 0, time.UTC),
				Email:      "testemail@mail.ru",
				Phone:      &phone_test,
				ImagePath:  &image_path_test,
				Password:   "qwerty",
				RoleID:     1,
			},
			mockBehavior: func(r *service_mocks.MockAuthorization, user app.User) {
				r.EXPECT().CreateUser(user).Return(1, nil).AnyTimes()
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            `{"name": "name"}`,
			inputUser:            app.User{},
			mockBehavior:         func(r *service_mocks.MockAuthorization, user app.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"Key: 'User.Password' Error:Field validation for 'Password' failed on the 'required' tag\nKey: 'User.Surname' Error:Field validation for 'Surname' failed on the 'required' tag\nKey: 'User.RoleID' Error:Field validation for 'RoleID' failed on the 'required' tag"}`,
		},
		{
			name: "Service Error",
			inputBody: `{
				"name": "Test Name", 
				"surname": "Test Surname", 
				"patronymic": "Test Patronymic", 
				"password": "qwerty",
				"birth_date": "2012-03-20T00:00:00Z",
				"email": "testemail@mail.ru",
				"image_path": "image_path",
				"phone": "+71234567890",
				"role_id": 1,
				"password": "qwerty"
				}`,
			inputUser: app.User{
				Name:       "Test Name",
				Surname:    "Test Surname",
				Patronymic: "Test Patronymic",
				BirthDate:  time.Date(2012, 03, 20, 0, 0, 0, 0, time.UTC),
				Email:      "testemail@mail.ru",
				Phone:      &phone_test,
				ImagePath:  &image_path_test,
				RoleID:     1,
				Password:   "qwerty",
			},
			mockBehavior: func(r *service_mocks.MockAuthorization, user app.User) {
				r.EXPECT().CreateUser(user).Return(0, errors.New("something went wrong")).AnyTimes()
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}{"id":0}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := service_mocks.NewMockAuthorization(c)
			test.mockBehavior(repo, test.inputUser)

			services := &service.Service{Authorization: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
