package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Alert struct {
	EventAlert
	GroupKey       string    `json:"group_key" bson:"group_key"`
	EventCount     int       `json:"event_count" bson:"event_count"`
	LastEventTime  time.Time `json:"last_event_time" bson:"last_event_time"`
	FirstEventTime time.Time `json:"first_event_time" bson:"first_event_time"`
}

func NewAlertFromEvent(evt *Event) *Alert {
	return &Alert{
		EventAlert: EventAlert{
			ID:       primitive.NewObjectID(),
			DedupKey: evt.DedupKey,
			Entity:   evt.Entity,
			Message:  evt.Message,
			Severity: evt.Severity,
			Status:   evt.Status,
			Details:  evt.Details,
		},
		GroupKey:       evt.DedupKey,
		EventCount:     1,
		LastEventTime:  evt.EventTime,
		FirstEventTime: evt.EventTime,
	}
}
