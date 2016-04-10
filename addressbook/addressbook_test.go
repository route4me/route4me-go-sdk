package addressbook

import (
	"math/rand"
	"reflect"
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
	}
	query := &Query{
		Limit:  0,
		Offset: 0,
	}
	_, total, err := service.Get(query)
	if err != nil {
		t.Error(err)
	}
	newContact, err := service.Add(contact)
	if err != nil {
		t.Error(err)
	}
	_, newTotal, err := service.Get(query)
	if err != nil {
		t.Error(err)
	}
	if newTotal-total != 1 {
		t.Error("Tried to add a contact, but number of contacts has not changed - rerun the test, there's a possiblity of others adding contacts to the test account.")
	}
	contact.ID = newContact.ID
	if !reflect.DeepEqual(contact, newContact) {
		t.Error("Contacts do not equal")
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
	contacts, total, err := service.Get(query)
	if err != nil {
		t.Error(err)
		return
	}
	removed, err := service.Delete([]string{strconv.FormatUint(contacts[0].ID, 10)})
	if err != nil {
		t.Error(err)
		return
	}
	_, newTotal, err := service.Get(query)
	if total-newTotal != 1 {
		t.Error("Tried to remove a contact, but number of contacts has not changed - rerun the test, there's a possiblity of others removing contacts from the test account.")
		return
	}
	if !removed {
		t.Error("Did not succeed in removing contacts")
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
	contact, err := service.Update(&contacts[0])
	if err != nil {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(&contacts[0], contact) {
		t.Error("Contacts do not equal")
	}
}
