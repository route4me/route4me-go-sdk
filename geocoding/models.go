package geocoding

type Address struct {
	ZipCode string `json:"zipcode,omitempty"`
	Name    string `json:"street_name,omitempty"`
}

type Coordinates struct {
	Lat float64 `json:"lat,omitempty"`
	Lng float64 `json:"lng,omitempty"`
}

type Geocoding struct {
	BBox          []Coordinates `json:"bbox"`
	Confidence    string        `json:"confidence,omitempty"`
	CountryRegion string        `json:"countryRegion,omitempty"`
	Coordinates   []Coordinates `json:"curbside_coordinates,omitempty"`
	Key           string        `json:"key,omitempty"`
	Lat           float64       `json:"lat"`
	Lng           float64       `json:"lng"`
	Name          string        `json:"name,omitempty"`
	PostalCode    string        `json:"postalCode,omitempty"`
	Type          string        `json:"type,omitempty"`
}
