package v1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Axel791/auth/internal/usecases/auth/dto"
	"github.com/Axel791/auth/internal/usecases/auth/scenarios/mock"

	"github.com/Axel791/appkit"
	log "github.com/sirupsen/logrus"
)

const registerUrl = "/public/api/v1/users/registration"

func TestRegistrationHandler_ServeHTTP_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRegistrationScenario := mock.NewMockRegistration(ctrl)

	handler := NewRegistrationHandler(
		mockRegistrationScenario,
		log.New(),
	)

	requestBody, _ := json.Marshal(map[string]string{
		"login":    "test_login",
		"password": "test_password",
	})

	req := httptest.NewRequest(http.MethodPost, registerUrl, bytes.NewReader(requestBody))
	rec := httptest.NewRecorder()

	mockRegistrationScenario.
		EXPECT().
		Execute(req.Context(), dto.UserDTO{
			Login:    "test_login",
			Password: "test_password",
		}).
		Return(nil)

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Empty(t, rec.Body.String(), "тело ответа ожидается пустым при 201")
}

func TestRegistrationHandler_ServeHTTP_BadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRegistrationScenario := mock.NewMockRegistration(ctrl)
	handler := NewRegistrationHandler(
		mockRegistrationScenario,
		log.New(),
	)

	req := httptest.NewRequest(
		http.MethodPost,
		"/public/api/v1/users/registration",
		bytes.NewReader([]byte("{invalid_json")),
	)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestRegistrationHandler_ServeHTTP_ScenarioError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRegistrationScenario := mock.NewMockRegistration(ctrl)
	handler := NewRegistrationHandler(
		mockRegistrationScenario,
		log.New(),
	)

	requestBody, _ := json.Marshal(map[string]string{
		"login":    "test_login",
		"password": "test_password",
	})

	req := httptest.NewRequest(http.MethodPost, registerUrl, bytes.NewReader(requestBody))
	rec := httptest.NewRecorder()

	err400 := appkit.BadRequestError("user login already exists")
	mockRegistrationScenario.
		EXPECT().
		Execute(req.Context(), dto.UserDTO{
			Login:    "test_login",
			Password: "test_password",
		}).
		Return(err400)

	handler.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "user login already exists")
}
