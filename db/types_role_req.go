package db

import (
	"database/sql"

	"gitlab.keelakay.com/keelakay/utils/errors"
)

type RoleReq struct {
	RoleID      string `structs:"role_id,omitempty" db:"role_id" json:"role_id,omitempty"`
	GroupID     string `structs:"group_id,omitempty" db:"group_id" json:"group_id,omitempty"`
	GreaterThan uint64 `structs:"greater_than,omitempty" db:"greater_than" json:"greater_than,omitempty"`
	LessThan    uint64 `structs:"less_than,omitempty" db:"less_than" json:"less_than,omitempty"`
}

func GetRoleReqByRoleID[H DBStruct](h H, roleID string) (*RoleReq, error) {
	const op errors.Op = "GetRoleReqByRoleID"
	const query = "select role_id, group_id, greater_than, less_than from role_req where role_id = ?"
	var r RoleReq

	err := h.Get(&r, query, roleID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "GetRoleReqs Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &r, nil
}

func GetRoleReqsByGroupID[H DBStruct](h H, groupID string) (*[]RoleReq, error) {
	const op errors.Op = "GetRoleReqsByGroupID"
	const query = "select role_id, group_id, greater_than, less_than from role_req where group_id = ?"
	var list []RoleReq

	err := h.Select(&list, query, groupID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "GetRoleReqs Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}
