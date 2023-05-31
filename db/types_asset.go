package db

import (
	"database/sql"
	"fmt"
	"strings"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
	"gitlab.algotoolchain.com/keelakay/utils/misc"
)

type AssetStandard string

const (
	AssetStandardASA      AssetStandard = "asa"
	AssetStandardNFD      AssetStandard = "nfd"
	AssetStandardNFTArc69 AssetStandard = "nft-arc69"
	AssetStandardNFTArc3  AssetStandard = "nft-arc3"
	AssetStandardNFTArc19 AssetStandard = "nft-arc19"
	AssetStandardNFTNone  AssetStandard = "nft-none"
	AssetStandardLPToken  AssetStandard = "lp"
)

type Asset struct {
	ID          uint64         `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	Name        *string        `structs:"name,omitempty" db:"name" json:"name,omitempty"`
	URL         *string        `structs:"url,omitempty" db:"url" json:"url,omitempty"`
	ImageURL    *string        `structs:"image_url,omitempty" db:"image_url" json:"image_url,omitempty"`
	Standard    *AssetStandard `structs:"standard,omitempty" db:"standard" json:"standard,omitempty"`
	Mime        *string        `structs:"mime,omitempty" db:"mime" json:"mime,omitempty"`
	UnitName    *string        `structs:"unit_name,omitempty" db:"unit_name" json:"unit_name,omitempty"`
	ClawBack    *string        `structs:"clawback,omitempty" db:"clawback" json:"clawback,omitempty"`
	Freeze      *string        `structs:"freeze,omitempty" db:"freeze" json:"freeze,omitempty"`
	Manager     *string        `structs:"manager,omitempty" db:"manager" json:"manager,omitempty"`
	Reserve     *string        `structs:"reserve,omitempty" db:"reserve" json:"reserve,omitempty"`
	Creator     string         `structs:"creator,omitempty" db:"creator" json:"creator,omitempty"`
	CreatedAt   *uint64        `structs:"created_at,omitempty" db:"created_at" json:"created_at,omitempty"`
	Deleted     bool           `structs:"deleted,omitempty" db:"deleted" json:"deleted,omitempty"`
	DestroyedAt *uint64        `structs:"destroyed_at,omitempty" db:"destroyed_at" json:"destroyed_at,omitempty"`
	Decimals    *uint64        `structs:"decimals,omitempty" db:"decimals" json:"decimals,omitempty"`
	Total       uint64         `structs:"total,omitempty" db:"total" json:"total,omitempty"`
}

var AssetTableKeySlices = map[TableKeySlice][]string{
	Tiny:  {"id", "name", "standard", "url", "image_url", "mime", "unit_name"},
	Small: {"id", "name", "standard", "url", "image_url", "mime", "unit_name", "reserve", "creator", "total"},
	Full:  {"id", "name", "standard", "url", "unit_name", "mime", "clawback", "freeze", "manager", "reserve", "creator", "created_at", "deleted", "destroyed_at", "decimals", "total"},
}

func AssetTableKeys(st TableKeySlice) []string {
	return AssetTableKeySlices[st]
}

func GetAsset[H Handle](h H, id uint64, st TableKeySlice) (*Asset, error) {
	const op errors.Op = "GetAsset"
	query := fmt.Sprintf("select %s from %s.asset where id = ?", strings.Join(AssetTableKeys(st), ","), NFDDatabase())

	var a Asset
	err := h.Get(&a, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Asset Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &a, nil
}

func GetAllAssetIDs[H Handle](h H) ([]uint64, error) {
	const op errors.Op = "GetAllAssetIDs"
	query := fmt.Sprintf("select id from %s.asset", NFDDatabase())
	var list []uint64
	err := h.Select(&list, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Asset Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return list, nil
}

func GetAssets[H Handle](h H, ids []uint64, st TableKeySlice) (*[]Asset, error) {
	const op errors.Op = "GetAssets"
	query := fmt.Sprintf("select %s from %s.asset where id in (%s)", strings.Join(AssetTableKeys(st), ","), NFDDatabase(), strings.Repeat("?, ", len(ids))[0:(len(ids)*3)-2])

	var list []Asset
	err := h.Select(&list, query, misc.ToInterfaceSlice(ids)...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Asset Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetAssetsByDetails[H Handle](h H, creators []string, prefixes []string, include []uint64, exclude []uint64, st TableKeySlice) (*[]Asset, error) {
	const op errors.Op = "GetAssetsByDetails"
	data := []interface{}{}
	query := fmt.Sprintf(
		"select %s from %s.asset where",
		strings.Join(AssetTableKeys(st), ","),
		NFDDatabase(),
	)

	if !(len(creators) > 0) {
		return nil, errors.E(pkg, op, errors.Database, fmt.Errorf("no creators provided"))
	}

	data = append(data, misc.ToInterfaceSlice(creators)...)
	// data = append(data, fmt.Sprintf("%v%%", prefix))
	creatorsQMarks := []rune(strings.Repeat("?, ", len(creators)))
	query += fmt.Sprintf(" creator in (%s)", string(creatorsQMarks[0:len(creatorsQMarks)-2]))

	query += " and ("

	if len(prefixes) > 0 {
		wildcardPrefixes := []string{}
		for i, prefix := range prefixes {
			wildcardPrefixes = append(wildcardPrefixes, fmt.Sprintf("%v%%", prefix))
			if i == 0 {
				query += " unit_name like ?"
			} else {
				query += " or unit_name like ?"
			}
		}
		data = append(data, misc.ToInterfaceSlice(wildcardPrefixes)...)
	}

	if len(include) > 0 {
		data = append(data, misc.ToInterfaceSlice(include)...)
		includeQMarks := []rune(strings.Repeat("?, ", len(include)))
		query += fmt.Sprintf(" or id in(%s)", string(includeQMarks[0:len(includeQMarks)-2]))
	}

	query += ")"

	if len(exclude) > 0 {
		data = append(data, misc.ToInterfaceSlice(exclude)...)
		excludeQMarks := []rune(strings.Repeat("?, ", len(exclude)))
		query += fmt.Sprintf(" and id not in(%s)", string(excludeQMarks[0:len(excludeQMarks)-2]))
	}

	query += " order by id asc"

	var list []Asset
	err := h.Select(&list, query, data...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Asset Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

// func GetAssetsByCollectionIDAndUnitPrefix[H Handle](h H, collectionID string, prefix string, st TableKeySlice) ([]*Asset, error) {
// 	const op errors.Op = "GetAssetsByCollectionID"
// 	query := fmt.Sprintf("select %s from %s.asset where collection_id = ? and unit_name like ?", strings.Join(AssetTableKeys(st), ","), NFDDatabase())

// 	var list []*Asset
// 	err := h.Select(&list, query, collectionID, fmt.Sprintf("%v%%", prefix))
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Asset Not Found")
// 		}
// 		return nil, errors.E(pkg, op, err)
// 	}
// 	return list, nil
// }
