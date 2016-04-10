package routing

//AlgorithmType is a type of algorithm used for optimization
type AlgorithmType uint

const (
	//TSP stands for single depot, single driver route
	TSP AlgorithmType = iota + 1
	//VRP stands for single depot, multiple driver, no constraints, no time windows, no capacities
	VRP
	//CVRP_TW_SD stands for single depot, multiple driver, capacitated, time windows
	CVRP_TW_SD
	//CVRP_TW_MD stands for multiple depot, multiple driver, capacitated, time windows
	CVRP_TW_MD
	//TSP_TW stands for single depot, single driver, time windows
	TSP_TW
	//TSP_TW_CR stands for single depot, single driver, time windows, continuous optimization (minimal location shifting)
	TSP_TW_CR
	//BBCVRP stands for shifts addresses from one route to another over time on a recurring schedule
	BBCVRP
)

//Metric defines what system is used for computing
type Metric uint

const (
	//Euclidean measures point to point distance as a straight line
	Euclidean = iota + 1
	//Manhattan measures point to point distance as taxicab geometry line
	Manhattan
	//Geodesic measures point to point distance approximating curvature of the earth
	Geodesic
	//Matrix measures point to point distance by traversing the actual road network
	Matrix
	//Exact2D measures point to point distance using 2d rectilinear distance
	Exact2D
)

//DeviceType is a type of the device making a request
type DeviceType string

const (
	Web           = "web"
	IPhone        = "iphone"
	IPad          = "ipad"
	AndroidPhone  = "android_phone"
	AndroidTablet = "android_tablet"
)

//OptimizationQuality decides how fast the optimization finishes and how accurate it is.
type OptimizationQuality uint

const (
	//Fast generates optimized routes as quickly as possible
	Fast OptimizationQuality = iota + 1
	//GoodLooking generates routes that look better on a map
	GoodLooking
	//Best generates the shortest and quickest possible routes
	Best
)

//DistanceUnit describes the unit which all measurements will be converted to
type DistanceUnit string

const (
	Kilometers DistanceUnit = "km"
	Miles      DistanceUnit = "mi"
)

type RouteParameters struct {

	//let the R4M api know if this sdk request is coming from a file upload within your environment (for analytics)
	IsUpload string `json:"is_upload,omitempty"`

	//the tour type of this route. rt is short for round trip, the optimization engine changes its behavior for round trip routes
	RT bool `json:"rt,omitempty"`

	//by disabling optimization, the route optimization engine will not resequence the stops in your
	DisableOptimization bool `json:"disable_optimization,omitempty"`

	//the name of this route. this route name will be accessible in the search API, and also will be displayed on the mobile device of a user
	RouteName string `json:"route_name,omitempty"`

	//the route start date in UTC, unix timestamp seconds.
	//used to show users when the route will begin, also used for reporting and analytics
	RouteDate uint64 `json:"route_date,omitempty"`

	//offset in seconds relative to the route start date (i.e. 9AM would be 60 * 60 * 9)
	RouteTime int `json:"route_time,omitempty"`

	Optimize string `json:"optimize,omitempty"`

	//when the tour type is not round trip (rt = false), enable lock last so that the final destination is fixed
	//example: driver leaves a depot, but must always arrive at home ( or a specific gas station) at the end of the route
	LockLast bool `json:"lock_last,omitempty"`

	VehicleCapacity string `json:"vehicle_capacity,omitempty"`

	VehicleMaxDistanceMI string `json:"vehicle_max_distance_mi,omitempty"`

	DistanceUnit DistanceUnit `json:"distance_unit,omitempty"`

	TravelMode string `json:"travel_mode,omitempty"`

	Avoid string `json:"avoid,omitempty"`

	VehicleID string `json:"vehicle_id,omitempty"`

	//the latitude of the device making this sdk request
	DevLatitude float64 `json:"dev_lat,omitempty"`

	//the longitude of the device making this sdk request
	DevLongitude float64 `json:"dev_lng,omitempty"`

	//when using a multiple driver algorithm, this is the maximum permissible duration of a generated route
	//the optimization system will automatically create more routes when the route_max_duration is exceeded for a route
	//however it will create an 'unrouted' list of addresses if the maximum number of drivers is exceeded
	RouteMaxDuration int `json:"route_max_duration,omitempty"`

	//the email address to notify upon completion of an optimization request
	RouteEmail string `json:"route_email,omitempty"`

	//type of route being created: ENUM(api,null)
	RouteType string `json:"route_type,omitempty"`

	//all routes are stored by default at this time
	StoreRoute bool `json:"store_route,omitempty"`

	Metric Metric `json:"metric,omitempty"`

	//the type of algorithm to use when optimizing the route
	AlgorithmType AlgorithmType `json:"algorithm_type,omitempty"`

	//in order for users in your organization to have routes assigned to them,
	//you must provide their member id within the route4me system
	//a list of member ids can be retrieved with view_users api method
	MemberID string `json:"member_id,omitempty"`

	//specify the ip address of the remote user making this optimization request
	IP string `json:"ip,omitempty"`

	//the method to use when compute the distance between the points in a route
	//1 = DEFAULT (R4M PROPRIETARY ROUTING)
	//2 = DEPRECRATED
	//3 = R4M TRAFFIC ENGINE
	//4 = DEPRECATED
	//5 = DEPRECATED
	//6 = TRUCKING
	DM int `json:"dm,omitempty"`

	//directions method
	//1 = DEFAULT (R4M PROPRIETARY INTERNAL NAVIGATION SYSTEM)
	//2 = DEPRECATED
	//3 = TRUCKING
	//4 = DEPRECATED
	Dirm int `json:"dirm,omitempty"`

	Parts int `json:"parts,omitempty"`

	//the type of device making this request
	//ENUM("web", "iphone", "ipad", "android_phone", "android_tablet")
	DeviceType DeviceType `json:"device_type,omitempty"`

	//for routes that have trucking directions enabled, directions generated
	//will ensure compliance so that road directions generated do not take the vehicle
	//where trailers are prohibited
	HasTrailer bool `json:"has_trailer,omitempty"`

	//for routes that have trucking directions enabled, directions generated
	//will ensure compliance so that road directions generated do not take the vehicle
	//on roads where the weight of the vehicle in tons exceeds this value
	TrailerWeightT float64 `json:"trailer_weight_t,omitempty"`

	LimitedWeightT float64 `json:"limited_weight_t,omitempty"`

	//for routes that have trucking directions enabled, directions generated
	//will ensure compliance so that road directions generated do not take the vehicle
	//where the weight per axle in tons exceeds this value
	WeightPerAxleT float64 `json:"weight_per_axle_t,omitempty"`

	//for routes that have trucking directions enabled, directions generated
	//will ensure compliance of this maximum height of truck when generating road network driving directions
	TruckHeightMeters int `json:"truck_height_meters,omitempty"`

	//for routes that have trucking directions enabled, directions generated
	//will ensure compliance of this width of the truck when generating road network driving directions
	TruckWidthMeters int `json:"truck_width_meters,omitempty"`

	//for routes that have trucking directions enabled, directions generated
	//will ensure compliance of this length of the truck when generating road network driving directions
	TruckLengthMeters int `json:"truck_length_meters,omitempty"`

	//the minimum number of stops permitted per created subroute
	MinTourSize int `json:"min_tour_size,omitempty"`

	//the maximum number of stops permitted per created subroute
	MaxTourSize int `json:"max_tour_size,omitempty"`

	OptimizationQuality OptimizationQuality `json:"optimization_quality,omitempty"`
}
