package routing

import (
	"reflect"
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

func TestIntegrationGetOptimizations(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	_, err := service.GetOptimizations(&RouteQuery{Limit: 5})
	if err != nil {
		t.Error(err)
	}
}

func TestIntegrationGetOptimization(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	optimizations, err := service.GetOptimizations(&RouteQuery{Limit: 1})
	if err != nil {
		t.Error("Error in external function (getOptimizations): ", err)
		return
	}
	if len(optimizations) < 1 {
		t.Skip("Not enough optimizations in the getOptimizations")
	}
	_, err = service.GetOptimization(&OptimizationParameters{ProblemID: optimizations[0].ProblemID})
	if err != nil {
		t.Error(err)
	}
}

func TestIntegrationUpdateOptimization(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	optimizations, err := service.GetOptimizations(&RouteQuery{Limit: 1})
	if err != nil {
		t.Error("Error in external function (getOptimizations): ", err)
		return
	}
	if len(optimizations) < 1 {
		t.Skip("Not enough optimizations in the getOptimizations")
	}

	updated, err := service.UpdateOptimization(&OptimizationParameters{ProblemID: optimizations[0].ProblemID, Reoptimize: true})
	if err != nil {
		t.Error(err)
		return
	}
	get, err := service.GetOptimization(&OptimizationParameters{ProblemID: optimizations[0].ProblemID})
	if err != nil {
		t.Error(err)
		return
	}
	updated.SentToBackground = false
	if !reflect.DeepEqual(get, updated) {
		t.Error("Optimizations do not match")
	}
}

func TestIntegrationRunOptimization(t *testing.T) {

}
