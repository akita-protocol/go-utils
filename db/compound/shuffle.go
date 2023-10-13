package compound

import (
	"gitlab.algotoolchain.com/keelakay/utils/db"
	"gitlab.algotoolchain.com/keelakay/utils/errors"
	"gitlab.algotoolchain.com/keelakay/utils/misc"
)

type Shuffle struct {
	*db.Shuffle
	Assets               *[]ShuffleAsset           `json:"assets,omitempty"`
	PaymentOptions       *[]ShufflePaymentOption   `json:"payment_options,omitempty"`
	WhitelistTokens      *[]ShuffleWhitelistToken  `json:"whitelist_tokens,omitempty"`
	ShuffleReserveAssets *[]db.ShuffleReserveAsset `json:"shuffle_reserve_assets,omitempty"`
	Sales                *[]db.ShuffleSale         `json:"sales,omitempty"`
	Creator              *Wallet                   `json:"creator,omitempty"`
	Available            *uint64                   `json:"available,omitempty"`
}

type ShuffleGetExclude string

const (
	ShuffleGetExcludeAssets               ShuffleGetExclude = "assets"
	ShuffleGetExcludeAssetProperties      ShuffleGetExclude = "asset_properties"
	ShuffleGetExcludePaymentOptions       ShuffleGetExclude = "payment_options"
	ShuffleGetExcludeWhitelistTokens      ShuffleGetExclude = "whitelist_tokens"
	ShuffleGetExcludeShuffleReserveAssets ShuffleGetExclude = "shuffle_reserve_assets"
	ShuffleGetExcludeSales                ShuffleGetExclude = "sales"
	ShuffleGetExcludeCreator              ShuffleGetExclude = "creator"
	ShuffleGetExcludeAvailable            ShuffleGetExclude = "available"
)

func GetShuffles[H db.Handle](h H, nfdID uint64, exclude ...ShuffleGetExclude) (*[]Shuffle, error) {
	const op errors.Op = "GetShuffles"

	var shuffles []Shuffle
	shufs, err := db.GetShufflesByNFDID(h, nfdID)
	if err != nil {
		return nil, errors.E(pkg, op, err)
	}

	for i := range *shufs {
		shuf := (*shufs)[i]
		shuffle := Shuffle{
			Shuffle: &shuf,
		}

		var buffer int
		if misc.InSlice(ShuffleGetExcludeAssetProperties, exclude) {
			buffer = (7 - (len(exclude) - 1))
		} else {
			buffer = (7 - len(exclude))
		}

		if buffer == 0 {
			continue
		}

		rChan := make(chan interface{}, buffer)
		defer close(rChan)

		if !misc.InSlice(ShuffleGetExcludeAssets, exclude) {
			go func() {
				if !misc.InSlice(ShuffleGetExcludeAssetProperties, exclude) {
					assets, err := GetShuffleAssets(h, shuf.ID)
					if err != nil {
						rChan <- err
						return
					}
					rChan <- assets
				} else {
					assets, err := db.GetShuffleAssets(h, shuf.ID)
					if err != nil {
						rChan <- err
						return
					}

					var cmpAssets []ShuffleAsset
					for i := range *assets {
						asset := (*assets)[i]
						cmpAssets = append(cmpAssets, ShuffleAsset{ShuffleAsset: &asset})
					}

					rChan <- &cmpAssets
				}
			}()
		}

		if !misc.InSlice(ShuffleGetExcludePaymentOptions, exclude) {
			go func() {
				paymentOptions, err := GetShufflePaymentOptions(h, shuf.ID)
				if err != nil {
					rChan <- err
					return
				}
				rChan <- paymentOptions
			}()
		}

		if !misc.InSlice(ShuffleGetExcludeWhitelistTokens, exclude) {
			go func() {
				whitelistTokens, err := GetShuffleWhitelistTokens(h, shuf.ID)
				if err != nil && !db.ErrNoRows(err) {
					rChan <- err
					return
				}
				rChan <- whitelistTokens
			}()
		}

		if !misc.InSlice(ShuffleGetExcludeShuffleReserveAssets, exclude) {
			go func() {
				shuffleReserveAssets, err := db.GetShuffleReserveAssets(h, shuf.ID)
				if err != nil && !db.ErrNoRows(err) {
					rChan <- err
					return
				}
				rChan <- shuffleReserveAssets
			}()
		}

		if !misc.InSlice(ShuffleGetExcludeSales, exclude) {
			go func() {
				sales, err := db.GetShuffleSales(h, shuf.ID)
				if err != nil && !db.ErrNoRows(err) {
					rChan <- err
					return
				}
				rChan <- sales
			}()
		}

		if !misc.InSlice(ShuffleGetExcludeAvailable, exclude) {
			go func() {
				count, err := db.GetShuffleSalesCount(h, shuf.ID)
				if err != nil && !db.ErrNoRows(err) {
					rChan <- err
					return
				}
				available := (*shuf.MaxSold - *count)
				rChan <- &available
			}()
		}

		if !misc.InSlice(ShuffleGetExcludeCreator, exclude) {
			go func() {
				nfd, err := GetNFDByAddress(h, shuf.Creator, db.TKSTiny)
				if err != nil && !db.ErrNoRows(err) {
					rChan <- err
					return
				}

				if nfd != nil {
					rChan <- &Wallet{
						Address: shuf.Creator,
						NFD:     nfd,
					}
				} else {
					rChan <- &Wallet{Address: shuf.Creator}
				}
			}()
		}

		for i := 0; i < buffer; i++ {
			r := <-rChan
			switch v := r.(type) {
			case error:
				return nil, errors.E(pkg, op, v)
			case *[]ShuffleAsset:
				shuffle.Assets = v
			case *[]ShufflePaymentOption:
				shuffle.PaymentOptions = v
			case *[]ShuffleWhitelistToken:
				shuffle.WhitelistTokens = v
			case *[]db.ShuffleReserveAsset:
				shuffle.ShuffleReserveAssets = v
			case *[]db.ShuffleSale:
				shuffle.Sales = v
			case *Wallet:
				shuffle.Creator = v
			case *uint64:
				shuffle.Available = v
			}
		}

		shuffles = append(shuffles, shuffle)
	}

	return &shuffles, nil
}

