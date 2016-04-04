package data

type Activity struct {
	ID        string `json:"activity_id"`
	Type      string `json:"activity_type"`
	Timestamp uint64 `json:"activity_timestamp,omitempty"`
}
