package db

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/akita-protocol/go-utils/errors"
)

type StakingSnapshotAsset struct {
	SnapshotID string  `structs:"snapshot_id,omitempty" db:"snapshot_id" json:"snapshot_id,omitempty"`
	UserID     string  `structs:"user_id,omitempty" db:"user_id" json:"user_id,omitempty"`
	AssetID    uint64  `structs:"asset_id,omitempty" db:"asset_id" json:"asset_id,omitempty"`
	Quantity   *uint64 `structs:"quantity,omitempty" db:"quantity" json:"quantity,omitempty"`
}

func SnapshotAssetTableKeys() []string {
	return []string{"snapshot_id", "user_id", "asset_id", "quantity"}
}

func GetSnapshotAssetsByID[H Handle](h H, snapshotID string) (*[]StakingSnapshotAsset, error) {
	const op errors.Op = "GetSnapshotAsset"
	query := fmt.Sprintf("select %s from snapshot_asset where snapshot_id = ?", strings.Join(SnapshotAssetTableKeys(), ","))

	var list []StakingSnapshotAsset
	err := h.Select(&list, query, snapshotID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Snapshot Assets Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetSnapshotAssetsByIDAndUserID[H Handle](h H, snapshotID string, userID string) (*[]StakingSnapshotAsset, error) {
	const op errors.Op = "GetSnapshotAsset"
	query := fmt.Sprintf("select %s from snapshot_asset where snapshot_id = ? user_id = ?", strings.Join(SnapshotAssetTableKeys(), ","))

	var list []StakingSnapshotAsset
	err := h.Select(&list, query, snapshotID, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Snapshot Assets Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetSnapshotAssetsByUserID[H Handle](h H, userID string) (*[]StakingSnapshotAsset, error) {
	const op errors.Op = "GetSnapshotAsset"
	query := fmt.Sprintf("select %s from snapshot_asset where user_id = ?", strings.Join(SnapshotAssetTableKeys(), ","))

	var list []StakingSnapshotAsset
	err := h.Select(&list, query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Snapshot Assets Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

// func GetTodaysSnapshotCount[H Handle](h H) (int64, error) {
// 	const op errors.Op = "GetTodaysSnapshotCount"
// 	query := "select count(*) from snapshot where ts > ? and ts < ?"
// }

func GetTodaysSnapshotAssetsByEntryID[H Handle](h H, entryID string) (*[]StakingSnapshotAsset, error) {
	const op errors.Op = "GetSnapshotAsset"
	query := fmt.Sprintf("select %s from snapshot_asset where snapshot_id in(select id from snapshot where entry_id = ? and ts > ? and ts < ?)", strings.Join(SnapshotAssetTableKeys(), ","))

	loc, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		return nil, errors.E(pkg, op, err)
	}

	year, month, day := time.Now().In(loc).Date()
	t := time.Date(year, month, day, 0, 0, 0, 0, loc)

	fmt.Println(entryID, t.Format(TimeFormat), t.Add(24*time.Hour).Format(TimeFormat))

	var list []StakingSnapshotAsset
	err = h.Select(&list, query, entryID, t.Format(TimeFormat), t.Add(24*time.Hour).Format(TimeFormat))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Snapshot Assets Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}
