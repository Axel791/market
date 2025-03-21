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
	"github.com/Axel791/order/internal/rest/v1/api"
	"github.com/Axel791/order/internal/usecases/order/dto"
	"github.com/Axel791/order/internal/usecases/order/scenarios/mock"
	log "github.com/sirupsen/logrus"
)

const orderUrl = "/api/v1/order"

func TestCreateOrderHandler_ServeHTTP_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mock.NewMockCreateOrderUseCase(ctrl)
	handler := NewCreateOrderHandler(log.New(), mockUseCase)

	requestBody, _ := json.Marshal(api.InputCreateOrder{
		UserID:     10,
		Code:       "ABC123",
		TotalPrice: 1500,
	})
	req := httptest.NewRequest(http.MethodPost, orderUrl, bytes.NewReader(requestBody))
	rec := httptest.NewRecorder()

	mockUseCase.
		EXPECT().
		Execute(req.Context(), dto.CreateOrder{
			UserID:     10,
			Code:       "ABC123",
			TotalPrice: 1500,
		}).
		Return(dto.Order{
			ID:         101,
			UserID:     10,
			Code:       "ABC123",
			TotalPrice: 1500,
		}, nil)

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	expectedResp := `{"id":101,"user_id":10,"total_price":1500,"code":"ABC123"}`
	assert.JSONEq(t, expectedResp, rec.Body.String())
}

func TestCreateOrderHandler_ServeHTTP_BadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mock.NewMockCreateOrderUseCase(ctrl)
	handler := NewCreateOrderHandler(log.New(), mockUseCase)

	req := httptest.NewRequest(http.MethodPost, orderUrl, bytes.NewReader([]byte(`{invalid_json`)))
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "invalid request body")
}

func TestCreateOrderHandler_ServeHTTP_UseCaseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mock.NewMockCreateOrderUseCase(ctrl)
	handler := NewCreateOrderHandler(log.New(), mockUseCase)

	requestBody, _ := json.Marshal(api.InputCreateOrder{
		UserID:     20,
		Code:       "XYZ",
		TotalPrice: 999,
	})
	req := httptest.NewRequest(http.MethodPost, orderUrl, bytes.NewReader(requestBody))
	rec := httptest.NewRecorder()

	errBad := appkit.BadRequestError("some create error")
	mockUseCase.
		EXPECT().
		Execute(req.Context(), dto.CreateOrder{
			UserID:     20,
			Code:       "XYZ",
			TotalPrice: 999,
		}).
		Return(dto.Order{}, errBad)

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "some create error")
}
