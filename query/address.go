package query

type Address struct {
	RouteID            string `http:"route_id,omitempty"`
	RouteDestinationID int    `http:"route_destination_id,omitempty"`
	Notes              bool   `http:"notes"`
}
