package compound

import "gitlab.algotoolchain.com/keelakay/utils/errors"

const pkg errors.Pkg = "compound"

type CompoundObject interface {
	*StakingPool | StakingPool | *StakingReward | StakingReward | *StakingSnapshot | StakingSnapshot
}
