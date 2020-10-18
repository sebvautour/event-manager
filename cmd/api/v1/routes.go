package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sebvautour/event-manager/cmd/api/v1/alerts"
	"github.com/sebvautour/event-manager/cmd/api/v1/events"
)

func InitRouteGroup(r *gin.RouterGroup) {
	alerts.InitRouteGroup(r.Group("/alerts"))
	alerts.InitRouteGroup(r.Group("/alert"))
	events.InitRouteGroup(r.Group("/events"))
}
