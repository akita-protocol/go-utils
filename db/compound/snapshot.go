package compound

import "github.com/akita-protocol/go-utils/db"

type StakingSnapshot struct {
	Snapshot *db.StakingSnapshot        `json:"snapshot"`
	Assets   *[]db.StakingSnapshotAsset `json:"assets"`
}
