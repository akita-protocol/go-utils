package compound

import "github.com/akita-protocol/go-utils/errors"

const pkg errors.Pkg = "compound"

type CompoundObject interface {
	*StakingPool | StakingPool | *Reward | Reward | *StakingSnapshot | StakingSnapshot
}
