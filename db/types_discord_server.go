package db

import (
	"database/sql"
	"fmt"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type DiscordServer struct {
	ID   string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	Name string `structs:"name,omitempty" db:"name" json:"name,omitempty"`
}

func GetDiscordServer[H DBStruct](h H, id string) (*DiscordServer, error) {
	const op errors.Op = "GetServer"
	query := fmt.Sprintf("select id, name from %s.server where id = ?", DiscordBotDatabase())
	var s DiscordServer

	err := h.Get(&s, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Server Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &s, nil
}

func GetDiscordServers[H DBStruct](h H) (*[]DiscordServer, error) {
	const op errors.Op = "GetServers"
	query := fmt.Sprintf("select id, name from %s.server", DiscordBotDatabase())
	var list []DiscordServer

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

func GetDiscordServersWithRolesUsingAsset[H DBStruct](h H, assetID uint64) (*[]DiscordServer, error) {
	const op errors.Op = "GetServersWithRolesUsingAsset"
	query := fmt.Sprintf("select id, name from %s.server where id in (select distinct(server_id) from %s.server_role where id in (select role_id from %s.role_req where group_id in (select id from %s.asset_group where asset_id = ?)))", DiscordBotDatabase(), DiscordBotDatabase(), DiscordBotDatabase(), DiscordBotDatabase())
	var list []DiscordServer

	err := h.Select(&list, query, assetID)
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
