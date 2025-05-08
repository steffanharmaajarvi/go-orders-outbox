package outbox

type OutboxAggregateType string

const (
	OrderOutboxAggregateType OutboxAggregateType = "order"
)

type OutboxEventType string

const (
	OutboxEventTypeOrderCreated OutboxEventType = "order_created"
)
