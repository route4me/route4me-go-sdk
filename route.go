package route4me

import (
        "net/url"
        "strings"
        "strconv"
)

type RouteResponse struct {
        Routes []Route
}

type OptimizeTypeEnum string

const (
        OPTIMIZE_DISTANCE OptimizeTypeEnum = "Distance"
        OPTIMIZE_TIME OptimizeTypeEnum = "Time"
        OPTIMIZE_TIMEWITHTRAFFIC OptimizeTypeEnum = "timeWithTraffic"
)

type DistanceUnitEnum string

const (
        DISTANCEUNIT_MI DistanceUnitEnum = "mi"
        DISTANCEUNIT_KM DistanceUnitEnum = "km"
)

type TravelTypeEnum string

const (
        TRAVELMODE_DRIVING = "Driving"
        TRAVELMODE_WALKING = "Walking"
        TRAVELMODE_TRUCKING = "Trucking"
)

type AvoidTypeEnum string

const (
        AVOIDTYPE_HIGHWAYS AvoidTypeEnum = "Highways"
        AVOIDTYPE_TOLLS AvoidTypeEnum = "Tolls"
        AVOIDTYPE_MINIMIZEHIGHWAYS AvoidTypeEnum = "minimizeHighways"
        AVOIDTYPE_MINIMIZETOLLS AvoidTypeEnum = "minimizeTolls"
)

type DeviceTypeEnum string

const (
        DEVICETYPE_WEB DeviceTypeEnum = "web"
        DEVICETYPE_IPHONE DeviceTypeEnum = "iphone"
        DEVICETYPE_IPAD DeviceTypeEnum = "ipad"
        DEVICETYPE_ANDROIDPHONE DeviceTypeEnum = "android_phone"
        DEVICETYPE_ANDROIDTABLET DeviceTypeEnum = "android_tablet"
)

type AlgorithmTypeEnum uint

const (
        _ = iota
        ALGORITHMTYPE_TSP AlgorithmTypeEnum = iota
        ALGORITHMTYPE_VRP
        ALGORITHMTYPE_CVRP_TW_SD
        ALGORITHMTYPE_CVRP_TW_MD
        ALGORITHMTYPE_TSP_TW
        ALGORITHMTYPE_TSP_TW_CR
        ALGORITHMTYPE_BBCVRP
)

type MetricEnum uint

const (
        _ = iota
        METRIC_EUCLIDEAN MetricEnum = iota
        METRIC_MANHATTAN
        METRIC_GEODESIC
        METRIC_MATRIX
        METRIC_EXACT_2D
)

type RouteParameters struct {
        Is_upload interface{} `json:"is_upload,omitempty"`
        Rt bool `json:"rt,omitempty"`
        Route_name string `json:"route_name,omitempty"`
        Route_date uint `json:"route_date,omitempty"`
        Route_time uint `json:"route_time,omitempty"`
        Shared_publicly interface{} `json:"shared_publicly,omitempty"`
        Disable_optimization bool `json:"disable_optimization,omitempty"`
        Optimize OptimizeTypeEnum `json:"optimize,omitempty"`
        Lock_last bool `json:"lock_last,omitempty"`
        Vehicle_capacity interface{} `json:"vehicle_capacity,omitempty"`
        Vehicle_max_distance interface{} `json:"vehicle_max_distance,omitempty"`
        Distance_unit DistanceUnitEnum `json:"distance_unit,omitempty"`
        Travel_mode TravelTypeEnum `json:"travel_mode,omitempty"`
        Avoid AvoidTypeEnum `json:"avoid,omitempty"`
        Vehicle_id string `json:"vehicle_id,omitempty"`
        Driver_id string `json:"driver_id,omitempty"`
        Dev_lat float32 `json:"dev_lat,omitempty"`
        Dev_lng float32 `json:"dev_lng,omitempty"`
        Route_max_duration uint `json:"route_max_duration,omitempty"`
        Route_email string `json:"route_email,omitempty"`
        Route_type string `json:"route_type,omitempty"`
        Store_route bool `json:"store_route,omitempty"`
        Metric MetricEnum `json:"metric,omitempty"`
        Algorithm_type interface{} `json:"algorithm,omitempty"`
        Member_id string `json:"member_id,omitempty"`
        Ip interface{} `json:"ip,omitempty"`
        Dm uint `json:"dm,omitempty"`
        Dirm uint `json:"dirm,omitempty"`
        Parts uint `json:"parts,omitempty"`
        Device_id string `json:"device_id,omitempty"`
        Device_type DeviceTypeEnum `json:"device_type,omitempty"`
        Has_trailer bool `json:"has_trailer,omitempty"`
        Trailer_weight_t float32 `json:"trailer_weight_t,omitempty"`
        Limited_weight_t float32 `json:"limited_weight_t,omitempty"`
        Weight_per_axle_t float32 `json:"weight_per_axle_t,omitempty"`
        Truck_height_meters uint `json:"truck_height_memters,omitempty"`
        Truck_width_meters uint `json:"truck_width_meters,omitempty"`
        Truck_hazardous_goods string `json:"truck_hazardous_goods,omitempty"`
}

