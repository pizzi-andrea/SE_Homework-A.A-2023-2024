package database

import (
	"database/sql"
	"errors"
)

// GetComment allow to get post specificated it id
func (db *appdbimpl) GetComment(commentId Id) (comment *Comment, err error) {
	var c Comment

	err = db.c.QueryRow("SELECT * FROM Comments WHERE commentId = ?", commentId).Scan(&c.CommentId, &c.Author.Uid, &commentId, &c.Text, &c.TimeStamp)

	if errors.Is(err, sql.ErrNoRows) {
		err = nil
		return
	}

	comment = &c
	return

}
