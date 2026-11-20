package db

import (
	"database/sql"
	"fmt"

	"github.com/akita-protocol/go-utils/errors"
)

type DiscordUser struct {
	ID        string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	DiscordID string `structs:"discord_id,omitempty" db:"discord_id" json:"discord_id,omitempty"`
}

func GetUserByID[H DBStruct](h H, id string) (*DiscordUser, error) {
	const op errors.Op = "GetUserByID"
	query := fmt.Sprintf("select id, discord_id from %s.user where id = ?", DiscordBotDatabase())
	var u DiscordUser

	err := h.Get(&u, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "User Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &u, nil
}

func GetUserByDiscordID[H DBStruct](h H, did string) (*DiscordUser, error) {
	const op errors.Op = "GetUserByDiscordID"
	query := fmt.Sprintf("select id, discord_id from %s.user where discord_id = ?", DiscordBotDatabase())
	var u DiscordUser

	err := h.Get(&u, query, did)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "User Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &u, nil
}

func GetUserByAddress[H DBStruct](h H, address string) (*DiscordUser, error) {
	const op errors.Op = "GetUserByAddress"
	query := fmt.Sprintf("select id, discord_id from %s.user where id = (select user_id from %s.wallet where address = ? and verified = 1)", DiscordBotDatabase(), DiscordBotDatabase())
	var u DiscordUser

	err := h.Get(&u, query, address)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "User Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &u, nil
}
