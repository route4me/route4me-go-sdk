package telematics

import (
	"net/http"
	"time"

	"github.com/route4me/route4me-go-sdk"
)

const (
	TelematicsBaseURL = "https://telematics.route4me.com"
	vendorsEndpoint   = "/api/vendors.php"
)

type TelematicsService struct {
	client *route4me.Client
}

func NewTelematicsService(APIKey string) *TelematicsService {
	return &TelematicsService{
		client: route4me.NewClientWithOptions(APIKey, route4me.DefaultTimeout, TelematicsBaseURL),
	}
}

func NewTelematicsServiceWithOptions(APIKey string, timeout time.Duration, baseURL string) *TelematicsService {
	return &TelematicsService{
		client: route4me.NewClientWithOptions(APIKey, timeout, baseURL),
	}
}

type getVendorsResponse struct {
	Vendors []Vendor `json:"vendors"`
}

func (s *TelematicsService) GetVendors() ([]Vendor, error) {
	response := &getVendorsResponse{}
	return response.Vendors, s.client.Do(http.MethodGet, vendorsEndpoint, nil, &response)
}

type getVendorRequest struct {
	VendorID int `http:"vendor_id"`
}

type getVendorResponse struct {
	Vendor *Vendor `json:"vendor"`
}

func (s *TelematicsService) GetVendor(vendor int) (*Vendor, error) {
	response := &getVendorResponse{}
	return response.Vendor, s.client.Do(http.MethodGet, vendorsEndpoint, &getVendorRequest{VendorID: vendor}, response)
}

func (s *TelematicsService) SearchVendors(query *VendorQuery) ([]Vendor, error) {
	response := &getVendorsResponse{}
	return response.Vendors, s.client.Do(http.MethodGet, vendorsEndpoint, nil, &response)
}

type compareVendorsRequest struct {
	Vendors []int `json:"vendors`
}

func (s *TelematicsService) CompareVendors(vendors ...int) ([]Vendor, error) {
	response := &getVendorsResponse{}
	return response.Vendors, s.client.Do(http.MethodGet, vendorsEndpoint, &compareVendorsRequest{Vendors: vendors}, response)
}
