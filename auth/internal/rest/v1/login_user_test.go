package v1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Axel791/appkit"
	"github.com/Axel791/auth/internal/usecases/auth/dto"
	"github.com/Axel791/auth/internal/usecases/auth/scenarios/mock"
	log "github.com/sirupsen/logrus"
)

const loginUrl = "/public/api/v1/users/login"

func TestLoginHandler_ServeHTTP_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoginScenario := mock.NewMockLogin(ctrl)

	handler := NewLoginHandler(
		log.New(),
		mockLoginScenario,
	)
	body, _ := json.Marshal(map[string]string{
		"login":    "test_login",
		"password": "test_password",
	})
	req := httptest.NewRequest(http.MethodPost, loginUrl, bytes.NewReader(body))
	rec := httptest.NewRecorder()

	mockLoginScenario.
		EXPECT().
		Execute(req.Context(), dto.UserDTO{
			Login:    "test_login",
			Password: "test_password",
		}).
		Return(dto.TokenDTO{Token: "some_jwt_token"}, nil)

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), `"token":"some_jwt_token"`)
}

func TestLoginHandler_ServeHTTP_BadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoginScenario := mock.NewMockLogin(ctrl)
	handler := NewLoginHandler(
		log.New(),
		mockLoginScenario,
	)

	req := httptest.NewRequest(http.MethodPost, loginUrl, bytes.NewReader([]byte("{invalid_json")))
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "invalid request body")
}

func TestLoginHandler_ServeHTTP_ScenarioError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoginScenario := mock.NewMockLogin(ctrl)
	handler := NewLoginHandler(
		log.New(),
		mockLoginScenario,
	)

	body, _ := json.Marshal(map[string]string{
		"login":    "bad_login",
		"password": "wrong_password",
	})
	req := httptest.NewRequest(http.MethodPost, loginUrl, bytes.NewReader(body))
	rec := httptest.NewRecorder()

	errBad := appkit.BadRequestError("invalid password")
	mockLoginScenario.
		EXPECT().
		Execute(req.Context(), dto.UserDTO{
			Login:    "bad_login",
			Password: "wrong_password",
		}).
		Return(dto.TokenDTO{}, errBad)

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "invalid password")
}
