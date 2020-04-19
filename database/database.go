package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //postgres driver : https://stackoverflow.com/questions/52789531/how-do-i-solve-panic-sql-unknown-driver-postgres-forgotten-import/52791919
)

// PostgresqlInfo : Stores the credentials of the database
type PostgresqlInfo struct {
	Dbname   string
	Host     string
	Username string
	Password string
	Port     int
}

// DSN : DNS(Data Source Name) returns database connection string using credentials
func DSN(data PostgresqlInfo) string {

	t := "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
	connectionString := fmt.Sprintf(t, data.Host, data.Port, data.Username, data.Password, data.Dbname)
	return connectionString
}

// Connect : creates connection
// returns connection pointer
func Connect(data PostgresqlInfo) *dbPointer {

	var err error
	SQL, err := sql.Open("postgres", DSN(data))
	if err != nil {
		panic(err)
	}

	err = SQL.Ping()
	if err != nil {
		panic(err)
	}
	return &dbPointer{SQL}

}
