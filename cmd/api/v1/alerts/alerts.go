package alerts

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebvautour/event-manager/cmd/api/v1/helpers"
	"github.com/sebvautour/event-manager/internal/db"
	"github.com/sebvautour/event-manager/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InitRouteGroup(r *gin.RouterGroup) {
	r.GET("", getAlertsHandler)
	r.POST("", getAlertsHandler)
	r.GET("/:id", getAlertIDHandler)
	r.GET("/:id/events", getAlertIDEventsHandler)
}

// @Summary Get alerts
// @Description Get alerts based on a given query
// @Tags alerts
// @Accept  json
// @Produce  json
// @Param filter query string false "filter as MongoDB JSON schema"
// @Param filter body string false "filter as MongoDB JSON schema"
// @Param encoding query string false "filter query param will be base64 decoded if encoding param value is base64"
// @Param grouped query string false "returns list of alerts grouped by their group_id if true"
// @Router /alerts [get]
func getAlertsHandler(c *gin.Context) {
	filter, err := queryFilter(c)
	if err != nil {
		helpers.Error(c, "failed to retrieve filter", http.StatusBadRequest, err.Error())
		return
	}

	alerts, err := db.AlertSearch(c.Request.Context(), filter)
	if err != nil {
		helpers.Error(c, "failed to query alerts", http.StatusInternalServerError, err.Error())
		return
	}

	if c.Query("grouped") != "" && c.Query("grouped") != "false" {
		c.JSON(200, model.NewAlertGroups(alerts))
		return
	}
	c.JSON(200, alerts)

}

func queryFilter(c *gin.Context) (filter bson.D, err error) {
	queryFilter := c.Query("filter")
	var queryFilterBytes []byte
	if queryFilter != "" {
		if c.Query("encoding") == "base64" {
			queryFilterBytes, err = base64.RawStdEncoding.DecodeString(queryFilter)
			if err != nil {
				return filter, fmt.Errorf("filter query param not base64 encoded: %w", err)

			}
		} else {
			queryFilterBytes = []byte(queryFilter)
		}
	} else {
		queryFilterBytes, err = c.GetRawData()
		if err != nil {
			return filter, fmt.Errorf("filter needs to be provided in base64 encoded as a query param, or as the request body: %w", err)
		}
	}

	if queryFilterBytes == nil {
		return filter, errors.New("filter needs to be provided in base64 encoded as a query param, or as the request body")
	}
	if string(queryFilterBytes) == "" {
		return filter, errors.New("filter needs to be provided in base64 encoded as a query param, or as the request body")
	}

	if err := bson.UnmarshalExtJSON(queryFilterBytes, true, &filter); err != nil {
		return filter, fmt.Errorf("failed to parse filter as JSON: %w", err)
	}
	return filter, nil
}

// @Summary Get alert
// @Description get alert by ID
// @Tags alerts
// @Produce  json
// @Param id path string true "Alert ID"
// @Router /alert/{id} [get]
func getAlertIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		helpers.Error(c, "id not provided", http.StatusBadRequest)
		return
	}

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		helpers.Error(c, "invalid id provided", http.StatusBadRequest, err.Error())
		return
	}

	alert, found, err := db.AlertByID(c.Request.Context(), id)
	if err != nil {
		helpers.Error(c, "failed to retrieve alert", http.StatusInternalServerError, err.Error())
		return
	}

	if !found {
		helpers.Error(c, "alert not found with given id", http.StatusBadRequest)
		return
	}

	c.JSON(200, alert)
}

// @Summary Get alert events
// @Description get array of events for a given alert ID
// @Tags alerts
// @Produce  json
// @Param id path string true "Alert ID"
// @Router /alert/{id}/events [get]
func getAlertIDEventsHandler(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		helpers.Error(c, "id not provided", http.StatusBadRequest)
		return
	}

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		helpers.Error(c, "invalid id provided", http.StatusBadRequest, err.Error())
		return
	}

	events, err := db.EventsForAlert(c.Request.Context(), id)
	if err != nil {
		helpers.Error(c, "failed to query for events", http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, events)
}
