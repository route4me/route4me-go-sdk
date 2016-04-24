package tracking

import (
	"net/http"

	"github.com/route4me/route4me-go-sdk"
	"github.com/route4me/route4me-go-sdk/routing"
)

const setEndpoint = "/track/set.php"
const lastLocationEndpoint = "/api.v4/route.php"

type Service struct {
	Client *route4me.Client
}

func (s *Service) SetGPS(data *GPS) (string, error) {
	byt, err := s.Client.DoNoDecode(http.MethodGet, setEndpoint, data)
	return string(byt), err
}

type getLastLocationRequest struct {
	RouteID               string `http:"route_id"`
	DeviceTrackingHistory int    `http:"device_tracking_history"`
}

func (s *Service) GetLastLocation(routeID string) (*routing.DataObject, error) {
	request := &getLastLocationRequest{RouteID: routeID, DeviceTrackingHistory: 1}
	resp := &routing.DataObject{}
	return resp, s.Client.Do(http.MethodGet, lastLocationEndpoint, request, resp)
}
