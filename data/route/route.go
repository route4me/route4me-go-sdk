package route

type Route struct {
	DataObject
	ID                        string  `json:"route_id,omitempty"`
	MemberID                  string  `json:"member_id,omitempty"`
	MemberEmail               string  `json:"member_email,omitempty"`
	VehicleAlias              string  `json:"vehicle_alias,omitempty"`
	DriverAlias               string  `json:"driver_alias,omitempty"`
	RouteCost                 float64 `json:"route_cost,omitempty"`
	RouteRevenue              float64 `json:"route_revenue,omitempty"`
	NetRevenuePerDistanceUnit float64 `json:"net_revenue_per_distance_unit,omitempty"`
	CreatedTimestamp          uint64  `json:"created_timestamp,omitempty"`
	Mpg                       string  `json:"mpg,omitempty"`
	TripDistance              float64 `json:"trip_distance,omitempty"`
	GasPrice                  float64 `json:"gas_price,omitempty"`
	RouteDurationSec          int     `json:"route_distance_sec,omitempty"`
}
