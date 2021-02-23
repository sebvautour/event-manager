package processor

import (
	"encoding/json"
	"fmt"

	"github.com/sebvautour/event-manager/internal/db"
	"github.com/sebvautour/event-manager/pkg/model"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (p *Processor) EventProcessor(reader *kafka.Reader) error {
	return p.run(p.processEventMsgFunc, reader)
}

func (p *Processor) processEventMsgFunc(msg kafka.Message) (retry bool, err error) {
	// create event obj
	var evt model.Event
	if err := json.Unmarshal(msg.Value, &evt); err != nil {
		return false, fmt.Errorf("json unmarshal event: %w", err)
	}
	evt.ID = primitive.NewObjectID()

	// check if we have an alert for it
	alert, found, err := db.AlertByDedupKey(p.s.Context, evt.DedupKey)
	if err != nil {
		return true, fmt.Errorf("query alert: %w", err)

	}

	if found {
		// existing alert
		db.UpdateAlertWithEvent(p.s.Context, evt)
		if err != nil {
			return true, fmt.Errorf("updating alert: %w", err)
		}

		p.s.Log.Infof("Updated alert %s\n", alert.ID.(primitive.ObjectID).Hex())
	} else {
		// new alert
		alert = model.NewAlertFromEvent(&evt)

		if err := db.AddAlert(p.s.Context, alert); err != nil {
			return true, fmt.Errorf("creating alert: %w", err)
		}

		p.s.Log.Infof("Created alert %s\n", alert.ID.(primitive.ObjectID).Hex())
	}

	evt.AlertID = alert.ID
	if err := db.AddEvent(p.s.Context, &evt); err != nil {
		return true, fmt.Errorf("creating event: %w", err)
	}

	p.s.Log.Infof("Created event %s: %q\n", evt.ID.(primitive.ObjectID).Hex(), msg.Value)

	return false, nil
}
