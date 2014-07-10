package main

import (
        "fmt"
        "github.com/route4me/route4me-go-sdk"
        "encoding/json"
        "io/ioutil"
)

func main() {
        r4m := route4me.NewClient("11111111111111111111111111111111")

        fileContents, err := ioutil.ReadFile("addresses.json")
        if err != nil {
                fmt.Print(err)
                return
        }

        var addresses []route4me.Address
        err = json.Unmarshal(fileContents, &addresses)
        if err != nil {
                fmt.Print(err)
                return
        }

        request := route4me.NewOptimizationNewParams()
        request.Parameters = &route4me.RouteParameters{
                Algorithm_type: route4me.ALGORITHMTYPE_TSP,
                Distance_unit: route4me.DISTANCEUNIT_MI,
                Device_type: route4me.DEVICETYPE_WEB,
                Optimize: route4me.OPTIMIZE_DISTANCE,
                Route_max_duration: 86400,
                Travel_mode: route4me.TRAVELMODE_DRIVING,
                Vehicle_capacity: 1,
                Vehicle_max_distance: 10000,
                Rt: true,
                Store_route: true,
        }
        request.Addresses = addresses[0:10]

        response, exception, err := r4m.NewOptimization(request)
        if err != nil {
                fmt.Print(err)
                return
        }
        if exception != nil {
                for _, error := range exception.Errors {
                        fmt.Print(error)
                }
                return
        }
        fmt.Printf("%+v", response)
}
