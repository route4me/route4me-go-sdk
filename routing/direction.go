package routing

import "github.com/route4me/route4me-go-sdk/geocoding"

type Direction struct {
	Location *Location `json:"location,omitempty"`
	Steps    []Steps   `json:"steps,omitempty"`
}

type Location struct {
	DirectionsError string  `json:"directions_error,omitempty"`
	EndLocation     string  `json:"end_location,omitempty"`
	ErrorCode       int     `json:"error_code,omitempty"`
	Name            string  `json:"name,omitempty"`
	SegmentDistance float64 `json:"segment_distance,omitempty"`
	StartLocation   string  `json:"start_location,omitempty"`
	Time            int     `json:"time,omitempty"`
}

type Steps struct {
	CompassDirection string                 `json:"compass_direction,omitempty"`
	Direction        string                 `json:"direction,omitempty"`
	Directions       string                 `json:"directions,omitempty"`
	Distance         float64                `json:"distance,omitempty"`
	DistanceUnit     string                 `json:"distance_unit,omitempty"`
	DurationSec      int                    `json:"duration_sec,omitempty"`
	ManeuverPoint    *geocoding.Coordinates `json:"maneuverPoint,omitempty"`
	ManeuverType     string                 `json:"maneuverType,omitempty"`
}
