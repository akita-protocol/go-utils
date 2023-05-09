package db

import (
	"database/sql"
	"fmt"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type SessionToken struct {
	UserID string `structs:"user_id,omitempty" db:"user_id" json:"user_id,omitempty"`
	Token  string `structs:"token,omitempty" db:"token" json:"token,omitempty"`
}

func GetSessionTokenByUserID[H DBStruct](h H, userID string) (*SessionToken, error) {
	const op errors.Op = "GetSessionTokenByUserID"
	query := fmt.Sprintf("select user_id, token from %s.user where user_id = ?", DiscordBotDatabase())
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
	query := fmt.Sprintf("select user_id, token from %s.user where token = ?", DiscordBotDatabase())
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
