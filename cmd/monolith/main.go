package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
	v1 "github.com/sebvautour/event-manager/cmd/api/v1"
	"github.com/sebvautour/event-manager/internal/db"
	"github.com/sebvautour/event-manager/internal/msgbus"
	"github.com/sebvautour/event-manager/internal/processor"
	"github.com/sebvautour/event-manager/internal/service"
	"github.com/sirupsen/logrus"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var exitedEventProcessors = make(chan error, 0)
var exitedActionProcessors = make(chan error, 0)

var ctx = context.Background()

var svc *service.Service

func init() {
	if err := db.Init(ctx); err != nil {
		log.Fatal("Failed to init db: " + err.Error())
	}

	msgbus.InitEventsWriter()

}

func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://localhost:8080/swagger/doc.json")))

	v1.Init(r.Group("/api/v1"))

	r.Static("/js", "frontend/dist/js")
	r.Static("/css", "frontend/dist/css")
	r.StaticFile("/", "frontend/dist/index.html")

	go svc.Log.Fatal(r.Run(":8080"))

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
