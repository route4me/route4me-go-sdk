package tracking

import (
	"testing"

	"github.com/route4me/route4me-go-sdk"
	"github.com/route4me/route4me-go-sdk/routing"
)

var client = route4me.NewClient("11111111111111111111111111111111")
var service = &Service{Client: client}

func TestSetGPS(t *testing.T) {
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
	query := &GPS{
		RouteID:         opt.Routes[0].ID,
		Latitude:        33.14384,
		Longitude:       -83.22466,
		Course:          1,
		Speed:           120,
		DeviceType:      routing.IPad,
		MemberID:        1,
		DeviceGUID:      "TEST_GPS",
		DeviceTimestamp: "2014-06-14 17:43:35",
	}
	_, err = service.SetGPS(query)
	if err != nil {
		t.Error(err)
		return
	}
}
