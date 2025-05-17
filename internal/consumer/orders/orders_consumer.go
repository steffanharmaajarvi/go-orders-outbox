package orders

import (
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"orders.go/m/internal/infrastructure/event"
	"orders.go/m/internal/usecases/orders"
	"time"
)

type OrderCreatedEvent struct {
	ID          string     `json:"id"`
	Status      string     `json:"status"`
	UserID      string     `json:"user_id"`
	CreatedAt   *time.Time `json:"created_at"` // или `string`, если null приходит без ISO-формата
	UpdatedAt   *time.Time `json:"updated_at"`
	TotalAmount float64    `json:"total_amount,string"` // строка → float64
}

type OrderConsumer struct {
	confirmOrderUseCase *orders.ConfirmOrderUseCase
	consumer            event.EventConsumer
	workerCount         int
}

func NewOrderConsumer(
	consumer event.EventConsumer,
	confirmOrderUseCase *orders.ConfirmOrderUseCase,
) *OrderConsumer {
	return &OrderConsumer{
		consumer:            consumer,
		workerCount:         5,
		confirmOrderUseCase: confirmOrderUseCase,
	}
}

func (c *OrderConsumer) Start(ctx context.Context) {
	workerChan := make(chan event.ConsumerMessage, c.workerCount)

	go func() {
		consumerMessageChan, err := c.consumer.Start(ctx)
		if err != nil {
			panic(err)
		}

		for {
			select {
			case <-ctx.Done():
				return
			case consumerMessage := <-consumerMessageChan:
				workerChan <- consumerMessage
			}
		}
	}()

	for range c.workerCount {
		go c.worker(ctx, workerChan)
	}
}

func (c *OrderConsumer) worker(ctx context.Context, channel <-chan event.ConsumerMessage) {
	for {
		select {
		case consumerMessage := <-channel:
			c.handleRecord(ctx, consumerMessage)
		case <-ctx.Done():
			return
		}
	}
}

func (c *OrderConsumer) handleRecord(ctx context.Context, msg event.ConsumerMessage) {
	var event OrderCreatedEvent
	err := json.Unmarshal(msg.Value, &event)
	if err != nil {
		log.Err(err).Msg("failed to unmarshal event as order_created_event")
		return
	}

	err = c.confirmOrderUseCase.Execute(ctx, orders.ConfirmOrderInput{OrderID: event.ID})
	if err != nil {
		log.Err(err).Msg("failed to confirm order")
		return
	}
}
