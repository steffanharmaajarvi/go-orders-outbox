package event

import (
	"context"
	"orders.go/m/internal/entities/outbox"
)

const (
	OrderTopic string = "order"
)

var AggregateTypeToTopic = map[string]string{
	string(outbox.OrderOutboxAggregateType): OrderTopic,
}

type EventPublisher interface {
	Publish(ctx context.Context, topic string, key string, payload []byte) error
	Close()
}
