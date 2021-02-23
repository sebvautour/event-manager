package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/sebvautour/event-manager/cmd/api/docs" // docs is generated by Swag CLI, you have to import it.

	v1 "github.com/sebvautour/event-manager/cmd/api/v1"
	"github.com/sebvautour/event-manager/internal/db"
	"github.com/sebvautour/event-manager/internal/msgbus"
	"github.com/sebvautour/event-manager/internal/service"
)

var ctx = context.Background()

var svc *service.Service

func init() {
	if err := db.Init(ctx); err != nil {
		log.Fatal("Failed to init db: " + err.Error())
	}

	msgbus.InitEventsWriter()

}

// @title Event Manager API
// @version 1.0
// @description This the API component of the Event Manager platform.

// @license.name MIT License
// @license.url https://github.com/sebvautour/event-manager/blob/master/LICENSE

// @host localhost:8080
// @BasePath /api/v1
func main() {

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://localhost:8080/swagger/doc.json")))

	v1.Init(r.Group("/api/v1"))

	svc.Log.Fatal(r.Run(":8080"))
}
