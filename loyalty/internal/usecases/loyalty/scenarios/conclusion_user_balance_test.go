package scenarios

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Axel791/loyalty/internal/usecases/loyalty/dto"
	"github.com/Axel791/loyalty/internal/usecases/loyalty/repositories/mock"
)

func TestConclusionUserBalanceHandler_Execute_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoyaltyRepo := mock.NewMockLoyaltyRepository(ctrl)
	mockHistoryRepo := mock.NewMockLoyaltyHistoryRepository(ctrl)
	mockUow := mock.NewMockUnitOfWork(ctrl)

	useCase := NewConclusionUserBalance(
		mockLoyaltyRepo,
		mockHistoryRepo,
		mockUow,
	)

	ctx := context.Background()
	inputDTO := dto.LoyaltyBalance{
		UserID: 10,
		Count:  100,
	}

	mockUow.
		EXPECT().
		Do(ctx, gomock.Any()).
		DoAndReturn(func(ctx context.Context, fn func(txContext context.Context) error) error {
			return fn(ctx)
		})

	mockLoyaltyRepo.
		EXPECT().
		ConclusionUserBalance(gomock.Any(), gomock.Any()).
		Return(nil)

	mockHistoryRepo.
		EXPECT().
		CreateLoyaltyHistory(gomock.Any(), gomock.Any()).
		Return(nil)

	err := useCase.Execute(ctx, inputDTO)
	assert.NoError(t, err)
}

func TestConclusionUserBalanceHandler_Execute_ValidationErrorUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoyaltyRepo := mock.NewMockLoyaltyRepository(ctrl)
	mockHistoryRepo := mock.NewMockLoyaltyHistoryRepository(ctrl)
	mockUow := mock.NewMockUnitOfWork(ctrl)

	useCase := NewConclusionUserBalance(
		mockLoyaltyRepo,
		mockHistoryRepo,
		mockUow,
	)
	inputDTO := dto.LoyaltyBalance{
		UserID: 0,
		Count:  50,
	}

	err := useCase.Execute(context.Background(), inputDTO)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user id must be > 0")
}

func TestConclusionUserBalanceHandler_Execute_ValidationErrorCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoyaltyRepo := mock.NewMockLoyaltyRepository(ctrl)
	mockHistoryRepo := mock.NewMockLoyaltyHistoryRepository(ctrl)
	mockUow := mock.NewMockUnitOfWork(ctrl)

	useCase := NewConclusionUserBalance(
		mockLoyaltyRepo,
		mockHistoryRepo,
		mockUow,
	)
	inputDTO := dto.LoyaltyBalance{
		UserID: 10,
		Count:  -100,
	}

	err := useCase.Execute(context.Background(), inputDTO)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "balance count must be > 0")
}

func TestConclusionUserBalanceHandler_Execute_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoyaltyRepo := mock.NewMockLoyaltyRepository(ctrl)
	mockHistoryRepo := mock.NewMockLoyaltyHistoryRepository(ctrl)
	mockUow := mock.NewMockUnitOfWork(ctrl)

	useCase := NewConclusionUserBalance(
		mockLoyaltyRepo,
		mockHistoryRepo,
		mockUow,
	)

	ctx := context.Background()
	inputDTO := dto.LoyaltyBalance{
		UserID: 10,
		Count:  100,
	}

	mockUow.
		EXPECT().
		Do(ctx, gomock.Any()).
		DoAndReturn(func(ctx context.Context, fn func(txContext context.Context) error) error {
			return fn(ctx)
		})

	mockLoyaltyRepo.
		EXPECT().
		ConclusionUserBalance(gomock.Any(), gomock.Any()).
		Return(errors.New("db error on conclusion"))

	err := useCase.Execute(ctx, inputDTO)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error conclusion balance")
}

func TestConclusionUserBalanceHandler_Execute_HistoryRepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoyaltyRepo := mock.NewMockLoyaltyRepository(ctrl)
	mockHistoryRepo := mock.NewMockLoyaltyHistoryRepository(ctrl)
	mockUow := mock.NewMockUnitOfWork(ctrl)

	useCase := NewConclusionUserBalance(
		mockLoyaltyRepo,
		mockHistoryRepo,
		mockUow,
	)

	ctx := context.Background()
	inputDTO := dto.LoyaltyBalance{
		UserID: 10,
		Count:  100,
	}

	mockUow.
		EXPECT().
		Do(ctx, gomock.Any()).
		DoAndReturn(func(ctx context.Context, fn func(txContext context.Context) error) error {
			return fn(ctx)
		})

	mockLoyaltyRepo.
		EXPECT().
		ConclusionUserBalance(gomock.Any(), gomock.Any()).
		Return(nil)

	mockHistoryRepo.
		EXPECT().
		CreateLoyaltyHistory(gomock.Any(), gomock.Any()).
		Return(errors.New("create history error"))

	err := useCase.Execute(ctx, inputDTO)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error create loyalty history")
}

func TestConclusionUserBalanceHandler_Execute_UnitOfWorkError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoyaltyRepo := mock.NewMockLoyaltyRepository(ctrl)
	mockHistoryRepo := mock.NewMockLoyaltyHistoryRepository(ctrl)
	mockUow := mock.NewMockUnitOfWork(ctrl)

	useCase := NewConclusionUserBalance(
		mockLoyaltyRepo,
		mockHistoryRepo,
		mockUow,
	)

	ctx := context.Background()
	inputDTO := dto.LoyaltyBalance{
		UserID: 10,
		Count:  100,
	}
	mockUow.
		EXPECT().
		Do(ctx, gomock.Any()).
		Return(fmt.Errorf("failed to start transaction"))

	err := useCase.Execute(ctx, inputDTO)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error conclusion balance")
	assert.Contains(t, err.Error(), "failed to start transaction")
}
