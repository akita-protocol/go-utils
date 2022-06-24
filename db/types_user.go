package db

import (
	"database/sql"

	"gitlab.keelakay.com/keelakay/utils/errors"
)

type User struct {
	ID        string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	DiscordID string `structs:"discord_id,omitempty" db:"discord_id" json:"discord_id,omitempty"`
}

func GetUserByID[H DBStruct](h H, id string) (*User, error) {
	const op errors.Op = "GetUserByID"

	const query = "select id, discord_id from user where id = ?"
	var u User

	err := h.Get(&u, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "User Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &u, nil
}

func GetUserByDiscordID[H DBStruct](h H, did uint64) (*User, error) {
	const op errors.Op = "GetUserByDiscordID"

	const query = "select id, discord_id from user where discord_id = ?"
	var u User

	err := h.Get(&u, query, did)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "User Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &u, nil
}

type UserNFD struct {
	UserID string `structs:"user_id,omitempty" db:"user_id" json:"user_id,omitempty"`
	NFDID  string `structs:"nfd_id" db:"nfd_id" json:"nfd_id,omitempty"`
}

func GetUserNFDS[H DBStruct](h H, userID string) (*[]UserNFD, error) {
	const op errors.Op = "GetUserNFDS"

	const query = "select user_id, nfd_id from user_nfd where user_id = ?"
	var list []UserNFD

	err := h.Select(&list, query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "User NFD Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}
