package activity

import (
	"net/http"

	"github.com/route4me/route4me-go-sdk"
	"github.com/route4me/route4me-go-sdk/utils"
)

const (
	getEndpoint = "/api/get_activities.php"
	logEndpoint = "/api.v4/activity_feed.php"
)

type Service struct {
	Client *route4me.Client
}

type getResponse struct {
	Results []Activity `json:"results"`
	Total   int        `json:"total"`
}

func (s *Service) Get(query *Query) ([]Activity, error) {
	resp := &getResponse{}
	return resp.Results, s.Client.Do(http.MethodGet, getEndpoint, query, resp)
}

func (s *Service) Log(message string, routeID string) error {
	req := &Activity{
		Message: message,
		RouteID: routeID,
		Type:    UserMessage,
	}
	resp := &utils.StatusResponse{}
	err := s.Client.Do(http.MethodPost, logEndpoint, req, resp)
	if err == nil && !resp.Status {
		return utils.ErrOperationFailed
	}
	return err
}
