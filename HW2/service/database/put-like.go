package database

import (
	"database/sql"
	"errors"
)

/*
PutLike allows users to put likes on the photo.
The operation must be idepotent therefore attempting to put a like whether it exists or not will always give a positive result.
The function will return unsuccessfully if an error occurs
*/
func (db *appdbimpl) PutLike(uid Id, photoId Id) (r bool, err error) {

	err = db.c.QueryRow("INSERT OR IGNORE INTO  Likes(user, photo)  VALUES(?, ?) RETURNING *", uid, photoId).Scan(&uid, &photoId)

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
