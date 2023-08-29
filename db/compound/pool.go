package compound

import (
	"fmt"

	"gitlab.algotoolchain.com/keelakay/utils/db"
	"gitlab.algotoolchain.com/keelakay/utils/errors"
	"gitlab.algotoolchain.com/keelakay/utils/misc"
)

type PoolStatus string

const (
	PoolStatusOpen     PoolStatus = "open"
	PoolStatusClosed   PoolStatus = "closed"
	PoolStatusUpcoming PoolStatus = "upcoming"
	PoolStatusEnded    PoolStatus = "ended"
	PoolStatusUnknown  PoolStatus = "unknown"
)

type QualifiedAssetTypes string

const (
	QualifiedAssetTypesNFT  QualifiedAssetTypes = "nft"
	QualifiedAssetTypesASA  QualifiedAssetTypes = "asa"
	QualifiedAssetTypesBoth QualifiedAssetTypes = "both"
)

type StakingPool struct {
	Pool            *db.StakingPool            `json:"pool"`
	Status          PoolStatus                 `json:"status"`
	QualifiedAssets QualifiedAssetTypes        `json:"qualified_assets"`
	Assets          *[]uint64                  `json:"assets"`
	Exclusions      *[]StakingExclusion        `json:"exclusions"`
	Requirements    *[]StakingEntryRequirement `json:"requirements"`
	EntryCount      uint64                     `json:"entry_count"`
	Entries         *[]StakingPoolEntry        `json:"entries"`
	Rewards         *[]Reward                  `json:"rewards"`
}

type PoolGetExclude string

const (
	PoolGetExcludeAssets       PoolGetExclude = "assets"
	PoolGetExcludeExclusions   PoolGetExclude = "exclusions"
	PoolGetExcludeRequirements PoolGetExclude = "requirements"
	PoolGetExcludeEntryCount   PoolGetExclude = "entry_count"
	PoolGetExcludeEntries      PoolGetExclude = "entries"
	PoolGetExcludeRewards      PoolGetExclude = "rewards"
)

// Get returns a pool from the database concurrently given a list of exclude fields
func GetStakingPool[H db.Handle](h H, id string, exclude ...PoolGetExclude) (*StakingPool, error) {
	const op errors.Op = "GetStakingPool"
	buffer := (7 - len(exclude))
	rChan := make(chan interface{}, buffer)
	defer close(rChan)
	var p StakingPool

	go func() {
		pool, err := db.GetStakingPool(h, id, db.Full)
		if err != nil {
			rChan <- err
			return
		}
		rChan <- pool
	}()

	if !misc.InSlice(PoolGetExcludeAssets, exclude) {
		go func() {
			assets, err := db.GetStakingPoolAssetIDs(h, id)
			if err != nil {
				rChan <- err
				return
			}
			rChan <- assets
		}()
	}

	if !misc.InSlice(PoolGetExcludeExclusions, exclude) {
		go func() {
			exclusions, err := db.GetStakingExclusionIDsByPoolID(h, id)
			if err != nil {
				rChan <- err
				return
			}

			cexclusions := &[]StakingExclusion{}
			for _, cid := range *exclusions {
				poolIDs, err := db.GetStakingExclusionPoolIDsByID(h, cid, id)
				if err != nil {
					rChan <- err
					return
				}

				pools, err := db.GetStakingPoolsIn(h, *poolIDs, db.Tiny)
				if err != nil {
					rChan <- err
					return
				}

				*cexclusions = append(*cexclusions, StakingExclusion{
					ID:    id,
					Pools: pools,
				})
			}

			rChan <- cexclusions
		}()
	}

	if !misc.InSlice(PoolGetExcludeRequirements, exclude) {
		go func() {
			requirements, err := GetStakingEntryRequirementsByPoolID(h, id)
			if err != nil {
				rChan <- err
				return
			}
			rChan <- requirements
		}()
	}

	if !misc.InSlice(PoolGetExcludeEntryCount, exclude) {
		go func() {
			entryCount, err := GetPoolEntryCount(h, id)
			if err != nil {
				rChan <- err
				return
			}
			rChan <- entryCount
		}()
	}

	if !misc.InSlice(PoolGetExcludeEntries, exclude) {
		go func() {
			entries, err := GetPoolEntriesByPoolID(h, id)
			if err != nil {
				rChan <- err
				return
			}
			rChan <- entries
		}()
	}

	if !misc.InSlice(PoolGetExcludeRewards, exclude) {
		go func() {
			rewards, err := GetRewardsByPoolID(h, id)
			if err != nil {
				rChan <- err
				return
			}
			rChan <- rewards
		}()
	}

	var errs []error
	for i := 0; i < buffer; i++ {
		data := <-rChan
		switch result := data.(type) {
		case *db.StakingPool:
			p.Pool = result

			// replace
			usr, err := db.GetStakingUser(h, p.Pool.Creator, db.Tiny)
			if err != nil {
				errs = append(errs, err)
				continue
			}

			p.Pool.Creator = usr.Address

		case *[]uint64:
			p.Assets = result
		case *[]StakingExclusion:
			p.Exclusions = result
		case *[]StakingEntryRequirement:
			p.Requirements = result
		case uint64:
			p.EntryCount = result
		case *[]StakingPoolEntry:
			p.Entries = result
		case *[]Reward:
			p.Rewards = result
		case error:
			err := result.(*errors.Error)
			if err.Kind == errors.DatabaseResultNotFound && err.Msg == "Pool Not Found" {
				errs = append(errs, err)
			} else {
				errs = append(errs, err)
			}
		}
	}

	if len(errs) > 0 {
		msg := ""
		for _, err := range errs {
			msg += err.Error() + "\n"
		}

		return nil, errors.E(op, pkg, fmt.Errorf(msg))
	}

	p.Status = PoolStatus(p.Pool.GetStakingPoolStatus())

	return &p, nil
}
