package compound

import (
	"github.com/akita-protocol/go-utils/db"
	"github.com/akita-protocol/go-utils/errors"
	"github.com/akita-protocol/go-utils/misc"
	"github.com/akita-protocol/go-utils/uuid"
)

type Reward struct {
	Reward        *db.Reward           `json:"reward"`
	Disbursements *[]db.Disbursement   `json:"disbursements,omitempty"`
	Codes         *[]db.RewardCode     `json:"code,omitempty"`
	Entry         *db.StakingPoolEntry `json:"entry,omitempty"`
}

func GetRewardsByPoolID[H db.Handle](h H, poolID string) (*[]Reward, error) {
	const op errors.Op = "GetRewardsByPoolID"
	rewards, err := db.GetRewardsByPoolID(h, poolID)
	if err != nil {
		return nil, errors.E(op, err)
	}

	fullRewards := []Reward{}
	for _, reward := range *rewards {
		fullRewards = append(fullRewards, Reward{
			Reward: &reward,
		})
	}
	return &fullRewards, nil
}

func GetRewardsByAddress[H db.Handle](h H, address string) (*[]Reward, error) {
	const op errors.Op = "GetRewardByAddress"
	rewards := []Reward{}

	user, err := db.GetStakingUserByAddress(h, address, db.TKSTiny)
	if err != nil {
		return nil, errors.E(op, err)
	}

	// TODO: paginate this api endpoint

	disbursements, err := db.GetUnclaimedDisbursementsByUserID(h, user.ID)
	if err != nil {
		return nil, errors.E(op, err)
	}

	if !(len(*disbursements) > 0) {
		return &rewards, nil
	}

	rewardIDs := []string{}
	rewardDisbursements := map[string][]db.Disbursement{}
	for _, dis := range *disbursements {
		if misc.InSlice(dis.RewardID, rewardIDs) {
			list := rewardDisbursements[dis.RewardID]
			list = append(list, dis)
			rewardDisbursements[dis.RewardID] = list
			continue
		}
		rewardIDs = append(rewardIDs, dis.RewardID)
		rewardDisbursements[dis.RewardID] = []db.Disbursement{dis}
	}

	dbrewards, err := db.GetRewardsByID(h, rewardIDs)
	if err != nil {
		return nil, errors.E(op, err)
	}

	for _, reward := range *dbrewards {
		realReward := reward
		if reward.Type != db.RewardTypeCode {
			disbursements := rewardDisbursements[reward.ID]
			rewards = append(rewards, Reward{
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

// 	claims, err := db.GetUnclaimedDisbursementsByUserAddress(h, address)
// 	if err != nil {
// 		return nil, errors.E(pkg, op, err)
// 	}

// 	for _, claim := range *claims {

// 	}
// }

func NewOneOffRewardAsset[H db.Handle](h H, name string, url string, assetID uint64, amount uint64, userID string, rewardAddress string) (rewardID string, disbursementID string, err error) {
	const op errors.Op = "NewReward"

	rewardID = uuid.New(uuid.Reward)
	reward := db.Reward{
		ID:               rewardID,
		Type:             db.RewardTypeAsset,
		DistributionType: db.RewardDistributionTypeEnd,
		RateType:         db.RewardRateTypeFlat,
		Rate:             1,
		Title:            name,
		RewardURL:        url,
		AssetID:          &assetID,
		Quantity:         &amount,
		Address:          &rewardAddress,
	}

	_, err = db.Insert(h, &reward)
	if err != nil {
		rewardID = ""
		err = errors.E(op, err)
		return
	}

	disbursementID = uuid.New(uuid.Disbursement)
	disbursement := db.Disbursement{
		ID:       disbursementID,
		GroupID:  uuid.New(uuid.StakingGroup),
		RewardID: rewardID,
		UserID:   userID,
		Quantity: &amount,
	}

	_, err = db.Insert(h, &disbursement)
	if err != nil {
		rewardID = ""
		disbursementID = ""
		err = errors.E(op, err)
		return
	}

	return
}
