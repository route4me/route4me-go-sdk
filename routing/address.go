package routing

type AddressQuery struct {
	RouteID            string `http:"route_id,omitempty"`
	RouteDestinationID int    `http:"route_destination_id,omitempty"`
	Notes              bool   `http:"notes"`
}

type Address struct {
	RouteDestinationID    int     `json:"route_destination_id,omitempty"`
	Alias                 string  `json:"alias"`
	MemberID              int64   `json:"member_id"`
	AddressString         string  `json:"address"`
	IsDepot               bool    `json:"is_depot,omitempty"`
	Latitude              float64 `json:"lat"`
	Longtitude            float64 `json:"lng"`
	RouteID               string  `json:"route_id,omitempty"`
	OriginalRouteID       string  `json:"original_route_id,omitempty"`
	SequenceNo            int     `json:"sequence_no,omitempty"`
	Geocoded              bool    `json:"geocoded,omitempty"`
	PreferredGeocoding    bool    `json:"preferred_geocoding,omitempty"`
	FailedGeocoding       bool    `json:"failed_geocoding,omitempty"`
	ContactID             int     `json:"contact_id,omitempty"`
	IsVisited             bool    `json:"is_visited,omitempty"`
	IsDeparted            bool    `json:"is_departed,omitempty"`
	TimestampLastVisited  uint64  `json:"timestamp_last_visited,omitempty"`
	TimestampLastDeparted uint64  `json:"timestamp_last_departed,omitempty"`

	CustomerPo  string  `json:"customer_po,omitempty"`
	InvoiceNo   string  `json:"invoide_no,omitempty"`
	ReferenceNo string  `json:"reference_no,omitempty"`
	OrderNo     string  `json:"order_no,omitempty"`
	Weight      float64 `json:"weight,omitempty"`
	Cost        float64 `json:"cost,omitempty"`
	Revenue     float64 `json:"revenue,omitempty"`
	Cube        float64 `json:"cube,omitempty"`
	Pieces      int     `json:"pieces,omitempty"`
	Email       string  `json:"email,omitempty"`
	Phone       string  `json:"phone,omitempty"`

	DestinationNoteCount       int     `json:"destination_note_count,omitempty"`
	DriveTimeToNextDestination int     `json:"drive_time_to_next_destination,omitempty"`
	DistanceToNextDestination  float64 `json:"distance_to_next_destination,omitempty"`

	GeneratedTimeWindowStart float64 `json:"generated_time_window_start,omitempty"`
	GeneratedTimeWindowEnd   float64 `json:"generated_time_window_end,omitempty"`

	ChannelName     string `json:"channel_name,omitempty"`
	TimeWindowStart int    `json:"time_window_start,omitempty"`
	TimeWindowEnd   int    `json:"time_window_end,omitempty"`
	Time            int    `json:"time,omitempty"`
	Notes           []Note `json:"notes,omitempty"`

	Priority          int     `json:"priority,omitempty"`
	CurbsideLatitude  float64 `json:"curbside_lat"`
	CurbsideLongitude float64 `json:"curbside_lng"`

	TimeWindowStart2 int `json:"time_window_start_2,omitempty"`
	TimeWindowEnd2   int `json:"time_window_end_2,omitempty"`

	CustomFields map[string]string `json:"custom_fields,omitempty"`
}
