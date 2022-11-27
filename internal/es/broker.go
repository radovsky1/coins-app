package es

import (
	"coins-app/internal/core"
	"coins-app/internal/es/kafka"
	"context"
)

type AccountMessageBroker interface {
	PublishAccountCreated(ctx context.Context, account core.Account) error
	PublishAccountUpdated(ctx context.Context, account core.Account) error
	PublishAccountDeleted(ctx context.Context, account core.Account) error
	PublishAccountLocked(ctx context.Context, account core.Account) error
	PublishAccountUnlocked(ctx context.Context, account core.Account) error
	PublishAccountBalanceUpdated(ctx context.Context, account core.Account) error
}

type TransferMessageBroker interface {
	PublishTransferCreated(ctx context.Context, transfer core.Transfer) error
}

type MessageBroker struct {
	Account  AccountMessageBroker
	Transfer TransferMessageBroker
}

func NewKafkaMessageBroker(writer *kafka.Writer) *MessageBroker {
	return &MessageBroker{
		Account:  kafka.NewAccountWriter(writer, writer.Topic),
		Transfer: kafka.NewTransferWriter(writer, writer.Topic),
	}
}
