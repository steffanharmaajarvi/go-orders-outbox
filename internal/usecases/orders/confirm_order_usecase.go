package orders

import (
	"context"
	"database/sql"
	"github.com/friendsofgo/errors"
	"github.com/rs/zerolog/log"
	"orders.go/m/internal/models"
	"orders.go/m/internal/repository/orders"
	"orders.go/m/internal/uow"
)

type ConfirmOrderUseCase struct {
	orderRepo *orders.OrdersRepo
	uow       *uow.UnitOfWork
}

func NewConfirmOrderUseCase(
	orderRepo *orders.OrdersRepo,
	uow *uow.UnitOfWork,
) *ConfirmOrderUseCase {
	return &ConfirmOrderUseCase{
		orderRepo: orderRepo,
		uow:       uow,
	}
}

type ConfirmOrderInput struct {
	OrderID string `json:"order_id" binding:"required"`
}

func (uc *ConfirmOrderUseCase) Execute(ctx context.Context, input ConfirmOrderInput) error {
	err := uc.uow.Do(ctx, func(tx *sql.Tx) (err error) {
		order, err := uc.orderRepo.GetOrderByID(ctx, tx, input.OrderID)

		if err != nil {
			return err
		}

		if order.Status != models.OrderStatusPending {
			return errors.New("order status is not pending")
		}

		order.Status = models.OrderStatusConfirmed
		if err := uc.orderRepo.Save(ctx, tx, order); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Err(err).Msg("failed to save order")

		return errors.New("failed to save order")
	}

	return nil
}
