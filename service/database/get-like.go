package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) GetLike(uid Id, photoId Id) (like *User, err error) {
	var u User

	if err = db.c.QueryRow("SELECT uid, username FROM Likes, Users WHERE user = ? AND photo = ?", uid, photoId).Scan(&u.Uid, &u.Username); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return like, nil
		} else {
			return nil, err
		}
	}

	like = &u
	return like, err

}
