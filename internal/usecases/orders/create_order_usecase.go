package orders

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/ericlagergren/decimal"
	"github.com/friendsofgo/errors"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
	"orders.go/m/internal/entities/outbox"
	"orders.go/m/internal/models"
	"orders.go/m/internal/repository/orders"
	outbox2 "orders.go/m/internal/repository/outbox"
	"orders.go/m/internal/uow"
	"time"
)

type CreateOrderUseCase struct {
	orderRepo  *orders.OrdersRepo
	outboxRepo *outbox2.OutboxRepo
	uow        *uow.UnitOfWork
}

func NewCreateOrderUseCase(
	orderRepo *orders.OrdersRepo,
	uow *uow.UnitOfWork,
	outboxRepo *outbox2.OutboxRepo,
) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		orderRepo:  orderRepo,
		uow:        uow,
		outboxRepo: outboxRepo,
	}
}

type CreateOrderInput struct {
	UserID      string  `json:"user_id"`
	TotalAmount float64 `json:"total_amount"`
}

func (uc *CreateOrderUseCase) Execute(ctx context.Context, input CreateOrderInput) error {
	totalAmountDecimal := new(decimal.Big).SetFloat64(input.TotalAmount)

	order := &models.Order{
		ID:          uuid.New().String(),
		UserID:      input.UserID,
		TotalAmount: types.NewDecimal(totalAmountDecimal),
		Status:      models.OrderStatusPending,
	}

	var result types.JSON
	orderPayload, err := json.Marshal(order)
	if err != nil {
		log.Err(err).Msg("failed to marshal order")

		return errors.New("failed to marshal order")
	}

	err = result.UnmarshalJSON(orderPayload)
	if err != nil {
		log.Err(err).Msg("failed to unmarshal order")

		return errors.New("failed to unmarshal order")
	}

	return uc.uow.Do(ctx, func(tx *sql.Tx) error {
		if err := uc.orderRepo.Save(ctx, tx, order); err != nil {
			return err
		}

		event := &models.Outbox{
			ID:            uuid.New().String(),
			AggregateType: string(outbox.OrderOutboxAggregateType),
			AggregateID:   order.ID,
			Type:          string(outbox.OutboxEventTypeOrderCreated),
			Payload:       orderPayload,
			SentAt:        null.TimeFrom(time.Now()),
		}

		return uc.outboxRepo.Save(ctx, tx, event)
	})
}
