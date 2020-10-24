package processor

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/sebvautour/event-manager/internal/db"
	"github.com/sebvautour/event-manager/pkg/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gocloud.dev/docstore"
	"gocloud.dev/pubsub"
)

func (p *Processor) EventProcessor(sub *pubsub.Subscription) error {
	return p.run(p.processEventMsgFunc, sub)
}

func (p *Processor) processEventMsgFunc(msg *pubsub.Message) (err error) {
	// create event obj
	var evt model.Event
	if err := json.Unmarshal(msg.Body, &evt); err != nil {
		return fmt.Errorf("json unmarshal event: %w", err)
	}
	evt.ID = primitive.NewObjectID()

	// check if we have an alert for it
	aiter := db.Alerts.Query().Where("DedupKey", "=", evt.DedupKey).Limit(1).Get(p.s.Context, "_id")
	defer aiter.Stop()

	alert := &model.Alert{}
	err = aiter.Next(p.s.Context, alert)
	if err == io.EOF {
		// new alert
		alert = model.NewAlertFromEvent(&evt)

		err = db.Alerts.Create(p.s.Context, alert)
		if err != nil {
			return fmt.Errorf("creating alert: %w", err)
		}

		p.s.Log.Infof("Created alert %s\n", alert.ID.(primitive.ObjectID).Hex())

	} else if err != nil {
		return fmt.Errorf("querying alert: %w", err)
	} else {
		// existing alert
		err := db.Alerts.Actions().Update(alert, docstore.Mods{
			"Entity":        evt.Entity,
			"Message":       evt.Message,
			"Severity":      evt.Severity,
			"Status":        evt.Status,
			"Details":       evt.Details,
			"EventCount":    docstore.Increment(1),
			"LastEventTime": evt.EventTime,
		}).Do(p.s.Context)

		if err != nil {
			return fmt.Errorf("updating alert: %w", err)
		}

		p.s.Log.Infof("Updated alert %s\n", alert.ID.(primitive.ObjectID).Hex())

	}

	evt.AlertID = alert.ID
	err = db.Events.Create(p.s.Context, &evt)
	if err != nil {
		return fmt.Errorf("creating event: %w", err)
	}

	p.s.Log.Infof("Created event %s: %q\n", evt.ID.(primitive.ObjectID).Hex(), msg.Body)
	// Messages must always be acknowledged with Ack.
	msg.Ack()
	return nil
}
