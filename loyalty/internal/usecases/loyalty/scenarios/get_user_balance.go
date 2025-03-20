package scenarios

import (
	"context"
	"fmt"

	"github.com/Axel791/loyalty/internal/usecases/loyalty/dto"
	"github.com/Axel791/loyalty/internal/usecases/loyalty/repositories"
)

type GetUserBalanceHandler struct {
	loyaltyRepository repositories.LoyaltyRepository
}

func NewGetUserBalance(
	loyaltyRepository repositories.LoyaltyRepository,
) *GetUserBalanceHandler {
	return &GetUserBalanceHandler{
		loyaltyRepository: loyaltyRepository,
	}
}

func (s *GetUserBalanceHandler) Execute(ctx context.Context, userID int64) (dto.LoyaltyBalance, error) {
	loyaltyDomain, err := s.loyaltyRepository.GetUserBalance(ctx, userID)
	if err != nil {
		return dto.LoyaltyBalance{}, fmt.Errorf("could not get user balance: %w", err)
	}
	var balance dto.LoyaltyBalance
	balance = dto.LoyaltyBalance{
		ID:     loyaltyDomain.ID,
		UserID: loyaltyDomain.UserID,
		Count:  loyaltyDomain.Count,
	}
	return balance, nil
}
