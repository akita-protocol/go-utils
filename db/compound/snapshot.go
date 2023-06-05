package compound

import "gitlab.algotoolchain.com/keelakay/utils/db"

type StakingSnapshot struct {
	Snapshot *db.StakingSnapshot        `json:"snapshot"`
	Assets   *[]db.StakingSnapshotAsset `json:"assets"`
}
