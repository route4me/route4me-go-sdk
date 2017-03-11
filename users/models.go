package users

import "github.com/route4me/route4me-go-sdk"

type MemberType string

const (
	PrimaryAccount            MemberType = "PRIMARY_ACCOUNT"
	SubAccountAdmin           MemberType = "SUB_ACCOUNT_ADMIN"
	SubAccountRegionalManager MemberType = "SUB_ACCOUNT_REGIONAL_MANAGER"
	SubAccountDispatcher      MemberType = "SUB_ACCOUNT_DISPATCHER"
	SubAccountDriver          MemberType = "SUB_ACCOUNT_DRIVER"
)

type MemberBase struct {
	FirstName string     `json:"member_first_name,omitempty"`
	LastName  string     `json:"member_last_name,omitempty"`
	Phone     string     `json:"member_phone,omitempty"`
	Type      MemberType `json:"member_type,omitempty"`
	Email     string     `json:"member_email,omitempty"`
	Password  string     `json:"member_password,omitempty"`
	ZIPCode   string     `json:"member_zipcode,omitempty"`

	PreferredUnits    string `json:"preferred_units,omitempty"`
	PreferredLanguage string `json:"preferred_language,omitempty"`
	Timezone          string `json:"timezone,omitempty"`

	RegionCountryID string `json:"user_reg_country_id,omitempty"`
	RegionStateID   string `json:"user_reg_state_id,omitempty"`

	HideRouteAddresses   route4me.StringBool `json:"HIDE_ROUTED_ADDRESSES,omitempty"`
	HideVisitedAddresses route4me.StringBool `json:"HIDE_VISITED_ADDRESSES,omitempty"`
	HideNonFutureRoutes  route4me.StringBool `json:"HIDE_NONFUTURE_ROUTES,omitempty"`
	ReadOnly             route4me.StringBool `json:"READONLY_USER,omitempty"`
	ShowAllDrivers       route4me.StringBool `json:"SHOW_ALL_DRIVERS,omitempty"`
	ShowAllVehicles      route4me.StringBool `json:"SHOW_ALLVehicles,omitempty"`
	DateOfBirth          string              `json:"date_of_birth,omitempty"`
}

type Member struct {
	MemberBase
	ID      int64 `json:"member_id,string,omitempty"`
	OwnerID int64 `json:"OWNER_MEMBER_ID,string,omitempty"`
}

type Session struct {
	Status                        bool          `json:"status"`
	Error                         string        `json:"error,omitempty"`
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

	HideRouteAddresses   route4me.Bool `json:"HIDE_ROUTED_ADDRESSES"`
	HideVisitedAddresses route4me.Bool `json:"HIDE_VISITED_ADDRESSES"`
	HideNonFutureRoutes  route4me.Bool `json:"HIDE_NONFUTURE_ROUTES"`
	ReadOnly             route4me.Bool `json:"READONLY_USER"`
	AutoLogoutTs         int           `json:"auto_logout_ts"`
}

type KeyValue struct {
	MemberID int    `json:"member_id,omitempty"`
	Key      string `json:"config_key,omitempty"`
	Value    string `json:"config_value,omitempty"`
}

type WebinarRegistration struct {
	EmailAddress string `json:"email_address"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	PhoneNumber  string `json:"phone_number"`
	CompanyName  string `json:"company_name"`
	MemberID     string `json:"member_id"`
	WebiinarDate string `json:"webiinar_date"`
}
