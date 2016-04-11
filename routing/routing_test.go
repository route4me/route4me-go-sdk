package routing

import (
	"math/rand"
	"reflect"
	"strconv"
	"testing"

	"github.com/route4me/route4me-go-sdk"
)

var client = route4me.NewClient("11111111111111111111111111111111")
var service = &Service{Client: client}

func TestIntegrationGetRoute(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	routes, err := service.GetTeamRoutes(&RouteQuery{Limit: 1})
	if err != nil {
		t.Error("Error in external service (GetTeamRoutes): ", err)
		return
	}
	if len(routes) != 1 {
		t.Skip("Not enough routes to run GetAddress")
	}
	_, err = service.GetRoute(&RouteQuery{ID: routes[0].ID})
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

func TestIntegrationGetRouteID(t *testing.T) {
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
	routeID, err := service.GetRouteID(optimizations[0].ProblemID)
	if err != nil {
		t.Error(err)
		return
	}
	if routeID == "" {
		t.Error("Empty route ID.")
	}

	_, err = service.GetRouteID("-")
	if err == nil {
		t.Error("Error not matching the expected one", err.Error())
		return
	}
}

func TestIntegrationDeleteRoutes(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	routes, err := service.GetTeamRoutes(&RouteQuery{Limit: 2})
	if err != nil {
		t.Error(err)
		return
	}
	if len(routes) < 2 {
		t.Skip("Not enough routes to test deleting.")
	}
	deleted, err := service.DeleteRoutes([]string{routes[0].ID, routes[1].ID})
	if err != nil {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(deleted[0].ID, routes[0].ID) || !reflect.DeepEqual(deleted[1].ID, routes[1].ID) {
		t.Error("Deleting routes failed.")
		return
	}
}

func TestIntegrationUpdateRoute(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	routes, err := service.GetTeamRoutes(&RouteQuery{Limit: 1})
	if err != nil {
		t.Error("Error in external service (GetTeamRoutes)", err)
		return
	}
	if len(routes) < 1 {
		t.Skip("Not enough routes to test deleting.")
	}
	get, err := service.GetRoute(&RouteQuery{ID: routes[0].ID})
	if err != nil {
		t.Error("Error in external service (GetRoute)", err)
		return
	}
	get.Parameters.Name = "Updated" + strconv.Itoa(rand.Int())
	route, err := service.UpdateRoute(get)
	if err != nil {
		t.Error(err)
		return
	}
	if route.Parameters.Name != get.Parameters.Name {
		t.Error("Updating route failed")
	}
}

func TestIntegrationDuplicateRoute(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	routes, err := service.GetTeamRoutes(&RouteQuery{Limit: 1})
	if err != nil {
		t.Error("Error in external service (GetTeamRoutes)", err)
		return
	}
	if len(routes) < 1 {
		t.Skip("Not enough routes to test deleting.")
	}
	duplicated, err := service.DuplicateRoute(routes[0].ID)
	if err != nil {
		t.Error(err)
		return
	}
	if duplicated == routes[0].ID {
		t.Error("Duplicated route received thesame ID")
		return
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

func TestIntegrationGetAddress(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	routes, err := service.GetTeamRoutes(&RouteQuery{Limit: 1})
	if err != nil {
		t.Error("Error in external service (GetTeamRoutes): ", err)
		return
	}
	if len(routes) != 1 {
		t.Skip("Not enough routes to run GetAddress")
	}
	get, err := service.GetRoute(&RouteQuery{ID: routes[0].ID})
	if err != nil {
		t.Error(err)
		return
	}
	if len(get.Addresses) < 1 {
		t.Skip("Not enough addresses to run GetAddress")
	}
	_, err = service.GetAddress(&AddressQuery{
		RouteID:            routes[0].ID,
		Notes:              true,
		RouteDestinationID: get.Addresses[0].RouteDestinationID.String(),
	})
	if err != nil {
		t.Error(err)
	}
}

func TestIntegrationUpdateAddress(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	routes, err := service.GetTeamRoutes(&RouteQuery{Limit: 1})
	if err != nil {
		t.Error("Error in external service (GetTeamRoutes): ", err)
		return
	}
	if len(routes) != 1 {
		t.Skip("Not enough routes to run UpdateAddress")
	}
	get, err := service.GetRoute(&RouteQuery{ID: routes[0].ID})
	if err != nil {
		t.Error(err)
		return
	}
	if len(get.Addresses) < 1 {
		t.Skip("Not enough addresses to run UpdateAddress")
	}
	get.Addresses[0].Alias = "Updated" + strconv.Itoa(rand.Int())
	updated, err := service.UpdateAddress(&get.Addresses[0])
	if err != nil {
		t.Error(err)
		return
	}
	if updated.Alias != get.Addresses[0].Alias {
		t.Error("Updated addresses do not equal")
	}
}
