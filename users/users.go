package users

import (
	"errors"
	"net/http"

	"github.com/route4me/route4me-go-sdk"
	"github.com/route4me/route4me-go-sdk/utils"
)

const (
	authEndpoint          = "/actions/authenticate.php"
	validateEndpoint      = "/datafeed/session/validate_session.php"
	registerEndpoint      = "/actions/register_action.php"
	webinarEndpoint       = "/actions/webinar_register.php"
	endpoint              = "/api.v4/user.php"
	configurationEndpoint = "/api.v4/configuration-settings.php"
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

func (s *Service) RegisterToWebinar(data *WebinarRegistration) (bool, error) {
	resp := &utils.StatusResponse{}
	return resp.Status, s.Client.Do(http.MethodPost, webinarEndpoint, data, resp)
}

func (s *Service) Register(member *MemberBase) (*Member, error) {
	resp := &Member{}
	return resp, s.Client.Do(http.MethodPost, endpoint, member, resp)
}

func (s *Service) Delete(memberID int64) (bool, error) {
	resp := &utils.StatusResponse{}
	return resp.Status, s.Client.Do(http.MethodDelete, endpoint, &struct {
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

type Account struct {
	Industry  string `form:"strIndustry"`
	FirstName string `form:"strFirstName"`
	LastName  string `form:"strLastName"`
	Email     string `form:"strEmail"`
	Password  string `form:"strPassword_1"`
	AcceptTOS bool   `form:"chkTerms"`
	Plan      string `http:"plan"`
}

type createRequest struct {
	Industry  string `form:"strIndustry"`
	FirstName string `form:"strFirstName"`
	LastName  string `form:"strLastName"`
	Email     string `form:"strEmail"`
	Password  string `form:"strPassword_1"`
	AcceptTOS bool   `form:"chkTerms"`
	Plan      string `http:"plan"`

	Password2  string `form:"strPassword_2"`
	DeviceType string `form:"device_type"`
	Format     string `form:"format"`
}

func (s *Service) Create(account *Account) (bool, error) {
	req := &createRequest{
		Industry:   account.Industry,
		FirstName:  account.FirstName,
		LastName:   account.LastName,
		Email:      account.Email,
		Password:   account.Password,
		Password2:  account.Password,
		AcceptTOS:  account.AcceptTOS,
		Plan:       account.Plan,
		DeviceType: "web",
		Format:     "json",
	}
	resp := utils.StatusResponse{}
	return resp.Status, s.Client.Do(http.MethodPost, registerEndpoint, req, &resp)
}

type configResponse struct {
	Result   string     `json:"result"`
	Affected int        `json:"affected"`
	Data     []KeyValue `json:"data,omitempty"`
}

func (s *Service) AddConfigEntry(key string, value string) (bool, error) {
	resp := &configResponse{}
	return resp.Result == "OK", s.Client.Do(http.MethodPost, configurationEndpoint, &KeyValue{Key: key, Value: value}, resp)
}

type getConfigEntryRequest struct {
	Key string `http:"config_key"`
}

func (s *Service) GetConfigEntry(key string) (string, error) {
	resp := &configResponse{}
	err := s.Client.Do(http.MethodGet, configurationEndpoint, &getConfigEntryRequest{Key: key}, resp)
	if err != nil {
		return "", err
	}
	if len(resp.Data) != 1 {
		return "", errors.New("Specific key has not been found")
	}
	return resp.Data[0].Value, err
}

func (s *Service) DeleteConfigEntry(key string) (bool, error) {
	resp := &configResponse{}
	return resp.Result == "OK", s.Client.Do(http.MethodDelete, configurationEndpoint, &KeyValue{Key: key}, resp)
}

func (s *Service) GetConfigValues() ([]KeyValue, error) {
	resp := &configResponse{}
	return resp.Data, s.Client.Do(http.MethodGet, configurationEndpoint, nil, resp)
}

func (s *Service) UpdateConfigEntry(key string, value string) (bool, error) {
	resp := &configResponse{}
	return resp.Result == "OK", s.Client.Do(http.MethodPut, configurationEndpoint, &KeyValue{Key: key, Value: value}, resp)
}
