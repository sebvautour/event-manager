package alerts

import (
	"encoding/base64"
	"log"
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
// @Router /alerts [get]
func getAlertsHandler(c *gin.Context) {
	queryFilter := c.Query("filter")
	var queryFilterBytes []byte
	var err error
	if queryFilter != "" {
		if c.Query("encoding") == "base64" {
			queryFilterBytes, err = base64.RawStdEncoding.DecodeString(queryFilter)
			if err != nil {
				helpers.Error(c, "filter query param not base64 encoded", http.StatusBadRequest, err.Error())
				return
			}
		} else {
			queryFilterBytes = []byte(queryFilter)
		}
	} else {
		queryFilterBytes, err = c.GetRawData()
		if err != nil {
			helpers.Error(c, "filter needs to be provided in base64 encoded as a query param, or as the request body", http.StatusBadRequest, err.Error())
			return
		}
	}

	if queryFilterBytes == nil {
		helpers.Error(c, "filter needs to be provided in base64 encoded as a query param, or as the request body", http.StatusBadRequest)
		return
	}
	if string(queryFilterBytes) == "" {
		helpers.Error(c, "filter needs to be provided in base64 encoded as a query param, or as the request body", http.StatusBadRequest)
		return
	}

	var filterBSON bson.D

	if err := bson.UnmarshalExtJSON(queryFilterBytes, true, &filterBSON); err != nil {
		helpers.Error(c, "failed to parse filter as JSON", http.StatusBadRequest, err.Error())
		return
	}

	cursor, err := db.AlertsMongo.Find(c.Request.Context(), filterBSON)
	if err != nil {
		helpers.Error(c, "failed to query alerts", http.StatusInternalServerError, err.Error())
		return
	}
	var alerts []bson.M
	if err = cursor.All(c.Request.Context(), &alerts); err != nil {
		helpers.Error(c, "failed to retrieve all alerts", http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, alerts)
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

	alert := &model.Alert{}
	alert.ID = id

	err = db.Alerts.Actions().Get(alert).Do(c.Request.Context())
	if err != nil {
		helpers.Error(c, "failed to retrieve alert", http.StatusInternalServerError, err.Error())
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

	log.Println(id.Hex())

	cursor, err := db.EventsMongo.Find(c.Request.Context(), bson.M{"AlertID": id})
	if err != nil {
		helpers.Error(c, "failed to query events", http.StatusInternalServerError, err.Error())
		return
	}
	var events []bson.M
	if err = cursor.All(c.Request.Context(), &events); err != nil {
		helpers.Error(c, "failed to retrieve all events", http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, events)
}
