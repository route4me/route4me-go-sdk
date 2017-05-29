package users

import (
	"testing"

	"strings"

	"github.com/route4me/route4me-go-sdk"
)

var client = route4me.NewClient("11111111111111111111111111111111")
var service = &Service{Client: client}

func TestIntegrationGetSubusers(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	_, err := service.GetSubusers()
	if err != nil {
		t.Error(err)
	}
}

func TestIntegrationRegisterToWebinar(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	_, err := service.RegisterToWebinar(&WebinarRegistration{
		EmailAddress: "oooooo@yahoo.com",
		FirstName:    "First Name",
		LastName:     "Last Name",
		PhoneNumber:  "454-454544",
		CompanyName:  "Company",
		MemberID:     "123456",
		WebiinarDate: "2016-06-05 10:00:00",
	})
	if err != nil {
		t.Error(err)
	}
}

func TestIntegrationGetUserByID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	users, err := service.GetSubusers()
	if err != nil {
		t.Error(err)
		return
	}
	user := users[len(users)-1]
	_, err = service.GetUserByID(user.ID)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationRegisterDeleteUser(t *testing.T) {
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
	if !resp {
		t.Error("Could not delete account, unknown error occured")
	}
}

func TestIntegrationAuthenticate(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	_, err := service.Authenticate("dddd@yahoo.com", "111111")
	if err != nil {
		t.Error(err)
	}
}

func TestIntegrationValidateSession(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	_, err := service.ValidateSession("4552222222", 787544566)
	if err != nil {
		t.Error(err)
	}
}

func TestIntegrationEditUser(t *testing.T) {
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

func TestIntegrationCreateAccount(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	_, err := service.Create(&Account{
		Industry:  "Gifting",
		FirstName: "Olman",
		LastName:  "Oland",
		Email:     "ololol@outlook.com",
		AcceptTOS: true,
		Password:  "111111",
		Plan:      "enterprise_plan",
	})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationGetConfigValues(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	_, err := service.GetConfigValues()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationAddConfigEntry(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	_, err := service.AddConfigEntry("config-test-key-go", "config-test-value-go")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationGetConfigEntry(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	_, err := service.GetConfigEntry("config-test-key-go")
	if err != nil {
		if strings.Contains(err.Error(), "Specific key has not been found") {
			t.Skip("No key to test against")
		}
		t.Error(err)
		return
	}
}

func TestIntegrationUpdateConfigEntry(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	_, err := service.UpdateConfigEntry("config-test-key-go", "config-test-value-go-1")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationDeleteConfigEntry(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	_, err := service.DeleteConfigEntry("config-test-key-go")
	if err != nil {
		t.Error(err)
		return
	}
}
