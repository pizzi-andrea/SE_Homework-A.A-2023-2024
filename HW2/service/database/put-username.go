package database

/*
SetUsername allow to update username of user identificated by uid value. If username is updated with success
the function return User with new username setted otherwise return nil value for user and not nil value for error variable.
*/
func (db *appdbimpl) SetUsername(uid Id, username string) (usr *User, err error) {
	var u User
	err = db.c.QueryRow("UPDATE Users SET username=? WHERE uid = ? RETURNING *", username, uid).Scan(&u.Uid, &u.Username)
	if err == nil {
		usr = &u
	}

	return

}
