package database

import (
	"database/sql"
	"errors"
)

// GetUserFromUser give username  and get user that use username if exist. If username not used function
// return nil value and not nil error value for err
func (db *appdbimpl) GetUserFromUser(username Username) (usr *User, err error) {
	var u User
	err = db.c.QueryRow(`
	SELECT * FROM Users 
	WHERE username = ?`, username).Scan(&u.Uid, &u.Username)

	if err == nil {
		usr = &u
	}
	if errors.Is(err, sql.ErrNoRows) {
		err = nil
	}
	return
}
