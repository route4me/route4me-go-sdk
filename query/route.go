package query

import "github.com/route4me/route4me-go-sdk/data/route"

type Route struct {
	// Route Identifier
	ID string `http:"route_id,omitempty"`
	// Pass True to return directions and the route path
	Directions bool `http:"directions,omitempty"`
	// "None" - no path output. "Points" - points path output
	RoutePathOutput string `http:"route_path_output,omitempty"`
	// Output route tracking data in response
	DeviceTrackingHistory bool `http:"device_tracking_history,omitempty"`
	// The number of existing routes that should be returned per response when looking at a list of all the routes.
	Limit uint `http:"limit,omitempty"`
	// The page number for route listing pagination. Increment the offset by the limit number to move to the next page.
	Offset uint `http:"offset,omitempty"`
	// Output addresses and directions in the original optimization request sequence. This is to allow us to compare routes before & after optimization.
	Original bool `http:"original,omitempty"`
	// Output route and stop-specific notes. The notes will have timestamps, note types, and geospatial information if available
	Notes bool `http:"notes,omitempty"`
	// Search query
	Query string `http:"query,omitempty"`
	// Updating a route supports the reoptimize=1 parameter, which reoptimizes only that route. Also supports the parameters from GET.
	Reoptimize bool `http:"reoptimize,omitempty"`
	// By sending recompute_directions=1 we request that the route directions be recomputed (note that this does happen automatically if certain properties of the route are updated, such as stop sequence_no changes or round-tripness)
	RecomputeDirections bool `http:"recompute_directions,omitempty"`
	// Route Parameters to update.
	// (After a PUT there is no guarantee that the route_destination_id values are preserved! It may create copies resulting in new destination IDs, especially when dealing with multiple depots.)
	Parameters route.Parameters `json:"parameters,omitempty"`
}
