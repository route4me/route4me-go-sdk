package main

import (
        "fmt"
        "route4me"
        "time"
)

func main() {
        r4m := route4me.NewClient("11111111111111111111111111111111")

        request := route4me.NewTrackSetParams()
        request.Route_id = "F2FEA85DA7EFCE180CAD70704816347A"
        request.Member_id = 1
        request.Course = 1
        request.Speed = 120
        request.Lat = 41.8927521
        request.Lng = -109.0803888
        request.Device_type = route4me.DEVICETYPE_IPHONE
        request.Device_guid = "qweqweqwe"
        request.Device_timestamp = time.Now()

        response, exception, err := r4m.SetTrack(request)
        if err != nil {
                fmt.Println("Error: " , err)
                return
        }
        if exception != nil {
                for _, error := range exception.Errors {
                        fmt.Println("Error: " , error)
                }
                return
        }


        fmt.Printf("%+v", response)

}
