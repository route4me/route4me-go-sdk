package routing

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

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
		t.Skip("Error in external service (GetTeamRoutes): ", err)
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
		t.Skip("Error in external function (getOptimizations): ", err)
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

func TestMarkAddressAsDetectedAndDeparted(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	routes, err := service.GetTeamRoutes(&RouteQuery{Limit: 1})
	if err != nil {
		t.Error(err)
		return
	}
	if len(routes) < 1 {
		t.Skip("Not enough routes to test deleting.")
	}
	route, err := service.GetRoute(&RouteQuery{ID: routes[0].ID})
	if err != nil {
		t.Error(err)
		return
	}
	route.Addresses[0].IsDeparted = true
	_, err = service.UpdateAddress(&route.Addresses[0])
	if err != nil {
		t.Error(err)
		return
	}
}

func TestMarkAddressAsDetectedAndVisited(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	routes, err := service.GetTeamRoutes(&RouteQuery{Limit: 1})
	if err != nil {
		t.Error(err)
		return
	}
	if len(routes) < 1 {
		t.Skip("Not enough routes to test deleting.")
	}
	route, err := service.GetRoute(&RouteQuery{ID: routes[0].ID})
	if err != nil {
		t.Error(err)
		return
	}
	route.Addresses[0].IsVisited = true
	_, err = service.UpdateAddress(&route.Addresses[0])
	if err != nil {
		t.Error(err)
		return
	}
}

func TestMarkAddressAsVisited(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	routes, err := service.GetTeamRoutes(&RouteQuery{Limit: 1})
	if err != nil {
		t.Error(err)
		return
	}
	if len(routes) < 1 {
		t.Skip("Not enough routes to test deleting.")
	}
	route, err := service.GetRoute(&RouteQuery{ID: routes[0].ID})
	if err != nil {
		t.Error(err)
		return
	}
	route.Addresses[0].IsVisited = true
	_, err = service.MarkAddressAsVisited(&route.Addresses[0])
	if err != nil {
		t.Error(err)
		return
	}
}

