package main

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

// Email type
type Email struct {

	// eg: steve.mcqueen@gmail.com
	Username string // steve.mcqueen
	Domain   string // gmail.com
	Valid    bool   // says the type is valid or not
}

// Person type
type Person struct {
	Name         string `json: "name"`
	EmailAddress Email  `json:"email"`
}

// Stringer implementation for Email
func (em *Email) String() string {
	if em.Valid {
		return fmt.Sprintf("%s@%s", em.Username, em.Domain)
	}
	return ""
}

// MarshalJSON method Email
func (em *Email) MarshalJSON() ([]byte, error) {

	if !em.Valid {
		return []byte("null"), nil
	}
	byteString, err := json.Marshal(em.String())
	return byteString, err
}

// UnmarshalJSON method for Email
func (em *Email) UnmarshalJSON(b []byte) error {

	if string(b) == `null` {
		em.Valid = false
		return nil
	}

	var fullEmailAddress string
	if err := json.Unmarshal(b, &fullEmailAddress); err != nil {
		em.Valid = false
		return err
	}

	mid := strings.Index(fullEmailAddress, "@")
	username := fullEmailAddress[:mid]
	domain := fullEmailAddress[mid+1:]

	*em = Email{Username: username, Domain: domain, Valid: true}

	return nil
}

//Scan method for type Email
func (em *Email) Scan(value interface{}) error {

	// if value in the database is nil, we will create the email obj with Valid flag as false
	if value == nil {
		*em = Email{Valid: false}
		return nil
	}

	mid := strings.Index(value.(string), "@")                    //spilt the string in database by '@'
	username := value.(string)[:mid]                             // find username
	domain := value.(string)[mid+1:]                             //find domain
	*em = Email{Username: username, Domain: domain, Valid: true} // Create email obj pointer
	return nil
}

//Value method for type Email
func (em *Email) Value() (driver.Value, error) {
	if !em.Valid { // if email is not valid return nil (NULL)
		return nil, nil
	}
	return em.String(), nil //for a valid email returns the combined string
}

func testJSON() {

	// json marshalling

	jsonBytes := []byte(`{
		"name": "Steve Mcqueen",
		"email": "steve.mcqueen@gmail.com"
		}`)

	// person pointer
	person := &Person{}
	if err := json.Unmarshal(jsonBytes, person); err != nil {
		fmt.Println(err)
	}

	fmt.Println(*person)

	//////////////////////////////////////////////////////////////////////

	// json marshal

	email := Email{Username: "james.garner", Domain: "gmail.com", Valid: true}
	personObj := &Person{Name: "James Garner", EmailAddress: email}

	MarshaljsonBytes, err := json.Marshal(&personObj)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(MarshaljsonBytes))

}

func testSQL() {

	// Connect to database
	//TODO : Set you database connection string
	db, err := sql.Open("postgres", "<add your connection string>")
	if err != nil {
		fmt.Print(err)
		panic(err)
	}

	///////////////////////////////////////////////////////////////////////

	// read from database
	emailRead := Email{}
	rows := db.QueryRow("SELECT email FROM public.t_usertypes limit 1")

	err = rows.Scan(&emailRead)
	if err != nil {
		panic(err)
	}
	fmt.Println(emailRead)

	////////////////////////////////////////////////////////////////////////

	// insert email
	emailWrite := Email{Username: "charles.bronson", Domain: "gmail.com", Valid: true}

	query := `INSERT INTO public.t_usertypes (email) VALUES ($1) `
	_, err = db.Exec(query, &emailWrite)
	if err != nil {
		panic(err)
	}

}

func main() {

	fmt.Print("running")
	testJSON()
	//testSQL()
}
