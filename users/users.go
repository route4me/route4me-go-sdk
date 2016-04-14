package users

import (
	"net/http"

	"github.com/route4me/route4me-go-sdk"
)

const endpoint = "/api/member/view_users.php"

type Service struct {
	Client *route4me.Client
}

func (s *Service) GetUsers() ([]User, error) {
	resp := []User{}
	return resp, s.Client.Do(http.MethodGet, endpoint, &struct{}{}, &resp)
}
