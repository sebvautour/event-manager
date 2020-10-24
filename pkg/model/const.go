package model

const (
	// SeverityCritical defines a critical severity alert or event
	SeverityCritical = "critical"
	// SeverityMajor defines a major severity alert or event
	SeverityMajor = "major"
	// SeverityMinor defines a minor severity alert or event
	SeverityMinor = "minor"
	// SeverityWarning defines a warning severity alert or event
	SeverityWarning = "warning"
	// SeverityInfo defines a informational severity alert or event
	SeverityInfo = "info"

	// SeverityDefault is used when the severity is unknown
	SeverityDefault = SeverityMinor
)

// Severities provides an array of severities, orders from most to least severe
var Severities = [...]string{SeverityCritical, SeverityMajor, SeverityMinor, SeverityWarning, SeverityInfo}

// SeverityIndex returns the index of the given severity based on the Severites const, or -1 if not found
func SeverityIndex(severity string) int {
	for i, s := range Severities {
		if s == severity {
			return i
		}
	}
	return -1
}

const (
	// StatusActive defines an alert or event that is still active / unresolved
	StatusActive = "active"
	// StatusResolved defines an alert that is no longer active, but could return active again if a new event arrives
	StatusResolved = "resolved"
	// StatusSilenced defines an alert that was silenced by a process or operator
	StatusSilenced = "silenced"
	// StatusClosed defines an alert that is no longer active and will not be active again
	StatusClosed = "closed"

	// StatusDefault is used when the status is unknown
	StatusDefault = StatusActive
)
