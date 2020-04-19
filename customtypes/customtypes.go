package customtypes

import (
	"strings"
)

type Email struct {
	Username string
	Domain   string
	Valid    bool
}

func (email *Email) Scan(value interface{}) error {

	if value == nil {
		*email = Email{Valid: false}
	}

	mid := strings.Index(value.(string), "@")
	username := value.(string)[:mid]
	domain := value.(string)[mid+1:]
	*email = Email{Username: username, Domain: domain, Valid: true}
	return nil
}
