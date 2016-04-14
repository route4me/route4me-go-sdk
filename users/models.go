package users

import "github.com/route4me/route4me-go-sdk"

type User struct {
	MemberID               int           `json:"member_id,string,omitempty"`
	AccountTypeID          int           `json:"account_type_id,string,omitempty"`
	MemberType             string        `json:"member_type,omitempty"`
	MemberFirstName        string        `json:"member_first_name"`
	MemberLastName         string        `json:"member_last_name"`
	MemberEmail            string        `json:"member_email"`
	PhoneNumber            string        `json:"phone_number"`
	ReadonlyUser           route4me.Bool `json:"readonly_user,omitempty"`
	ShowSuperUserAddresses route4me.Bool `json:"show_superuser_addresses,omitempty"`
}
