package routing

import (
	"encoding/json"

	"github.com/route4me/route4me-go-sdk/geocoding"
)

type AddressQuery struct {
	RouteID            string `http:"route_id"`
	RouteDestinationID string `http:"route_destination_id"`
	Notes              bool   `http:"notes"`
}

type Address struct {
	AbnormalTrafficTimeToNextDestination int                   `json:"abnormal_traffic_time_to_next_destination,omitempty"`
	AccountNo                            string                `json:"account_no,omitempty"`
	AddressString                        string                `json:"address,omitempty"`
	AddressStopType                      string                `json:"address_stop_type,omitempty"`
	Alias                                string                `json:"alias,omitempty"`
	ChannelName                          string                `json:"channel_name,omitempty"`
	ContactID                            int                   `json:"contact_id,omitempty"`
	Cost                                 float64               `json:"cost,omitempty"`
	Cube                                 float64               `json:"cube,omitempty"`
	CurbsideLat                          float64               `json:"curbside_lat,omitempty"`
	CurbsideLng                          float64               `json:"curbside_lng,omitempty"`
	CustomFields                         map[string]string     `json:"custom_fields,omitempty"`
	CustomFieldsConfig                   []string              `json:"custom_fields_config,omitempty"`
	CustomFieldsConfigStrJSON            string                `json:"custom_fields_config_str_json,omitempty"`
	CustomFieldsStrJSON                  string                `json:"custom_fields_str_json,omitempty"`
	CustomerPo                           string                `json:"customer_po,omitempty"`
	DepartedLat                          float64               `json:"departed_lat,omitempty"`
	DepartedLng                          float64               `json:"departed_lng,omitempty"`
	DestinationNoteCount                 int                   `json:"destination_note_count,omitempty"`
	Directions                           []Direction           `json:"directions,omitempty"`
	DistanceToNextDestination            float64               `json:"distance_to_next_destination,omitempty"`
	DriveTimeToNextDestination           int                   `json:"drive_time_to_next_destination,omitempty"`
	Email                                string                `json:"email,omitempty"`
	FailedGeocoding                      bool                  `json:"failed_geocoding,omitempty"`
	FirstName                            string                `json:"first_name,omitempty"`
	GeneratedTimeWindowEnd               int                   `json:"generated_time_window_end,omitempty"`
	GeneratedTimeWindowStart             int                   `json:"generated_time_window_start,omitempty"`
	Geocoded                             bool                  `json:"geocoded,omitempty"`
	Geocodings                           []geocoding.Geocoding `json:"geocodings,omitempty"`
	GeofenceDetectedDepartedLat          float64               `json:"geofence_detected_departed_lat,omitempty"`
	GeofenceDetectedDepartedLng          float64               `json:"geofence_detected_departed_lng,omitempty"`
	GeofenceDetectedDepartedTimestamp    int                   `json:"geofence_detected_departed_timestamp,omitempty"`
	GeofenceDetectedServiceTime          int                   `json:"geofence_detected_service_time,omitempty"`
	GeofenceDetectedVisitedLat           float64               `json:"geofence_detected_visited_lat,omitempty"`
	GeofenceDetectedVisitedLng           float64               `json:"geofence_detected_visited_lng,omitempty"`
	GeofenceDetectedVisitedTimestamp     int                   `json:"geofence_detected_visited_timestamp,omitempty"`
	Group                                string                `json:"group,omitempty"`
	InvoiceNo                            string                `json:"invoice_no,omitempty"`
	IsDeparted                           bool                  `json:"is_departed,omitempty"`
	IsDepot                              bool                  `json:"is_depot,omitempty"`
	IsVisited                            bool                  `json:"is_visited,omitempty"`
	LastName                             string                `json:"last_name,omitempty"`
	Latitude                             float64               `json:"lat,omitempty"`
	Longitude                            float64               `json:"lng,omitempty"`
	Manifest                             *Manifest             `json:"manifest,omitempty"`
	MemberID                             int                   `json:"member_id,omitempty"`
	Notes                                []Note                `json:"notes,omitempty"`
	OptimizationProblemID                string                `json:"optimization_problem_id,omitempty"`
	OrderID                              int                   `json:"order_id,omitempty"`
	OrderNo                              string                `json:"order_no,omitempty"`
	OriginalRouteID                      string                `json:"original_route_id,omitempty"`
	PathToNext                           []PathToNext          `json:"path_to_next,omitempty"`
	Phone                                string                `json:"phone,omitempty"`
	Pieces                               int                   `json:"pieces,omitempty"`
	PreferredGeocoding                   int                   `json:"preferred_geocoding,omitempty"`
	Priority                             int                   `json:"priority,omitempty"`
	ReferenceNo                          string                `json:"reference_no,omitempty"`
	Revenue                              float64               `json:"revenue,omitempty"`
	RouteDestinationID                   json.Number           `json:"route_destination_id,omitempty" http:"route_destination_id"`
	RouteID                              string                `json:"route_id,omitempty" http:"route_id"`
	RouteName                            string                `json:"route_name,omitempty"`
	SequenceNo                           int                   `json:"sequence_no,omitempty"`
	Time                                 int                   `json:"time,omitempty"`
	TimeWindowEnd                        int                   `json:"time_window_end,omitempty"`
	TimeWindowEnd2                       int                   `json:"time_window_end_2,omitempty"`
	TimeWindowStart                      int                   `json:"time_window_start,omitempty"`
	TimeWindowStart2                     int                   `json:"time_window_start_2,omitempty"`
	TimeframeViolationRate               float64               `json:"timeframe_violation_rate,omitempty"`
	TimeframeViolationState              int                   `json:"timeframe_violation_state,omitempty"`
	TimeframeViolationTime               int                   `json:"timeframe_violation_time,omitempty"`
	TimestampLastDeparted                int                   `json:"timestamp_last_departed,omitempty"`
	TimestampLastVisited                 int                   `json:"timestamp_last_visited,omitempty"`
	TrackingNumber                       string                `json:"tracking_number,omitempty"`
	TrafficTimeToNextDestination         int                   `json:"traffic_time_to_next_destination,omitempty"`
	UncongestedTimeToNextDestination     int                   `json:"uncongested_time_to_next_destination,omitempty"`
	VisitedLat                           float64               `json:"visited_lat,omitempty"`
	VisitedLng                           float64               `json:"visited_lng,omitempty"`
	WaitTimeToNextDestination            int                   `json:"wait_time_to_next_destination,omitempty"`
	Weight                               float64               `json:"weight,omitempty"`
}

