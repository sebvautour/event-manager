package main

import (
	"context"

	"github.com/prometheus/common/log"
	"github.com/sebvautour/event-manager/internal/db"
	"github.com/sebvautour/event-manager/internal/msgbus"
	"github.com/sebvautour/event-manager/internal/processor"
	"github.com/sebvautour/event-manager/internal/service"
	"github.com/sirupsen/logrus"
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
	msgbus.InitEventsReader()

	log.Info("starting events processor")
	exitedEventProcessors <- processor.New(&service.Service{
		Context: ctx,
		Log:     logrus.WithField("cmp", "event-processor"),
	}).EventProcessor(msgbus.EventsReader)
}

func launchActionProcessor() {
	msgbus.InitActionsReader()

	log.Info("starting actions processor")
	exitedActionProcessors <- processor.New(&service.Service{
		Context: ctx,
		Log:     logrus.WithField("cmp", "action-processor"),
	}).ActionProcessor(msgbus.ActionsReader)
}
