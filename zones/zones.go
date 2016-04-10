package zones

import (
	"net/http"

	"github.com/route4me/route4me-go-sdk"
)

const endpoint = "/api.v4/avoidance.php"

type Service struct {
	Client *route4me.Client
}

func (s *Service) Get(query *Query) (*AvoidanceZone, error) {
	resp := &AvoidanceZone{}
	return resp, s.Client.Do(http.MethodGet, endpoint, query, resp)
}

func (s *Service) GetAll(query *Query) ([]AvoidanceZone, error) {
	resp := []AvoidanceZone{}
	return resp, s.Client.Do(http.MethodGet, endpoint, query, resp)
}

func (s *Service) Add(data *AvoidanceZone) (*AvoidanceZone, error) {
	resp := &AvoidanceZone{}
	return resp, s.Client.Do(http.MethodPost, endpoint, data, resp)
}

func (s *Service) Update(data *AvoidanceZone) (*AvoidanceZone, error) {
	resp := &AvoidanceZone{}
	return resp, s.Client.Do(http.MethodPut, endpoint, data, resp)
}

// func Delete(c *route4me.Client, data Query) (*AvoidanceZone, error) {

// }
