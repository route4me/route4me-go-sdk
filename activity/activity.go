package activity

import (
	"net/http"

	"github.com/route4me/route4me-go-sdk"
)

const endpoint = "/api/get_activities.php"

type Service struct {
	Client *route4me.Client
}

type getResponse struct {
	Results []Activity `json:"results"`
}

func (s *Service) Get(query *Query) ([]Activity, error) {
	resp := &getResponse{}
	return resp.Results, s.Client.Do(http.MethodGet, endpoint, query, resp)
}
