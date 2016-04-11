package routing

import (
	"errors"
	"net/http"
	"strings"

	"github.com/route4me/route4me-go-sdk"
)

const (
	addressEndpoint        = "/api.v4/address.php"
	routeEndpoint          = "/api.v4/route.php"
	optimizationEndpoint   = "/api.v4/optimization_problem.php"
	duplicateRouteEndpoint = "/actions/duplicate_route.php"
)

type Service struct {
	Client *route4me.Client
}

//Optimization
func (s *Service) GetOptimization(parameters *OptimizationParameters) (*DataObject, error) {
	resp := &DataObject{}
	return resp, s.Client.Do(http.MethodGet, optimizationEndpoint, parameters, resp)
}

type getOptimizationsResponse struct {
	Optimizations []DataObject `json:"optimizations"`
}

func (s *Service) GetOptimizations(parameters *RouteQuery) ([]DataObject, error) {
	resp := &getOptimizationsResponse{}
	return resp.Optimizations, s.Client.Do(http.MethodGet, optimizationEndpoint, parameters, resp)
}

func (s *Service) RunOptimization(parameters *OptimizationParameters) (*DataObject, error) {
	resp := &DataObject{}
	return resp, s.Client.Do(http.MethodPost, optimizationEndpoint, parameters, resp)
}

func (s *Service) UpdateOptimization(parameters *OptimizationParameters) (*DataObject, error) {
	resp := &DataObject{}
	return resp, s.Client.Do(http.MethodPut, optimizationEndpoint, parameters, resp)
}

//Addresses
func (s *Service) GetAddress(query *AddressQuery) (*Address, error) {
	resp := &Address{}
	return resp, s.Client.Do(http.MethodGet, addressEndpoint, query, resp)
}

func (s *Service) UpdateAddress(data *Address) (*Address, error) {
	resp := &Address{}
	return resp, s.Client.Do(http.MethodPut, addressEndpoint, data, resp)
}

//Routes
func (s *Service) GetRoute(query *RouteQuery) (*Route, error) {
	resp := &Route{}
	return resp, s.Client.Do(http.MethodGet, routeEndpoint, query, resp)
}

type getRouteIDRequest struct {
	ProblemID         string `http:"optimization_problem_id"`
	WaitForFinalState int    `http:"wait_for_final_state,string"`
}

func (s *Service) GetRouteID(problemID string) (string, error) {
	request := &getRouteIDRequest{ProblemID: problemID, WaitForFinalState: 1}
	response := &DataObject{}
	err := s.Client.Do(http.MethodGet, optimizationEndpoint, request, response)
	if err != nil {
		return "", err
	}
	if len(response.Routes) > 0 {
		return response.Routes[0].ID, nil
	}
	return "", errors.New("Could not find requested route")
}

type duplicateRouteResponse struct {
	ProblemID string `json:"optimization_problem_id,omitempty"`
	Success   bool   `json:"success"`
}

type duplicateRouteRequest struct {
	RouteID string `http:"route_id"`
	To      string `http:"to"`
}

func (s *Service) DuplicateRoute(routeID string) (string, error) {
	request := &duplicateRouteRequest{RouteID: routeID, To: "none"}
	response := &duplicateRouteResponse{}
	err := s.Client.Do(http.MethodGet, duplicateRouteEndpoint, request, response)
	if err != nil {
		return "", err
	}
	if response.Success && response.ProblemID != "" {
		return s.GetRouteID(response.ProblemID)
	}
	return "", errors.New("Could not find requested route")
}

func (s *Service) GetTeamRoutes(query *RouteQuery) ([]Route, error) {
	resp := []Route{}
	return resp, s.Client.Do(http.MethodGet, routeEndpoint, query, &resp)
}

func (s *Service) UpdateRoute(route *Route) (*Route, error) {
	resp := &Route{}
	return resp, s.Client.Do(http.MethodPut, routeEndpoint, route, resp)
}

type deleteRequest struct {
	RouteID string `json:"route_id"`
}

func (s *Service) DeleteRoutes(routeIDs []string) ([]Route, error) {
	request := &deleteRequest{
		RouteID: strings.Join(routeIDs, ","),
	}
	resp := []Route{}
	return resp, s.Client.Do(http.MethodGet, routeEndpoint, request, &resp)
}

//TODO: Add Notes and Removing/Moving destinations
