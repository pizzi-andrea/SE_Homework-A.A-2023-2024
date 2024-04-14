package database

import (
	"bytes"
	"encoding/base64"
	"time"
)

// CreatePost allow the users to upload post with photo. Only new post is provided to photo description and image uploaded
// and more addictional informations like:
//
//   - "time to update"
//
//   - list of comments on post
//
//   - location of image uploaded
//
//   - empty lists on comments and likes
func (db *appdbimpl) CreatePost(owner Id, img []byte, description string) (post *Post, err error) {
	var p Post
	s := base64.NewEncoder(base64.RawStdEncoding, bytes.NewBuffer(img))
	defer s.Close()
	if err = db.c.QueryRow("INSERT INTO Photos(owner, descriptionImg, imageData, timeUpdate) VALUES (?, ?, ?, ?) RETURNING *", owner, description, img, time.Now().Format(time.RFC3339)).Scan(&p.Refer, &owner, &p.DescriptionImg, &img, &p.TimeUpdate); err != nil {
		return
	}

	post = &p
	return

}
