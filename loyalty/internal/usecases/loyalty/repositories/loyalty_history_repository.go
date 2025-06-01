package repositories

import (
	"context"
	"fmt"

	"github.com/Axel791/loyalty/internal/domains"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

// SqlLoyaltyHistoryRepository - структура репозитория истории лояльности
type SqlLoyaltyHistoryRepository struct {
	db *sqlx.DB
}

// NewSqlLoyaltyHistoryRepository - конструктор репозитория истории лояльности
func NewSqlLoyaltyHistoryRepository(db *sqlx.DB) *SqlLoyaltyHistoryRepository {
	return &SqlLoyaltyHistoryRepository{
		db: db,
	}
}

// GetUserHistory - получение истории пополнения системы лояльности пользователя
func (r *SqlLoyaltyHistoryRepository) GetUserHistory(
	ctx context.Context,
	userID int64,
) ([]domains.LoyaltyHistory, error) {
	builder := sq.Select(
		"id",
		"user_id",
		"order_id",
		"count",
	).
		From("loyalty_history").
		Where(sq.Eq{"user_id": userID}).PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error build query: %w", err)
	}
	rows, err := r.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query all metrics: %w", err)
	}
	defer rows.Close()

	var loyaltyHistories []domains.LoyaltyHistory

	for rows.Next() {
		var loyaltyHistory domains.LoyaltyHistory
		if err := rows.StructScan(&loyaltyHistory); err != nil {
			return nil, fmt.Errorf("rows scan failed: %w", err)
		}
		loyaltyHistories = append(loyaltyHistories, loyaltyHistory)
	}
	return loyaltyHistories, nil
}

// CreateLoyaltyHistory - создание истории для пользователя
func (r *SqlLoyaltyHistoryRepository) CreateLoyaltyHistory(
	ctx context.Context,
	history domains.LoyaltyHistory,
) error {
	insertBuilder := sq.
		Insert("loyalty_history").
		Columns("user_id", "order_id", "count").
		Values(history.UserID, history.OrderID, history.Count).PlaceholderFormat(sq.Dollar)
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
