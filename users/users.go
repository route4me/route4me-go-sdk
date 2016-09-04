package users

import (
	"net/http"

	"github.com/route4me/route4me-go-sdk"
)

const (
	subUsersEndpoint = "/api/member/view_users.php"
	authEndpoint     = "/actions/authenticate.php"
	validateEndpoint = "/datafeed/session/validate_session.php"
	registerEndpoint = "/actions/register_action.php"
)

type Service struct {
	Client *route4me.Client
}

func (s *Service) GetUsers() ([]User, error) {
	resp := []User{}
	return resp, s.Client.Do(http.MethodGet, subUsersEndpoint, nil, &resp)
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

func (s *Service) RegisterAccount(details *AccountDetails) (*AccountPlan, error) {
	details.Format = "json"
	resp := &AccountPlan{}
	return resp, s.Client.Do(http.MethodPost, registerEndpoint, details, resp)
}
