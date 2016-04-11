package route4me

import (
	"reflect"
	"testing"
)

type testStruct struct {
	String    string  `http:"string"`
	Int       int     `http:"int"`
	Int16     int32   `http:"int16"`
	Int32     int32   `http:"int32"`
	Int64     int64   `http:"int64"`
	UInt      uint    `http:"uint"`
	UInt16    int32   `http:"uint16"`
	UInt32    int32   `http:"uint32"`
	UInt64    int64   `http:"uint64"`
	Float32   float32 `http:"float32"`
	Float64   float64 `http:"float64"`
	Float64E  float64 `http:"float64e"`
	NoConvert string
}

func TestStructToURLValues(t *testing.T) {
	test := &testStruct{
		String:    "foo",
		Int:       423,
		Int16:     552,
		Int32:     -12555,
		Int64:     -423424,
		UInt:      45,
		UInt16:    52,
		UInt32:    155,
		UInt64:    4242,
		Float64:   -4242,
		Float32:   15,
		NoConvert: "hey",
	}
	urlValues := structToURLValues(test)
	if urlValues.Get("NoConvert") != "" {
		t.Error("Value that should not be converted somehow ended up being.")
	}
	if urlValues.Get("float64") != "-4242.0000" {
		t.Error("Wrong float64 value.")
	}
	if urlValues.Get("uint") != "45" {
		t.Error("Wrong uint value.")
	}
	if urlValues.Get("string") != "foo" {
		t.Error("Wrong string value.")
	}
	if urlValues.Get("float32") != "15.0000" {
		t.Error("Wrong float32 value.")
	}
	if urlValues.Get("float64e") != "" {
		t.Error("Empty value marshalled")
	}
}

func TestIsEmptyValue(t *testing.T) {
	if isEmptyValue(reflect.ValueOf(true)) || !isEmptyValue(reflect.ValueOf(false)) {
		t.Error("Boolean empty value check failed")
	}
	if isEmptyValue(reflect.ValueOf("e")) || !isEmptyValue(reflect.ValueOf("")) {
		t.Error("String empty value check failed")
	}
	if isEmptyValue(reflect.ValueOf(5)) || isEmptyValue(reflect.ValueOf(-2.5)) || !isEmptyValue(reflect.ValueOf(-0)) {
		t.Error("Numbers empty value check failed")
	}
	if !isEmptyValue(reflect.ValueOf((*testStruct)(nil))) || isEmptyValue(reflect.ValueOf(&testStruct{String: "abc"})) {
		t.Error("Interface empty value check failed")
	}
}
