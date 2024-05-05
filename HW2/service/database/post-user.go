package database

// PostUser allow to sig-in into the system. If user is signed with success the method return a point to new user othewer if occured
// error return nil pointer and the error. If user just exist ( other user have just keep username) function return nil pointer e nil error.
func (db *appdbimpl) PostUser(username Username) (usr *User, err error) {
	var u User
	err = db.c.QueryRow("INSERT INTO Users (username) VALUES (?) RETURNING *", username).Scan(&u.Uid, &u.Username)
	if err == nil {
		usr = &u
	}
	return
}
