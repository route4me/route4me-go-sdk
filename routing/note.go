package routing

type Note struct {
	//NoteID
	ID                 int     `json:"note_it,omitempty"`
	RouteID            string  `json:"route_id,omitempty"`
	RouteDestinationID int     `json:"route_destination_id,omitempty"`
	UploadID           string  `json:"upload_id"`
	TimestampAdded     uint64  `json:"ts_added,omitempty"`
	Latitude           float64 `json:"lat"`
	Longitude          float64 `json:"lng"`
	ActivityType       string  `json:"activity_type"`
	Contents           string  `json:"contents"`
	UploadType         string  `json:"upload_type"`
	UploadURL          string  `json:"upload_url"`
	UploadExtension    string  `json:"upload_extension"`
	DeviceType         string  `json:"device_type"`
}

type NoteQuery struct {
	RouteID      string  `json:"route_id,omitempty"`
	AddressID    int     `json:"address_id,omitempty"`
	Latitude     float64 `json:"dev_lat"`
	Longitude    float64 `json:"dev_lng"`
	DeviceType   string  `json:"device_type"`
	ActivityType string  `json:"strUpdateType"`
}
