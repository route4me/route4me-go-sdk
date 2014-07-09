package main

import (
        "fmt"
        "route4me"
)

func main() {
        r4m := route4me.NewClient("11111111111111111111111111111111")

        response, exception, err := r4m.GetRoute("C154F6B4739B14FE46A2DE43E7A9A833")
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
