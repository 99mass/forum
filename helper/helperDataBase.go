package helper

import "database/sql"

func createDatabase() (*sql.DB, error) {
	db, err:= sql.Open("sqlite3", "../static/rsc/forum.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createTable(db *sql.DB) error {
	_,err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			age INT
		)
	`)
	return err
}

