package model

import (
	"errors"
	"time"
)

// AlertGroups holds an array of AlertGroup with methods to interact with it
type AlertGroups struct {
	Groups []AlertGroup `json:"groups"`
}

// AlertGroup holds a group of alerts with the same group key
type AlertGroup struct {
	GroupKey       string    `json:"group_key"`
	AlertCount     int       `json:"alert_count"`
	LastEventTime  time.Time `json:"last_event_time"`
	FirstEventTime time.Time `json:"first_event_time"`
	PrimaryAlert   Alert     `json:"primary_alert"`
	Alerts         []Alert   `json:"alerts"`
}

// NewAlertGroups creates a new AlertGroups for the given array of Alert
func NewAlertGroups(alerts []Alert) (ags *AlertGroups) {
	ags = &AlertGroups{}
	for _, a := range alerts {
		ags.AddAlert(a)
	}
	// err is warning
	_ = ags.UpdateAllGroupInfo(AlertGroupPrimaryMethodSeverity)

	return ags
}

// AddAlert adds a new alert to an existing AlertGroup, or creates a new one
func (ags *AlertGroups) AddAlert(alert Alert) {
	i := ags.groupIndex(alert.GroupKey)

	if i == -1 {
		ags.Groups = append(ags.Groups, AlertGroup{
			GroupKey:       alert.GroupKey,
			AlertCount:     1,
			LastEventTime:  alert.FirstEventTime,
			FirstEventTime: alert.LastEventTime,
			PrimaryAlert:   alert,
			Alerts:         []Alert{alert},
		})
		return
	}

	ags.Groups[i].Alerts = append(ags.Groups[i].Alerts, alert)

}

func (ags *AlertGroups) groupIndex(groupKey string) int {
	for i, a := range ags.Groups {
		if a.GroupKey == groupKey {
			return i
		}
	}
	return -1
}

const (
	// AlertGroupPrimaryMethodSeverity is used for the primaryMethod param, it will make the primary alert be picked based on the most severe severity
	AlertGroupPrimaryMethodSeverity = "severity"
	// AlertGroupPrimaryMethodLastEvent is used for the primaryMethod param, it will make the primary alert be picked based on the last event time
	AlertGroupPrimaryMethodLastEvent = "last_event"
	// AlertGroupPrimaryMethodFirst is used for the primaryMethod param, it will pick the first alert from the alertgroup as the primary alert
	AlertGroupPrimaryMethodFirst = "first"
)

// ErrUnknownPrimaryMethod is sent when the primaryMethod param is an unknown AlertGroupPrimaryMethod
var ErrUnknownPrimaryMethod = errors.New("Unknown alert group primary method")

// UpdateGroupInfo updates the AlertGroup info and primary alert
// primaryMethod is passed to determine how the primary alert is selected
// an error means the primaryMethod given is invalid, this can be treated as a warning since it will default to AlertGroupPrimaryMethodSeverity
func (ag *AlertGroup) UpdateGroupInfo(primaryMethod string) error {
	// base
	if len(ag.Alerts) == 0 {
		return errors.New("No alerts in group")
	}

	ag.GroupKey = ag.Alerts[0].GroupKey
	ag.AlertCount = len(ag.Alerts)
	ag.PrimaryAlert = ag.Alerts[0]

	ag.LastEventTime = ag.Alerts[0].LastEventTime
	ag.FirstEventTime = ag.Alerts[0].FirstEventTime
	for _, alert := range ag.Alerts {
		if !alert.FirstEventTime.After(ag.FirstEventTime) {
			ag.FirstEventTime = alert.FirstEventTime
		}

		if alert.LastEventTime.After(ag.LastEventTime) {
			ag.LastEventTime = alert.LastEventTime
		}
	}

	switch primaryMethod {
	case AlertGroupPrimaryMethodFirst:
		// the base part covers this
	case AlertGroupPrimaryMethodSeverity:
		ag.updateGroupInfoSev()
	case AlertGroupPrimaryMethodLastEvent:
		ag.updateGroupInfoLastEvent()
	default:
		// default
		ag.updateGroupInfoSev()
		return ErrUnknownPrimaryMethod
	}
	return nil

}

func (ag *AlertGroup) updateGroupInfoSev() {
	for _, alert := range ag.Alerts {
		if SeverityIndex(alert.Severity) < SeverityIndex(ag.PrimaryAlert.Severity) {
			ag.PrimaryAlert = alert
		}
	}
}

func (ag *AlertGroup) updateGroupInfoLastEvent() {
	for _, alert := range ag.Alerts {
		if alert.LastEventTime.After(ag.PrimaryAlert.LastEventTime) {
			ag.PrimaryAlert = alert
		}
	}
}

// UpdateAllGroupInfo runs UpdateGroupInfo for all AlertGroup items in the AlertGroups
func (ags *AlertGroups) UpdateAllGroupInfo(primaryMethod string) error {
	for i := range ags.Groups {
		if err := ags.Groups[i].UpdateGroupInfo(primaryMethod); err != nil {
			return err
		}
	}
	return nil
}
