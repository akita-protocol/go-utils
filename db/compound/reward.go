package compound

import (
	"gitlab.algotoolchain.com/keelakay/utils/db"
	"gitlab.algotoolchain.com/keelakay/utils/errors"
	"gitlab.algotoolchain.com/keelakay/utils/misc"
)

type StakingReward struct {
	Reward        *db.StakingReward         `json:"reward"`
	Disbursements *[]db.StakingDisbursement `json:"disbursements,omitempty"`
	Codes         *[]db.StakingRewardCode   `json:"code,omitempty"`
	Entry         *db.StakingPoolEntry      `json:"entry,omitempty"`
}

func GetStakingRewardsByPoolID[H db.Handle](h H, poolID string) (*[]StakingReward, error) {
	const op errors.Op = "GetRewardsByPoolID"
	rewards, err := db.GetStakingRewardsByPoolID(h, poolID)
	if err != nil {
		return nil, errors.E(op, err)
	}

	fullRewards := []StakingReward{}
	for _, reward := range *rewards {
		fullRewards = append(fullRewards, StakingReward{
			Reward: &reward,
		})
	}
	return &fullRewards, nil
}

func GetStakingRewardsByAddress[H db.Handle](h H, address string) (*[]StakingReward, error) {
	const op errors.Op = "GetRewardByAddress"
	rewards := []StakingReward{}

	user, err := db.GetStakingUserByAddress(h, address, db.Tiny)
	if err != nil {
		return nil, errors.E(op, err)
	}

	// TODO: paginate this api endpoint

	disbursements, err := db.GetUnclaimedStakingDisbursementsByUserID(h, user.ID)
	if err != nil {
		return nil, errors.E(op, err)
	}

	if !(len(*disbursements) > 0) {
		return &rewards, nil
	}

	rewardIDs := []string{}
	rewardDisbursements := map[string][]db.StakingDisbursement{}
	for _, dis := range *disbursements {
		if misc.InSlice(dis.RewardID, rewardIDs) {
			list := rewardDisbursements[dis.RewardID]
			list = append(list, dis)
			rewardDisbursements[dis.RewardID] = list
			continue
		}
		rewardIDs = append(rewardIDs, dis.RewardID)
		rewardDisbursements[dis.RewardID] = []db.StakingDisbursement{dis}
	}

	dbrewards, err := db.GetStakingRewardsByID(h, rewardIDs)
	if err != nil {
		return nil, errors.E(op, err)
	}

	for _, reward := range *dbrewards {
		realReward := reward
		if reward.Type != db.StakingRewardTypeCode {
			disbursements := rewardDisbursements[reward.ID]
			rewards = append(rewards, StakingReward{
				Reward:        &realReward,
				Disbursements: &disbursements,
			})
			continue
		}

		// TODO: codes implementation
		// codes, err := db.GetRewardCodesByRewardID(h, reward.ID)
		// if err != nil {
		// 	return nil, errors.E(op, err)
		// }

		// rewards = append(rewards, Reward{
		// 	Reward: &reward,
		// 	Disbursements: disbursements,
		// 	Codes:         codes,
		// })
	}
	return &rewards, nil
}

// func GetUnclaimedRewardsByUserAddress[H db.Handle](h H, address string) (*[]Reward, error) {
// 	const op errors.Op = "GetUnclaimedRewardsByUserAddress"

// 	claims, err := db.GetUnclaimedRewardDisbursementsByUserAddress(h, address)
// 	if err != nil {
// 		return nil, errors.E(pkg, op, err)
// 	}

// 	for _, claim := range *claims {

// 	}
// }
