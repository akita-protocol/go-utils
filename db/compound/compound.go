package compound

import "gitlab.algotoolchain.com/keelakay/utils/errors"

const pkg errors.Pkg = "compound"

type CompoundObject interface {
	*StakingPool | StakingPool | *Reward | Reward | *StakingSnapshot | StakingSnapshot
}
