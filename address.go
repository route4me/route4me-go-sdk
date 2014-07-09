package route4me

import (
        "net/url"
        "strconv"
 )

type Address struct {
        Route_destination_id uint `json:"route_destination_id,omitempty"`
        Alias string `json:"alias,omitempty"`
        Member_id interface{} `json:"member_id,omitempty"`
        Address string `json:"address,omitempty"`
        Is_depot bool `json:"is_depot,omitempty"`
        Lat float64` json:"lat,omitempty"`
        Lng float64 `json:"lng,omitempty"`
        Route_id string `json:"route_id,omitempty"`
        Original_route_id string `json:"original_route_id,omitempty"`
        Optimization_problem_id string `json:"optimization_problem_id,omitempty"`
        Sequence_no uint `json:"sequence_no,omitempty"`
        Geocoded bool `json:"geocoded,omitempty"`
        Preferred_geocoding uint `json:"preferred_geocoding,omitempty"`
        Failed_geocoding bool `json:"failed_geocoding,omitempty"`
        Geocodings []string` json:"geocodings,omitempty"`
        Contact_id uint `json:"contact_id,omitempty"`
        Is_visited bool `json:"is_visited,omitempty"`
        Time uint `json:"time,omitempty"`
}

func (r4m *Route4Me) GetAddress(route_id string, route_destination_id uint) (Address, *Exception, error) {
        var container Address

        requestParams := url.Values{}
        requestParams.Set("route_id", route_id)
        requestParams.Set("route_destination_id", strconv.FormatUint(uint64(route_destination_id), 10))

        res, err := r4m.Get("/api.v4/address.php", requestParams)
        response, exception, err := processResponse(res, err, &container)
        if response, ok := response.(*Address); ok {
                return *response, exception, err
        } else {
                return container, exception, err
        }
}
