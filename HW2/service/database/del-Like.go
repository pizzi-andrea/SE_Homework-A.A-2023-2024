package database

import (
	"database/sql"
	"errors"
)

/*
DelLike allows users to remove likes on the photo.
The operation must be idepotent therefore attempting to delete a like whether it exists or not will always give a positive result.
The function will return unsuccessfully if an error occurs
*/
func (db *appdbimpl) DelLike(uid Id, photoId Id) (r bool, err error) {

	err = db.c.QueryRow("DELETE FROM Likes WHERE user = ? AND photo = ? RETURNING *", uid, photoId).Scan(&uid, &photoId)

	if err == nil {
		r = true
	} else if errors.Is(err, sql.ErrNoRows) {
		r = false
		err = nil
	} else {
		r = false
	}

	return

}
