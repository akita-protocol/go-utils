package db

import (
	"database/sql"

	"gitlab.keelakay.com/keelakay/utils/errors"
)

type SessionToken struct {
	UserID string `structs:"user_id,omitempty" db:"user_id" json:"user_id,omitempty"`
	Token  string `structs:"token,omitempty" db:"token" json:"token,omitempty"`
}

func GetSessionTokenByUserID[H DBStruct](h H, userID string) (*SessionToken, error) {
	const op errors.Op = "GetSessionTokenByUserID"

	const query = "select user_id, token from user where user_id = ?"
	var s SessionToken

	err := h.Get(&s, query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Session Token Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &s, nil
}

func GetSessionTokenByToken[H DBStruct](h H, token string) (*SessionToken, error) {
	const op errors.Op = "GetSessionTokenByToken"

	const query = "select user_id, token from user where token = ?"
	var s SessionToken

	err := h.Get(&s, query, token)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Session Token Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &s, nil
}
