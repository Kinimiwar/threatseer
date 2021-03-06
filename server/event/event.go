package event

import (
	api "github.com/capsule8/capsule8/api/v0"
)

// Event contains information about the source,
// the original event, and any indicators gathered
// up to the point in time
type Event struct {
	Event      *api.TelemetryEvent
	Indicators []Indicator
	ClientAddr string
}

// Indicator is an individual result from an engine
type Indicator struct {
	Engine        string
	RuleName      string
	IndicatorType string
	Description   string
	ExtraInfo     string
	Score         int
}
