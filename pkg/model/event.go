package model

import (
	"errors"
	"time"
)

// ErrorMissingDedupKey is returned when the given event payload does not contain a dedup_key
var ErrorMissingDedupKey = errors.New("missing dedup_key")

// Event represents a single monitoring event
type Event struct {
	ID               interface{}            `json:"id" bson:"_id"`
	DedupKey         string                 `json:"dedup_key"`
	SpecificGroupKey string                 `json:"specific_group_key"`
	Entity           string                 `json:"entity" `
	Message          string                 `json:"message" `
	Severity         string                 `json:"severity"`
	Status           string                 `json:"status" `
	Details          map[string]interface{} `json:"details"`
	AlertID          interface{}            `json:"alert_id"`
	EventTime        time.Time              `json:"event_time"`
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
