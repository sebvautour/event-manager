package model

import (
	"errors"
	"time"
)

var ErrorMissingDedupKey = errors.New("missing dedup_key")

type Event struct {
	EventAlert
	AlertID   interface{} `json:"alert_id" bson:"alert_id"`
	EventTime time.Time   `json:"event_time" bson:"event_time"`
}

// EventAlert contains field common in events and alerts
type EventAlert struct {
	ID       interface{}            `json:"id" bson:"_id"`
	DedupKey string                 `json:"dedup_key" bson:"dedup_key"`
	Entity   string                 `json:"entity" bson:"entity"`
	Message  string                 `json:"message" bson:"message"`
	Severity string                 `json:"severity" bson:"severity"`
	Status   string                 `json:"status" bson:"status"`
	Details  map[string]interface{} `json:"details" bson:"details"`
}

// Validate ensures an Event has all the required information and fills certain fields with default values where possible
func (evt *Event) Validate() error {
	if evt.DedupKey == "" {
		return ErrorMissingDedupKey
	}

	if evt.Severity == "" {
		evt.Severity = SeverityDefault
	}

	if evt.Status == "" {
		evt.Status = StatusDefault
	}

	if evt.Details == nil {
		evt.Details = make(map[string]interface{})
	}

	if evt.EventTime.IsZero() {
		evt.EventTime = time.Now()
	}

	return nil
}
