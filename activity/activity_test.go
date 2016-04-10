package activity

import (
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
		RouteID: "D2B71CDCA0550779664952407DFF8712",
	}
	_, err := service.Get(query)
	if err != nil {
		t.Error(err)
	}
}
