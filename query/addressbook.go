package query

type AddressBook struct {
	AddressID string `http:"address_id,omitempty"`
	Limit     uint   `http:"limit,omitempty"`
	Offset    uint   `http:"offset,omitempty"`
	Start     uint   `http:"start,omitempty"`
	Query     string `http:"query,omitempty"`
	Fields    string `http:"fields,omitempty"`
	Display   string `http:"display,omitempty"`
}
