package utils

import (
	"net/url"
	"reflect"
	"strconv"
)

//https://gist.github.com/tonyhb/5819315#file-main-go-L32
// codebeat:disable[ABC]
func StructToURLValues(tag string, i interface{}) url.Values {
	values := url.Values{}
	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)
		if isEmptyValue(f) {
			continue
		}
		tag := typ.Field(i).Tag.Get(tag)
		if tag == "" {
			continue
		}
		// Convert each type into a string for the url.Values string map
		var v string
		switch f.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v = strconv.FormatInt(f.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			v = strconv.FormatUint(f.Uint(), 10)
		case reflect.Float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case reflect.Float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case reflect.String:
			v = f.String()
		}
		values.Set(tag, v)
	}
	return values
}
// codebeat:enable[ABC]

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}
