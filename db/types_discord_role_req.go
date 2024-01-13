package db

import (
	"database/sql"
	"fmt"
	"strings"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type DiscordRoleReq struct {
	RoleID         string  `structs:"role_id,omitempty" db:"role_id" json:"role_id,omitempty"`
	GroupID        string  `structs:"group_id,omitempty" db:"group_id" json:"group_id,omitempty"`
	GreaterThan    *uint64 `structs:"greater_than,omitempty" db:"greater_than" json:"greater_than,omitempty"`
	LessThan       uint64  `structs:"less_than,omitempty" db:"less_than" json:"less_than,omitempty"`
	DistinctAssets *bool   `structs:"distinct_assets,omitempty" db:"distinct_assets" json:"distinct_assets,omitempty"`
}

func DiscordRoleReqTableKeys() []string {
	return []string{"role_id", "group_id", "greater_than", "less_than", "distinct_assets"}
}

func GetDiscordRoleReqByRoleID[H DBStruct](h H, roleID string) (*DiscordRoleReq, error) {
	const op errors.Op = "GetRoleReqByRoleID"
	query := fmt.Sprintf("select %s from %s.role_req where role_id = ?", strings.Join(DiscordRoleReqTableKeys(), ","), DiscordBotDatabase())
	var r DiscordRoleReq

	err := h.Get(&r, query, roleID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "role_req not found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &r, nil
}

func GetDiscordRoleReqsByGroupID[H DBStruct](h H, groupID string) (*[]DiscordRoleReq, error) {
	const op errors.Op = "GetRoleReqsByGroupID"
	query := fmt.Sprintf("select %s from %s.role_req where group_id = ?", strings.Join(DiscordRoleReqTableKeys(), ","), DiscordBotDatabase())
	var list []DiscordRoleReq

	err := h.Select(&list, query, groupID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "role_reqs not found")
		}
		return nil, errors.E(pkg, op, err)
	}

	if !(len(list) > 0) {
		return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, fmt.Errorf("role_reqs not found"))
	}

	return &list, nil
}
