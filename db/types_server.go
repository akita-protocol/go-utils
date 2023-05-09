package db

import (
	"database/sql"
	"fmt"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type Server struct {
	ID   string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	Name string `structs:"name,omitempty" db:"name" json:"name,omitempty"`
}

func GetServer[H DBStruct](h H, id string) (*Server, error) {
	const op errors.Op = "GetServer"
	query := fmt.Sprintf("select id, name from %s.server where id = ?", DiscordBotDatabase())
	var s Server

	err := h.Get(&s, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Server Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &s, nil
}

func GetServers[H DBStruct](h H) (*[]Server, error) {
	const op errors.Op = "GetServers"
	query := fmt.Sprintf("select id, name from %s.server", DiscordBotDatabase())
	var list []Server

	err := h.Select(&list, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "servers not found")
		}
		return nil, errors.E(pkg, op, err)
	}

	if !(len(list) > 0) {
		return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, fmt.Errorf("servers not found"))
	}

	return &list, nil
}
