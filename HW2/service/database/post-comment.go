package database

import (
	"time"
)

// PutPhoto allow the users to put comments on photos phosted from other users.If comment is created with success will return it
// otherwise will return nil value and error not nil.
func (db *appdbimpl) PostComment(from Id, text string, to Id) (com *Comment, err error) {
	var c Comment

	if err = db.c.QueryRow("INSERT INTO Comments(author, photo, text_, timeStamp_) VALUES (?, ?, ?, ?) RETURNING *", from, to, text, time.Now().Format(time.RFC3339)).Scan(&c.CommentId, &c.Author.Uid, &to, &c.Text, &c.TimeStamp); err != nil {
		return
	}

	if err = db.c.QueryRow("SELECT * FROM Users WHERE uid = ?", c.Author.Uid).Scan(&c.Author.Uid, &c.Author.Username); err != nil {
		return
	}

	com = &c
	return

}
