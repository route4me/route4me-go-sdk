package query

type GPS struct {
	Format          string  `http:"format"`
	MemberID        int     `http:"member_id"`
	RouteID         string  `http:"route_id"`
	TxID            string  `http:"tx_id"`
	VehicleID       int     `http:"vehicle_id"`
	Course          int     `http:"course"`
	Speed           float64 `http:"speed"`
	Latitude        float64 `http:"lat"`
	Longitude       float64 `http:"lng"`
	Altitude        float64 `http:"altitude,omitempty"`
	DeviceType      string  `http:"device_type"`
	DeviceGUID      string  `http:"device_guid"`
	DeviceTimestamp string  `http:"device_timestamp,omitempty"`
	AppVersion      string  `http:"app_version,omitempty"`
}
