package scenarios

import (
	"context"

	"github.com/Axel791/auth/internal/usecases/auth/dto"
)

type Registration interface {
	Execute(ctx context.Context, userDTO dto.UserDTO) error
}

type Login interface {
	Execute(ctx context.Context, userDTO dto.UserDTO) (dto.TokenDTO, error)
}

type Validate interface {
	Execute(ctx context.Context, token string) error
}
