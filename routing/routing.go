package routing

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/route4me/route4me-go-sdk"
	"github.com/route4me/route4me-go-sdk/utils"
)

const (
	addressEndpoint        = "/api.v4/address.php"
	routeEndpoint          = "/api.v4/route.php"
	optimizationEndpoint   = "/api.v4/optimization_problem.php"
	duplicateRouteEndpoint = "/actions/duplicate_route.php"
	notesEndpoint          = "/actions/addRouteNotes.php"
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

func (s *Service) GetAddressNotes(query *NoteQuery) ([]Note, error) {
	addressQuery := &AddressQuery{
		RouteID:            query.RouteID,
		RouteDestinationID: query.AddressID,
		Notes:              true,
	}
	addr, err := s.GetAddress(addressQuery)
	return addr.Notes, err
}

type addAddressNoteResponse struct {
	Status bool  `json:"status"`
	Note   *Note `json:"note,omitempty"`
}

func (s *Service) AddAddressNote(query *NoteQuery, noteContents string) (*Note, error) {
	strUpdateType := "unclassified"
	if query.ActivityType != "" {
		strUpdateType = string(query.ActivityType)
	}
	getValues := utils.StructToURLValues(query)
	getValues.Add("api_key", s.Client.APIKey)

	bodyValues := &url.Values{}
	bodyValues.Add("strUpdateType", strUpdateType)
	bodyValues.Add("strNoteContents", noteContents)
	response := &addAddressNoteResponse{}

	request, err := http.NewRequest("POST", s.Client.BaseURL+notesEndpoint, strings.NewReader(getValues.Encode()))
	if err != nil {
		return nil, err
	}
	request.URL.RawQuery = getValues.Encode()
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	read, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(read, response)
	if err != nil {
		return nil, err
	}
	if !response.Status {
		return nil, errors.New("Note not added.")
	}
	return response.Note, nil
}

//TODO: Removing/Moving destinations