func GetShuffle[H db.Handle](h H, shuffleID string, exclude ...ShuffleGetExclude) (*Shuffle, error) {
	const op errors.Op = "GetShuffles"

	shuf, err := db.GetShuffle(h, shuffleID)
	if err != nil {
		return nil, errors.E(pkg, op, err)
	}

	shuffle := Shuffle{
		Shuffle: shuf,
	}

	var buffer int
	if misc.InSlice(ShuffleGetExcludeAssetProperties, exclude) {
		buffer = (7 - (len(exclude) - 1))
	} else {
		buffer = (7 - len(exclude))
	}

	if buffer == 0 {
		return &shuffle, nil
	}

	rChan := make(chan interface{}, buffer)
	defer close(rChan)

	if !misc.InSlice(ShuffleGetExcludeAssets, exclude) {
		go func() {
			if !misc.InSlice(ShuffleGetExcludeAssetProperties, exclude) {
				assets, err := GetShuffleAssets(h, shuf.ID)
				if err != nil {
					rChan <- err
					return
				}
				rChan <- assets
			} else {
				assets, err := db.GetShuffleAssets(h, shuf.ID)
				if err != nil {
					rChan <- err
					return
				}

				var cmpAssets []ShuffleAsset
				for i := range *assets {
					asset := (*assets)[i]
					cmpAssets = append(cmpAssets, ShuffleAsset{ShuffleAsset: &asset})
				}

				rChan <- &cmpAssets
			}
		}()
	}

	if !misc.InSlice(ShuffleGetExcludePaymentOptions, exclude) {
		go func() {
			paymentOptions, err := GetShufflePaymentOptions(h, shuf.ID)
			if err != nil {
				rChan <- err
				return
			}
			rChan <- paymentOptions
		}()
	}

	if !misc.InSlice(ShuffleGetExcludeWhitelistTokens, exclude) {
		go func() {
			whitelistTokens, err := GetShuffleWhitelistTokens(h, shuf.ID)
			if err != nil && !db.ErrNoRows(err) {
				rChan <- err
				return
			}
			rChan <- whitelistTokens
		}()
	}

	if !misc.InSlice(ShuffleGetExcludeShuffleReserveAssets, exclude) {
		go func() {
			shuffleReserveAssets, err := db.GetShuffleReserveAssets(h, shuf.ID)
			if err != nil && !db.ErrNoRows(err) {
				rChan <- err
				return
			}
			rChan <- shuffleReserveAssets
		}()
	}

	if !misc.InSlice(ShuffleGetExcludeSales, exclude) {
		go func() {
			sales, err := db.GetShuffleSales(h, shuf.ID)
			if err != nil && !db.ErrNoRows(err) {
				rChan <- err
				return
			}
			rChan <- sales
		}()
	}

	if !misc.InSlice(ShuffleGetExcludeAvailable, exclude) {
		go func() {
			count, err := db.GetShuffleSalesCount(h, shuf.ID)
			if err != nil && !db.ErrNoRows(err) {
				rChan <- err
				return
			}
			available := (*shuf.MaxSold - *count)
			rChan <- &available
		}()
	}

	if !misc.InSlice(ShuffleGetExcludeCreator, exclude) {
		go func() {
			nfd, err := GetNFDByAddress(h, shuf.Creator, db.TKSTiny)
			if err != nil && !db.ErrNoRows(err) {
				rChan <- err
				return
			}

			if nfd != nil {
				rChan <- &Wallet{
					Address: shuf.Creator,
					NFD:     nfd,
				}
			} else {
				rChan <- &Wallet{Address: shuf.Creator}
			}
		}()
	}

	for i := 0; i < buffer; i++ {
		r := <-rChan
		switch v := r.(type) {
		case error:
			return nil, errors.E(pkg, op, v)
		case *[]ShuffleAsset:
			shuffle.Assets = v
		case *[]ShufflePaymentOption:
			shuffle.PaymentOptions = v
		case *[]ShuffleWhitelistToken:
			shuffle.WhitelistTokens = v
		case *[]db.ShuffleReserveAsset:
			shuffle.ShuffleReserveAssets = v
		case *[]db.ShuffleSale:
			shuffle.Sales = v
		case *Wallet:
			shuffle.Creator = v
		case *uint64:
			shuffle.Available = v
		}
	}

	return &shuffle, nil
}
