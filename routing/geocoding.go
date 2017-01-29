package routing

type CurbsideCoordinates struct {
	Lat float64 `json:"lat,omitempty"`
	Lng float64 `json:"lng,omitempty"`
}

type Geocoding struct {
	BBox                []CurbsideCoordinates `json:"bbox"`
	Confidence          string                `json:"confidence,omitempty"`
	CountryRegion       string                `json:"countryRegion,omitempty"`
	CurbsideCoordinates []CurbsideCoordinates `json:"curbside_coordinates,omitempty"`
	Key                 string                `json:"key,omitempty"`
	Lat                 float64               `json:"lat"`
	Lng                 float64               `json:"lng"`
	Name                string                `json:"name,omitempty"`
	PostalCode          string                `json:"postalCode,omitempty"`
	Type                string                `json:"type,omitempty"`
}
