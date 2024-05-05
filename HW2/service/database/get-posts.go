package database

import (
	"database/sql"
)

type OrderBy int

const ( //   sort by value
	timeUpdate OrderBy = iota
)

type Ordering int

const ( //   Ordering parameters
	asc Ordering = iota
	desc
)

/*
GetPhotos accepts the user ID (uid) as input and extracts all photos uploaded by users. Users can charge up or more
photos (deleted photos will not be listed), in fact the photo stream may be returned empty. If a value of zero is returned
for photos, an error occurred
*/
func (db *appdbimpl) GetPosts(uid Id, by []OrderBy, ord ...Ordering) (posts Stream, err error) {
	var photo Photo
	var postPhoto *Post

	var rows *sql.Rows
	var ordy string
	var oord string
	if len(by) == 0 {
		by = append(by, timeUpdate)
	}
	if len(ord) == 0 {
		ord = append(ord, desc)
	}

	switch ord[0] {
	case asc:
		oord = "ASC"
	case desc:
		oord = "DESC"
	default:
		oord = "DESC"
	}

	switch by[0] {
	case timeUpdate:
		ordy = "timeUpdate"

	}

	if rows, err = db.c.Query("SELECT photoId FROM Photos WHERE owner = ? ORDER BY ?, ?", uid, ordy, oord); err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if rows.Err() != nil {
			return
		}
		if err = rows.Scan(&photo.PhotoId); err != nil {
			return
		}

		if postPhoto, err = db.GetPost(photo.PhotoId); err != nil {
			return
		}

		posts = append(posts, *postPhoto)

	}
	return

}
