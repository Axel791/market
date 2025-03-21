package scenarios

import (
	"context"
	"errors"
	"testing"

	"github.com/Axel791/loyalty/internal/usecases/loyalty/repositories/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateLoyaltyBalanceHandler_Execute_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockLoyaltyRepository(ctrl)
	useCase := NewCreateLoyaltyBalance(mockRepo)

	ctx := context.Background()
	userID := int64(123)

	mockRepo.
		EXPECT().
		CreateUserBalance(ctx, userID).
		Return(nil)

	err := useCase.Execute(ctx, userID)
	assert.NoError(t, err)
}

func TestCreateLoyaltyBalanceHandler_Execute_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockLoyaltyRepository(ctrl)
	useCase := NewCreateLoyaltyBalance(mockRepo)

	ctx := context.Background()
	userID := int64(123)

	mockRepo.
		EXPECT().
		CreateUserBalance(ctx, userID).
		Return(errors.New("db error"))

	err := useCase.Execute(ctx, userID)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error create user balance failed")
	assert.Contains(t, err.Error(), "db error")
}
