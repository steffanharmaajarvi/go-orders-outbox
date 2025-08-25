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
	"orders.go/m/internal/infrastructure/event"
	"orders.go/m/internal/models"
	"orders.go/m/internal/repository/orders"
	outbox2 "orders.go/m/internal/repository/outbox"
	"orders.go/m/internal/uow"
	"time"
)

type CreateOrderUseCase struct {
	orderRepo      *orders.OrdersRepo
	outboxRepo     *outbox2.OutboxRepo
	uow            *uow.UnitOfWork
	eventPublisher event.EventPublisher
}

func NewCreateOrderUseCase(
	orderRepo *orders.OrdersRepo,
	uow *uow.UnitOfWork,
	outboxRepo *outbox2.OutboxRepo,
	eventPublisher event.EventPublisher,
) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		orderRepo:      orderRepo,
		uow:            uow,
		outboxRepo:     outboxRepo,
		eventPublisher: eventPublisher,
	}
}

type CreateOrderInput struct {
	UserID      string  `json:"user_id" binding:"required"`
	TotalAmount float64 `json:"total_amount" binding:"required"`
}

func (uc *CreateOrderUseCase) Execute(ctx context.Context, input CreateOrderInput) error {
	totalAmountDecimal := new(decimal.Big).SetFloat64(input.TotalAmount)

	order := &models.Order{
		ID:          uuid.New().String(),
		UserID:      input.UserID,
		TotalAmount: types.NewDecimal(totalAmountDecimal),
		Status:      models.OrderStatusPending,
	}

	orderPayload, err := json.Marshal(order)
	if err != nil {
		log.Err(err).Msg("failed to marshal order")

		return errors.New("failed to marshal order")
	}

	err = uc.uow.Do(ctx, func(tx *sql.Tx) (err error) {
		if err := uc.orderRepo.Save(ctx, tx, order); err != nil {
			return err
		}

		event := &models.Outbox{
			ID:            uuid.New().String(),
			AggregateType: string(outbox.OrderOutboxAggregateType),
			AggregateID:   order.ID,
			Type:          string(outbox.OutboxEventTypeOrderCreated),
			Payload:       orderPayload,
			SentAt:        null.TimeFromPtr(nil),
			OccurredAt:    null.TimeFrom(time.Now()),
		}

		return uc.outboxRepo.Save(ctx, tx, event)
	})

	if err != nil {
		log.Err(err).Msg("failed to save order")

		return errors.New("failed to save order")
	}

	return nil
}
