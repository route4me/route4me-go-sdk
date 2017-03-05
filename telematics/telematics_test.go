package telematics

import (
	"testing"
	"time"
)

var service = NewTelematicsService("11111111111111111111111111111111")

func TestTelematicsServiceCreation(t *testing.T) {
	Telematics := NewTelematicsService("123")
	if Telematics.client.APIKey != "123" {
		t.Error("API key not forwarded to underlying client.")
	}
	if Telematics.client.BaseURL != TelematicsBaseURL {
		t.Error("BaseURL not forwarded to underlying client.")
	}
	Telematics = NewTelematicsServiceWithOptions("1235", 5*time.Second, "https://example.com")
	if Telematics.client.APIKey != "1235" {
		t.Error("API key not forwarded to underlying client.")
	}
	if Telematics.client.BaseURL != "https://example.com" {
		t.Error("BaseURL not forwarded to underlying client.")
	}
}

func TestIntegrationGetVendors(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	if _, err := service.GetVendors(); err != nil {
		t.Error(err)
	}
}

func TestIntegrationGetVendor(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	if _, err := service.GetVendor(153); err != nil {
		t.Error(err)
	}
}

func TestIntegrationSearchVendors(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	if _, err := service.SearchVendors(&VendorQuery{
		Integrated: true,
		Feature:    "Satellite",
		Country:    "GB",
	}); err != nil {
		t.Error(err)
	}
}
