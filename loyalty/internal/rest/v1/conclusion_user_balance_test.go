package v1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Axel791/appkit"
	"github.com/Axel791/loyalty/internal/rest/v1/api"
	"github.com/Axel791/loyalty/internal/usecases/loyalty/dto"
	"github.com/Axel791/loyalty/internal/usecases/loyalty/scenarios/mock"
	log "github.com/sirupsen/logrus"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const conclusionUrl = "/api/v1/loyalty/balance"

func TestConclusionUserBalanceHandler_ServeHTTP_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mock.NewMockConclusionUserBalanceUseCase(ctrl)
	handler := NewConclusionUserBalanceHandler(log.New(), mockUseCase)

	bodyData, _ := json.Marshal(api.LoyaltyBalance{
		UserID: 10,
		Count:  100,
	})
	req := httptest.NewRequest(http.MethodPost, conclusionUrl, bytes.NewReader(bodyData))
	rec := httptest.NewRecorder()

	mockUseCase.
		EXPECT().
		Execute(req.Context(), dto.LoyaltyBalance{UserID: 10, Count: 100}).
		Return(nil)

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Empty(t, rec.Body.String())
}

func TestConclusionUserBalanceHandler_ServeHTTP_BadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mock.NewMockConclusionUserBalanceUseCase(ctrl)
	handler := NewConclusionUserBalanceHandler(log.New(), mockUseCase)

	req := httptest.NewRequest(http.MethodPost, conclusionUrl, bytes.NewReader([]byte("{bad_json")))
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "invalid request body")
}

func TestConclusionUserBalanceHandler_ServeHTTP_UseCaseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mock.NewMockConclusionUserBalanceUseCase(ctrl)
	handler := NewConclusionUserBalanceHandler(log.New(), mockUseCase)

	bodyData, _ := json.Marshal(api.LoyaltyBalance{
		UserID: 5,
		Count:  50,
	})
	req := httptest.NewRequest(http.MethodPost, conclusionUrl, bytes.NewReader(bodyData))
	rec := httptest.NewRecorder()

	errUseCase := appkit.BadRequestError("count must be > 0") // допустим, вернулось 400
	mockUseCase.
		EXPECT().
		Execute(req.Context(), dto.LoyaltyBalance{UserID: 5, Count: 50}).
		Return(errUseCase)

	handler.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "count must be > 0")
}
