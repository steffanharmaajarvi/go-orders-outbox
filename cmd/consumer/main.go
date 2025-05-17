package main

import (
	"context"
	log2 "github.com/rs/zerolog/log"
	"orders.go/m/internal/consumer/orders"
	"orders.go/m/internal/infrastructure/event/kafka"
	orders3 "orders.go/m/internal/repository/orders"
	"orders.go/m/internal/setup"
	uow2 "orders.go/m/internal/uow"
	orders2 "orders.go/m/internal/usecases/orders"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	container, err := setup.BuildContainer()
	if err != nil {
		panic(err)
	}

	consumer, err := kafka.NewKafkaConsumer(kafka.KafkaConsumerConfig{
		Topic:   container.Config.GetString("kafka.topic"),
		Host:    container.Config.GetString("kafka.host"),
		GroupID: container.Config.GetString("kafka.consumer_group"),
	})
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	// Graceful shutdown
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(sig)

		select {
		case <-sig:
			log2.Info().Msg("Shutdown via signal")
		case <-ctx.Done():
			log2.Info().Msg("Context canceled before signal")
		}

		cancel()
	}()

	orderDb, err := container.Database.GetConnection("orders")
	if err != nil {
		panic(err)
	}

	uow := uow2.NewUnitOfWork(orderDb)

	ordersRepo := orders3.NewOrdersRepo(orderDb)

	confirmOrderUseCase := orders2.NewConfirmOrderUseCase(ordersRepo, uow)

	orderConsumer := orders.NewOrderConsumer(consumer, confirmOrderUseCase)
	go orderConsumer.Start(ctx)

	<-ctx.Done()
}
