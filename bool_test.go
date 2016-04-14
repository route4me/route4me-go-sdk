package route4me

import (
	"encoding/json"
	"testing"
)

func TestBoolUnmarshalling(t *testing.T) {
	str := &struct {
		FromStringInteger Bool `json:"fsi"`
		FromInteger       Bool `json:"fi"`
		FromBoolean       Bool `json:"fb"`
		FromStringBoolean Bool `json:"fsb"`
	}{}
	err := json.Unmarshal([]byte(`{"fsi":"1","fi":1,"fb": true,"fsb":"true"}`), str)
	if err != nil {
		t.Error("Could not unmarshall integer to boolean: ", err)
	}
	if str.FromBoolean != true || str.FromInteger != true || str.FromStringBoolean != true || str.FromStringInteger != true {
		t.Error("Error occured while unmarshalling bool")
	}
	err = json.Unmarshal([]byte(`{"fsi":"0","fi":0,"fb": false,"fsb":"false"}`), str)
	if err != nil {
		t.Error("Could not unmarshall integer to boolean: ", err)
	}
	if str.FromBoolean != false || str.FromInteger != false || str.FromStringBoolean != false || str.FromStringInteger != false {
		t.Error("Error occured while unmarshalling bool")
	}
}

func TestBoolMarshalling(t *testing.T) {
	str := &struct {
		True  Bool `json:"true"`
		False Bool `json:"false"`
	}{True: true, False: false}
	byt, err := json.Marshal(str)
	if err != nil {
		t.Error("Error occured while marshalling bool: ", err)
	}
	if string(byt) != `{"true":true,"false":false}` {
		t.Error("Error occured while marshallign bool.")
	}
}
