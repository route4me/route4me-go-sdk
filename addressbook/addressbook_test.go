package addressbook

import (
	"reflect"
	"testing"

	"github.com/route4me/route4me-go-sdk"
)

var client = route4me.NewClient("11111111111111111111111111111111")
var service = &Service{Client: client}

func TestUnitGetMinimal(t *testing.T) {

}

func TestUnitGetFull(t *testing.T) {

}

func TestIntegrationGet(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	query := &Query{
		Limit:  10,
		Offset: 0,
	}

	_, err := service.Get(query)
	if err != nil {
		t.Error(err)
	}
}

func TestIntegrationAdd(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	contact := &Contact{
		FirstName:   "John",
		Alias:       "johny",
		Address1:    "Some address",
		CachedLat:   38.024654,
		CachedLng:   -77.338814,
		Email:       "john@smith.com",
		PhoneNumber: "000-000-000",
		StateID:     "5",
		CountryID:   "3",
		City:        "City",
		ZIP:         "00-000",
	}
	newContact, err := service.Add(contact)
	if err != nil {
		t.Error(err)
	}
	contact.ID = newContact.ID
	if !reflect.DeepEqual(contact, newContact) {
		t.Error("Contacts do not equal")
	}
}
