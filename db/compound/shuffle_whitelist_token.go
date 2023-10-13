package compound

import (
	"gitlab.algotoolchain.com/keelakay/utils/db"
	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type ShuffleWhitelistToken struct {
	*db.ShuffleWhitelistToken
	Asset *db.Asset `json:"asset,omitempty"`
}

func GetShuffleWhitelistTokens[H db.Handle](h H, shuffleID string) (*[]ShuffleWhitelistToken, error) {
	const op errors.Op = "GetShuffleWhitelistTokens"

	tokens, err := db.GetShuffleWhitelistTokens(h, shuffleID)
	if err != nil {
		return nil, errors.E(pkg, op, err)
	}

	var swts []ShuffleWhitelistToken
	for i := range *tokens {
		swt := (*tokens)[i]
		asset, err := db.GetAsset(h, swt.AssetID, db.TKSTiny)
		if err != nil {
			return nil, errors.E(pkg, op, err)
		}

		swts = append(swts, ShuffleWhitelistToken{
			ShuffleWhitelistToken: &swt,
			Asset:                 asset,
		})
	}

	return &swts, nil
}
