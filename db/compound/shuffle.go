package compound

import (
	"gitlab.algotoolchain.com/keelakay/utils/db"
	"gitlab.algotoolchain.com/keelakay/utils/errors"
	"gitlab.algotoolchain.com/keelakay/utils/misc"
)

type Shuffle struct {
	*db.Shuffle
	Assets               *[]ShuffleAsset            `json:"assets,omitempty"`
	PaymentOptions       *[]db.ShufflePaymentOption `json:"payment_options,omitempty"`
	WhitelistTokens      *[]ShuffleWhitelistToken   `json:"whitelist_tokens,omitempty"`
	ShuffleReserveAssets *[]db.ShuffleReserveAsset  `json:"shuffle_reserve_assets,omitempty"`
	Sales                *[]db.ShuffleSale          `json:"sales,omitempty"`
	Creator              *Wallet                    `json:"creator,omitempty"`
	Available            uint64                     `json:"available,omitempty"`
}

type ShuffleGetExclude string

const (
	ShuffleGetExcludeAssets               ShuffleGetExclude = "assets"
	ShuffleGetExcludePaymentOptions       ShuffleGetExclude = "payment_options"
	ShuffleGetExcludeWhitelistTokens      ShuffleGetExclude = "whitelist_tokens"
	ShuffleGetExcludeShuffleReserveAssets ShuffleGetExclude = "shuffle_reserve_assets"
	ShuffleGetExcludeSales                ShuffleGetExclude = "sales"
	ShuffleGetExcludeCreator              ShuffleGetExclude = "creator"
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
		buffer := (6 - len(exclude))
		if buffer == 0 {
			continue
		}

		rChan := make(chan interface{}, buffer)
		defer close(rChan)

		if !misc.InSlice(ShuffleGetExcludeAssets, exclude) {
			go func() {
				assets, err := GetShuffleAssets(h, shuf.ID)
				if err != nil {
					rChan <- err
					return
				}
				rChan <- assets
			}()
		}

		if !misc.InSlice(ShuffleGetExcludePaymentOptions, exclude) {
			go func() {
				paymentOptions, err := db.GetShufflePaymentOptions(h, shuf.ID)
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

		if !misc.InSlice(ShuffleGetExcludeCreator, exclude) {
			go func() {
				nfds, err := db.GetNFDsByAddress(h, shuf.Creator, db.Tiny)
				if err != nil && !db.ErrNoRows(err) {
					rChan <- err
					return
				}

				if nfds != nil && len(*nfds) > 0 {
					rChan <- &Wallet{
						Address: shuf.Creator,
						NFD:     &(*nfds)[0],
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
			case *[]db.ShufflePaymentOption:
				shuffle.PaymentOptions = v
			case *[]ShuffleWhitelistToken:
				shuffle.WhitelistTokens = v
			case *[]db.ShuffleReserveAsset:
				shuffle.ShuffleReserveAssets = v
			case *[]db.ShuffleSale:
				shuffle.Sales = v
			case *Wallet:
				shuffle.Creator = v
			}
		}

		if !misc.InSlice(ShuffleGetExcludeSales, exclude) {
			shuffle.Available = *shuf.MaxSold - uint64(len((*shuffle.Sales)))
		}

		shuffles = append(shuffles, shuffle)
	}

	return &shuffles, nil
}
