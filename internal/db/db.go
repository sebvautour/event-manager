package db

import (
	"context"
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// MongoClient gives direct access to the Mongo client
	MongoClient *mongo.Client
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
	MongoClient, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	err = MongoClient.Connect(ctx)
	if err != nil {
		return err
	}

	emDB := MongoClient.Database("event-manager")
	EventsCollection = emDB.Collection("events")
	AlertsCollection = emDB.Collection("alerts")

	return nil
}

// close db package
func close(ctx context.Context) error {

	if err := MongoClient.Disconnect(ctx); err != nil {
		return fmt.Errorf("disconnecting mongo client: %w", err)
	}
	return nil
}
