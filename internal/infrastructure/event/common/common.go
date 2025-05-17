package common

import "orders.go/m/internal/entities/outbox"

const (
	OrderTopic string = "order"
)

var AggregateTypeToTopic = map[string]string{
	string(outbox.OrderOutboxAggregateType): OrderTopic,
}
