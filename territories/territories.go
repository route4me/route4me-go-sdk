package territories

import (
	"net/http"

	"github.com/route4me/route4me-go-sdk"
	"github.com/route4me/route4me-go-sdk/utils"
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
	return resp, s.Client.Do(http.MethodGet, endpoint, query, &resp)
}

func (s *Service) Add(data *AvoidanceZone) (*AvoidanceZone, error) {
	resp := &AvoidanceZone{}
	return resp, s.Client.Do(http.MethodPost, endpoint, data, resp)
}

func (s *Service) Update(data *AvoidanceZone) (*AvoidanceZone, error) {
	resp := &AvoidanceZone{}
	return resp, s.Client.Do(http.MethodPut, endpoint, data, resp)
}

func (s *Service) Delete(data *Query) error {
	resp := &utils.StatusResponse{}
	err := s.Client.Do(http.MethodDelete, endpoint, data, resp)
	if err == nil && !resp.Status {
		return utils.ErrOperationFailed
	}
	return err
}
