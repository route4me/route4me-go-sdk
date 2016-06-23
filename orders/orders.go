package orders

import (
	"net/http"

	"github.com/route4me/route4me-go-sdk"
)

const endpoint = "/api.v4/order.php"

type Service struct {
	Client *route4me.Client
}

func (s *Service) Add(order *Order) (*Order, error) {
	response := &Order{}
	return response, s.Client.Do(http.MethodPost, endpoint, order, response)
}

type singleGetRequest struct {
	ID uint64 `http:"order_id"`
}

func (s *Service) Get(orderID uint64) (*Order, error) {
	response := &Order{}
	return response, s.Client.Do(http.MethodGet, endpoint, &singleGetRequest{ID: orderID}, response)
}

type getResponse struct {
	Results []Order `json:"results"`
	Total   int     `json:"total"`
}

func (s *Service) GetAll(query *Query) ([]Order, int, error) {
	response := &getResponse{}
	return response.Results, response.Total, s.Client.Do(http.MethodGet, endpoint, query, response)
}

func (s *Service) Update(order *Order) (*Order, error) {
	response := &Order{}
	return response, s.Client.Do(http.MethodPut, endpoint, order, response)
}

type deleteRequest struct {
	IDs []uint64 `json:"order_ids"`
}

type deleteResponse struct {
	Status bool `json:"status"`
}

func (s *Service) Delete(IDs []uint64) (bool, error) {
	status := &deleteResponse{}
	return status.Status, s.Client.Do(http.MethodDelete, endpoint, &deleteRequest{IDs: IDs}, status)
}
