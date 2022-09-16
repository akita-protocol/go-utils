package db

import (
	"database/sql"

	"gitlab.keelakay.com/keelakay/utils/errors"
)

type Server struct {
	ID   string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	Name string `structs:"name,omitempty" db:"name" json:"name,omitempty"`
}

func GetServer[H DBStruct](h H, id string) (*Server, error) {
	const op errors.Op = "GetServer"
	const query = "select id, name from server where id = ?"
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
	const query = "select id, name from server"
	var list []Server

	err := h.Select(&list, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Servers Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}
