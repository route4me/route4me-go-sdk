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
	if len(opts) < 1 {
		t.Skip("Not enough routes to test activity stream")
	}
	opt, err := oService.GetOptimization(&routing.OptimizationParameters{
		ProblemID: opts[0].ProblemID,
	})
	if len(opt.Routes) < 1 {
		t.Skip("Not enough routes to test activity stream")
	}
	query := &Query{
		RouteID: opt.Routes[0].ID,
	}
	_, _, err = service.Get(query)
	if err != nil {
		t.Error(err)
	}
}

func TestIntegrationLog(t *testing.T) {
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
	if len(opts) < 1 {
		t.Skip("Not enough routes to test activity stream")
	}
	opt, err := oService.GetOptimization(&routing.OptimizationParameters{
		ProblemID: opts[0].ProblemID,
	})
	if len(opt.Routes) < 1 {
		t.Skip("Not enough routes to test activity stream")
	}
	err = service.Log("TestMessage", opt.Routes[0].ID)
	if err != nil {
		t.Error(err)
	}
}
