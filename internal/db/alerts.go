package db

import (
	"context"

	"github.com/sebvautour/event-manager/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	// AlertsCollection represents the Mongo alerts collection
	AlertsCollection *mongo.Collection
)

// AlertByDedupKey returns the Alert based off of the given dedup_key
func AlertByDedupKey(ctx context.Context, dedupKey string) (alert *model.Alert, found bool, err error) {
	alert = &model.Alert{}
	q := AlertsCollection.FindOne(ctx, bson.M{"dedup_key": dedupKey})

	err = q.Decode(alert)

	if err == nil {
		return alert, true, nil
	}

	if err == mongo.ErrNoDocuments {
		return nil, false, nil
	}
	return nil, false, err
}

// AlertByID returns the Alert based off of the given alert ID
func AlertByID(ctx context.Context, id primitive.ObjectID) (alert *model.Alert, found bool, err error) {
	alert = &model.Alert{}
	q := AlertsCollection.FindOne(ctx, bson.M{"_id": id})

	err = q.Decode(alert)

	if err == nil {
		return alert, true, nil
	}

	if err == mongo.ErrNoDocuments {
		return nil, false, nil
	}
	return nil, false, err
}

// AlertSearch returns the alerts that match a given filter
func AlertSearch(ctx context.Context, filter primitive.D) (alerts []model.Alert, err error) {
	cursor, err := AlertsCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &alerts); err != nil {
		return nil, err
	}
	return alerts, nil
}

// AddAlert adds a new alert in the DB
func AddAlert(ctx context.Context, alert *model.Alert) (err error) {
	_, err = AlertsCollection.InsertOne(ctx, alert)
	return err
}

// UpdateAlertWithEvent updates the alert with the given event info
func UpdateAlertWithEvent(ctx context.Context, evt model.Event) (err error) {
	_, err = AlertsCollection.UpdateOne(ctx,
		bson.M{"dedup_key": evt.DedupKey},
		bson.D{
			{"$set", bson.M{
				"labels":          evt.Labels,
				"severity":        evt.Severity,
				"status":          evt.Status,
				"last_event_time": evt.EventTime,
			}},
			{"$inc", bson.M{
				"event_count": 1,
			}},
		})

	return err
}
