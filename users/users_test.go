package users

import (
	"testing"

	"github.com/route4me/route4me-go-sdk"
)

var client = route4me.NewClient("11111111111111111111111111111111")
var service = &Service{Client: client}

func TestGetUsers(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	_, err := service.GetUsers()
	if err != nil {
		t.Error(err)
	}
}

func TestAuth(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode.")
	}
	t.Skip("Getting 'invalid character '<' looking for beginning of value'")
	_, err := service.Authenticate("dddd@yahoo.com", "111111")
	if err != nil {
		t.Error(err)
	}
}
