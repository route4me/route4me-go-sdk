package users

import (
	"net/http"

	"github.com/route4me/route4me-go-sdk"
	"github.com/route4me/route4me-go-sdk/utils"
)

const (
	authEndpoint     = "/actions/authenticate.php"
	validateEndpoint = "/datafeed/session/validate_session.php"
	endpoint         = "/api.v4/user.php"
)

type Service struct {
	Client *route4me.Client
}

func (s *Service) GetUserByID(id int64) (*Member, error) {
	mem := &Member{}
	return mem, s.Client.Do(http.MethodGet, endpoint, &struct {
		MemberID int64 `http:"member_id"`
	}{MemberID: id}, mem)
}

type usersResponse struct {
	Results []*Member `json:"results"`
	Total   int       `json:"total"`
}

func (s *Service) GetSubusers() ([]*Member, error) {
	resp := &usersResponse{}
	return resp.Results, s.Client.Do(http.MethodGet, endpoint, nil, resp)
}

func (s *Service) Register(member *MemberBase) (*Member, error) {
	resp := &Member{}
	return resp, s.Client.Do(http.MethodPost, endpoint, member, resp)
}

func (s *Service) Delete(memberID int64) (*utils.StatusResponse, error) {
	resp := &utils.StatusResponse{}
	return resp, s.Client.Do(http.MethodDelete, endpoint, &struct {
		MemberID int64 `json:"member_id"`
	}{MemberID: memberID}, resp)
}

func (s *Service) Edit(member *Member) (*Member, error) {
	resp := &Member{}
	return resp, s.Client.Do(http.MethodPut, endpoint, member, resp)
}

type authenticationRequest struct {
	Email    string `form:"strEmail"`
	Password string `form:"strPassword"`
	Format   string `form:"format"`
}

func (s *Service) Authenticate(email string, password string) (*Session, error) {
	req := &authenticationRequest{
		Email:    email,
		Password: password,
		Format:   "json",
	}
	resp := &Session{}
	return resp, s.Client.Do(http.MethodPost, authEndpoint, req, resp)
}

type validateRequest struct {
	SessionGUID string `http:"session_guid"`
	MemberID    int64  `http:"member_id"`
	Format      string `http:"format"`
}

type validateResponse struct {
	Authenticated bool `json:"authenticated"`
}

func (s *Service) ValidateSession(guid string, memberID int64) (bool, error) {
	req := &validateRequest{
		SessionGUID: guid,
		MemberID:    memberID,
		Format:      "json",
	}
	resp := &validateResponse{}
	return resp.Authenticated, s.Client.Do(http.MethodGet, validateEndpoint, req, resp)
}
