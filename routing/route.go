package routing

//OptimizationState describes an optimization problem can be at one state at any given time
//every state change invokes a socket notification to the associated member id
//every state change invokes a callback webhook event invocation if it was provided during the initial optimization
type OptimizationState uint

const (
	Initial OptimizationState = iota + 1
	MatrixProcessing
	Optimizing
	Optimized
	Error
	ComputingDirections
)

type TrackingHistory struct {
	Speed     float64 `json:"s,omitempty"`
	Latitude  float64 `json:"lt,string,omitempty"`
	Longitude float64 `json:"lg,string,omitempty"`
	Direction int     `json:"d,omitempty"`
	//TODO: Check format
	Timestamp         uint64 `json:"ts,string,omitempty"`
	FriendlyTimestamp string `json:"ts_friendly,omitempty"`
}

type DataObject struct {
	ProblemID        string            `json:"optimization_problem_id,omitempty"`
	State            OptimizationState `json:"state,omitempty"`
	UserErrors       []string          `json:"user_errors,omitempty"`
	SentToBackground bool              `json:"sent_to_background,omitempty"`
	Parameters       RouteParameters   `json:"parameters,omitempty"`
	Addresses        []Address         `json:"addresses,omitempty"`
	Routes           []Route           `json:"routes,omitempty"`
	Links            Links             `json:"links,omitempty"`
	TrackingHistory  []TrackingHistory `json:"tracking_history,omitempty"`
}

type Route struct {
	DataObject
	ID                        string  `json:"route_id,omitempty" http:"route_id"`
	MemberID                  uint64  `json:"member_id,omitempty" http:"member_id"`
	MemberEmail               string  `json:"member_email,omitempty"`
	VehicleAlias              string  `json:"vehicle_alias,omitempty"`
	DriverAlias               string  `json:"driver_alias,omitempty"`
	RouteCost                 float64 `json:"route_cost,omitempty"`
	RouteRevenue              float64 `json:"route_revenue,omitempty"`
	NetRevenuePerDistanceUnit float64 `json:"net_revenue_per_distance_unit,omitempty"`
	CreatedTimestamp          uint64  `json:"created_timestamp,omitempty"`
	Mpg                       float64 `json:"mpg,omitempty"`
	TripDistance              float64 `json:"trip_distance,omitempty"`
	GasPrice                  float64 `json:"gas_price,omitempty"`
	RouteDurationSec          int     `json:"route_distance_sec,omitempty"`
}

type Links struct {
	Route                 string `json:",omitempty"`
	View                  string `json:"view,omitempty"`
	OptimizationProblemID string `json:"optimization_problem_id,omitempty"`
}

type Optimize string

const (
	Distance        Optimize = "Distance"
	Time            Optimize = "Time"
	TimeWithTraffic Optimize = "TimeWithTraffic"
)

type TravelMode string

const (
	Driving  TravelMode = "Driving"
	Walking  TravelMode = "Walking"
	Trucking TravelMode = "Trucking"
)

type PathOutput string

const (
	None   PathOutput = "None"
	Points PathOutput = "Points"
)

type RouteQuery struct {
	// Route Identifier
	ID string `http:"route_id"`
	// Pass True to return directions and the route path
	Directions bool `http:"directions"`
	// "None" - no path output. "Points" - points path output
	PathOutput PathOutput `http:"route_path_output"`
	//
	TravelMode TravelMode `http:"travel_mode"`
	// Output route tracking data in response
	DeviceTrackingHistory bool `http:"device_tracking_history"`
	// The number of existing routes that should be returned per response when looking at a list of all the routes.
	Limit uint `http:"limit"`
	// The page number for route listing pagination. Increment the offset by the limit number to move to the next page.
	Offset uint `http:"offset"`
	// Output addresses and directions in the original optimization request sequence. This is to allow us to compare routes before & after optimization.
	Original bool `http:"original"`
	// Output route and stop-specific notes. The notes will have timestamps, note types, and geospatial information if available
	Notes bool `http:"notes"`
	// Search query
	Query string `http:"query"`
	// Updating a route supports the reoptimize=1 parameter, which reoptimizes only that route. Also supports the parameters from GET.
	Reoptimize bool `http:"reoptimize"`
	// By sending recompute_directions=1 we request that the route directions be recomputed (note that this does happen automatically if certain properties of the route are updated, such as stop sequence_no changes or round-tripness)
	RecomputeDirections bool `http:"recompute_directions"`
	// Route Parameters to update.
	// (After a PUT there is no guarantee that the route_destination_id values are preserved! It may create copies resulting in new destination IDs, especially when dealing with multiple depots.)
	Parameters RouteParameters `json:"parameters,omitempty"`
}
