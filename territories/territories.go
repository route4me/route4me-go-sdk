package territories

import (
	"net/http"

	"github.com/route4me/route4me-go-sdk"
	"github.com/route4me/route4me-go-sdk/utils"
)

const (
	avoidanceEndpoint = "/api.v4/avoidance.php"
	territoryEndpoint = "/api.v4/territory.php"
)

type Service struct {
	Client *route4me.Client
}

func (s *Service) GetAvoidanceZone(query *Query) (*AvoidanceZone, error) {
	resp := &AvoidanceZone{}
	return resp, s.Client.Do(http.MethodGet, avoidanceEndpoint, query, resp)
}

func (s *Service) GetAvoidanceZones(query *Query) ([]AvoidanceZone, error) {
	resp := []AvoidanceZone{}
	return resp, s.Client.Do(http.MethodGet, avoidanceEndpoint, query, &resp)
}

func (s *Service) AddAvoidanceZone(data *AvoidanceZone) (*AvoidanceZone, error) {
	resp := &AvoidanceZone{}
	return resp, s.Client.Do(http.MethodPost, avoidanceEndpoint, data, resp)
}

func (s *Service) UpdateAvoidanceZone(data *AvoidanceZone) (*AvoidanceZone, error) {
	resp := &AvoidanceZone{}
	return resp, s.Client.Do(http.MethodPut, avoidanceEndpoint, data, resp)
}

func (s *Service) DeleteAvoidanceZone(data *Query) error {
	resp := &utils.StatusResponse{}
	err := s.Client.Do(http.MethodDelete, avoidanceEndpoint, data, resp)
	if err == nil && !resp.Status {
		return utils.ErrOperationFailed
	}
	return err
}

func (s *Service) GetTerritory(query *Query) (*Territory, error) {
	resp := &Territory{}
	return resp, s.Client.Do(http.MethodGet, territoryEndpoint, query, resp)
}

func (s *Service) GetTerritories(query *Query) ([]Territory, error) {
	resp := []Territory{}
	return resp, s.Client.Do(http.MethodGet, territoryEndpoint, query, &resp)
}
func (s *Service) AddTerritory(data *Territory) (*Territory, error) {
	resp := &Territory{}
	return resp, s.Client.Do(http.MethodPost, territoryEndpoint, data, resp)
}

func (s *Service) UpdateTerritory(data *Territory) (*Territory, error) {
	resp := &Territory{}
	return resp, s.Client.Do(http.MethodPut, territoryEndpoint, data, resp)
}

func (s *Service) DeleteTerritory(data *Query) error {
	resp := &utils.StatusResponse{}
	err := s.Client.Do(http.MethodDelete, territoryEndpoint, data, resp)
	if err == nil && !resp.Status {
		return utils.ErrOperationFailed
	}
	return err
}
