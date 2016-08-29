package zones

import (
	"math/rand"
	"reflect"
	"strconv"
	"testing"

	"github.com/route4me/route4me-go-sdk"
)

var client = route4me.NewClient("11111111111111111111111111111111")
var service = &Service{Client: client}

func TestIntegrationAdd(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	zone := &AvoidanceZone{
		Name:  "John" + strconv.Itoa(rand.Int()),
		Color: "beeeee",
		Territory: Territory{
			Type: Circle,
			Data: []string{"37.569752822786455,-77.47833251953125", "5000"},
		},
	}
	newZone, err := service.Add(zone)
	if err != nil {
		t.Error(err)
		return
	}
	zone.ID = newZone.ID
	zone.MemberID = newZone.MemberID
	if !reflect.DeepEqual(newZone, zone) {
		t.Error("Zones do not match")
	}
}

func TestIntegrationGet(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}

	zones, err := service.GetAll(&Query{})
	if err != nil {
		t.Error(err)
		return
	}
	if len(zones) < 1 {
		t.Skip("Not enough avoidance zones to test get 1.")
	}
	zone, err := service.Get(&Query{ID: zones[0].ID})
	if err != nil {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(zone, &zones[0]) {
		t.Error("Zones do not match")
	}
}

func TestIntegrationRemove(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	zones, err := service.GetAll(&Query{})
	if err != nil {
		t.Error(err)
		return
	}
	if len(zones) < 1 {
		t.Skip("Not enough avoidance zones to test remove.")
	}
	err = service.Delete(&Query{ID: zones[0].ID})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	zones, err := service.GetAll(&Query{})
	if err != nil {
		t.Error(err)
		return
	}
	if len(zones) < 1 {
		t.Skip("Not enough avoidance zones to test remove.")
	}
	zones[0].Name = "Johny" + strconv.Itoa(rand.Int())
	contact, err := service.Update(&zones[0])
	if err != nil {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(&zones[0], contact) {
		t.Error("Zones do not equal")
	}
}
