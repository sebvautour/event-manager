package msgbus

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/sebvautour/event-manager/pkg/model"
	"gocloud.dev/pubsub"
	_ "gocloud.dev/pubsub/kafkapubsub"
)

var EventsTopic *pubsub.Topic

func InitEventsTopic(ctx context.Context) (err error) {
	os.Setenv("KAFKA_BROKERS", "192.168.2.12:9092")

	EventsTopic, err = pubsub.OpenTopic(ctx, "kafka://events")
	return err
}

func AddEvent(ctx context.Context, evt *model.Event) (err error) {
	if err := evt.Validate(); err != nil {
		return fmt.Errorf("invalid event: %w", err)
	}

	b, err := json.Marshal(evt)
	if err != nil {
		return fmt.Errorf("JSON marshal: %w", err)
	}

	if err = EventsTopic.Send(ctx, &pubsub.Message{Body: b}); err != nil {
		return fmt.Errorf("publish event: %w", err)
	}

	return nil
}
