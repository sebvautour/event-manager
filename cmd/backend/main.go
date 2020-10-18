package main

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/prometheus/common/log"
	"github.com/sebvautour/event-manager/internal/db"
	"github.com/sebvautour/event-manager/internal/processor"
	"github.com/sebvautour/event-manager/internal/service"
	"github.com/sebvautour/event-manager/pkg/client"
	"github.com/sirupsen/logrus"
	"gocloud.dev/pubsub"

	//_ "gocloud.dev/pubsub/mempubsub"
	_ "gocloud.dev/pubsub/kafkapubsub"
)

var api *client.Client
var exitedEventProcessors = make(chan error, 0)
var exitedActionProcessors = make(chan error, 0)
var ctx = context.Background()

func init() {
	if err := db.Init(ctx); err != nil {
		log.Fatal("Failed to init db: " + err.Error())
	}

	api = client.New()

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

func runExample() {
	// Create a topic.

	//eventsTopic, err := pubsub.OpenTopic(ctx, "mem://events")
	log.Info("opening events topic")
	eventsTopic, err := pubsub.OpenTopic(ctx, "kafka://events")
	if err != nil {
		logrus.Fatal(err)
	}
	//defer eventsTopic.Shutdown(ctx)
	// Create a topic.
	//actionsTopic, err := pubsub.OpenTopic(ctx, "mem://actions")
	log.Info("opening actions topic")
	actionsTopic, err := pubsub.OpenTopic(ctx, "kafka://actions")
	if err != nil {
		logrus.Fatal(err)
	}
	//defer actionsTopic.Shutdown(ctx)

	for i := 1; i < 10; i++ {

		err = eventsTopic.Send(ctx, &pubsub.Message{
			Body: []byte("Event Message " + strconv.Itoa(i)),
			// Metadata is optional and can be nil.
			Metadata: map[string]string{
				// These are examples of metadata.
				// There is nothing special about the key names.
				"language":   "en",
				"importance": "high",
			},
		})
		if err != nil {
			log.Fatalln(err)
		}
		log.Info("event sent")
		err = actionsTopic.Send(ctx, &pubsub.Message{
			Body: []byte("Action Message " + strconv.Itoa(i)),
			// Metadata is optional and can be nil.
			Metadata: map[string]string{
				// These are examples of metadata.
				// There is nothing special about the key names.
				"language":   "en",
				"importance": "high",
			},
		})
		if err != nil {
			log.Fatalln(err)
		}

		time.Sleep(time.Second)
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
		API:     api,
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
		API:     api,
		Log:     logrus.WithField("cmp", "action-processor"),
	}).ActionProcessor(subscription)
}
