package events

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebvautour/event-manager/cmd/api/v1/helpers"
	"github.com/sebvautour/event-manager/internal/msgbus"
	"github.com/sebvautour/event-manager/pkg/model"
)

func InitRouteGroup(r *gin.RouterGroup) {
	r.POST("", postEventHandler)
}

// @Summary Post event
// @Description Adds an event to the message bus
// @Tags events
// @Accept  json
// @Param event body string false "event payload"
// @Router /events [post]
func postEventHandler(c *gin.Context) {
	var evt model.Event
	if err := c.BindJSON(&evt); err != nil {
		helpers.Error(c, "failed to decode JSON body", http.StatusBadRequest, err.Error())
		return
	}

	if err := msgbus.AddEvent(c.Request.Context(), &evt); err != nil {
		helpers.Error(c, "failed to add event", http.StatusInternalServerError, err.Error())
		return
	}
}
