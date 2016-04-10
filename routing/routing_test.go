package routing

import (
	"testing"

	"github.com/route4me/route4me-go-sdk"
)

var client = route4me.NewClient("11111111111111111111111111111111")
var service = &Service{Client: client}

func TestIntegrationGetRoute(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	_, err := service.GetRoute(&RouteQuery{ID: "D2B71CDCA0550779664952407DFF8712"})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationGetTeamRoutes(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	_, err := service.GetTeamRoutes(&RouteQuery{Limit: 10, Offset: 5})
	if err != nil {
		t.Error(err)
		return
	}

	//Lets try and get an array as a single route (shouldn't work)
	_, err = service.GetRoute(&RouteQuery{Limit: 10, Offset: 5})
	if err == nil {
		t.Error("Array unmarshalled into a non-array type. This shouldn't happen.")
	}
}