type Manifest struct {
	ActualArrivalTimeTs      int     `json:"actual_arrival_time_ts,omitempty"`
	ActualDepartureTimeTs    int     `json:"actual_departure_time_ts,omitempty"`
	EstimatedArrivalTimeTs   int     `json:"estimated_arrival_time_ts,omitempty"`
	EstimatedDepartureTimeTs int     `json:"estimated_departure_time_ts,omitempty"`
	FuelCostFromStart        float64 `json:"fuel_cost_from_start,omitempty"`
	FuelFromStart            float64 `json:"fuel_from_start,omitempty"`
	ProjectedArrivalTimeTs   int     `json:"projected_arrival_time_ts,omitempty"`
	ProjectedDepartureTimeTs int     `json:"projected_departure_time_ts,omitempty"`
	RunningDistance          float64 `json:"running_distance,omitempty"`
	RunningServiceTime       int     `json:"running_service_time,omitempty"`
	RunningTravelTime        int     `json:"running_travel_time,omitempty"`
	RunningWaitTime          int     `json:"running_wait_time,omitempty"`
	TimeImpact               int     `json:"time_impact,omitempty"`
}

type PathToNext struct {
	Lat float64 `json:"lat,omitempty"`
	Lng float64 `json:"lng,omitempty"`
}
