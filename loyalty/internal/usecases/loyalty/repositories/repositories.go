package repositories

import (
	"context"

	"github.com/Axel791/loyalty/internal/domains"
)

type LoyaltyRepository interface {
	GetUserBalance(ctx context.Context, userID int64) (domains.LoyaltyBalance, error)
	ConclusionUserBalance(ctx context.Context, userBalance domains.LoyaltyBalance) error
	InputUserBalance(ctx context.Context, userBalance domains.LoyaltyBalance) error
	CreateUserBalance(ctx context.Context, userID int64) error
}

type LoyaltyHistoryRepository interface {
	GetUserHistory(ctx context.Context, userID int64) ([]domains.LoyaltyHistory, error)
	CreateLoyaltyHistory(ctx context.Context, history domains.LoyaltyHistory) error
}

type UnitOfWork interface {
	Do(ctx context.Context, fn func(ctx context.Context) error) error
}
