package activity

import (
	"testing"

	"github.com/route4me/route4me-go-sdk"
	"github.com/route4me/route4me-go-sdk/routing"
)

var client = route4me.NewClient("11111111111111111111111111111111")
var service = &Service{Client: client}

func TestIntegrationGet(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	oService := &routing.Service{Client: client}
	opts, err := oService.GetOptimizations(&routing.RouteQuery{
		Limit: 1,
	})
	if err != nil {
		t.Error("Error occured in external service:", err)
		return
	}
	if len(opts) < 1 || len(opts[0].Routes) < 1 {
		t.Skip("Not enough routes to test setGPS")
	}
	query := &Query{
		RouteID: opts[0].Routes[0].ID,
	}
	_, err = service.Get(query)
	if err != nil {
		t.Error(err)
	}
}
