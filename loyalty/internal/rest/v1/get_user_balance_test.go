package v1

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Axel791/loyalty/internal/usecases/loyalty/dto"
	"github.com/Axel791/loyalty/internal/usecases/loyalty/scenarios/mock"
	log "github.com/sirupsen/logrus"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const getUserBalanceUrl = "/api/v1/loyalty/balance/conclusion"

func TestGetUserBalanceHandler_ServeHTTP_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mock.NewMockGetUserBalanceUseCase(ctrl)

	handler := NewGetUserBalanceHandler(
		log.New(),
		mockUseCase,
	)

	req := httptest.NewRequest(http.MethodGet, getUserBalanceUrl, nil)
	rec := httptest.NewRecorder()

	mockUseCase.
		EXPECT().
		Execute(req.Context(), int64(42)).
		Return(dto.LoyaltyBalance{
			ID:     100,
			UserID: 42,
			Count:  500,
		}, nil)

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	expectedJSON := `{"id":100,"user_id":42,"count":500}`
	assert.JSONEq(t, expectedJSON, rec.Body.String())
}

func TestGetUserBalanceHandler_ServeHTTP_NoIDParameter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mock.NewMockGetUserBalanceUseCase(ctrl)
	handler := NewGetUserBalanceHandler(log.New(), mockUseCase)

	req := httptest.NewRequest(http.MethodGet, getUserBalanceUrl, nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Параметр id не указан")
}

func TestGetUserBalanceHandler_ServeHTTP_InvalidIDFormat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mock.NewMockGetUserBalanceUseCase(ctrl)
	handler := NewGetUserBalanceHandler(log.New(), mockUseCase)

	req := httptest.NewRequest(http.MethodGet, getUserBalanceUrl, nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Неверный формат id")
}

func TestGetUserBalanceHandler_ServeHTTP_UseCaseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mock.NewMockGetUserBalanceUseCase(ctrl)
	handler := NewGetUserBalanceHandler(log.New(), mockUseCase)

	req := httptest.NewRequest(http.MethodGet, "/balance?id=50", nil)
	rec := httptest.NewRecorder()

	mockUseCase.
		EXPECT().
		Execute(req.Context(), int64(50)).
		Return(dto.LoyaltyBalance{}, fmt.Errorf("some error"))

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "some error")
}
