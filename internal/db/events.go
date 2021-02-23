package db

import (
	"context"

	"github.com/sebvautour/event-manager/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	// EventsCollection represent the Mongo events collection
	EventsCollection *mongo.Collection
)

// AddEvent adds a new event to the DB
func AddEvent(ctx context.Context, evt *model.Event) (err error) {
	_, err = EventsCollection.InsertOne(ctx, evt)
	return err
}

// EventsForAlert returns the events for a given alert ID
func EventsForAlert(ctx context.Context, alertID primitive.ObjectID) (events []model.Event, err error) {
	cursor, err := EventsCollection.Find(ctx, bson.M{"alert_id": alertID})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &events); err != nil {
		return nil, err
	}
	return events, nil
}
