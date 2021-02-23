package model

import (
	"errors"
	"time"
)

// ErrorMissingDedupKey is returned when the given event payload does not contain a dedup_key
var ErrorMissingDedupKey = errors.New("missing dedup_key")

// Event represents a single monitoring event
type Event struct {
	ID               interface{}            `bson:"_id" json:"id"`
	DedupKey         string                 `bson:"dedup_key" json:"dedup_key"`
	Severity         string                 `bson:"severity" json:"severity"`
	Status           string                 `bson:"status" json:"status"`
	Labels           map[string]interface{} `bson:"labels" json:"labels"`
	SpecificGroupKey string                 `bson:"specific_group_key" json:"specific_group_key"`
	AlertID          interface{}            `bson:"alert_id" json:"alert_id"`
	EventTime        time.Time              `bson:"event_time" json:"event_time"`
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

	if evt.Labels == nil {
		evt.Labels = make(map[string]interface{})
	}

	if evt.EventTime.IsZero() {
		evt.EventTime = time.Now()
	}

	return nil
}
