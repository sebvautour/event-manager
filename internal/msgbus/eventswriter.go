package msgbus

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sebvautour/event-manager/pkg/model"
	kafka "github.com/segmentio/kafka-go"
)

var eventsWriter *kafka.Writer

func InitEventsWriter() {
	eventsWriter = kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"192.168.2.12:9092"},
		Topic:   "events",
	})

}

func AddEvent(ctx context.Context, evt *model.Event) (err error) {
	if err := evt.Validate(); err != nil {
		return fmt.Errorf("invalid event: %w", err)
	}

	b, err := json.Marshal(evt)
	if err != nil {
		return fmt.Errorf("JSON marshal: %w", err)
	}

	return eventsWriter.WriteMessages(ctx, kafka.Message{
		Key:   []byte(evt.DedupKey),
		Value: b,
	})

}
