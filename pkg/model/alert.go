package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Alert is a dedup of one or more events with the same dedup key
type Alert struct {
	ID             interface{}            `bson:"_id" json:"id"`
	DedupKey       string                 `bson:"dedup_key" json:"dedup_key"`
	Severity       string                 `bson:"severity" json:"severity"`
	Status         string                 `bson:"status" json:"status"`
	Labels         map[string]interface{} `bson:"labels" json:"labels"`
	GroupKey       string                 `bson:"group_key" json:"group_key"`
	EventCount     int                    `bson:"event_count" json:"event_count"`
	LastEventTime  time.Time              `bson:"last_event_time" json:"last_event_time"`
	FirstEventTime time.Time              `bson:"first_event_time" json:"first_event_time"`
}

// NewAlertFromEvent creates a new Alert for a given Event
func NewAlertFromEvent(evt *Event) *Alert {
	a := &Alert{

		ID:             primitive.NewObjectID(),
		DedupKey:       evt.DedupKey,
		Severity:       evt.Severity,
		Status:         evt.Status,
		Labels:         evt.Labels,
		GroupKey:       evt.SpecificGroupKey,
		EventCount:     1,
		LastEventTime:  evt.EventTime,
		FirstEventTime: evt.EventTime,
	}
	if a.GroupKey == "" {
		a.GroupKey = a.DedupKey
	}
	return a
}
