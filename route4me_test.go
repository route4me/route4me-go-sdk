package route4me

import "testing"

func getClient() *Route4Me {
        return NewClient("11111111111111111111111111111111")
}

func TestOptimization(t *testing.T) {
        r4m := getClient()

        request := NewOptimizationNewParams()
        request.Parameters = &RouteParameters{
                Algorithm_type: ALGORITHMTYPE_TSP,
                Shared_publicly: false,
                Store_route: false,
                Route_time: 0,
                Route_max_duration: 86400,
                Vehicle_capacity: 1,
                Vehicle_max_distance: 10000,
                Route_name: "Single Driver Round Trip",
                Optimize: OPTIMIZE_DISTANCE,
                Distance_unit: DISTANCEUNIT_MI,
                Device_type: DEVICETYPE_WEB,
                Travel_mode: TRAVELMODE_DRIVING,
        }
        request.Addresses = []Address{
                Address {
                        Lng: -83.244743347168,
                        Lat: 33.132675170898,
                        Is_depot: true,
                        Time: 0,
                        Address: "151 Arbor Way Milledgeville GA 31061",
                },
                Address {
                        Lng: -83.24577331543,
                        Lat: 33.129695892334,
                        Time: 0,
                        Address: "230 Arbor Way Milledgeville GA 31061",
                },
        }

        response, exception, err := r4m.NewOptimization(request)
        if err != nil {
                t.Error(err)
                return
        }
        if exception != nil {
                for _, error := range exception.Errors {
                        t.Error(error)
                }
                return
        }

        if response.State != OPTIMIZATIONSTATE_OPTIMIZED {
                t.Error("Expected 4, got ", response.State, " (", response.Links.View , ")")
        }
        return
}


func TestRoute(t *testing.T) {
        r4m := getClient()

        request := NewRouteSearchParams()
        _, exception, err := r4m.SearchRoutes(request)
        if err != nil {
                t.Error(err)
                return
        }
        if exception != nil {
                for _, error := range exception.Errors {
                        t.Error(error)
                }
                return
        }

        // Don't need to check type of response, because it will always
        //be an array of Route
}
