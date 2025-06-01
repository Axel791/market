package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	userDomain "github.com/Axel791/auth/internal/domains"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

// SqlUserRepository - структура репозитория
type SqlUserRepository struct {
	db *sqlx.DB
}

// NewUserRepository - конструктор репозитория пользователя
func NewUserRepository(db *sqlx.DB) *SqlUserRepository {
	return &SqlUserRepository{db: db}
}

// CreateUser - создание пользователя
func (r *SqlUserRepository) CreateUser(ctx context.Context, user userDomain.User) (userDomain.User, error) {
	query, args, err := sq.Insert("users").
		Columns("login", "password").
		Values(user.Login, user.Password).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return userDomain.User{}, fmt.Errorf("failed to build SQL query: %w", err)
	}

	err = r.db.QueryRowxContext(ctx, query, args...).Scan(&user.ID)
	if err != nil {
		return userDomain.User{}, fmt.Errorf("failed to execute insert query: %w", err)
	}

	return user, nil
}

// GetUserById - получение пользователя по ID
func (r *SqlUserRepository) GetUserById(ctx context.Context, userID int64) (userDomain.User, error) {
	query, args, err := sq.Select("id", "login", "password").
		From("users").
		Where(sq.Eq{"id": userID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return userDomain.User{}, fmt.Errorf("failed to build SQL query: %w", err)
	}

	var user userDomain.User
	err = r.db.GetContext(ctx, &user, query, args...)
	if err != nil {
		return userDomain.User{}, fmt.Errorf("failed to execute select query: %w", err)
	}

	return user, nil
}

// GetUserByLogin - получение пользователя по логину
func (r *SqlUserRepository) GetUserByLogin(ctx context.Context, login string) (userDomain.User, error) {
	query, args, err := sq.Select("id", "login", "password").
		From("users").
		Where(sq.Eq{"login": login}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return userDomain.User{}, fmt.Errorf("failed to build SQL query: %w", err)
	}

	var user userDomain.User
	err = r.db.GetContext(ctx, &user, query, args...)

	if errors.Is(err, sql.ErrNoRows) {
		return userDomain.User{}, err
	}
	if err != nil {
		return userDomain.User{}, fmt.Errorf("failed to execute select query: %w", err)
	}

	return user, nil
}
