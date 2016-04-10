package tracking

import (
	"net/http"

	"github.com/route4me/route4me-go-sdk"
)

const setEndpoint = "/track/set.php"

type Service struct {
	Client *route4me.Client
}

func (s *Service) SetGPS(data *GPS) (string, error) {
	byt, err := s.Client.DoNoDecode(http.MethodGet, setEndpoint, data)
	return string(byt), err
}
