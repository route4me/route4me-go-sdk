package routing

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/route4me/route4me-go-sdk"
	"github.com/route4me/route4me-go-sdk/utils"
)

const (
	addressEndpoint              = "/api.v4/address.php"
	routeEndpoint                = "/api.v4/route.php"
	optimizationEndpoint         = "/api.v4/optimization_problem.php"
	duplicateRouteEndpoint       = "/actions/duplicate_route.php"
	notesEndpoint                = "/actions/addRouteNotes.php"
	moveRouteDestinationEndpoint = "/actions/route/move_route_destination.php"
	mergeRoutesEndpoint          = "/actions/merge_routes.php"
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

type deleteOptimizationRequest struct {
	OptimizationProblemID string `http:"optimization_problem_id"`
}

func (s *Service) DeleteOptimization(optimizationProblemID string) error {
	return s.Client.Do(http.MethodDelete, optimizationEndpoint, &deleteOptimizationRequest{OptimizationProblemID: optimizationProblemID}, nil)
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

type deleteAddressRequest struct {
	OptimizationProblemID string      `http:"optimization_problem_id"`
	RouteDestinationID    json.Number `http:"route_destination_id"`
}

type deleteAddressResponse struct {
	Deleted            bool         `json:"deleted"`
	RouteDestinationID *json.Number `json:"route_destination_id,omitempty"`
}

// DeleteAddress removes a destination (an address) with specified route_destination_id from an optimization problem with specified optimization_problem_id.
func (s *Service) DeleteAddress(optimizationID string, routeDestinationID string) (*json.Number, error) {
	req := &deleteAddressRequest{
		OptimizationProblemID: optimizationID,
		RouteDestinationID:    json.Number(routeDestinationID),
	}
	resp := &deleteAddressResponse{}
	err := s.Client.Do(http.MethodDelete, addressEndpoint, req, resp)
	if err == nil && !resp.Deleted {
		return nil, utils.ErrOperationFailed
	}
	return resp.RouteDestinationID, err
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
	RouteID string `http:"route_id"`
}

func (s *Service) DeleteRoutes(routeIDs ...string) ([]Route, error) {
	request := &deleteRequest{
		RouteID: strings.Join(routeIDs, ","),
	}
	resp := []Route{}
	return resp, s.Client.Do(http.MethodDelete, routeEndpoint, request, &resp)
}

type mergeRequest struct {
	RouteIDs []string `json:"route_ids"`
}

func (s *Service) MergeRoutes(routeIDs ...string) error {
	request := &mergeRequest{
		RouteIDs: routeIDs,
	}
	resp := &utils.StatusResponse{}
	err := s.Client.Do(http.MethodPost, mergeRoutesEndpoint, request, resp)
	if err == nil && !resp.Status {
		return utils.ErrOperationFailed
	}
	return err
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

type addAddressNoteRequest struct {
	*NoteQuery
	NoteContents string `form:"strNoteContents"`
	UpdateType   string `form:"strUpdateType"`
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

	request := &addAddressNoteRequest{
		NoteQuery:    query,
		UpdateType:   strUpdateType,
		NoteContents: noteContents,
	}
	resp := &addAddressNoteResponse{}
	err := s.Client.Do(http.MethodPost, notesEndpoint, request, resp)
	// always returns false (needs clarification)
	// if err == nil && !resp.Status {
	// 	return nil, utils.ErrOperationFailed
	// }
	return resp.Note, err
}

type addRouteDestinationsRequest struct {
	RouteID   string    `http:"route_id"`
	Addresses []Address `json:"addresses"`
}

func (s *Service) AddRouteDestinations(routeID string, addresses []Address) (*DataObject, error) {
	request := &addRouteDestinationsRequest{
		RouteID:   routeID,
		Addresses: addresses,
	}
	resp := &DataObject{}
	return resp, s.Client.Do(http.MethodPut, routeEndpoint, request, resp)
}

type removeRouteDestinationRequest struct {
	RouteID            string `http:"route_id"`
	RouteDestinationID string `http:"route_destination_id"`
}

type removeRouteDestinationResposne struct {
	Deleted bool `json:"deleted"`
}

func (s *Service) RemoveRouteDestination(routeID string, destinationID string) (bool, error) {
	request := &removeRouteDestinationRequest{
		RouteID:            routeID,
		RouteDestinationID: destinationID,
	}
	resp := &removeRouteDestinationResposne{}
	return resp.Deleted, s.Client.Do(http.MethodDelete, routeEndpoint, request, resp)
}

type DestinationMoveRequest struct {
	ToRouteID          string `form:"to_route_id"`
	RouteDestinationID string `form:"route_destination_id"`
	AfterDestinationID string `form:"after_destination_id"`
}

type moveDestinationToRouteResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

func (s *Service) MoveDestinationToRoute(query *DestinationMoveRequest) error {
	resp := &moveDestinationToRouteResponse{}
	err := s.Client.Do(http.MethodPost, moveRouteDestinationEndpoint, query, resp)
	if err == nil && !resp.Success {
		return utils.ErrOperationFailed
	}
	return err
}
