package repositories

import (
	"context"

	userDomain "github.com/Axel791/auth/internal/domains"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user userDomain.User) (userDomain.User, error)
	GetUserById(ctx context.Context, userID int64) (userDomain.User, error)
	GetUserByLogin(ctx context.Context, login string) (userDomain.User, error)
}
