package database

import (
	"database/sql"
	"errors"
)

// PutBan allow user "from" banned user "to". User "from" not ban his infact "from" will be different from "to".
// If ban ended with success, function return true and nil values, if ban value is just present function return false and nil otherwise
// if occuring errers function return false and error occured
func (db *appdbimpl) PutBan(from, to Id) (r bool, err error) {
	r = false
	err = db.c.QueryRow("INSERT OR IGNORE INTO BANS (from_, to_) VALUES (?, ?) RETURNING *", from, to).Scan(&from, &to)

	if err == nil {
		r = true
		return
	}

	if errors.Is(err, sql.ErrNoRows) {
		err = nil
		return
	}

	return

}
