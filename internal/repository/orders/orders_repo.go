package orders

import (
	"context"
	"database/sql"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"orders.go/m/internal/models"
)

type OrdersRepo struct {
	db *sql.DB
}

func NewOrdersRepo(db *sql.DB) *OrdersRepo {
	return &OrdersRepo{db: db}
}

func (repo *OrdersRepo) Save(ctx context.Context, exec boil.ContextExecutor, order *models.Order) error {
	err := order.Upsert(ctx, exec, true, []string{"id"}, boil.Whitelist("id", "status", "total_amount"), boil.Infer())
	if err != nil {
		log.Err(err).Msg("failed to upsert order")

		return err
	}

	return nil
}
