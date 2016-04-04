package route

type Parameters struct {

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
	RouteDate long `json:"route_date,omitempty"`

	//offset in seconds relative to the route start date (i.e. 9AM would be 60 * 60 * 9)
	RouteTime object `json:"route_time,omitempty"`

	//deprecated
	//specify if the route can be viewed by unauthenticated users
	SharedPublicly string `json:"shared_publicly,omitempty"`

	Optimize string `json:"optimize,omitempty"`

	//when the tour type is not round trip (rt = false), enable lock last so that the final destination is fixed
	//example: driver leaves a depot, but must always arrive at home ( or a specific gas station) at the end of the route
	LockLast bool `json:"lock_last,omitempty"`

	VehicleCapacity string `json:"vehicle_capacity,omitempty"`

	VehicleMaxDistanceMI string `json:"vehicle_max_distance_mi,omitempty"`

	//km or mi, the route4me api will convert all measurements into these units
	DistanceUnit string `json:"distance_unit,omitempty"`

	TravelMode string `json:"travel_mode,omitempty"`

	Avoid string `json:"avoid,omitempty"`

	VehicleId string `json:"vehicle_id,omitempty"`

	//deprecated, all new routes should be assigned to a member_id
	DriverId string `json:"driver_id,omitempty"`

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

	//deprecated
	//all routes are stored by default at this time
	StoreRoute bool `json:"store_route,omitempty"`

	//1 = ROUTE4ME_METRIC_EUCLIDEAN (use euclidean distance when computing point to point distance)
	//2 = ROUTE4ME_METRIC_MANHATTAN (use manhattan distance (taxicab geometry) when computing point to point distance)
	//3 = ROUTE4ME_METRIC_GEODESIC (use geodesic distance when computing point to point distance)
	//#4 is the default and suggested metric
	//4 = ROUTE4ME_METRIC_MATRIX (use road network driving distance when computing point to point distance)
	//5 = ROUTE4ME_METRIC_EXACT_2D (use exact rectilinear distance)
	Metric Metric `json:"metric,omitempty"`

	//the type of algorithm to use when optimizing the route
	AlgorithmType AlgorithmType `json:"algorithm_type,omitempty"`

	//in order for users in your organization to have routes assigned to them,
	//you must provide their member id within the route4me system
	//a list of member ids can be retrieved with view_users api method
	MemberId string `json:"member_id,omitempty"`

	//specify the ip address of the remote user making this optimization request
	Ip string `json:"ip,omitempty"`

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

	//deprecated
	DeviceID object `json:"device_id,omitempty"`

	//the type of device making this request
	//ENUM("web", "iphone", "ipad", "android_phone", "android_tablet")
	DeviceType string `json:"device_type,omitempty"`

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

	//there are 3 types of optimization qualities that are optimizations goals
	//1 - Generate Optimized Routes As Quickly as Possible
	//2 - Generate Routes That Look Better On A Map
	//3 - Generate The Shortest And Quickest Possible Routes

	OptimizationQuality int `json:"optimization_quality,omitempty"`
}
