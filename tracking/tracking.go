package tracking

import (
	"net/http"

	"github.com/route4me/route4me-go-sdk"
	"github.com/route4me/route4me-go-sdk/routing"
)

const (
	setEndpoint            = "/track/set.php"
	lastLocationEndpoint   = "/api.v4/route.php"
	deviceLocationEndpoint = "/api/track/get_device_location.php"
	statusEndpoint         = "/api.v4/status.php"
)

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

type getDeviceLocationHistoryResponse struct {
	Data []routing.TrackingHistory `json:"data"`
}

func (s *Service) GetDeviceLocationHistory(query *TrackingHistoryQuery) ([]routing.TrackingHistory, error) {
	resp := &getDeviceLocationHistoryResponse{}
	return resp.Data, s.Client.Do(http.MethodGet, deviceLocationEndpoint, query, resp)
}

func (s *Service) TrackAssets(tracking string) (*AssetTracking, error) {
	resp := &AssetTracking{}
	return resp, s.Client.Do(http.MethodGet, statusEndpoint, &struct {
		Tracking string `http:"tracking"`
	}{tracking}, resp)
}
