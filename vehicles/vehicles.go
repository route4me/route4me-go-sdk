package vehicles

import (
	"net/http"

	"github.com/route4me/route4me-go-sdk"
)

const viewVehiclesEndpoint = "/api/vehicles/view_vehicles.php"

type Service struct {
	Client *route4me.Client
}

func (s *Service) GetVehicles() ([]Vehicle, error) {
	resp := []Vehicle{}
	return resp, s.Client.Do(http.MethodGet, viewVehiclesEndpoint, nil, &resp)
}
