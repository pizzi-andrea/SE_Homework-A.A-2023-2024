package database

import (
	"database/sql"
)

/*
returns users based on the specified username.
If largeSearch is false then the user who has exactly the specified username is returned,
otherwise all users who have a similar username are returned.

To return all users of the system set username="" and largeSearch=true
*/
func (db *appdbimpl) GetUsers(username Username, largeSearch bool) (users []User, err error) {
	var rows *sql.Rows
	var uid Id

	if largeSearch {
		rows, err = db.c.Query("SELECT * FROM Users WHERE username LIKE '%" + username + "%'")
	} else {
		rows, err = db.c.Query("SELECT * FROM Users WHERE username = ?", username)
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
			return users, err
		}
		users = append(users, User{
			Uid:      uid,
			Username: username,
		})
	}

	err = rows.Err()
	return

}
