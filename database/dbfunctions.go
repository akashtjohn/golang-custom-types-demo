package database

import (
	"database/sql"
	. "golangcustomtype/customtypes"
)

type dbPointer struct {
	*sql.DB
}

func (db *dbPointer) ReadEmail(id uint) (*Email, error) {

	query := `SELECT "data" FROM public.t_usertypes where id = $1`

	row := db.QueryRow(query, id)

	email := &Email{}

	if err := row.Scan(email); err != nil {
		return &Email{}, err
	}

	return email, nil
}

func (db dbPointer) InsertToDatabase(email *Email) error {

	query := `INSERT INTO public.t_usertypes ("data") VALUES ($1) `
	_, err := db.Exec(query, email)
	if err != nil {
		return err
	}
	return nil
}
