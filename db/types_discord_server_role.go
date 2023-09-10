package db

import (
	"database/sql"
	"fmt"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type DiscordServerRole struct {
	ID       string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	ServerID string `structs:"server_id,omitempty" db:"server_id" json:"server_id,omitempty"`
	RoleID   string `structs:"role_id,omitempty" db:"role_id" json:"role_id,omitempty"`
	Name     string `structs:"name,omitempty" db:"name" json:"name,omitempty"`
}

func GetServerRoles[H DBStruct](h H, serverID string) (*[]DiscordServerRole, error) {
	const op errors.Op = "GetServerRoles"
	query := fmt.Sprintf("select id, server_id, role_id, name from %s.server_role where server_id = ?", DiscordBotDatabase())
	var list []DiscordServerRole

	err := h.Select(&list, query, serverID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "server_role not found")
		}
		return nil, errors.E(pkg, op, err)
	}

	if !(len(list) > 0) {
		return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, fmt.Errorf("server_role not found"))
	}

	return &list, nil
}
