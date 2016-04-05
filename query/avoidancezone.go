package query

import "github.com/route4me/route4me-go-sdk/data"

type AvoidanceZoneBase struct {
	DeviceID    string `http:"device_id,omitempty"`
	TerritoryID string `http:"territory_id,omitempty" json:"territory_id,omitempty"`
}

type AvoidanceZone struct {
	AvoidanceZoneBase
	TerritoryName  string         `json:"territory_name"`
	TerritoryColor string         `json:"territory_color"`
	MemberID       string         `json:"member_id"`
	Territory      data.Territory `json:"territory"`
}
