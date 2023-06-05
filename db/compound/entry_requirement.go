package compound

import (
	"gitlab.algotoolchain.com/keelakay/utils/db"
	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type StakingEntryRequirement struct {
	*db.StakingEntryRequirement
	Selection *db.StakingEntryRequirementSelection `json:"selection,omitempty"`
	Asset     *db.StakingEntryRequirementAsset     `json:"asset,omitempty"`
}

func GetStakingEntryRequirementsByPoolID[H db.Handle](h H, poolID string) (*[]StakingEntryRequirement, error) {
	const op errors.Op = "GetStakingEntryRequirementsByPoolID"

	er, err := db.GetStakingEntryRequirementsByPoolID(h, poolID)
	if err != nil {
		return nil, errors.E(pkg, op, err)
	}

	var list []StakingEntryRequirement
	for _, e := range *er {
		switch e.Type {
		case db.StakingEntryRequirementTypeSelection:
			selection, err := db.GetStakingEntryRequirementSelection(h, e.ID)
			if err != nil {
				return nil, errors.E(pkg, op, err)
			}

			list = append(list, StakingEntryRequirement{
				StakingEntryRequirement: &e,
				Selection:               selection,
			})
		case db.StakingEntryRequirementTypeAsset:
			asset, err := db.GetStakingEntryRequirementAsset(h, e.ID)
			if err != nil {
				return nil, errors.E(pkg, op, err)
			}

			list = append(list, StakingEntryRequirement{
				StakingEntryRequirement: &e,
				Asset:                   asset,
			})
		default:
			return nil, errors.E(pkg, op, "Invalid Entry Requirement Type")
		}
	}
	return &list, nil
}
