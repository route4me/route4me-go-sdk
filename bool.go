package route4me

import (
	"errors"
	"strings"
)

func UnmarshalJSON(bytes []byte) (bool, error) {
	str := string(bytes)
	if strings.HasPrefix(str, `"`) && strings.HasSuffix(str, `"`) {
		str = str[1 : len(str)-1]
	}
	if strings.ToLower(str) == "true" || str == "1" {
		return true, nil
	} else if strings.ToLower(str) == "false" || str == "0" || str == "null" || str == "" {
		return false, nil
	} else {
		return false, errors.New("Can't unmarshall unknown format to boolean " + str)
	}
}

type Bool bool

func (b *Bool) UnmarshalJSON(bytes []byte) error {
	res, err := UnmarshalJSON(bytes)
	if err != nil {
		return err
	}
	*b = Bool(res)
	return nil
}

func (b *Bool) MarshalJSON() ([]byte, error) {
	if *b == true {
		return []byte("true"), nil
	}
	return []byte("false"), nil
}

type StringBool Bool

func (b *StringBool) UnmarshalJSON(bytes []byte) error {
	res, err := UnmarshalJSON(bytes)
	if err != nil {
		return err
	}
	*b = StringBool(res)
	return nil
}

func (b *StringBool) MarshalJSON() ([]byte, error) {
	if *b == true {
		return []byte("\"TRUE\""), nil
	}
	return []byte("\"FALSE\""), nil
}
