package database

import (
	"database/sql"
)

// GetMyStream return user's stream or part of this. The function allow to search photos in stream by username owner, is largeSearch flag is true
// will get all photos of users that have similiar username. If flag is false will return photos of users exactly username gived in input.
func (db *appdbimpl) GetLikes(photoId Id) (likes []User, err error) {
	var rows *sql.Rows
	var u User

	if rows, err = db.c.Query("SELECT uid, username FROM Likes, Users WHERE user = uid AND photo = ?", photoId); err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		if rows.Err() != nil {
			return
		}

		if err = rows.Scan(&u.Uid, &u.Username); err != nil {
			return
		}

		likes = append(likes, u)

	}
	return

}
