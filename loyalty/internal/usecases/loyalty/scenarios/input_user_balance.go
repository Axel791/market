package scenarios

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Axel791/loyalty/internal/domains"
	"github.com/Axel791/loyalty/internal/usecases/loyalty/dto"
	"github.com/Axel791/loyalty/internal/usecases/loyalty/repositories"
)

type InputUserBalanceHandler struct {
	loyaltyRepository        repositories.LoyaltyRepository
	loyaltyHistoryRepository repositories.LoyaltyHistoryRepository
	unitOfWork               repositories.UnitOfWork
}

func NewInputUserBalance(
	loyaltyRepository repositories.LoyaltyRepository,
	loyaltyHistoryRepository repositories.LoyaltyHistoryRepository,
	unitOfWork repositories.UnitOfWork,
) *InputUserBalanceHandler {
	return &InputUserBalanceHandler{
		loyaltyRepository:        loyaltyRepository,
		loyaltyHistoryRepository: loyaltyHistoryRepository,
		unitOfWork:               unitOfWork,
	}
}

func (s *InputUserBalanceHandler) Execute(ctx context.Context, orderID int64, balance dto.LoyaltyBalance) error {
	var domainLoyaltyBalance domains.LoyaltyBalance
	domainLoyaltyBalance = domains.LoyaltyBalance{
		UserID: balance.UserID,
		Count:  balance.Count,
	}
	if err := domainLoyaltyBalance.ValidateCount(); err != nil {
		return err
	}

	if err := domainLoyaltyBalance.ValidateUserID(); err != nil {
		return err
	}

	loyaltyBalance := domainLoyaltyBalance.Count / 2
	domainLoyaltyBalance.Count = loyaltyBalance

	var loyaltyHistory domains.LoyaltyHistory
	loyaltyHistory = domains.LoyaltyHistory{
		UserID: balance.UserID,
		Count:  balance.Count,
		OrderID: sql.NullInt64{
			Int64: orderID,
			Valid: true,
		},
	}
	err := s.unitOfWork.Do(ctx, func(txContext context.Context) error {

		if err := s.loyaltyRepository.InputUserBalance(txContext, domainLoyaltyBalance); err != nil {
			return fmt.Errorf("error conclusion balance: %w", err)
		}

		if err := s.loyaltyHistoryRepository.CreateLoyaltyHistory(txContext, loyaltyHistory); err != nil {
			return fmt.Errorf("error create loyalty history: %w", err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("error conclusion balance: %w", err)
	}

	return nil
}
