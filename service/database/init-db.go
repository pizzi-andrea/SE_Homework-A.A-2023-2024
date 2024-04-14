package database

import (
	"bytes"
	"database/sql"
	"os"
	"path/filepath"
)

/*
Init db schema loading sql script called createTable.sql
*/
func initDb(db *sql.DB) (_error error) {
	var path string
	if path, _error = filepath.Abs("service/database/createTable.sql"); _error != nil {
		return
	}

	byteQuery, _error := os.ReadFile(path)
	if _error != nil {
		return
	}

	query := bytes.NewBuffer(byteQuery).String()
	_, _error = db.Exec(query)

	return

}