func TestMarkAddressAsDeparted(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	routes, err := service.GetTeamRoutes(&RouteQuery{Limit: 1})
	if err != nil {
		t.Error(err)
		return
	}
	if len(routes) < 1 {
		t.Skip("Not enough routes to test deleting.")
	}
	route, err := service.GetRoute(&RouteQuery{ID: routes[0].ID})
	if err != nil {
		t.Error(err)
		return
	}
	route.Addresses[0].IsDeparted = true
	_, err = service.MarkAddressAsDeparted(&route.Addresses[0])
	if err != nil {
		t.Error(err)
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
	_, err = service.DeleteRoutes(routes[0].ID, routes[1].ID)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationUpdateRoute(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	routes, err := service.GetTeamRoutes(&RouteQuery{Limit: 1})
	if err != nil {
		t.Skip("Error in external service (GetTeamRoutes)", err)
		return
	}
	if len(routes) < 1 {
		t.Skip("Not enough routes to test deleting.")
	}
	get, err := service.GetRoute(&RouteQuery{ID: routes[0].ID})
	if err != nil {
		t.Skip("Error in external service (GetRoute)", err)
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
		t.Skip("Error in external service (GetTeamRoutes)", err)
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

func TestIntegrationMergeRoutes(t *testing.T) {
	t.Skip("Cannot test under test account. 'Point is not allowed for test account'")
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	routes, err := service.GetTeamRoutes(&RouteQuery{Limit: 2})
	if err != nil {
		t.Error(err)
		return
	}
	if len(routes) < 2 {
		t.Skip("Not enough routes to test merging.")
	}
	err = service.MergeRoutes(&MergeRequest{
		RouteIDs:       routes[0].ID + "," + routes[1].ID,
		RemoveOrigin:   false,
		DepotAddress:   "10180 Dyer St, El Paso, TX 79924, USA",
		DepotLatitude:  31.9061405,
		DepotLongitude: -106.4033899,
	})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationShareRoute(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	routes, err := service.GetTeamRoutes(&RouteQuery{Limit: 1})
	if err != nil {
		t.Error(err)
		return
	}
	if len(routes) < 1 {
		t.Skip("Not enough routes to test sharing.")
	}
	err = service.ShareRoute(routes[0].ID, "oooooo@gmail.com")
	if err != nil {
		t.Error(err)
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
		t.Skip("Error in external function (getOptimizations): ", err)
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

func TestIntegrationDeleteOptimization(t *testing.T) {
	//t.Skip("Currently doesn't work")
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	optimizations, err := service.GetOptimizations(&RouteQuery{Limit: 1})
	if err != nil {
		t.Skip("Error in external function (getOptimizations): ", err)
		return
	}
	if len(optimizations) < 1 {
		t.Skip("Not enough optimizations in the getOptimizations")
	}
	err = service.DeleteOptimization(optimizations[0].ProblemID)
	if err != nil {
		t.Error(err)
	}
}

func TestIntegrationRunOptimization(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	addresses := []*Address{
		&Address{AddressString: "3634 W Market St, Fairlawn, OH 44333", //all possible originating locations are depots, should be marked as true
			//stylistically we recommend all depots should be at the top of the destinations list
			IsDepot:   true,
			Latitude:  41.135762259364,
			Longitude: -81.629313826561,

			//the number of seconds at destination
			Time: 300,

			//together these two specify the time window of a destination
			//seconds offset relative to the route start time for the open availability of a destination
			TimeWindowStart: 28800,

			//seconds offset relative to the route end time for the open availability of a destination
			TimeWindowEnd: 29465},

		&Address{AddressString: "1218 Ruth Ave, Cuyahoga Falls, OH 44221",
			Latitude:        41.135762259364,
			Longitude:       -81.629313826561,
			Time:            300,
			TimeWindowStart: 29465,
			TimeWindowEnd:   30529},

		&Address{AddressString: "512 Florida Pl, Barberton, OH 44203",
			Latitude:        41.003671512008,
			Longitude:       -81.598461046815,
			Time:            300,
			TimeWindowStart: 30529,
			TimeWindowEnd:   33779},

		&Address{AddressString: "512 Florida Pl, Barberton, OH 44203",
			Latitude:        41.003671512008,
			Longitude:       -81.598461046815,
			Time:            100,
			TimeWindowStart: 33779,
			TimeWindowEnd:   33944},

		&Address{AddressString: "3495 Purdue St, Cuyahoga Falls, OH 44221",
			Latitude:        41.162971496582,
			Longitude:       -81.479049682617,
			Time:            300,
			TimeWindowStart: 33944,
			TimeWindowEnd:   34801},

		&Address{AddressString: "1659 Hibbard Dr, Stow, OH 44224",
			Latitude:        41.194505989552,
			Longitude:       -81.443351581693,
			Time:            300,
			TimeWindowStart: 34801,
			IsDepot:         false,
			TimeWindowEnd:   36366},

		&Address{AddressString: "2705 N River Rd, Stow, OH 44224",
			Latitude:        41.145240783691,
			Longitude:       -81.410247802734,
			Time:            300,
			TimeWindowStart: 36366,
			TimeWindowEnd:   39173},

		&Address{AddressString: "10159 Bissell Dr, Twinsburg, OH 44087",
			Latitude:        41.340042114258,
			Longitude:       -81.421226501465,
			Time:            300,
			TimeWindowStart: 39173,
			TimeWindowEnd:   41617},

		&Address{AddressString: "367 Cathy Dr, Munroe Falls, OH 44262",
			Latitude:        41.148578643799,
			Longitude:       -81.429229736328,
			Time:            300,
			TimeWindowStart: 41617,
			TimeWindowEnd:   43660},

		&Address{AddressString: "367 Cathy Dr, Munroe Falls, OH 44262",
			Latitude:        41.148578643799,
			Longitude:       -81.429229736328,
			Time:            300,
			TimeWindowStart: 43660,
			TimeWindowEnd:   46392},

		&Address{AddressString: "512 Florida Pl, Barberton, OH 44203",
			Latitude:        41.003671512008,
			Longitude:       -81.598461046815,
			Time:            300,
			TimeWindowStart: 46392,
			TimeWindowEnd:   48389},

		&Address{AddressString: "559 W Aurora Rd, Northfield, OH 44067",
			Latitude:        41.315116882324,
			Longitude:       -81.558746337891,
			Time:            50,
			TimeWindowStart: 48389,
			TimeWindowEnd:   48449},

		&Address{AddressString: "3933 Klein Ave, Stow, OH 44224",
			Latitude:        41.169467926025,
			Longitude:       -81.429420471191,
			Time:            300,
			TimeWindowStart: 48449,
			TimeWindowEnd:   50152},

		&Address{AddressString: "2148 8th St, Cuyahoga Falls, OH 44221",
			Latitude:        41.136692047119,
			Longitude:       -81.493492126465,
			Time:            300,
			TimeWindowStart: 50152,
			TimeWindowEnd:   51982},

		&Address{AddressString: "3731 Osage St, Stow, OH 44224",
			Latitude:        41.161357879639,
			Longitude:       -81.42293548584,
			Time:            100,
			TimeWindowStart: 51982,
			TimeWindowEnd:   52180},

		&Address{AddressString: "3731 Osage St, Stow, OH 44224",
			Latitude:        41.161357879639,
			Longitude:       -81.42293548584,
			Time:            300,
			IsDepot:         false,
			TimeWindowStart: 52180,
			TimeWindowEnd:   54379},
	}
	routeParams := &RouteParameters{
		AlgorithmType:        CVRP_TW_MD,
		Name:                 "Multiple Depot, Multiple Driver",
		RouteDate:            time.Now().Unix(),
		RouteTime:            60 * 60 * 7,
		RouteMaxDuration:     86400,
		VehicleCapacity:      1,
		VehicleMaxDistanceMI: 1000,
		Optimize:             Distance,
		DistanceUnit:         Miles,
		DeviceType:           Web,
		TravelMode:           Driving,
	}

	optParams := &OptimizationParameters{
		Addresses:  addresses,
		Parameters: routeParams,
	}
	_, err := service.RunOptimization(optParams)
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
		t.Skip("Error in external function (getOptimizations): ", err)
		return
	}
	if len(optimizations) < 1 {
		t.Skip("Not enough optimizations in the getOptimizations")
	}
	_, err = service.UpdateOptimization(&OptimizationParameters{ProblemID: optimizations[0].ProblemID, Reoptimize: true})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationGetAddress(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	routes, err := service.GetTeamRoutes(&RouteQuery{Limit: 1})
	if err != nil {
		t.Skip("Error in external service (GetTeamRoutes): ", err)
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
		t.Skip("Error in external service (GetTeamRoutes): ", err)
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

func TestIntegrationDeleteAddress(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	routes, err := service.GetTeamRoutes(&RouteQuery{Limit: 64})
	if err != nil {
		t.Skip("Error in external service (GetTeamRoutes): ", err)
		return
	}
	if len(routes) == 0 {
		t.Skip("Not enough routes to run DeleteAddress")
	}
	var addr Address
	for i := 0; i < len(routes) && !addr.IsDepot; i++ {
		get, err := service.GetRoute(&RouteQuery{ID: routes[i].ID})
		if err != nil {
			t.Error(err)
			return
		}
		if len(get.Addresses) < 1 {
			continue
		}
		for j := 0; j < len(get.Addresses) && !addr.IsDepot; j++ {
			addr = get.Addresses[j]
		}
	}
	if addr.IsDepot {
		t.Skip("No non-depot addresses to run DeleteAddress test against")
	}
	routeID, err := service.DeleteAddress(addr.OptimizationProblemID, addr.RouteDestinationID.String())
	if err != nil {
		t.Error(err)
		return
	}
	if routeID.String() != addr.RouteDestinationID.String() {
		t.Error("Delete response returned different route destination id: " + routeID.String())
	}
}

func TestIntegrationGetAddressNotes(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	routes, err := service.GetTeamRoutes(&RouteQuery{Limit: 1})
	if err != nil {
		t.Skip("Error in external service (GetTeamRoutes): ", err)
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
	query := &NoteQuery{
		RouteID:   get.ID,
		AddressID: get.Addresses[0].RouteDestinationID.String(),
	}
	_, err = service.GetAddressNotes(query)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationAddAddressNote(t *testing.T) {
	//	t.Skip()
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	routes, err := service.GetTeamRoutes(&RouteQuery{Limit: 1})
	if err != nil {
		t.Skip("Error in external service (GetTeamRoutes): ", err)
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
	query := &NoteQuery{
		RouteID:      get.ID,
		AddressID:    get.Addresses[0].RouteDestinationID.String(),
		Latitude:     33.132675170898,
		Longitude:    -83.244743347168,
		ActivityType: DropOff,
	}
	_, err = service.AddAddressNote(query, "TestNoteContents")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationAddRouteDestinations(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	routes, err := service.GetTeamRoutes(&RouteQuery{Limit: 1})
	if err != nil {
		t.Skip("Error in external service (GetTeamRoutes): ", err)
		return
	}
	if len(routes) != 1 {
		t.Skip("Not enough routes to run UpdateAddress")
	}
	route := routes[0]
	addresses := []Address{
		//address to be added
		Address{
			AddressString: "717 5th Ave New York, NY 10021",
			Alias:         "Giorgio Armani",
			Latitude:      40.7669692,
			Longitude:     -73.9693864,
			Time:          0,
		},
	}
	_, err = service.AddRouteDestinations(route.ID, addresses)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationRemoveRouteDestination(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	routes, err := service.GetTeamRoutes(&RouteQuery{Limit: 1})
	if err != nil {
		t.Skip("Error in external service (GetTeamRoutes): ", err)
		return
	}
	if len(routes) != 1 {
		t.Skip("Not enough routes to run RemoveRouteDestination")
	}
	get, err := service.GetRoute(&RouteQuery{ID: routes[0].ID})
	if err != nil {
		t.Error(err)
		return
	}
	if len(get.Addresses) < 1 {
		t.Skip("Not enough addresses to run RemoveRouteDestination")
	}
	_, err = service.RemoveRouteDestination(get.ID, get.Addresses[0].RouteDestinationID.String())
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationResequenceRoute(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	_, err := service.ResequenceRoute("CA902292134DBC134EAF8363426BD247")
	if err != nil {
		t.Error(err)
	}
}
