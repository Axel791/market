package scenarios

import (
	"context"
	"fmt"

	"github.com/Axel791/loyalty/internal/usecases/loyalty/repositories"
)

type CreateLoyaltyBalanceHandler struct {
	loyaltyRepository repositories.LoyaltyRepository
}

func NewCreateLoyaltyBalance(
	loyaltyRepository repositories.LoyaltyRepository,
) *CreateLoyaltyBalanceHandler {
	return &CreateLoyaltyBalanceHandler{
		loyaltyRepository: loyaltyRepository,
	}
}

func (s *CreateLoyaltyBalanceHandler) Execute(ctx context.Context, userID int64) error {
	err := s.loyaltyRepository.CreateUserBalance(ctx, userID)
	if err != nil {
		return fmt.Errorf("error create user balance failed: %w", err)
	}
	return nil
}
