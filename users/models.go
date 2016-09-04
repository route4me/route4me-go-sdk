package users

import "github.com/route4me/route4me-go-sdk"

type User struct {
	MemberID               int64         `json:"member_id,string,omitempty"`
	AccountTypeID          int           `json:"account_type_id,string,omitempty"`
	MemberType             string        `json:"member_type,omitempty"`
	MemberFirstName        string        `json:"member_first_name,omitempty"`
	MemberLastName         string        `json:"member_last_name,omitempty"`
	MemberEmail            string        `json:"member_email,omitempty"`
	PhoneNumber            string        `json:"phone_number,omitempty"`
	ReadonlyUser           route4me.Bool `json:"readonly_user,omitempty"`
	ShowSuperUserAddresses route4me.Bool `json:"show_superuser_addresses,omitempty"`
}

type Session struct {
	Status                        bool          `json:"status"`
	GeocodingService              string        `json:"geocoding_service"`
	SessionID                     int64         `json:"session_id"`
	SessionGUID                   string        `json:"session_guid"`
	MemberID                      string        `json:"member_id"`
	APIKey                        string        `json:"api_key"`
	TrackingTTL                   int           `json:"tracking_ttl"`
	GeofencePolygonShape          string        `json:"geofence_polygon_shape"`
	GeofencePolygonSize           int           `json:"geofence_polygon_size"`
	GeofenceTimeOnsiteTriggerSecs int           `json:"geofence_time_onsite_trigger_secs"`
	GeofenceMinimumTriggerSpeed   int           `json:"geofence_minimum_trigger_speed"`
	IsSubscriptionPastDue         route4me.Bool `json:"is_subscription_past_due"`
	VisitedDepartedEnabled        route4me.Bool `json:"visited_departed_enabled"`
	LongPressEnabled              route4me.Bool `json:"long_press_enabled"`
	AccountTypeID                 string        `json:"account_type_id"`
	AccountTypeAlias              string        `json:"account_type_alias"`
	MemberType                    string        `json:"member_type"`
	MaxStopsPerRoute              string        `json:"max_stops_per_route"`
	MaxRoutes                     string        `json:"max_routes"`
	RoutesPlanned                 string        `json:"routes_planned"`
	PreferredUnits                string        `json:"preferred_units"`
	PreferredLanguage             string        `json:"preferred_language"`
	HideRouteAddresses            route4me.Bool `json:"HIDE_ROUTED_ADDRESSES"`
	HideVisitedAddresses          route4me.Bool `json:"HIDE_VISITED_ADDRESSES"`
	HideOnFutureRoutes            route4me.Bool `json:"HIDE_NONFUTURE_ROUTES"`
	ReadOnly                      route4me.Bool `json:"READONLY_USER"`
	AutoLogoutTs                  int           `json:"auto_logout_ts"`
}

type AccountDetails struct {
	Plan      string `form:"plan"`
	Industry  string `form:"strIndustry"`
	FirstName string `form:"strFirstName"`
	LastName  string `form:"strLastName"`
	Email     string `form:"strEmail"`
	//if equal to 1, user agrees to registration terms
	AcceptTerms int    `form:"chkTerms"`
	DeviceType  string `form:"device_type"`
	Password1   string `form:"strPassword_1"`
	Password2   string `form:"strPassword_2"`

	Format string `form:"format"`
}

type AccountPlan struct {
	Status           bool   `json:"status"`
	SessionID        string `json:"session_id"`
	SessionGUID      string `json:"session_guid"`
	MemberID         int    `json:"member_id"`
	APIKey           string `json:"api_key"`
	AccountTypeID    int    `json:"account_type_id"`
	AccountTypeAlias int    `json:"account_type_alias"`
	MaxStopsPerRoute int    `json:"max_stops_per_route"`
	MaxRoutes        int    `json:"max_routes"`
	RoutesPlanned    int    `json:"routes_planned"`
}
