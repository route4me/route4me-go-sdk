package orders

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
	order := &Order{CachedLatitude: 48.335991, CachedLongitude: 31.18287, Address1: "1358 E Luzerne St, Philadelphia, PA 19124, US", AddressAlias: "Auto test address"}
	newOrder, err := service.Add(order)
	if err != nil {
		t.Error(err)
		return
	}
	order.Created = newOrder.Created
	order.ID = newOrder.ID
	order.DateAdded = newOrder.DateAdded
	order.CustomData = newOrder.CustomData
	order.CurbsideLatitude = newOrder.CurbsideLatitude
	order.CurbsideLongitude = newOrder.CurbsideLongitude
	order.Pending = newOrder.Pending
	order.MemberID = newOrder.MemberID
	if !reflect.DeepEqual(*newOrder, *order) {
		t.Error("Orders do not match")
	}
}

func TestIntegrationGet(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}

	orders, _, err := service.GetAll(&Query{})
	if err != nil {
		t.Error(err)
		return
	}
	if len(orders) < 1 {
		t.Skip("Not enough orders to test get 1.")
	}
	order, err := service.Get(orders[0].ID)
	if err != nil {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(order, &orders[0]) {
		t.Error("Orders do not match")
	}
}

func TestIntegrationRemove(t *testing.T) {
	//t.Skip("Skipping Removal integration test. Looks like the endpoint is broken.")
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	orders, _, err := service.GetAll(&Query{})
	if err != nil {
		t.Error(err)
		return
	}
	if len(orders) < 1 {
		t.Skip("Not enough orders to test remove.")
	}
	success, err := service.Delete([]uint64{orders[0].ID})
	if err != nil {
		t.Error(err)
		return
	}
	if !success {
		t.Error("Deleting order failed")
		return
	}
}

func TestIntegrationUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	orders, _, err := service.GetAll(&Query{})
	if err != nil {
		t.Error(err)
		return
	}
	if len(orders) < 1 {
		t.Skip("Not enough avoidance zones to test remove.")
	}
	orders[0].Address1 = "Some random" + strconv.Itoa(rand.Int())
	order, err := service.Update(&orders[0])
	if err != nil {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(&orders[0], order) {
		t.Error("Zones do not equal")
	}
}
