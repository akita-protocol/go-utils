package db

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type StakingPoolEntry struct {
	ID        string    `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	PoolID    string    `structs:"pool_id,omitempty" db:"pool_id" json:"pool_id,omitempty"`
	UserID    string    `structs:"user_id,omitempty" db:"user_id" json:"user_id,omitempty"`
	EntryDate time.Time `structs:"entry_date,omitempty" db:"entry_date" json:"entry_date,omitempty"`
	Txn       *string   `structs:"txn,omitempty" db:"txn" json:"txn,omitempty"`
}

func StakingPoolEntryTableKeys() []string {
	return []string{"id", "pool_id", "user_id", "entry_date", "txn"}
}

func GetStakingPoolEntryByID[H Handle](h H, id string) (*StakingPoolEntry, error) {
	const op errors.Op = "GetStakingPoolEntryByID"
	query := fmt.Sprintf("select %s from %s.pool_entry where id = ? limit 1", strings.Join(StakingPoolEntryTableKeys(), ","), StakingDatabase())
	var entry StakingPoolEntry
	err := h.Get(&entry, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Pool Entry Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &entry, nil
}

func GetStakingPoolEntries[H Handle](h H, poolID string) (*[]StakingPoolEntry, error) {
	const op errors.Op = "GetStakingPoolEntries"
	query := fmt.Sprintf("select %s from %s.pool_entry where pool_id = ?", strings.Join(StakingPoolEntryTableKeys(), ","), StakingDatabase())
	var list []StakingPoolEntry
	err := h.Select(&list, query, poolID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Pool Entry Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetValidStakingPoolEntries[H Handle](h H, poolID string) (*[]StakingPoolEntry, error) {
	const op errors.Op = "GetValidStakingPoolEntries"
	query := fmt.Sprintf("select %s from %s.pool_entry where pool_id = ? and txn is not null", strings.Join(StakingPoolEntryTableKeys(), ","), StakingDatabase())
	var list []StakingPoolEntry
	err := h.Select(&list, query, poolID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Pool Entry Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetStakingPoolEntriesByUserID[H Handle](h H, userID string) (*[]StakingPoolEntry, error) {
	const op errors.Op = "GetStakingPoolEntriesByUserID"
	query := fmt.Sprintf("select %s from %s.pool_entry where user_id = ?", strings.Join(StakingPoolEntryTableKeys(), ","), StakingDatabase())

	var list []StakingPoolEntry
	err := h.Select(&list, query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Pool Entry Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetStakingPoolEntryUserIDs[H Handle](h H, poolID string) (*[]string, error) {
	const op errors.Op = "GetStakingPoolEntryUserIDs"
	query := fmt.Sprintf("select user_id from %s.pool_entry where pool_id = ?", StakingDatabase())
	var list []string
	err := h.Select(&list, query, poolID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Pool Entry Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &list, nil
}

func GetStakingPoolEntryByPoolIDAndUserID[H Handle](h H, poolID string, userID string) (*StakingPoolEntry, error) {
	const op errors.Op = "GetStakingPoolEntryByPoolIDAndUserID"
	query := fmt.Sprintf("select %s from %s.pool_entry where pool_id = ? and user_id = ? limit 1", strings.Join(StakingPoolEntryTableKeys(), ","), StakingDatabase())
	var entry StakingPoolEntry
	err := h.Get(&entry, query, poolID, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Pool Entry Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &entry, nil
}

func GetStakingPoolEntryCountByPoolID[H Handle](h H, poolID string) (uint64, error) {
	const op errors.Op = "GetStakingPoolEntryCountByPoolID"
	query := fmt.Sprintf("select count(*) from %s.pool_entry where pool_id = ?", StakingDatabase())
	var count uint64
	err := h.Get(&count, query, poolID)
	if err != nil {
		return 0, errors.E(pkg, op, err)
	}
	return count, nil
}
