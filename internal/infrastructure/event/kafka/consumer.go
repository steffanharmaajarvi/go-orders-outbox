package kafka

import (
	"context"
	log2 "github.com/rs/zerolog/log"
	"github.com/twmb/franz-go/pkg/kgo"
	"orders.go/m/internal/infrastructure/event"
)

type KafkaConsumer struct {
	client *kgo.Client
}

type KafkaConsumerConfig struct {
	Topic   string
	Host    string
	GroupID string
}

func NewKafkaConsumer(cfg KafkaConsumerConfig) (*KafkaConsumer, error) {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(cfg.Host),
		kgo.ConsumeTopics(cfg.Topic),
		kgo.ConsumerGroup(cfg.GroupID),
		kgo.AutoCommitMarks(),
	)
	if err != nil {
		panic(err)
	}

	return &KafkaConsumer{
		client: client,
	}, nil
}

func (c *KafkaConsumer) Start(ctx context.Context) (<-chan event.ConsumerMessage, error) {
	messageChan := make(chan event.ConsumerMessage)

	go func() {
		defer func() {
			close(messageChan)
			log2.Info().Msg("Consumer stopped")
		}()

		for {
			records := c.client.PollFetches(ctx).Records()
			if ctx.Err() != nil {
				log2.Err(ctx.Err()).Msgf("Consumer stopped with error")
				return
			}

			for _, record := range records {
				message := event.ConsumerMessage{
					Key:   record.Key,
					Value: record.Value,
				}

				select {
				case messageChan <- message:
				case <-ctx.Done():
					return
				}

			}

			c.client.MarkCommitRecords(records...)
		}
	}()

	return messageChan, nil
}

func (c *KafkaConsumer) Stop() {
	c.client.Close()
}
