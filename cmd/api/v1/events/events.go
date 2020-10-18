package events

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebvautour/event-manager/cmd/api/v1/helpers"
	"github.com/sebvautour/event-manager/internal/msgbus"
	"github.com/sebvautour/event-manager/pkg/model"
)

func InitRouteGroup(r *gin.RouterGroup) {
	r.POST("", postEventHandler)
	r.POST("/alertmanager", postAlertManagerEventHandler)
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

// @Summary Post AlertManager event
// @Description Adds an event to the message bus
// @Tags events
// @Accept  json
// @Param event body string false "Alertmanager event payload"
// @Router /events/alertmanager [post]
func postAlertManagerEventHandler(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		helpers.Error(c, "failed get raw data", http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(string(b))

	helpers.NotImplemented(c)
}
