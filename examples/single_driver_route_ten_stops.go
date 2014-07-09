package main

import (
        "fmt"
        "route4me"
)

func main() {
        r4m := route4me.NewClient("11111111111111111111111111111111")


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
                Store_route: true,
        }
        request.Addresses = []route4me.Address{
                route4me.Address {
                        Lng: -85.757308,
                        Lat: 38.251698,
                        Is_depot: true,
                        Time: 300,
                        Sequence_no: 0,
                        Address: "455 S 4th St, Louisville, KY 40202",
                },
                route4me.Address {
                        Lng: -85.793846,
                        Lat: 38.141598,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 1,
                        Address: "1604 PARKRIDGE PKWY, Louisville, KY, 40214",
                },
                route4me.Address {
                        Lng: -85.786514,
                        Lat: 38.202496,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 2,
                        Address: "1407 MCCOY, Louisville, KY, 40215",
                },
                route4me.Address {
                        Lng: -85.774864,
                        Lat: 38.178844,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 3,
                        Address: "4805 BELLEVUE AVE, Louisville, KY, 40215",
                },
                route4me.Address {
                        Lng: -85.821121,
                        Lat: 38.248684,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 4,
                        Address: "730 CECIL AVENUE, Louisville, KY, 40211",
                },
                route4me.Address {
                        Lng: -85.800034,
                        Lat: 38.251923,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 5,
                        Address: "650 SOUTH 29TH ST UNIT 315, Louisville, KY, 40211",
                },
                route4me.Address {
                        Lng: -85.824638,
                        Lat: 38.176067,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 6,
                        Address: "4629 HILLSIDE DRIVE, Louisville, KY, 40216",
                },
                route4me.Address {
                        Lng: -85.775558,
                        Lat: 38.179806,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 7,
                        Address: "4738 BELLEVUE AVE, Louisville, KY, 40215",
                },
                route4me.Address {
                        Lng: -85.815094,
                        Lat: 38.259335,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 8,
                        Address: "318 SO. 39TH STREET, Louisville, KY, 40212",
                },
                route4me.Address {
                        Lng: -85.785118,
                        Lat: 38.179253,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 9,
                        Address: "1324 BLUEGRASS AVE, Louisville, KY, 40215",
                },
        }

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
