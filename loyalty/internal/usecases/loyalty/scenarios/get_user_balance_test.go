package scenarios

import (
	"context"
	"errors"
	"testing"

	"github.com/Axel791/loyalty/internal/domains"
	"github.com/Axel791/loyalty/internal/usecases/loyalty/dto"
	"github.com/Axel791/loyalty/internal/usecases/loyalty/repositories/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetUserBalanceHandler_Execute_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockLoyaltyRepository(ctrl)
	useCase := NewGetUserBalance(mockRepo)

	ctx := context.Background()
	userID := int64(42)

	domainBalance := domains.LoyaltyBalance{
		ID:     100,
		UserID: userID,
		Count:  500,
	}

	mockRepo.
		EXPECT().
		GetUserBalance(ctx, userID).
		Return(domainBalance, nil)

	balanceDTO, err := useCase.Execute(ctx, userID)
	assert.NoError(t, err)
	assert.Equal(t, int64(100), balanceDTO.ID)
	assert.Equal(t, int64(42), balanceDTO.UserID)
	assert.Equal(t, int64(500), balanceDTO.Count)
}

func TestGetUserBalanceHandler_Execute_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockLoyaltyRepository(ctrl)
	useCase := NewGetUserBalance(mockRepo)

	ctx := context.Background()
	userID := int64(99)

	mockRepo.
		EXPECT().
		GetUserBalance(ctx, userID).
		Return(domains.LoyaltyBalance{}, errors.New("database failure"))

	balanceDTO, err := useCase.Execute(ctx, userID)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "could not get user balance")

	assert.Equal(t, dto.LoyaltyBalance{}, balanceDTO)
}
