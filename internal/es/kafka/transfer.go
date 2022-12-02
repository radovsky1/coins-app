package kafka

import (
	"bytes"
	"coins-app/internal/core"
	"context"
	"encoding/json"
)

type TransferWriter struct {
	writer    *Writer
	topicName string
}

type event struct {
	Type  string
	Value core.Transfer
}

func NewTransferWriter(writer *Writer, topicName string) *TransferWriter {
	return &TransferWriter{
		writer:    writer,
		topicName: topicName,
	}
}

func (w *TransferWriter) PublishTransferCreated(ctx context.Context, transfer core.Transfer) error {
	return w.publish(ctx, "transfer.event.created", transfer)
}

func (w *TransferWriter) publish(ctx context.Context, msgType string, transfer core.Transfer) error {
	var b bytes.Buffer

	evt := event{
		Type:  msgType,
		Value: transfer,
	}

	if err := json.NewEncoder(&b).Encode(evt); err != nil {
		return err
	}

	return w.writer.WriteMessages(ctx, w.topicName, b.Bytes())
}
