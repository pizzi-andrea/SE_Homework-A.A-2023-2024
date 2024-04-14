package database

import (
	"database/sql"
	"errors"
)

// GetPost return a single post create by user. If post not exist will return
// null post pointer and nil error value. If post exist will return it. If occurend an error will return nil pointer and
// error value
func (db *appdbimpl) GetPost(photoId Id) (post *Post, err error) {
	var p Post

	if p.Comments, err = db.GetComments(photoId, "", false); err != nil {
		return nil, err
	}

	if p.Likes, err = db.GetLikes(photoId); err != nil {
		return nil, err
	}

	err = db.c.QueryRow("SELECT photoId, owner, descriptionImg, timeUpdate  FROM Photos WHERE photoId = ?", photoId).Scan(&p.Refer, &p.Owner, &p.DescriptionImg, &p.TimeUpdate)

	if errors.Is(err, sql.ErrNoRows) {
		err = nil
	} else if err != nil {
		return nil, err
	}

	post = &p
	return

}
