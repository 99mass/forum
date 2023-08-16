package helper

import (
	"database/sql"
	"io/ioutil"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./database/forum.db")
	if err != nil {
		return nil, err
	}
	//defer db.Close()
	return db, nil
}

func CreateTables(db *sql.DB) error {
	schema, err := ioutil.ReadFile("./database/structure.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return err
	}
	
	return nil
}
