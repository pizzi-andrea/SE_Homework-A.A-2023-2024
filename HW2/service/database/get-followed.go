package database

import "database/sql"

// GetFollowed allow to list followed by user. Is possible to reserach followed with similar username
func (db *appdbimpl) GetFollowed(uid Id, username Username, largeSearch bool) (followed []User, err error) {
	var rows *sql.Rows

	if largeSearch {
		rows, err = db.c.Query("SELECT uid, username FROM Followers, Users WHERE from_ = ? AND to_ = uid AND username LIKE '%"+username+"%'", uid)

	} else {
		rows, err = db.c.Query(`SELECT uid, username FROM Followers, Users WHERE from_ = ? AND to_ = uid AND username = ?`, uid, username)

	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if rows.Err() != nil {
			return
		}
		if err = rows.Scan(&uid, &username); err != nil {
			return
		}

		followed = append(followed, User{
			Uid:      uid,
			Username: username,
		})
	}
	return

}
