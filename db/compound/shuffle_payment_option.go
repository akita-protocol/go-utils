package compound

import (
	"gitlab.algotoolchain.com/keelakay/utils/db"
	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type ShufflePaymentOption struct {
	*db.ShufflePaymentOption
	Asset *db.Asset `json:"asset,omitempty"`
}

func GetShufflePaymentOptions[H db.Handle](h H, shuffleID string) (*[]ShufflePaymentOption, error) {
	const op errors.Op = "GetShufflePaymentOptions"

	paymentOptions, err := db.GetShufflePaymentOptions(h, shuffleID)
	if err != nil {
		return nil, errors.E(pkg, op, err)
	}

	var spos []ShufflePaymentOption
	for i := range *paymentOptions {
		spo := (*paymentOptions)[i]
		asset, err := db.GetAsset(h, *spo.AssetID, db.Tiny)
		if err != nil {
			return nil, errors.E(pkg, op, err)
		}

		spos = append(spos, ShufflePaymentOption{
			ShufflePaymentOption: &spo,
			Asset:                asset,
		})
	}

	return &spos, nil
}
