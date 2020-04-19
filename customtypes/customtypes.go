package customtypes

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type Email struct {
	Username string
	Domain   string
	Valid    bool
}

func (email *Email) String() string {
	if email.Valid {
		return fmt.Sprintf("%s@%s", email.Username, email.Domain)
	}
	return "Oh!! such empty"
}

//Scan method for type Email
func (email *Email) Scan(value interface{}) error {

	if value == nil {
		*email = Email{Valid: false}
		return nil
	}

	mid := strings.Index(value.(string), "@")
	username := value.(string)[:mid]
	domain := value.(string)[mid+1:]
	*email = Email{Username: username, Domain: domain, Valid: true}
	return nil
}

//Value method for type Email
func (email *Email) Value() (driver.Value, error) {
	if !email.Valid {
		return nil, nil
	}
	return email.String(), nil
}


