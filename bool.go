package route4me

import (
	"errors"
	"strings"
)

type Bool bool

func (b *Bool) UnmarshalJSON(bytes []byte) error {
	str := string(bytes)
	if strings.HasPrefix(str, `"`) && strings.HasSuffix(str, `"`) {
		str = str[1 : len(str)-1]
	}
	if strings.ToLower(str) == "true" || str == "1" {
		*b = true
	} else if strings.ToLower(str) == "false" || str == "0" || str == "null" {
		*b = false
	} else {
		return errors.New("Can't unmarshall unknown format to boolean " + str)
	}
	return nil
}

func (b *Bool) MarshalJSON() ([]byte, error) {
	if *b == true {
		return []byte("true"), nil
	}
	return []byte("false"), nil
}
