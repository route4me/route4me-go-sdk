package vehicles

import "encoding/json"

type Vehicle struct {
	ID           string      `json:"vehicle_id"`
	CreatedTime  string      `json:"created_time"`
	MemberID     int64       `json:"member_id"`
	Alias        string      `json:"vehicle_alias,omitempty"`
	VIN          string      `json:"vehicle_vin,omitempty"`
	RegState     string      `json:"vehicle_reg_state,omitempty"`
	RegStateID   string      `json:"vehicle_reg_state_id,omitempty"`
	RegCountry   string      `json:"vehicle_reg_country,omitempty"`
	RegCountryID string      `json:"vehicle_reg_country_id,omitempty"`
	LicensePlate string      `json:"vehicle_license_plate,omitempty"`
	Make         string      `json:"vehicle_make,omitempty"`
	ModelYear    string      `json:"vehicle_model_year,omitempty"`
	Model        string      `json:"vehicle_model,omitempty"`
	YearAcquired string      `json:"vehicle_year_acquired,omitempty"`
	AxleCount    json.Number `json:"vehicle_axle_count,omitempty"`
	MPGCity      json.Number `json:"mpg_city,omitempty"`
	MPGHighway   json.Number `json:"mpg_highway,omitempty"`
	FuelType     string      `json:"fuel_type,omitempty"`
	HeightInches json.Number `json:"height_inches,omitempty"`
	WeightLB     json.Number `json:"weight_lb,omitempty"`
}
