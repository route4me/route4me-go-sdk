package query

type Note struct {
	RouteID      string  `json:"route_id,omitempty"`
	AddressID    int     `json:"address_id,omitempty"`
	Latitude     float64 `json:"dev_lat"`
	Longitude    float64 `json:"dev_lng"`
	DeviceType   string  `json:"device_type"`
	ActivityType string  `json:"strUpdateType"`
}
