package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) GetUserBanned(from, to Id) (banned *User, err error) {

	var ban User

	if err = db.c.QueryRow(`SELECT uid, username FROM Users u, Bans b 
	WHERE b.from_ = ? AND u.uid = b.to_ AND b.to_ = ? `, from, to).Scan(&ban.Uid, &ban.Username); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	banned = &ban
	return banned, err

}
