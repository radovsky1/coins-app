package kafka

import (
	"context"
	kafkago "github.com/segmentio/kafka-go"
)

type Writer struct {
	writer *kafkago.Writer
	Topic  string
}

type Config struct {
	Address string
	Topic   string
}

func NewKafkaWriter(config Config) *Writer {
	writer := &kafkago.Writer{
		Addr:  kafkago.TCP(config.Address),
		Topic: config.Topic,
	}
	return &Writer{
		writer: writer,
		Topic:  config.Topic,
	}
}

func (w *Writer) WriteMessages(ctx context.Context, topic string, messages ...[]byte) error {
	for _, message := range messages {
		err := w.writer.WriteMessages(ctx, kafkago.Message{
			Topic: topic,
			Value: message,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
