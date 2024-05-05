package database

import (
	"database/sql"
	"errors"
)

/*
The IsBanned function allows you to check whether user from has banned user a.
*/
func (db *appdbimpl) IsBanned(from Id, to Id) (res bool, err error) {

	err = db.c.QueryRow("SELECT * FROM Bans WHERE from_ = ? AND to_ = ?", from, to).Scan(&from, &to)
	if errors.Is(err, sql.ErrNoRows) {
		res = false
		err = nil
	} else {
		res = true
	}
	return

}
