package query

type Activity struct {
	RouteID  string `http:"route_id,omitempty"`
	DeviceID string `http:"device_id,omitempty"`
	MemberID int    `http:"member_id,omitempty"`
	Limit    uint   `http:"limit,omitempty"`
	Offset   uint   `http:"offset,omitempty"`
	Start    uint   `http:"start,omitempty"`
	End      uint   `http:"end,omitempty"`
}
