package database

import (
	"database/sql"
	"errors"
)

/*
The IsFollower function allow to verify if user from follow user to.
*/
func (db *appdbimpl) IsFollower(from Id, to Id) (res bool, err error) {

	err = db.c.QueryRow("SELECT * FROM Followers WHERE from_ = ? AND to_ = ?", from, to).Scan(&from, &to)
	if errors.Is(err, sql.ErrNoRows) {
		res = false
		err = nil
	} else {
		res = true
	}
	return

}
