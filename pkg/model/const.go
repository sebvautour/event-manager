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
