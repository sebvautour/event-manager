package main

import (
	"context"
	"os"

	"github.com/prometheus/common/log"
	"github.com/sebvautour/event-manager/internal/db"
	"github.com/sebvautour/event-manager/internal/processor"
	"github.com/sebvautour/event-manager/internal/service"
	"github.com/sirupsen/logrus"
	"gocloud.dev/pubsub"

	//_ "gocloud.dev/pubsub/mempubsub"
	_ "gocloud.dev/pubsub/kafkapubsub"
)

var exitedEventProcessors = make(chan error, 0)
var exitedActionProcessors = make(chan error, 0)
var ctx = context.Background()

func init() {
	if err := db.Init(ctx); err != nil {
		log.Fatal("Failed to init db: " + err.Error())
	}

}

func main() {
	os.Setenv("KAFKA_BROKERS", "192.168.2.12:9092")

	for i := 0; i < 1; i++ {
		go launchEventProcessor()
	}

	for i := 0; i < 1; i++ {
		go launchActionProcessor()
	}

	for {
		select {
		case err := <-exitedEventProcessors:
			if err != nil {
				log.Error(err)
			}
			go launchEventProcessor()
		case err := <-exitedActionProcessors:
			if err != nil {
				log.Error(err)
			}
			go launchActionProcessor()
		default:
		}
	}

}

func launchEventProcessor() {
	// Create a subscription connected to that topic.
	//subscription, err := pubsub.OpenSubscription(ctx, "mem://events")
	log.Info("opening events sub")
	subscription, err := pubsub.OpenSubscription(ctx, "kafka://eventsprocessor?topic=events")
	if err != nil {
		logrus.Error(err)
	}

	log.Info("starting events processor")
	exitedEventProcessors <- processor.New(&service.Service{
		Context: ctx,
		Log:     logrus.WithField("cmp", "event-processor"),
	}).EventProcessor(subscription)
}

func launchActionProcessor() {
	// Create a subscription connected to that topic.
	//subscription, err := pubsub.OpenSubscription(ctx, "mem://actions")
	log.Info("opening actions sub")
	subscription, err := pubsub.OpenSubscription(ctx, "kafka://actionsprocessor?topic=actions")
	if err != nil {
		logrus.Error(err)
	}

	log.Info("starting actions processor")
	exitedActionProcessors <- processor.New(&service.Service{
		Context: ctx,
		Log:     logrus.WithField("cmp", "action-processor"),
	}).ActionProcessor(subscription)
}
