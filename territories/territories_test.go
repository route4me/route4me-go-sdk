package territories

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/route4me/route4me-go-sdk"
)

var client = route4me.NewClient("11111111111111111111111111111111")
var service = &Service{Client: client}

func TestIntegrationAddAvoidanceZone(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	zone := &AvoidanceZone{
		Name:  "John" + strconv.Itoa(rand.Int()),
		Color: "beeeee",
		Territory: TerritoryShape{
			Type: Circle,
			Data: []string{"37.569752822786455,-77.47833251953125", "5000"},
		},
	}
	_, err := service.AddAvoidanceZone(zone)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationGetAvoidanceZones(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}

	zones, err := service.GetAvoidanceZones(&Query{})
	if err != nil {
		t.Error(err)
		return
	}
	if len(zones) < 1 {
		t.Skip("Not enough avoidance zones to test get 1.")
	}
	_, err = service.GetAvoidanceZone(&Query{ID: zones[0].ID})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationRemoveAvoidanceZone(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	zones, err := service.GetAvoidanceZones(&Query{})
	if err != nil {
		t.Error(err)
		return
	}
	if len(zones) < 1 {
		t.Skip("Not enough avoidance zones to test remove.")
	}
	err = service.DeleteAvoidanceZone(&Query{ID: zones[0].ID})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationUpdateAvoidanceZone(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	zones, err := service.GetAvoidanceZones(&Query{})
	if err != nil {
		t.Error(err)
		return
	}
	if len(zones) < 1 {
		t.Skip("Not enough avoidance zones to test remove.")
	}
	zones[0].Name = "Johny" + strconv.Itoa(rand.Int())
	_, err = service.UpdateAvoidanceZone(&zones[0])
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationAddTerritory(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	zone := &Territory{
		Name:  "John" + strconv.Itoa(rand.Int()),
		Color: "beeeee",
		Territory: TerritoryShape{
			Type: Circle,
			Data: []string{"37.569752822786455,-77.47833251953125", "5000"},
		},
	}
	_, err := service.AddTerritory(zone)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationGetTerritories(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}

	zones, err := service.GetTerritories(&Query{})
	if err != nil {
		t.Error(err)
		return
	}
	if len(zones) < 1 {
		t.Skip("Not enough territories to test get 1.")
	}
	_, err = service.GetTerritory(&Query{ID: zones[0].ID})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationRemoveTerritories(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	zones, err := service.GetTerritories(&Query{})
	if err != nil {
		t.Error(err)
		return
	}
	if len(zones) < 1 {
		t.Skip("Not enough territories to test remove.")
	}
	err = service.DeleteTerritory(&Query{ID: zones[0].ID})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationUpdateTerritory(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	zones, err := service.GetTerritories(&Query{})
	if err != nil {
		t.Error(err)
		return
	}
	if len(zones) < 1 {
		t.Skip("Not enough territories to test remove.")
	}
	zones[0].Name = "Johny" + strconv.Itoa(rand.Int())
	_, err = service.UpdateTerritory(&zones[0])
	if err != nil {
		t.Error(err)
		return
	}
}
