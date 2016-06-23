package orders

import "encoding/json"

type Display string

const (
	All      Display = "all"
	Routed   Display = "routed"
	Unrouted Display = "unrouted"
)

type Order struct {
	Created           uint64           `json:"created_timestamp,omitempty"`
	ID                uint64           `json:"order_id,omitempty"`
	StatusID          int              `json:"order_status_id,omitempty"`
	DateAdded         string           `json:"day_added_YYMMDD,omitempty"`
	DateScheduled     string           `json:"day_scheduled_for_YYMMDD,omitempty"`
	AddressAlias      string           `json:"address_alias,omitempty"`
	Address1          string           `json:"address_1"`
	Address2          string           `json:"address_2,omitempty"`
	MemberID          int64            `json:"member_id,omitempty"`
	FirstName         string           `json:"EXT_FIELD_first_name,omitempty"`
	LastName          string           `json:"EXT_FIELD_last_name,omitempty"`
	Email             string           `json:"EXT_FIELD_email,omitempty"`
	Phone             string           `json:"EXT_FIELD_phone,omitempty"`
	CustomData        *json.RawMessage `json:"EXT_FIELD_custom_data,omitempty"`
	City              string           `json:"address_city,omitempty"`
	StateID           string           `json:"address_state_id,omitempty"`
	CachedLatitude    float64          `json:"cached_lat"`
	CachedLongitude   float64          `json:"cached_lng"`
	CurbsideLatitude  float64          `json:"curbside_lat,omitempty"`
	CurbsideLongitude float64          `json:"curbside_lng,omitempty"`
	InRouteCount      int              `json:"in_route_count,omitempty"`
	LastVisited       uint64           `json:"last_visited_timestamp,omitempty"`
	LastRouted        uint64           `json:"last_routed_timestamp,omitempty"`

	LocalTimeWindowStart  uint64 `json:"local_time_window_start,omitempty"`
	LocalTimeWindowEnd    uint64 `json:"local_time_window_end,omitempty"`
	LocalTimeWindowStart2 uint64 `json:"local_time_window_start_2,omitempty"`
	LocalTimeWindowEnd2   uint64 `json:"local_time_window_end_2,omitempty"`
	ServiceTime           int    `json:"service_time,omitempty"`
	TimezoneString        string `json:"local_timezone_string,omitempty"`
	Color                 string `json:"color,omitempty"`
	OrderIcon             string `json:"order_icon,omitempty"`

	Validated bool `json:"is_validated,omitempty"`
	Pending   bool `json:"is_pending,omitempty"`
	Accepted  bool `json:"is_accepted,omitempty"`
	Started   bool `json:"is_started,omitempty"`
	Completed bool `json:"is_completed,omitempty"`
}

type Query struct {
	//Used for retrieving by order id / ids
	//Comma separated
	IDs     string  `http:"order_ids"`
	Limit   int     `http:"limit"`
	Offset  int     `http:"offset"`
	Display Display `http:"display"`
	//in format sdate=YYYY-MM-DD
	DateInserted string `http:"day_added_YYMMDD"`
	//in format sdate=YYYY-MM-DD
	DateScheduled string `http:"scheduled_for_YYMMDD"`
	//Main parameter, search orders by their field values
	Query string `http:"query"`
	//Used to get specific fields
	Fields string `http:"fields"`
}
