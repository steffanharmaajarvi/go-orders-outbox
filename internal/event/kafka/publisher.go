package kafka

import (
	"context"
	"fmt"
	"github.com/twmb/franz-go/pkg/kgo"
	"log"
	"time"
)

type KafkaPublisher struct {
	client *kgo.Client
}

type KafkaConfig struct {
	Host    string
	GroupID string
}

func NewKafkaPublisher(cfg KafkaConfig) *KafkaPublisher {
	fmt.Printf("%s %s", cfg.Host, cfg.GroupID)

	client, err := kgo.NewClient(
		kgo.SeedBrokers(cfg.Host),
		kgo.ProducerLinger(100*time.Millisecond),
		kgo.RecordPartitioner(kgo.StickyKeyPartitioner(nil)),
	)
	if err != nil {
		panic(err)
	}

	return &KafkaPublisher{
		client: client,
	}
}

func (p *KafkaPublisher) Close() {
	p.client.Close()
}

func (p *KafkaPublisher) Publish(ctx context.Context, topic string, key string, payload []byte) error {
	record := &kgo.Record{
		Topic: topic,
		Key:   []byte(key),
		Value: payload,
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	err := p.client.ProduceSync(ctx, record).FirstErr()
	if err != nil {
		log.Fatalf("failed to produce: %v", err)

		return err
	}

	return nil
}
