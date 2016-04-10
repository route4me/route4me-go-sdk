package data

type Activity struct {
	ID        string `json:"activity_id"`
	Type      string `json:"activity_type"`
	Timestamp uint64 `json:"activity_timestamp,omitempty"`
	Message   string `json:"activity_message,omitempty"`
	RouteID   string `json:"route_destination_id,omitempty"`
}

type ActivityType string

const (
	DeleteDestination       ActivityType = "delete-destination"
	InsertDestination       ActivityType = "insert-destination"
	DeleteDestination       ActivityType = "delete-destination"
	InsertDestination       ActivityType = "insert-destination"
	MarkDestinationDeparted ActivityType = "mark-destination-departed"
	MarkDestinationVisited  ActivityType = "mark-destination-visited"
	MemberCreated           ActivityType = "member-created"
	MemberDeleted           ActivityType = "member-deleted"
	MemberModified          ActivityType = "member-modified"
	MoveDestination         ActivityType = "move-destination"
	NoteInsert              ActivityType = "note-insert"
	RouteDelete             ActivityType = "route-delete"
	RouteOptimized          ActivityType = "route-optimized"
	RouteOwnerChanged       ActivityType = "route-owner-changed"
	UpdateDestinations      ActivityType = "update-destinations"
	AreaAdded               ActivityType = "area-added"
	AreaRemoved             ActivityType = "area-removed"
	AreaUpdated             ActivityType = "area-updated"
	DestinationOutSequence  ActivityType = "destination-out-sequence"
	DriverArrivedEarly      ActivityType = "driver-arrived-early"
	DriverArrivedOnTime     ActivityType = "driver-arrived-on-time"
	DriverArrivedLate       ActivityType = "driver-arrived-late"
	GeofenceEntered         ActivityType = "geofence-entered"
	GeofenceLeft            ActivityType = "geofence-left"
)

type Query struct {
	RouteID  string       `http:"route_id"`
	DeviceID string       `http:"device_id"`
	Type     ActivityType `http:"activity_type"`
	MemberID int          `http:"member_id"`
	//TODO: Check if it's boolean of string boolean, documentation has two versions
	Team   bool `http:"team"`
	Limit  uint `http:"limit"`
	Offset uint `http:"offset"`
	Start  uint `http:"start"`
	End    uint `http:"end"`
}
