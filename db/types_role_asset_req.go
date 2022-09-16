package db

import (
	"database/sql"

	"gitlab.keelakay.com/keelakay/utils/errors"
)

type RoleAssetReq struct {
	RoleID  string `structs:"role_id,omitempty" db:"role_id" json:"role_id,omitempty"`
	AssetID uint64 `structs:"asset_id,omitempty" db:"asset_id" json:"asset_id,omitempty"`
	Amount  uint64 `structs:"amount,omitempty" db:"amount" json:"amount,omitempty"`
}

func GetRoleAssetReqs[H DBStruct](h H, roleID string) (*[]RoleAssetReq, error) {
	const op errors.Op = "GetRoleAssetReqs"
	const query = "select role_id, asset_id, amount from role_asset_req where role_id = ?"
	var list []RoleAssetReq

	err := h.Select(&list, query, roleID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "RoleAssetReq Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}
