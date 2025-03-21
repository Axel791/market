package scenarios

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Axel791/loyalty/internal/domains"
	"github.com/Axel791/loyalty/internal/usecases/loyalty/dto"
	"github.com/Axel791/loyalty/internal/usecases/loyalty/repositories/mock"
)

func TestInputUserBalanceHandler_Execute_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoyaltyRepo := mock.NewMockLoyaltyRepository(ctrl)
	mockHistoryRepo := mock.NewMockLoyaltyHistoryRepository(ctrl)
	mockUoW := mock.NewMockUnitOfWork(ctrl)

	useCase := NewInputUserBalance(mockLoyaltyRepo, mockHistoryRepo, mockUoW)
	ctx := context.Background()

	orderID := int64(999)
	inputDTO := dto.LoyaltyBalance{
		UserID: 10,
		Count:  100,
	}

	mockUoW.
		EXPECT().
		Do(ctx, gomock.Any()).
		DoAndReturn(func(txCtx context.Context, fn func(context.Context) error) error {
			return fn(txCtx)
		})

	mockLoyaltyRepo.
		EXPECT().
		InputUserBalance(gomock.Any(), domains.LoyaltyBalance{UserID: 10, Count: 50}).
		Return(nil)

	mockHistoryRepo.
		EXPECT().
		CreateLoyaltyHistory(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, hist domains.LoyaltyHistory) error {
			assert.Equal(t, int64(10), hist.UserID)
			assert.Equal(t, int64(100), hist.Count)
			assert.Equal(t, orderID, hist.OrderID.Int64)
			assert.True(t, hist.OrderID.Valid)
			return nil
		})

	err := useCase.Execute(ctx, orderID, inputDTO)
	assert.NoError(t, err)
}

func TestInputUserBalanceHandler_Execute_ValidateCountError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoyaltyRepo := mock.NewMockLoyaltyRepository(ctrl)
	mockHistoryRepo := mock.NewMockLoyaltyHistoryRepository(ctrl)
	mockUoW := mock.NewMockUnitOfWork(ctrl)

	useCase := NewInputUserBalance(mockLoyaltyRepo, mockHistoryRepo, mockUoW)
	ctx := context.Background()

	inputDTO := dto.LoyaltyBalance{
		UserID: 10,
		Count:  0,
	}

	err := useCase.Execute(ctx, 999, inputDTO)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "balance count must be > 0")
}

func TestInputUserBalanceHandler_Execute_ValidateUserIDError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoyaltyRepo := mock.NewMockLoyaltyRepository(ctrl)
	mockHistoryRepo := mock.NewMockLoyaltyHistoryRepository(ctrl)
	mockUoW := mock.NewMockUnitOfWork(ctrl)

	useCase := NewInputUserBalance(mockLoyaltyRepo, mockHistoryRepo, mockUoW)
	ctx := context.Background()

	inputDTO := dto.LoyaltyBalance{
		UserID: 0,
		Count:  100,
	}

	err := useCase.Execute(ctx, 999, inputDTO)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user id must be > 0")
}

func TestInputUserBalanceHandler_Execute_UnitOfWorkError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoyaltyRepo := mock.NewMockLoyaltyRepository(ctrl)
	mockHistoryRepo := mock.NewMockLoyaltyHistoryRepository(ctrl)
	mockUoW := mock.NewMockUnitOfWork(ctrl)

	useCase := NewInputUserBalance(mockLoyaltyRepo, mockHistoryRepo, mockUoW)
	ctx := context.Background()

	inputDTO := dto.LoyaltyBalance{
		UserID: 10,
		Count:  100,
	}

	mockUoW.
		EXPECT().
		Do(ctx, gomock.Any()).
		Return(fmt.Errorf("transaction error"))

	err := useCase.Execute(ctx, 999, inputDTO)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error conclusion balance")
	assert.Contains(t, err.Error(), "transaction error")
}

func TestInputUserBalanceHandler_Execute_InputUserBalanceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoyaltyRepo := mock.NewMockLoyaltyRepository(ctrl)
	mockHistoryRepo := mock.NewMockLoyaltyHistoryRepository(ctrl)
	mockUoW := mock.NewMockUnitOfWork(ctrl)

	useCase := NewInputUserBalance(mockLoyaltyRepo, mockHistoryRepo, mockUoW)
	ctx := context.Background()

	inputDTO := dto.LoyaltyBalance{
		UserID: 10,
		Count:  100,
	}

	mockUoW.
		EXPECT().
		Do(ctx, gomock.Any()).
		DoAndReturn(func(txCtx context.Context, fn func(context.Context) error) error {
			return fn(txCtx)
		})

	mockLoyaltyRepo.
		EXPECT().
		InputUserBalance(gomock.Any(), domains.LoyaltyBalance{UserID: 10, Count: 50}).
		Return(errors.New("db input error"))

	mockHistoryRepo.
		EXPECT().
		CreateLoyaltyHistory(gomock.Any(), gomock.Any()).
		Times(0)

	err := useCase.Execute(ctx, 999, inputDTO)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error conclusion balance")
	assert.Contains(t, err.Error(), "db input error")
}

func TestInputUserBalanceHandler_Execute_CreateLoyaltyHistoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoyaltyRepo := mock.NewMockLoyaltyRepository(ctrl)
	mockHistoryRepo := mock.NewMockLoyaltyHistoryRepository(ctrl)
	mockUoW := mock.NewMockUnitOfWork(ctrl)

	useCase := NewInputUserBalance(mockLoyaltyRepo, mockHistoryRepo, mockUoW)
	ctx := context.Background()

	inputDTO := dto.LoyaltyBalance{
		UserID: 10,
		Count:  100,
	}

	mockUoW.
		EXPECT().
		Do(ctx, gomock.Any()).
		DoAndReturn(func(txCtx context.Context, fn func(context.Context) error) error {
			return fn(txCtx)
		})

	mockLoyaltyRepo.
		EXPECT().
		InputUserBalance(gomock.Any(), domains.LoyaltyBalance{UserID: 10, Count: 50}).
		Return(nil)

	mockHistoryRepo.
		EXPECT().
		CreateLoyaltyHistory(gomock.Any(), gomock.Any()).
		Return(errors.New("db history error"))

	err := useCase.Execute(ctx, 999, inputDTO)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error create loyalty history")
	assert.Contains(t, err.Error(), "db history error")
}
