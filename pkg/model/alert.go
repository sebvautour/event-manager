package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Alert is a dedup of one or more events with the same dedup key
type Alert struct {
	ID             interface{}            `json:"id" bson:"_id"`
	DedupKey       string                 `json:"dedup_key" `
	Entity         string                 `json:"entity" `
	Message        string                 `json:"message" `
	Severity       string                 `json:"severity"`
	Status         string                 `json:"status" `
	Details        map[string]interface{} `json:"details"`
	GroupKey       string                 `json:"group_key"` // bson:"group_key"
	EventCount     int                    `json:"event_count"`
	LastEventTime  time.Time              `json:"last_event_time"`
	FirstEventTime time.Time              `json:"first_event_time"`
}

// NewAlertFromEvent creates a new Alert for a given Event
func NewAlertFromEvent(evt *Event) *Alert {
	a := &Alert{

		ID:       primitive.NewObjectID(),
		DedupKey: evt.DedupKey,
		Entity:   evt.Entity,
		Message:  evt.Message,
		Severity: evt.Severity,
		Status:   evt.Status,
		Details:  evt.Details,

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
