package data

type AvoidanceZone struct {
	//AvoidanceZone ID
	ID       string `json:"territory_id"`
	Name     string `json:"territory_name"`
	Color    string `json:"territory_color"`
	MemberID string `json:"member_id"`
	//Territory parameters
	Territory Territory `json:"territory"`
}
