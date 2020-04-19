package database

import (
	"database/sql"
	. "golangcustomtype/customtypes"
)

// func (db sql.DB)InsertToDatabase(date *Date) error {

// 	query := "INSERT INTO public.t_usertypes VALUES ()"
// 	db.Exec()

// }

type dbPointer struct {
	*sql.DB
}

func (db *dbPointer) ReadEmail(id uint) (*Email, error) {

	query := `SELECT "data" FROM public.t_usertypes where id = &1`

	row := db.QueryRow(query, id)

	email := &Email{}

	if err := row.Scan(email); err != nil {
		return &Email{}, err
	}

	return email, nil
}
