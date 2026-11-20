package compound

import (
	"github.com/akita-protocol/go-utils/db"
	"github.com/akita-protocol/go-utils/errors"
)

type ShuffleAsset struct {
	*db.ShuffleAsset
	AssetProperties *[]db.ShuffleAssetProperty `json:"asset_properties,omitempty"`
}

func GetShuffleAssets[H db.Handle](h H, shuffleID string) (*[]ShuffleAsset, error) {
	const op errors.Op = "GetShuffleAsset"
	var shuffleAssets []ShuffleAsset

	assets, err := db.GetShuffleAssets(h, shuffleID)
	if err != nil {
		return nil, errors.E(pkg, op, err)
	}

	if len(*assets) == 0 {
		return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Shuffle Asset Not Found")
	}

	for i := range *assets {
		asset := (*assets)[i]
		shuffleAsset := ShuffleAsset{
			ShuffleAsset: &asset,
		}
		assetProperties, err := db.GetShuffleAssetProperties(h, asset.AssetID)
		if err != nil {
			return nil, errors.E(pkg, op, err)
		}

		shuffleAsset.AssetProperties = assetProperties
		shuffleAssets = append(shuffleAssets, shuffleAsset)
	}

	return &shuffleAssets, nil
}
