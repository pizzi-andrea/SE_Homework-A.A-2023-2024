package database

import (
	"database/sql"
	"errors"
)

// DelPhoto allow to delete photo update in the system. When photo is deleted all comments and like associated will be deleted.
// if photo deleted with success function return true and nil error for error if photo not exist function return false and nil value
// if occured error function return false and not-nil error value.
func (db *appdbimpl) DelPhoto(id Id) (r bool, err error) {
	_, err = db.c.Exec("DELETE FROM Photos WHERE  photoId = ? ", id)

	if err == nil {
		r = true
		return
	}
	if errors.Is(err, sql.ErrNoRows) {
		err = nil
		r = false
	} else {
		r = false
	}
	return
}
