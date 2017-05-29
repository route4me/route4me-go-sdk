package addressbook

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/route4me/route4me-go-sdk"
)

var client = route4me.NewClient("11111111111111111111111111111111")
var service = &Service{Client: client}

func TestIntegrationGet(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	query := &Query{
		Limit:  10,
		Offset: 0,
	}

	_, _, err := service.Get(query)
	if err != nil {
		t.Error(err)
	}
}

func TestIntegrationAdd(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	contact := &Contact{
		FirstName:   "John" + strconv.Itoa(rand.Int()),
		Alias:       "johny" + strconv.Itoa(rand.Int()),
		Address1:    "Some address" + strconv.Itoa(rand.Int()),
		CachedLat:   38.024654,
		CachedLng:   -77.338814,
		Email:       "john@smith.com",
		PhoneNumber: "000-000-000",
		StateID:     "5",
		CountryID:   "3",
		City:        "City",
		ZIP:         "00-000",
		CurbsideLat: 38.024654,
		CurbsideLng: -77.338814,
		Color:       "fffeee",
	}

	_, err := service.Add(contact)
	if err != nil {
		t.Error(err)
	}
}

func TestIntegrationRemove(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	query := &Query{
		Limit:  1,
		Offset: 0,
	}
	contacts, _, err := service.Get(query)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = service.Delete([]string{strconv.FormatUint(contacts[0].ID, 10)})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	query := &Query{
		Limit:  1,
		Offset: 0,
	}
	contacts, _, err := service.Get(query)
	if err != nil {
		t.Error(err)
		return
	}
	contacts[0].FirstName = "EditedName"
	_, err = service.Update(&contacts[0])
	if err != nil {
		t.Error(err)
		return
	}
}
