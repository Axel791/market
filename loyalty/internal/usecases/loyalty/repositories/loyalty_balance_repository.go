package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Axel791/loyalty/internal/domains"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

// SqlLoyaltyBalanceRepository - структура репозитория
type SqlLoyaltyBalanceRepository struct {
	db *sqlx.DB
}

// NewSqlLoyaltyRepository - конструктор репозитория системы лояльности
func NewSqlLoyaltyRepository(db *sqlx.DB) *SqlLoyaltyBalanceRepository {
	return &SqlLoyaltyBalanceRepository{
		db: db,
	}
}

// GetUserBalance - получение баланса пользователя
func (r *SqlLoyaltyBalanceRepository) GetUserBalance(
	ctx context.Context,
	userID int64,
) (domains.LoyaltyBalance, error) {
	builder := sq.
		Select("id", "user_id", "count").
		From("loyalty_balance").
		Where(sq.Eq{"user_id": userID}).PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return domains.LoyaltyBalance{}, fmt.Errorf("failed to build query: %w", err)
	}

	var lb domains.LoyaltyBalance
	err = r.db.QueryRowContext(ctx, query, args...).Scan(&lb.ID, &lb.UserID, &lb.Count)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domains.LoyaltyBalance{}, nil
		}
		return domains.LoyaltyBalance{}, fmt.Errorf("failed to execute query: %w", err)
	}

	return lb, nil
}

// CreateUserBalance - создание бонусного баланса пользователю
func (r *SqlLoyaltyBalanceRepository) CreateUserBalance(ctx context.Context, userID int64) error {
	insertBuilder := sq.
		Insert("loyalty_balance").
		Columns("user_id", "count").
		Values(userID, 0).PlaceholderFormat(sq.Dollar)

	insertQuery, insertArgs, err := insertBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build insert query: %w", err)
	}

	_, err = r.db.ExecContext(ctx, insertQuery, insertArgs...)
	if err != nil {
		return fmt.Errorf("failed to insert balance: %w", err)
	}

	return nil
}

// ConclusionUserBalance - «вывод» баланса (уменьшение Count)
func (r *SqlLoyaltyBalanceRepository) ConclusionUserBalance(
	ctx context.Context,
	userBalance domains.LoyaltyBalance,
) error {
	builder := sq.
		Update("loyalty_balance").
		Set("count", sq.Expr("count - ?", userBalance.Count)).
		Where(sq.Eq{"user_id": userBalance.UserID}).PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update balance: %w", err)
	}

	return nil
}

// InputUserBalance - «ввод» баланса (увеличение Count)
func (r *SqlLoyaltyBalanceRepository) InputUserBalance(ctx context.Context, userBalance domains.LoyaltyBalance) error {
	updateBuilder := sq.
		Update("loyalty_balance").
		Set("count", sq.Expr("count + ?", userBalance.Count)).
		Where(sq.Eq{"user_id": userBalance.UserID}).PlaceholderFormat(sq.Dollar)

	updateQuery, updateArgs, err := updateBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build update query: %w", err)
	}

	res, err := r.db.ExecContext(ctx, updateQuery, updateArgs...)
	if err != nil {
		return fmt.Errorf("failed to execute update: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get RowsAffected: %w", err)
	}

	if rowsAffected == 0 {
		insertBuilder := sq.
			Insert("loyalty_balance").
			Columns("user_id", "count").
			Values(userBalance.UserID, userBalance.Count).PlaceholderFormat(sq.Dollar)

		insertQuery, insertArgs, err := insertBuilder.ToSql()
		if err != nil {
			return fmt.Errorf("failed to build insert query: %w", err)
		}

		_, err = r.db.ExecContext(ctx, insertQuery, insertArgs...)
		if err != nil {
			return fmt.Errorf("failed to insert balance: %w", err)
		}
	}
	return nil
}
