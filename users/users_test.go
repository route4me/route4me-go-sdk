package users

import (
	"reflect"
	"testing"

	"github.com/route4me/route4me-go-sdk"
)

var client = route4me.NewClient("11111111111111111111111111111111")
var service = &Service{Client: client}

func TestGetSubusers(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	_, err := service.GetSubusers()
	if err != nil {
		t.Error(err)
	}
}

func TestGetUserByID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	users, err := service.GetSubusers()
	if err != nil {
		t.Error(err)
		return
	}
	user := users[len(users)-1]
	singleUser, err := service.GetUserByID(user.ID)
	if err != nil {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(user, singleUser) {
		t.Error("Users do not match")
	}
}

func TestRegisterDeleteUser(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	account, err := service.Register(&MemberBase{
		FirstName:            "John",
		LastName:             "Example",
		Email:                "newdispatcher+something@gmail.com",
		Type:                 SubAccountDispatcher,
		Password:             "123",
		HideRouteAddresses:   true,
		HideVisitedAddresses: true,
	})
	if err != nil {
		t.Error(err)
		return
	}
	resp, err := service.Delete(account.ID)
	if err != nil {
		t.Error(err)
		return
	}
	if !resp.Status {
		t.Error("Could not delete account, unknown error occured")
	}
}

func TestEditUser(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	users, err := service.GetSubusers()
	if err != nil {
		t.Error(err)
		return
	}
	user := users[len(users)-1]
	user.Phone = "123452"
	member, err := service.Edit(user)
	if err != nil {
		t.Error(err)
		return
	}
	if member.Phone != "123452" {
		t.Error("Edit failed")
		return
	}
}
