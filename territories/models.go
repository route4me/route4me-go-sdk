package territories

import (
	"encoding/json"

	"github.com/route4me/route4me-go-sdk/routing"
)

type AvoidanceZone struct {
	//AvoidanceZone ID
	ID       string      `json:"territory_id"`
	Name     string      `json:"territory_name"`
	Color    string      `json:"territory_color"`
	MemberID json.Number `json:"member_id"`
	//Territory parameters
	Territory TerritoryShape `json:"territory"`
}

type Territory struct {
	//Territory ID
	ID       string      `json:"territory_id"`
	Name     string      `json:"territory_name"`
	Color    string      `json:"territory_color"`
	MemberID json.Number `json:"member_id"`
	//Territory parameters
	Territory TerritoryShape    `json:"territory"`
	Addresses []routing.Address `json:"addresses"`
}

type Query struct {
	ID       string `http:"territory_id"`
	DeviceID string `http:"device_id"`
}

type Type string

const (
	Circle    Type = "circle"
	Polygonal Type = "poly"
	Rectangle Type = "rect"
)

type TerritoryShape struct {
	Type Type     `json:"type"`
	Data []string `json:"data"`
}
