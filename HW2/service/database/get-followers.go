package database

import "database/sql"

// GetFollowers allow to list followers of user. When photo is deleted all comments and like associated will be deleted.
// if photo deleted with success function return true and nil error for error if photo not exist function return false and nil value
// if occured error function return false and not-nil error value.
func (db *appdbimpl) GetFollowers(uid Id, username Username, largeSearch bool) (followers []User, err error) {
	var rows *sql.Rows

	if largeSearch {
		rows, err = db.c.Query("SELECT uid, username FROM Followers, Users WHERE to_ = ? AND from_ = uid AND username LIKE '%"+username+"%'", uid)

	} else {
		rows, err = db.c.Query(`SELECT uid, username FROM Followers, Users WHERE to_ = ? AND from_ = uid AND username = ?`, uid, username)

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

		followers = append(followers, User{
			Uid:      uid,
			Username: username,
		})
	}
	return

}
