package main

import (
	"fmt"
	"os"

	. "golangcustomtype/customtypes"
	"golangcustomtype/database"
)

func main() {

	dbCredentials := database.PostgresqlInfo{
		Dbname:   os.Getenv("Dbname"),
		Host:     os.Getenv("Host"),
		Username: os.Getenv("Username"),
		Password: os.Getenv("Password"),
		Port:     5432}

	db := database.Connect(dbCredentials)

	// //fetch null
	// email, err := db.ReadEmail(uint(1))
	// //fetch valid
	// email, err := db.ReadEmail(uint(2))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	//insert email
	//email := &Email{Username: "nitya", Domain: "mahta", Valid: true}
	//insert null email
	email := &Email{Valid: false}
	err := db.InsertToDatabase(email)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(email)

	fmt.Println("done !!")
}
