package scenarios

import (
	"context"

	"github.com/Axel791/loyalty/internal/usecases/loyalty/dto"
)

type InputUserBalanceUseCase interface {
	Execute(ctx context.Context, orderID int64, balance dto.LoyaltyBalance) error
}

type GetUserBalanceUseCase interface {
	Execute(ctx context.Context, userID int64) (dto.LoyaltyBalance, error)
}

type ConclusionUserBalanceUseCase interface {
	Execute(ctx context.Context, balance dto.LoyaltyBalance) error
}

type CreateUserBalanceUseCase interface {
	Execute(ctx context.Context, userID int64) error
}
