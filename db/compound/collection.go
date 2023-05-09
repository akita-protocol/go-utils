package compound

import "gitlab.algotoolchain.com/keelakay/utils/db"

type Collection struct {
	*db.Collection
	Prefixes       []string              `json:"prefixes,omitempty"`
	Assets         []uint64              `json:"assets,omitempty"`
	ExcludedAssets []uint64              `json:"excluded_assets,omitempty"`
	Artists        []string              `json:"artists,omitempty"`
	Properties     []Property            `json:"properties,omitempty"`
	Extras         []db.CollectionExtras `json:"extras,omitempty"`
}
