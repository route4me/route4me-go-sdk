package addressbook

import (
	"net/http"

	"github.com/route4me/route4me-go-sdk"
)

const endpoint = "/api.v4/address_book.php"

type Service struct {
	Client *route4me.Client
}

type getResponse struct {
	Results []Contact `json:"results"`
	Total   int       `json:"total"`
}

func (s *Service) Get(query *Query) ([]Contact, int, error) {
	resp := &getResponse{}
	return resp.Results, resp.Total, s.Client.Do(http.MethodGet, endpoint, query, resp)
}

func (s *Service) Add(data *Contact) (*Contact, error) {
	resp := &Contact{}
	err := s.Client.Do(http.MethodPost, endpoint, data, resp)
	return resp, err
}

func (s *Service) Update(data *Contact) (*Contact, error) {
	resp := &Contact{}
	return resp, s.Client.Do(http.MethodPut, endpoint, data, resp)
}

type deleteRequest struct {
	AddressIDs []string `json:"address_ids"`
}

type deleteResponse struct {
	Status bool `json:"status"`
}

func (s *Service) Delete(ids []string) (bool, error) {
	request := &deleteRequest{AddressIDs: ids}
	resp := &deleteResponse{}
	return resp.Status, s.Client.Do(http.MethodDelete, endpoint, request, resp)
}
