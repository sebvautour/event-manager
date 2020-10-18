package db

import (
	"context"
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"gocloud.dev/docstore"
	"gocloud.dev/docstore/mongodocstore"
)

var (
	// MongoClient gives direct access to the Mongo client
	MongoClient *mongo.Client

	// Events docstore collection
	Events      *docstore.Collection
	EventsMongo *mongo.Collection

	// Alerts collection
	Alerts      *docstore.Collection
	AlertsMongo *mongo.Collection
)

var lck = sync.Mutex{}
var initDone = false
var closeDone = false

// Init db package
func Init(ctx context.Context) error {
	lck.Lock()
	defer lck.Unlock()

	if initDone {
		return nil
	}

	if err := initDB(ctx); err != nil {
		return err
	}

	initDone = true
	return nil
}

// Close db package
func Close(ctx context.Context) error {
	lck.Lock()
	defer lck.Unlock()

	if closeDone {
		return nil
	}

	if err := close(ctx); err != nil {
		return err
	}

	closeDone = true
	return nil
}

func initDB(ctx context.Context) (err error) {
	MongoClient, err = mongodocstore.Dial(ctx, "mongodb://localhost:27017")
	if err != nil {
		return fmt.Errorf("connecting to mongo: %w", err)
	}

	emDB := MongoClient.Database("event-manager")
	EventsMongo = emDB.Collection("events")
	AlertsMongo = emDB.Collection("alerts")

	Events, err = mongodocstore.OpenCollection(EventsMongo, "ID", nil)
	if err != nil {
		return fmt.Errorf("opening events collection: %w", err)
	}

	Alerts, err = mongodocstore.OpenCollection(AlertsMongo, "ID", nil)
	if err != nil {
		return fmt.Errorf("opening alerts collection: %w", err)
	}
	return nil
}

// close db package
func close(ctx context.Context) error {
	if err := Alerts.Close(); err != nil {
		return fmt.Errorf("closing alerts collection: %w", err)
	}

	if err := Events.Close(); err != nil {
		return fmt.Errorf("closing alerts collection: %w", err)
	}

	if err := MongoClient.Disconnect(ctx); err != nil {
		return fmt.Errorf("disconnecting mongo client: %w", err)
	}
	return nil
}
