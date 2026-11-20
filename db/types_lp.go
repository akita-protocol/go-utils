package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/akita-protocol/go-utils/errors"
	"github.com/akita-protocol/go-utils/misc"
)

type LPType string

const (
	LPTypeTinymanOne       LPType = "tinyman_1"
	LPTypeTinymanOnedotOne LPType = "tinyman_1.1"
	LPTypeTinymanTwo       LPType = "tinyman_2"
	LPTypeHumble           LPType = "humble"
	LPTypePact             LPType = "pact"
)

type LP struct {
	ID       uint64  `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	Type     LPType  `structs:"type,omitempty" db:"type" json:"type,omitempty"`
	AssetOne *uint64 `structs:"asset_one,omitempty" db:"asset_one" json:"asset_one,omitempty"`
	AssetTwo *uint64 `structs:"asset_two,omitempty" db:"asset_two" json:"asset_two,omitempty"`
	Address  string  `structs:"address,omitempty" db:"address" json:"address,omitempty"`
	Issued   *uint64 `structs:"issued,omitempty" db:"issued" json:"issued,omitempty"`
}

func LPTableKeys() []string {
	return []string{"id", "type", "asset_one", "asset_two", "address", "issued"}
}

func GetAllLPs[H Handle](h H) (*[]LP, error) {
	const op errors.Op = "GetAllLPs"
	query := fmt.Sprintf("select %s from %s.lp", strings.Join(LPTableKeys(), ","), NFDDatabase())

	var list []LP
	err := h.Select(&list, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "LP Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &list, nil
}

func GetLP[H Handle](h H, id uint64) (*LP, error) {
	const op errors.Op = "GetLP"
	query := fmt.Sprintf("select %s from %s.lp where id = ?", strings.Join(LPTableKeys(), ","), NFDDatabase())

	var a LP
	err := h.Get(&a, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "LP Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &a, nil
}

func GetLPs[H Handle](h H, ids []uint64) (*[]LP, error) {
	const op errors.Op = "GetLPs"
	query := fmt.Sprintf("select %s from %s.lp where id in (%s)", strings.Join(LPTableKeys(), ","), NFDDatabase(), strings.Repeat("?, ", len(ids))[0:(len(ids)*3)-2])

	var list []LP
	err := h.Select(&list, query, misc.ToInterfaceSlice(ids)...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "LP Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetLPByAsset[H Handle](h H, assetID uint64) (*[]LP, error) {
	const op errors.Op = "GetLPByAsset"
	query := fmt.Sprintf("select %s from %s.lp where asset_one = ? or asset_two = ?", strings.Join(LPTableKeys(), ","), NFDDatabase())

	var list []LP
	err := h.Select(&list, query, assetID, assetID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "LP Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}
