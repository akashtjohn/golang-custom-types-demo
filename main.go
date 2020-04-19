package main

import (
	"fmt"
	"os"

	"golangcustomtype/database"
)

func main() {

	dbCredentials := database.PostgresqlInfo{
		os.Getenv("Dbname"),
		os.Getenv("Host"),
		os.Getenv("Username"),
		os.Getenv("Password"),
		5432}

	db := database.Connect(dbCredentials)
	fmt.Println(db)

	fmt.Println("done !!")
}
