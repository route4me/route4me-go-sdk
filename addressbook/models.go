package addressbook

type Contact struct {
	ID          uint64  `json:"address_id,omitempty"`
	Group       string  `json:"address_group,omitempty"`
	Alias       string  `json:"address_alias,omitempty"`
	Address1    string  `json:"address_1"`
	Address2    string  `json:"address_2,omitempty"`
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
}

type Query struct {
	//Comma separated list of ids
	AddressID string `http:"address_id"`
	Limit     uint   `http:"limit"`
	Offset    uint   `http:"offset"`
	Start     uint   `http:"start"`
	Query     string `http:"query"`
	Fields    string `http:"fields"`
	Display   string `http:"display"`
}
