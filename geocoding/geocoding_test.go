package geocoding

import (
	"testing"
	"time"

	"encoding/json"

	"github.com/route4me/route4me-go-sdk"
)

var rapidService = NewRapidService("11111111111111111111111111111111")
var service = &Service{Client: route4me.NewClient("11111111111111111111111111111111")}

func TestIntegrationForwardAddress(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	if geo, err := service.ForwardAddress("Los20%Angeles20%International20%Airport,20%CA"); err != nil {
		t.Error(err)
	} else if len(geo) == 0 {
		t.Error("Received empty result set, wrong parameters(?)")
	}
}

func TestIntegrationForwardBulk(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	data := []byte(`[
        {
            "address": "6817 Harrison Rd, Fredericksburg, VA 22407",
            "email": "MirandaJCohen@dayrep.com",
            "username": "Reste1982",
            "web-site": "arridea.com",
            "phone": "404-317-9869",
            "first_name": "Miranda",
            "last_name": "Cohen"
        },
        {
            "address": "7404 Drew Ln, Fredericksburg, VA 22407",
            "email": "WilliamCBennett@rhyta.com",
            "username": "Enton1954",
            "phone": "912-852-2180",
            "first_name": "William",
            "last_name": "Bennett"
        },
        {
            "address": "12316 Willow Woods Dr, Fredericksburg, VA 22407",
            "email": "GeorgeENicholson@armyspy.com",
            "username": "Smis1967",
            "phone": "912-852-2180",
            "first_name": "George",
            "last_name": "Nicholson"
        }
    ]
`)
	rows := []Row{}
	err := json.Unmarshal(data, &rows)
	if err != nil {
		t.Error(err)
	}
	if _, err := service.ForwardBulk(rows); err != nil {
		t.Error(err)
	}
}

func TestIntegrationReverseAddress(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	if geo, err := service.ReverseAddress(33.945705, -118.391105); err != nil {
		t.Error(err)
	} else if len(geo) == 0 {
		t.Error("Received empty result set, wrong parameters(?)")
	}
}

func TestRapidServiceCreation(t *testing.T) {
	rapid := NewRapidService("123")
	if rapid.client.APIKey != "123" {
		t.Error("API key not forwarded to underlying client.")
	}
	if rapid.client.BaseURL != RapidBaseURL {
		t.Error("BaseURL not forwarded to underlying client.")
	}
	rapid = NewRapidServiceWithOptions("1235", 5*time.Second, "https://example.com")
	if rapid.client.APIKey != "1235" {
		t.Error("API key not forwarded to underlying client.")
	}
	if rapid.client.BaseURL != "https://example.com" {
		t.Error("BaseURL not forwarded to underlying client.")
	}
}

func TestIntegrationGetSingleAddress(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	if _, err := rapidService.GetSingleAddress(1); err != nil {
		t.Error(err)
	}
}

func TestIntegrationGetAddresses(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	if _, err := rapidService.GetAddresses(); err != nil {
		t.Error(err)
	}
}

func TestIntegrationGetLimitedAddresses(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	if addrs, err := rapidService.GetLimitedAddresses(10, 5); err != nil {
		t.Error(err)
	} else if len(addrs) != 10 {
		t.Error("Invalid number of addresses returned")
	}
}

func TestIntegrationGetAddressesByZipcode(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	if addrs, err := rapidService.GetAddressesByZipcode("00601"); err != nil {
		t.Error(err)
	} else if len(addrs) == 0 {
		t.Error("Empty result set has been returned")
	}
}

func TestIntegrationGetLimitedAddressesByZipcode(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	if addrs, err := rapidService.GetLimitedAddressesByZipcode("00601", 20, 0); err != nil {
		t.Error(err)
	} else if len(addrs) == 0 {
		t.Error("Empty result set has been returned")
	}
}

func TestIntegrationGetAddressesByZipcodeAndHousenumber(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	if addrs, err := rapidService.GetAddressesByZipcodeAndHousenumber("00601", 17); err != nil {
		t.Error(err)
	} else if len(addrs) == 0 {
		t.Error("Empty result set has been returned")
	}
}

func TestIntegrationGetLimitedAddressesByZipcodeAndHousenumber(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	if _, err := rapidService.GetLimitedAddressesByZipcodeAndHousenumber("00601", 17, 0, 20); err != nil {
		t.Error(err)
	}
}
