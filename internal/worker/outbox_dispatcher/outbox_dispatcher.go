package outbox_dispatcher

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/null/v8"
	"orders.go/m/internal/infrastructure/event"
	"orders.go/m/internal/infrastructure/event/common"
	"orders.go/m/internal/models"
	"orders.go/m/internal/repository/outbox"
	"time"
)

type OutboxDispatcher struct {
	eventPublisher event.EventPublisher
	outboxRepo     *outbox.OutboxRepo
	pollInterval   time.Duration
	workerCount    int
}

func NewOutboxDispatcher(
	eventPublisher event.EventPublisher,
	outboxRepo *outbox.OutboxRepo,
) *OutboxDispatcher {
	return &OutboxDispatcher{
		eventPublisher: eventPublisher,
		pollInterval:   time.Second,
		outboxRepo:     outboxRepo,
		workerCount:    5,
	}
}

func (o *OutboxDispatcher) Execute() error {
	log.Info().Msg("Starting outbox dispatcher")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ticker := time.NewTicker(o.pollInterval)
	defer ticker.Stop()

	workerChan := make(chan models.Outbox, o.workerCount)
	for range o.workerCount {
		go o.worker(ctx, workerChan)
	}

	for {
		select {
		case <-ticker.C:
			log.Info().Msg("Fetching outboxes")
			events, _ := o.outboxRepo.FetchUnsent(ctx)

			for _, event := range events {
				workerChan <- *event
			}
		default:
			log.Warn().Msg("worker pool full, skipping for now")
			time.Sleep(time.Second * 5)
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return nil
}

func (o *OutboxDispatcher) worker(ctx context.Context, channel <-chan models.Outbox) {
	for {
		select {
		case outboxEvent := <-channel:
			o.process(ctx, outboxEvent)
		case <-ctx.Done():
			return
		}

	}

}

func (o *OutboxDispatcher) process(ctx context.Context, eventOutbox models.Outbox) {
	log.Info().Msg("Processing outbox...")

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	topic, ok := common.AggregateTypeToTopic[eventOutbox.AggregateType]
	if !ok {
		return
	}

	key := fmt.Sprintf("%s-%s", eventOutbox.AggregateType, eventOutbox.ID)
	err := o.eventPublisher.Publish(ctx, topic, key, eventOutbox.Payload)
	if err != nil {
		log.Err(err).Msg("failed to publish outbox")
	}

	eventOutbox.SentAt = null.TimeFrom(time.Now())

	err = o.outboxRepo.Save(ctx, nil, &eventOutbox)
	if err != nil {
		log.Err(err).Msg("failed to save outbox")
	}

	log.Info().Msg("outbox processed...")

}
