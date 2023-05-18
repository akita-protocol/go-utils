package db

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type StakingSnapshot struct {
	ID        string    `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	EntryID   string    `structs:"entry_id,omitempty" db:"entry_id" json:"entry_id,omitempty"`
	Timestamp time.Time `structs:"ts,omitempty" db:"ts" json:"timestamp,omitempty"`
}

func StakingSnapshotTableKeys() []string {
	return []string{"id", "entry_id", "ts"}
}

func GetSnapshots[H Handle](h H, entryID string) (*[]StakingSnapshot, error) {
	const op errors.Op = "GetSnapshots"
	query := fmt.Sprintf("select %s from snapshot where entry_id = ?", strings.Join(StakingSnapshotTableKeys(), ","))
	var snapshots []StakingSnapshot
	err := h.Select(&snapshots, query, entryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Snapshots Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &snapshots, nil
}

func GetSnapshotsByDay[H Handle](h H, entryID string, date time.Time) (*[]StakingSnapshot, error) {
	const op errors.Op = "GetSnapshotsByDay"
	query := fmt.Sprintf("select %s from snapshot where entry_id = ? and ts >= ? and ts < ?", strings.Join(StakingSnapshotTableKeys(), ","))
	year, month, day := date.Date()
	new_date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	var snapshots []StakingSnapshot
	err := h.Select(&snapshots, query, entryID, new_date, new_date.Add(24*time.Hour))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Snapshots Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &snapshots, nil
}
