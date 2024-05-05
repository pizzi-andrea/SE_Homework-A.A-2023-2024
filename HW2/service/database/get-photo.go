package database

import (
	"database/sql"
	"errors"
)

// The GetPhoto method allows you to get a photo. If the photo exists it will return the photo otherwise
// it will return a value of zero for the photo.
func (db *appdbimpl) GetPhoto(id Id) (img *Photo, err error) {
	var photo Photo

	if err = db.c.QueryRow("SELECT photoId, imageData FROM Photos WHERE photoId = ?", id).Scan(&photo.PhotoId, &photo.ImageData); err != nil {
		return

	}

	if errors.Is(err, sql.ErrNoRows) {
		err = nil
	}

	img = &photo
	return

}
