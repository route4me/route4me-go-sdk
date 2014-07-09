package route4me

import (
        "net/url"
        "strconv"
        "time"
)


/*
func (t JsonDateTime) MarshalJSON() ([]byte, error) {
        if y := t.Year(); y < 0 || y >= 10000 {
                return nil, errors.New("JsonDateTime.MarshalJSON: year outside of range [0,9999]")
        }
        return []byte(t.Format(`"` + "2006-01-02 15:04:05" + `"`)), nil
}
*/

type TrackResponse struct {
        Status bool `json:"status"`
}

type TrackSetParams struct {
        //Format string
        Member_id uint64
        Route_id string
        Tx_id string
        Vehicle_id uint64
        Course uint64
        Speed float64
        Lat float64
        Lng float64
        Altitude float64
        Device_guid string
        Device_type DeviceTypeEnum
        Device_timestamp time.Time
        App_version string
}


func NewTrackSetParams() (TrackSetParams) {
        return TrackSetParams{}
}

func (r4m *Route4Me) SetTrack(request TrackSetParams) (bool, *Exception, error) {
        var container TrackResponse

        requestParams := url.Values{}

        requestParams.Set("format", "serialized")
        requestParams.Set("member_id", strconv.FormatUint(request.Member_id, 10))
        requestParams.Set("route_id", request.Route_id)
        requestParams.Set("vehicle_id", strconv.FormatUint(request.Vehicle_id, 10))
        requestParams.Set("lat", strconv.FormatFloat(request.Lat, 'g', -1, 64))
        requestParams.Set("lng", strconv.FormatFloat(request.Lng, 'g', -1, 64))
        requestParams.Set("altitude", strconv.FormatFloat(request.Altitude, 'g', -1, 64))
        requestParams.Set("device_type", string(request.Device_type))
        requestParams.Set("device_guid", request.Device_guid)
        requestParams.Set("Device_timestamp", request.Device_timestamp.Format("2006-01-02 15:04:05"))
        requestParams.Set("app_version", request.App_version)


        res, err := r4m.Get("/track/set.php", requestParams)

        response, exception, err := processResponse(res, err, &container)
        if response, ok := response.(*TrackResponse); ok {
                return (*response).Status, exception, err
        } else {
                return false, exception, err
        }
}
