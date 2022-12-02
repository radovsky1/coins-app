package kafka

import (
	"bytes"
	"coins-app/internal/core"
	"context"
	"encoding/json"
)

type AccountWriter struct {
	writer    *Writer
	topicName string
}

type AccountEvent struct {
	Type  string
	Value core.Account
}

func NewAccountWriter(writer *Writer, topicName string) *AccountWriter {
	return &AccountWriter{
		writer:    writer,
		topicName: topicName,
	}
}

func (w *AccountWriter) PublishAccountCreated(ctx context.Context, account core.Account) error {
	return w.publish(ctx, "account.event.created", account)
}

func (w *AccountWriter) PublishAccountUpdated(ctx context.Context, account core.Account) error {
	return w.publish(ctx, "account.event.updated", account)
}

func (w *AccountWriter) PublishAccountDeleted(ctx context.Context, account core.Account) error {
	return w.publish(ctx, "account.event.deleted", account)
}

func (w *AccountWriter) PublishAccountLocked(ctx context.Context, account core.Account) error {
	return w.publish(ctx, "account.event.locked", account)
}

func (w *AccountWriter) PublishAccountUnlocked(ctx context.Context, account core.Account) error {
	return w.publish(ctx, "account.event.unlocked", account)
}

func (w *AccountWriter) PublishAccountBalanceUpdated(ctx context.Context, account core.Account) error {
	return w.publish(ctx, "account.event.balance.updated", account)
}

func (w *AccountWriter) publish(ctx context.Context, msgType string, account core.Account) error {
	var b bytes.Buffer

	evt := AccountEvent{
		Type:  msgType,
		Value: account,
	}

	if err := json.NewEncoder(&b).Encode(evt); err != nil {
		return err
	}

	return w.writer.WriteMessages(ctx, w.topicName, b.Bytes())
}
