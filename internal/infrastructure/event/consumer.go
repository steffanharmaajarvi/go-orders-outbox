package event

import "context"

type ConsumerMessage struct {
	Key   []byte
	Value []byte
}

type EventConsumer interface {
	Start(ctx context.Context) (<-chan ConsumerMessage, error)
	Stop()
}
