package zones

type AvoidanceZone struct {
	//AvoidanceZone ID
	ID       string `json:"territory_id" http:"territory_id"`
	Name     string `json:"territory_name" http:"territory_name"`
	Color    string `json:"territory_color" http:"territory_color"`
	MemberID string `json:"member_id" http:"member_id"`
	//Territory parameters
	Territory Territory `json:"territory"`
}

type Query struct {
	AvoidanceZone
	DeviceID string `http:"device_id"`
}

type Territory struct {
	Type string   `json:"type"`
	Data []string `json:"data"`
}
