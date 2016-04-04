package data

type TrackingHistory struct {
	Speed     float64 `json:"s"`
	Latitude  float64 `json:"lt"`
	Longitude float64 `json:"lg"`
	Direction string  `json:"d"`
	//TODO: Check format
	Timestamp         uint64 `json:"ts,string"`
	FriendlyTimestamp string `json:"ts_friendly"`
}
