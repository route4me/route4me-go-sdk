package vehicles

import (
	"testing"

	"github.com/route4me/route4me-go-sdk"
)

var client = route4me.NewClient("11111111111111111111111111111111")
var service = &Service{Client: client}

func TestIntegrationGetVehicles(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	_, err := service.GetVehicles()
	if err != nil {
		t.Error(err)
	}
}
