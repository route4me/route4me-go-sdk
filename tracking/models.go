package tracking

import "github.com/route4me/route4me-go-sdk/routing"

type TimePeriod string

const (
	TimePeriodToday     TimePeriod = "today"
	TimePeriodYesterday TimePeriod = "yesterday"
	TimePeriodThisMonth TimePeriod = "thismonth"
	TimePeriod7Days     TimePeriod = "7days"
	TimePeriod14Days    TimePeriod = "14days"
	TimePeriod30Days    TimePeriod = "30days"
	TimePeriod60Days    TimePeriod = "60days"
	TimePeriod90Days    TimePeriod = "90days"
	TimePeriodAllTime   TimePeriod = "all_time"
	TimePeriodCustom    TimePeriod = "custom"
)

type TrackingHistoryQuery struct {
	EndDate      uint64     `http:"end_date,omitempty"`
	LastPosition bool       `http:"last_position,omitempty"`
	RouteID      string     `http:"route_id"`
	StartDate    uint64     `http:"start_date,omitempty"`
	TimePeriod   TimePeriod `http:"time_period,omitempty"`
}

type GPS struct {
	Format          string             `http:"format"`
	MemberID        int                `http:"member_id"`
	RouteID         string             `http:"route_id"`
	TxID            string             `http:"tx_id"`
	VehicleID       int                `http:"vehicle_id"`
	Course          int                `http:"course"`
	Speed           float64            `http:"speed"`
	Latitude        float64            `http:"lat"`
	Longitude       float64            `http:"lng"`
	Altitude        float64            `http:"altitude"`
	DeviceType      routing.DeviceType `http:"device_type"`
	DeviceGUID      string             `http:"device_guid"`
	DeviceTimestamp string             `http:"device_timestamp"`
	AppVersion      string             `http:"app_version"`
}
