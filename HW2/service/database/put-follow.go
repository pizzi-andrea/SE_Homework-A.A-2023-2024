package database

import (
	"database/sql"
	"errors"
)

// PutFollow allow user to follow other user. Operation terminate when added follow or user from just follow to user. In
// any other cases function return false value for r and not nil error value for err
func (db *appdbimpl) PutFollow(from Id, to Id) (r bool, err error) {

	err = db.c.QueryRow("INSERT OR IGNORE INTO Followers (from_, to_) VALUES (?, ?) RETURNING *", from, to).Scan(&from, &to)

	if err == nil {
		r = true
		return
	}

	if errors.Is(err, sql.ErrNoRows) {
		err = nil

	}
	r = false

	return
}