type Route struct {
        Route_id string `json:"route_id,omitempty"`
        Optimization_problem_id string `json:"optimization_problem_id,omitempty"`
        Vehicle_alias string `json:"vehicle_alias,omitempty"`
        Driver_alias string `json:"driver_alias,omitempty"`
        Trip_distance float32 `json:"trip_distance,omitempty"`
        Mpg float32 `json:"mpg,omitempty"`
        Gas_price string `json:"gas_price,omitempty"`
        Route_duration_sec uint `json:"route_duration_sec,omitempty"`
        Parameters RouteParameters `json:"parameters,omitempty"`
        Addresses []Address `json:"addresses,omitempty"`
        Links struct {
                Optimization_problem_id string `json:"optimization_problem_id,omitempty"`
                Route string `json:"route,omitempty"`
        } `json:"llinks,omitempty"`
}

type RouteSearchParams struct {
        Route_id []string
        Directions bool
        Route_path_output string
        Device_tracking_history bool
        Limit uint
        Offset uint
}

func (request *RouteSearchParams) ToParams() url.Values {
       requestParams := url.Values{}

        if request.Route_id != nil {
                requestParams.Set("route_id", strings.Join(request.Route_id, ","))
        }

        if request.Directions {
                requestParams.Set("directions", "1")
        }

        if request.Route_path_output != "None" {
                requestParams.Set("route_path_output", request.Route_path_output)
        }

        if request.Limit != 10 {
                requestParams.Set("limit", strconv.Itoa(int(request.Limit)))
        }

        if request.Offset != 0 {
                requestParams.Set("offset", strconv.Itoa(int(request.Offset)))
        }

        return requestParams
}

func NewRouteSearchParams() (RouteSearchParams) {
        return RouteSearchParams{
                Route_id: nil,
                Directions: false,
                Route_path_output: "None",
                Device_tracking_history: false,
                Limit: 10,
                Offset: 0}
}

func (r4m *Route4Me) SearchRoutes(request RouteSearchParams) ([]Route, *Exception, error) {
        var container []Route

        requestParams :=  request.ToParams()
        res, err := r4m.Get("/api.v4/route.php", requestParams)

        response, exception, err := processResponse(res, err, &container)
        if response, ok := response.(*[]Route); ok {
                return *response, exception, err
        } else {
                return nil, exception, err
        }
}


func (r4m *Route4Me) GetRoute(route_id string) (Route, *Exception, error) {

        var container Route

        requestParams := url.Values{}
        requestParams.Set("route_id", route_id)

        res, err := r4m.Get("/api.v4/route.php", requestParams)

        response, exception, err := processResponse(res, err, &container)
        if response, ok := response.(*Route); ok {
                return *response, exception, err
        } else {
                return container, exception, err
        }
}
