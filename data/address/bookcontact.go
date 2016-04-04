package address

type BookContact struct {
	ID          string  `json:"address_id,omitempty"`
	Group       string  `json:"address_group,omitempty"`
	Alias       string  `json:"address_alias,omitempty"`
	Address1    string  `json:"address_1"`
	Address2    string  `json:"address_2"`
	FirstName   string  `json:"first_name,omitempty"`
	LastName    string  `json:"last_name,omitempty"`
	Email       string  `json:"address_email,omitempty"`
	PhoneNumber string  `json:"address_phone_number,omitempty"`
	City        string  `json:"address_city,omitempty"`
	StateID     string  `json:"address_state_id,omitempty"`
	CountryID   string  `json:"address_country_id,omitempty"`
	ZIP         string  `json:"address_zip,omitempty"`
	CachedLat   float64 `json:"cached_lat"`
	CachedLng   float64 `json:"cached_lng"`
	Color       string  `json:"color"`
}
