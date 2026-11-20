package compound

import (
	"fmt"
	"strings"

	"github.com/akita-protocol/go-utils/db"
	"github.com/akita-protocol/go-utils/errors"
	"github.com/akita-protocol/go-utils/uuid"
)

type StakingPoolEntry struct {
	PoolEntry *db.StakingPoolEntry `json:"pool_entry,omitempty"`
	AssetIDs  *[]uint64            `json:"asset_ids,omitempty"`
}

func GetPoolEntryByID[H db.Handle](h H, id string) (*StakingPoolEntry, error) {
	const op errors.Op = "GetPoolEntryByID"

	pe, err := db.GetStakingPoolEntryByID(h, id)
	if err != nil {
		return nil, errors.E(pkg, op, err)
	}

	assets, err := db.GetStakingEntryAssets(h, pe.ID)
	if err != nil {
		return nil, errors.E(pkg, op, err)
	}

	return &StakingPoolEntry{
		PoolEntry: pe,
		AssetIDs:  assets,
	}, nil
}

func GetPoolEntriesByPoolID[H db.Handle](h H, poolID string) (*[]StakingPoolEntry, error) {
	const op errors.Op = "GetPoolEntriesByPoolID"

	entries, err := db.GetStakingPoolEntries(h, poolID)
	if err != nil {
		return nil, errors.E(pkg, op, err)
	}

	results := []StakingPoolEntry{}
	for _, entry := range *entries {
		assets, err := db.GetStakingEntryAssets(h, entry.ID)
		if err != nil {
			return nil, errors.E(pkg, op, err)
		}

		results = append(results, StakingPoolEntry{
			PoolEntry: &entry,
			AssetIDs:  assets,
		})
	}

	return &results, nil
}

func GetPoolEntryCount[H db.Handle](h H, poolID string) (uint64, error) {
	const op errors.Op = "GetPoolEntryCount"

	entries, err := GetPoolEntriesByPoolID(h, poolID)
	if err != nil {
		unwrappedError := err.(*errors.Error)
		if unwrappedError.Kind != errors.DatabaseResultNotFound {
			return 0, errors.E(pkg, op, err)
		}
		return 0, nil
	}

	requirements, err := GetStakingEntryRequirementsByPoolID(h, poolID)
	if err != nil {
		unwrappedError := err.(*errors.Error)
		if unwrappedError.Kind != errors.DatabaseResultNotFound {
			return 0, errors.E(pkg, op, err)
		}
	}

	entryCount := 0
	for _, entry := range *entries {
		if entry.AssetIDs != nil {
			entryCount += len(*entry.AssetIDs)
		}
	}

	if requirements == nil {
		return uint64(entryCount), nil
	}

	for _, requirement := range *requirements {
		if requirement.Type == "selection" {
			return uint64(entryCount) / requirement.Selection.Quantity, nil
		}
	}

	return uint64(entryCount), nil
}

func GetPoolEntriesByUserID[H db.Handle](h H, userID string) (*[]StakingPoolEntry, error) {
	const op errors.Op = "GetPoolEntriesByUserID"

	entries, err := db.GetStakingPoolEntriesByUserID(h, userID)
	if err != nil {
		return nil, errors.E(pkg, op, err)
	}

	results := make([]StakingPoolEntry, len(*entries))
	for _, entry := range *entries {
		assets, err := db.GetStakingEntryAssets(h, entry.ID)
		if err != nil {
			return nil, errors.E(pkg, op, err)
		}

		results = append(results, StakingPoolEntry{
			PoolEntry: &entry,
			AssetIDs:  assets,
		})
	}

	return &results, nil
}

func GetPoolEntryByPoolIDAndUserID[H db.Handle](h H, poolID string, userID string) (*StakingPoolEntry, error) {
	const op errors.Op = "GetPoolEntryByPoolIDAndUserID"

	pe, err := db.GetStakingPoolEntryByPoolIDAndUserID(h, poolID, userID)
	if err != nil {
		unwrappedError := err.(*errors.Error)
		if unwrappedError.Kind == errors.DatabaseResultNotFound {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err)
		}
		return nil, errors.E(pkg, op, err)
	}

	assets, err := db.GetStakingEntryAssets(h, pe.ID)
	if err != nil {
		unwrappedError := err.(*errors.Error)
		if unwrappedError.Kind == errors.DatabaseResultNotFound {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err)
		}
		return nil, errors.E(pkg, op, err)
	}

	return &StakingPoolEntry{
		PoolEntry: pe,
		AssetIDs:  assets,
	}, nil
}

func DeletePoolEntryAndAssetsByPoolIDAndUserID[H db.Handle](h H, poolID string, userID string) error {
	const op errors.Op = "DeletePoolEntryAndAssetsByPoolIDAndUserID"
	asset_query := fmt.Sprintf("delete from %s.entry_asset where entry_id = (select id from %s.pool_entry where pool_id = ? and user_id = ? limit 1)", db.StakingDatabase(), db.StakingDatabase())
	entry_query := fmt.Sprintf("delete from %s.pool_entry where pool_id = ? and user_id = ?", db.StakingDatabase())

	asset_stmt, err := h.Preparex(asset_query)
	if err != nil {
		return errors.E(pkg, op, err)
	}
	defer asset_stmt.Close()

	entry_stmt, err := h.Preparex(entry_query)
	if err != nil {
		return errors.E(pkg, op, err)
	}
	defer entry_stmt.Close()

	_, err = asset_stmt.Exec(poolID, userID)
	if err != nil {
		return errors.E(pkg, op, err)
	}

	_, err = entry_stmt.Exec(poolID, userID)
	if err != nil {
		return errors.E(pkg, op, err)
	}

	return nil
}

func NewStakingPoolEntry[H db.Handle](h H, poolID string, userID string, assets []uint64) (string, error) {
	const op errors.Op = "NewStakingPoolEntry"

	entry := &db.StakingPoolEntry{
		ID:     uuid.New(uuid.StakingPoolEntry),
		PoolID: poolID,
		UserID: userID,
	}

	_, err := db.Insert(h, entry)
	if err != nil {
		return "", errors.E(pkg, op, err)
	}

	for _, asset := range assets {
		asset := &db.StakingEntryAsset{
			EntryID: entry.ID,
			AssetID: asset,
		}

		_, err := db.Insert(h, asset)
		if err != nil {
			return "", errors.E(pkg, op, err)
		}
	}

	return entry.ID, nil
}

func IsPoolEntryChangeNecessary[H db.Handle](h H, poolID string, userID string, assets []uint64) (bool, error) {
	const op errors.Op = "CheckPoolEntryChangeNecessary"
	assetQuery := fmt.Sprintf("select entry_id from entry_asset where entry_id = (select id from pool_entry where pool_id = ? and user_id = ? and txn is not null limit 1) and asset_id in (%s)", strings.Repeat("?, ", len(assets))[0:(len(assets)*3)-2])
	var err error
	matches := []string{}

	params := []interface{}{
		poolID,
		userID,
	}
	for _, asset := range assets {
		params = append(params, asset)
	}

	err = h.Select(&matches, assetQuery, params...)
	if err != nil {
		return true, errors.E(op, err)
	}

	return !(len(assets) == len(matches)), nil
}
