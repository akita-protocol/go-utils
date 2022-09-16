package db

import (
	"database/sql"

	"gitlab.keelakay.com/keelakay/utils/errors"
)

type ServerRole struct {
	ServerID string `structs:"server_id,omitempty" db:"server_id" json:"server_id,omitempty"`
	RoleID   string `structs:"role_id,omitempty" db:"role_id" json:"role_id,omitempty"`
}

func GetServerRoles[H DBStruct](h H, serverID string) (*[]ServerRole, error) {
	const op errors.Op = "GetServerRoles"
	const query = "select server_id, role_id from server_role where server_id = ?"
	var list []ServerRole

	err := h.Select(&list, query, serverID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "ServerRole Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}
