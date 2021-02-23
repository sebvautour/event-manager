package msgbus

import (
	kafka "github.com/segmentio/kafka-go"
)

var EventsReader *kafka.Reader
var ActionsReader *kafka.Reader

func InitEventsReader() {
	// make a new reader that consumes from topic-A
	EventsReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"192.168.2.12:9092"},
		GroupID:  "event-manager",
		Topic:    "events",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

}

func InitActionsReader() {
	// make a new reader that consumes from topic-A
	ActionsReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"192.168.2.12:9092"},
		GroupID:  "event-manager",
		Topic:    "actions",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

}
