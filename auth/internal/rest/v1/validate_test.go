package v1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Axel791/appkit"
	"github.com/Axel791/auth/internal/rest/v1/api"
	"github.com/Axel791/auth/internal/usecases/auth/scenarios/mock"
	log "github.com/sirupsen/logrus"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const validateUrl = "/public/api/v1/users/validate"

func TestValidationHandler_ServeHTTP_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockValidateScenario := mock.NewMockValidate(ctrl)
	handler := NewValidationHandler(log.New(), mockValidateScenario)

	body, _ := json.Marshal(api.Token{Token: "valid_token"})
	req := httptest.NewRequest(http.MethodPost, validateUrl, bytes.NewReader(body))
	rec := httptest.NewRecorder()

	mockValidateScenario.
		EXPECT().
		Execute(req.Context(), "valid_token").
		Return(nil)

	handler.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Empty(t, rec.Body.String())
}

func TestValidationHandler_ServeHTTP_BadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockValidateScenario := mock.NewMockValidate(ctrl)
	handler := NewValidationHandler(log.New(), mockValidateScenario)

	req := httptest.NewRequest(http.MethodPost, validateUrl, bytes.NewReader([]byte("{invalid_json")))
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "invalid request body")
}

func TestValidationHandler_ServeHTTP_ScenarioError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockValidateScenario := mock.NewMockValidate(ctrl)
	handler := NewValidationHandler(log.New(), mockValidateScenario)

	body, _ := json.Marshal(api.Token{Token: "invalid_token"})
	req := httptest.NewRequest(http.MethodPost, validateUrl, bytes.NewReader(body))
	rec := httptest.NewRecorder()

	errUnauthorized := appkit.UnauthorizedError("invalid token")
	mockValidateScenario.
		EXPECT().
		Execute(req.Context(), "invalid_token").
		Return(errUnauthorized)

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Contains(t, rec.Body.String(), "invalid token")
}
