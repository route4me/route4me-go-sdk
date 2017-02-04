package geocoding

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/route4me/route4me-go-sdk"
)

const (
	RapidBaseURL       = "https://rapid.route4me.com"
	streetDataEndpoint = "/street_data/"
	geocoderEndpoint   = "/api/geocoder.php"
)

type Service struct {
	Client *route4me.Client
}

type geocodeRequest struct {
	Addresses string `http:"addresses" json:"-"`
	Format    string `http:"format" json:"-"`
}

func (s *Service) ForwardAddress(address string) ([]Geocoding, error) {
	geo := []Geocoding{}
	return geo, s.Client.Do(http.MethodPost, geocoderEndpoint, &geocodeRequest{Addresses: address, Format: "json"}, &geo)
}

func (s *Service) ReverseAddress(latitude float64, longitude float64) ([]Geocoding, error) {
	geo := []Geocoding{}
	return geo, s.Client.Do(http.MethodPost, geocoderEndpoint, &geocodeRequest{Addresses: fmt.Sprintf("%f,%f", latitude, longitude), Format: "json"}, &geo)
}

type RapidService struct {
	client *route4me.Client
}

func NewRapidService(APIKey string) *RapidService {
	return &RapidService{
		client: route4me.NewClientWithOptions(APIKey, route4me.DefaultTimeout, RapidBaseURL),
	}
}

func NewRapidServiceWithOptions(APIKey string, timeout time.Duration, baseURL string) *RapidService {
	return &RapidService{
		client: route4me.NewClientWithOptions(APIKey, timeout, baseURL),
	}
}

func (s *RapidService) GetAddresses() ([]Address, error) {
	response := []Address{}
	return response, s.client.Do(http.MethodGet, streetDataEndpoint, nil, &response)
}

func (s *RapidService) GetSingleAddress(pk int) (*Address, error) {
	response := &Address{}
	return response, s.client.Do(http.MethodGet, streetDataEndpoint+strconv.Itoa(pk), nil, response)
}

func (s *RapidService) GetLimitedAddresses(limit int, offset int) ([]Address, error) {
	response := []Address{}
	return response, s.client.Do(http.MethodGet, fmt.Sprintf("%s%d/%d/", streetDataEndpoint, offset, limit), nil, &response)
}

func (s *RapidService) GetAddressesByZipcode(zipcode string) ([]Address, error) {
	response := []Address{}
	return response, s.client.Do(http.MethodGet, streetDataEndpoint+"zipcode/"+zipcode, nil, &response)
}

func (s *RapidService) GetLimitedAddressesByZipcode(zipcode string, limit int, offset int) ([]Address, error) {
	response := []Address{}
	return response, s.client.Do(http.MethodGet, fmt.Sprintf("%s%s/%s/%d/%d/", streetDataEndpoint, "zipcode", zipcode, offset, limit), nil, &response)
}

func (s *RapidService) GetAddressesByZipcodeAndHousenumber(zipcode string, housenumber int) ([]Address, error) {
	response := []Address{}
	return response, s.client.Do(http.MethodGet, fmt.Sprintf("%s%s/%s/%d/", streetDataEndpoint, "service", zipcode, housenumber), nil, &response)
}

func (s *RapidService) GetLimitedAddressesByZipcodeAndHousenumber(zipcode string, housenumber int, limit int, offset int) ([]Address, error) {
	response := []Address{}
	return response, s.client.Do(http.MethodGet, fmt.Sprintf("%s%s/%s/%d/%d/%d/", streetDataEndpoint, "service", zipcode, housenumber, offset, limit), nil, &response)
}
