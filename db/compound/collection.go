package compound

import (
	"fmt"

	"gitlab.algotoolchain.com/keelakay/utils/db"
	"gitlab.algotoolchain.com/keelakay/utils/errors"
	"gitlab.algotoolchain.com/keelakay/utils/misc"
)

type Collection struct {
	*db.Collection
	BannerURL      *string           `json:"banner_url,omitempty"`
	BannerMime     *string           `json:"banner_mime,omitempty"`
	AvatarURL      *string           `json:"avatar_url,omitempty"`
	AvatarMime     *string           `json:"avatar_mime,omitempty"`
	Prefixes       []string          `json:"prefixes,omitempty"`
	Assets         []uint64          `json:"assets,omitempty"`
	ExcludedAssets []uint64          `json:"excluded_assets,omitempty"`
	Artists        []string          `json:"artists,omitempty"`
	Properties     []Property        `json:"properties,omitempty"`
	Extras         map[string]string `json:"extras,omitempty"`
}

type LabeledAsset struct {
	Asset *db.Asset `json:"asset"`
	Label string    `json:"label"`
}

type CollectionGetExclude string

const (
	CollectionGetExcludeBanner         CollectionGetExclude = "banner"
	CollectionGetExcludeAvatar         CollectionGetExclude = "avatar"
	CollectionGetExcludePrefixes       CollectionGetExclude = "prefixes"
	CollectionGetExcludeAssets         CollectionGetExclude = "assets"
	CollectionGetExcludeExcludedAssets CollectionGetExclude = "excluded_assets"
	CollectionGetExcludeArtists        CollectionGetExclude = "artists"
	CollectionGetExcludeProperties     CollectionGetExclude = "properties"
	CollectionGetExcludeExtras         CollectionGetExclude = "extras"
)

func GetCollectionsByNFDID[H db.Handle](h H, nfdID uint64, exclude ...CollectionGetExclude) (*[]Collection, error) {
	const op errors.Op = "GetCollectionsByNFDID"

	var collections []Collection
	cols, err := db.GetCollectionsByNFDID(h, nfdID)
	if err != nil {
		return nil, errors.E(op, err)
	}

	for i := range *cols {
		col := (*cols)[i]
		compCollection := Collection{
			Collection: &col,
		}
		buffer := (8 - len(exclude))
		if buffer == 0 {
			continue
		}

		rChan := make(chan interface{}, buffer)
		defer close(rChan)

		if !misc.InSlice(CollectionGetExcludeBanner, exclude) {
			if col.Banner != nil {
				go func() {
					bannerAsset, err := db.GetAsset(h, *col.Banner, db.Tiny)
					if err != nil && !db.ErrNoRows(err) {
						rChan <- err
						return
					}

					if !db.ErrNoRows(err) {
						rChan <- LabeledAsset{
							Asset: bannerAsset,
							Label: "banner",
						}
						return
					}

					rChan <- nil
				}()
			} else {
				fmt.Println("banner is nil")
				buffer--
			}
		}

		if !misc.InSlice(CollectionGetExcludeAvatar, exclude) {
			if col.Avatar != nil {
				go func() {
					avatarAsset, err := db.GetAsset(h, *col.Avatar, db.Tiny)
					if err != nil && !db.ErrNoRows(err) {
						rChan <- err
						return
					}

					if !db.ErrNoRows(err) {
						rChan <- LabeledAsset{
							Asset: avatarAsset,
							Label: "avatar",
						}
						return
					}

					rChan <- nil
				}()
			} else {
				fmt.Println("avatar is nil")
				buffer--
			}
		}

		if !misc.InSlice(CollectionGetExcludePrefixes, exclude) {
			go func() {
				prefixes, err := db.GetCollectionPrefixes(h, col.ID)
				if err != nil && !db.ErrNoRows(err) {
					rChan <- err
					return
				}
				rChan <- prefixes
			}()
		}

		if !misc.InSlice(CollectionGetExcludeAssets, exclude) {
			go func() {
				assets, err := db.GetCollectionAssets(h, col.ID)
				if err != nil && !db.ErrNoRows(err) {
					rChan <- err
					return
				}
				rChan <- assets
			}()
		}

		if !misc.InSlice(CollectionGetExcludeExcludedAssets, exclude) {
			go func() {
				excludedAssets, err := db.GetCollectionExcludedAssets(h, col.ID)
				if err != nil && !db.ErrNoRows(err) {
					rChan <- err
					return
				}
				rChan <- excludedAssets
			}()
		}

		if !misc.InSlice(CollectionGetExcludeArtists, exclude) {
			go func() {
				artists, err := db.GetCollectionArtistByCollection(h, col.ID)
				if err != nil && !db.ErrNoRows(err) {
					rChan <- err
					return
				}
				rChan <- artists
			}()
		}

		if !misc.InSlice(CollectionGetExcludeProperties, exclude) {
			go func() {
				properties, err := GetProperties(h, col.ID)
				if err != nil && !db.ErrNoRows(err) {
					rChan <- err
					return
				}
				rChan <- properties
			}()
		}

		if !misc.InSlice(CollectionGetExcludeExtras, exclude) {
			go func() {
				extras, err := db.GetCollectionExtras(h, col.ID)
				if err != nil && !db.ErrNoRows(err) {
					rChan <- err
					return
				}
				rChan <- extras
			}()
		}

		var errs []error
		for i := 0; i < buffer; i++ {
			data := <-rChan
			switch result := data.(type) {
			case *[]db.CollectionPrefix:
				prefixes := []string{}
				for i := range *result {
					prefix := (*result)[i]
					prefixes = append(prefixes, prefix.Prefix)
				}
				compCollection.Prefixes = prefixes
			case *[]db.CollectionAsset:
				assets := []uint64{}
				for i := range *result {
					asset := (*result)[i]
					assets = append(assets, asset.AsaID)
				}
				compCollection.Assets = assets
			case *[]db.CollectionExcludedAsset:
				assets := []uint64{}
				for i := range *result {
					asset := (*result)[i]
					assets = append(assets, asset.AsaID)
				}
				compCollection.ExcludedAssets = assets
			case *[]db.CollectionArtist:
				artists := []string{}
				for i := range *result {
					artist := (*result)[i]
					artists = append(artists, artist.Address)
				}
				compCollection.Artists = artists
			case *[]Property:
				compCollection.Properties = *result
			case *[]db.CollectionExtras:
				extras := map[string]string{}
				for i := range *result {
					extra := (*result)[i]
					extras[extra.Key] = extra.Value
				}

				compCollection.Extras = extras
			case LabeledAsset:
				switch result.Label {
				case "banner":
					compCollection.BannerMime = result.Asset.Mime
					if result.Asset.ImageURL != nil && *result.Asset.ImageURL != "" {
						compCollection.BannerURL = result.Asset.ImageURL
					} else {
						compCollection.BannerURL = result.Asset.URL
					}
				case "avatar":
					compCollection.AvatarMime = result.Asset.Mime
					if result.Asset.ImageURL != nil && *result.Asset.ImageURL != "" {
						compCollection.AvatarURL = result.Asset.ImageURL
					} else {
						compCollection.AvatarURL = result.Asset.URL
					}
				}
			case error:
				errs = append(errs, result)
			default:
				fmt.Println(op, "defaulted")
				fmt.Println(result)
			}
		}

		if len(errs) > 0 {

			msg := ""
			for _, err := range errs {
				msg += err.Error() + "\n"
			}

			return nil, errors.E(op, fmt.Errorf(msg))
		}
		collections = append(collections, compCollection)
	}

	return &collections, nil
}

func GetCollectionCriteria[H db.Handle](h H, collectionID string) (*Collection, error) {
	const op errors.Op = "GetCollectionCriteria"
	collection := Collection{}

	buffer := 3
	rChan := make(chan interface{}, buffer)
	defer close(rChan)

	go func() {
		prefixes, err := db.GetCollectionPrefixes(h, collectionID)
		if err != nil && !db.ErrNoRows(err) {
			rChan <- err
			return
		}
		rChan <- prefixes
	}()

	go func() {
		assets, err := db.GetCollectionAssets(h, collectionID)
		if err != nil && !db.ErrNoRows(err) {
			rChan <- err
			return
		}
		rChan <- assets
	}()

	go func() {
		excludedAssets, err := db.GetCollectionExcludedAssets(h, collectionID)
		if err != nil && !db.ErrNoRows(err) {
			rChan <- err
			return
		}
		rChan <- excludedAssets
	}()

	var errs []error
	for i := 0; i < buffer; i++ {
		data := <-rChan
		switch result := data.(type) {
		case *[]db.CollectionPrefix:
			prefixes := []string{}
			for i := range *result {
				prefix := (*result)[i]
				prefixes = append(prefixes, prefix.Prefix)
			}
			collection.Prefixes = prefixes
		case *[]db.CollectionAsset:
			assets := []uint64{}
			for i := range *result {
				asset := (*result)[i]
				assets = append(assets, asset.AsaID)
			}
			collection.Assets = assets
		case *[]db.CollectionExcludedAsset:
			assets := []uint64{}
			for i := range *result {
				asset := (*result)[i]
				assets = append(assets, asset.AsaID)
			}
			collection.ExcludedAssets = assets
		case error:
			errs = append(errs, result)
		}
	}

	if len(errs) > 0 {
		msg := ""
		for _, err := range errs {
			msg += err.Error() + "\n"
		}

		return nil, errors.E(op, fmt.Errorf(msg))
	}

	return &collection, nil
}

func GetCollectionAssets[H db.Handle](h H, nfdID uint64, collectionID string) (*[]db.Asset, error) {
	const op errors.Op = "GetCollectionAssets"

	buffer := 2
	rChan := make(chan interface{}, buffer)
	defer close(rChan)

	go func() {
		wallets, err := db.GetNFDVerifiedWallets(h, nfdID)
		if err != nil && !db.ErrNoRows(err) {
			rChan <- err
			return
		}
		rChan <- wallets
	}()

	go func() {
		col, err := GetCollectionCriteria(h, collectionID)
		if err != nil && !db.ErrNoRows(err) {
			rChan <- err
			return
		}
		rChan <- col
	}()

	creators := []string{}
	var criteria *Collection
	var errs []error

	for i := 0; i < buffer; i++ {
		data := <-rChan
		switch result := data.(type) {
		case *[]db.NFDWallet:
			for i := range *result {
				wallet := (*result)[i]
				creators = append(creators, wallet.Address)
			}
		case *Collection:
			criteria = result
		case error:
			errs = append(errs, result)
		}
	}

	if len(errs) > 0 {
		msg := ""
		for _, err := range errs {
			msg += err.Error() + "\n"
		}

		return nil, errors.E(op, fmt.Errorf(msg))
	}

	assets, err := db.GetAssetsByDetails(h, creators, criteria.Prefixes, criteria.Assets, criteria.ExcludedAssets, db.Tiny)
	if err != nil {
		return nil, errors.E(op, err)
	}

	return assets, nil
}

func GetCollectionAssetsByNFDName[H db.Handle](h H, name string, collectionID string) (*[]db.Asset, error) {
	const op errors.Op = "GetCollectionAssetsByNFDName"

	buffer := 2
	rChan := make(chan interface{}, buffer)
	defer close(rChan)

	nfdData, err := db.GetNFDByName(h, name, db.Tiny)
	if err != nil {
		return nil, errors.E(op, err)
	}

	go func() {
		wallets, err := db.GetNFDVerifiedWallets(h, nfdData.ID)
		if err != nil && !db.ErrNoRows(err) {
			rChan <- err
			return
		}
		rChan <- wallets
	}()

	go func() {
		col, err := GetCollectionCriteria(h, collectionID)
		if err != nil && !db.ErrNoRows(err) {
			rChan <- err
			return
		}
		rChan <- col
	}()

	creators := []string{}
	var criteria *Collection
	var errs []error

	for i := 0; i < buffer; i++ {
		data := <-rChan
		switch result := data.(type) {
		case *[]db.NFDWallet:
			for i := range *result {
				wallet := (*result)[i]
				creators = append(creators, wallet.Address)
			}
		case *Collection:
			criteria = result
		case error:
			errs = append(errs, result)
		}
	}

	if len(errs) > 0 {
		msg := ""
		for _, err := range errs {
			msg += err.Error() + "\n"
		}

		return nil, errors.E(op, fmt.Errorf(msg))
	}

	assets, err := db.GetAssetsByDetails(h, creators, criteria.Prefixes, criteria.Assets, criteria.ExcludedAssets, db.Tiny)
	if err != nil {
		return nil, errors.E(op, err)
	}

	return assets, nil
}

func GetCollection[H db.Handle](h H, collectionID string) (*[]db.Asset, error) {
	const op errors.Op = "GetCollection"

	buffer := 2
	rChan := make(chan interface{}, buffer)
	defer close(rChan)

	col, err := db.GetCollection(h, collectionID)
	if err != nil {
		return nil, errors.E(op, err)
	}

	go func() {
		wallets, err := db.GetNFDVerifiedWallets(h, col.NFDID)
		if err != nil && !db.ErrNoRows(err) {
			rChan <- err
			return
		}
		rChan <- wallets
	}()

	go func() {
		col, err := GetCollectionCriteria(h, collectionID)
		if err != nil && !db.ErrNoRows(err) {
			rChan <- err
			return
		}
		rChan <- col
	}()

	creators := []string{}
	var criteria *Collection
	var errs []error

	for i := 0; i < buffer; i++ {
		data := <-rChan
		switch result := data.(type) {
		case *[]db.NFDWallet:
			for i := range *result {
				wallet := (*result)[i]
				creators = append(creators, wallet.Address)
			}
		case *Collection:
			criteria = result
		case error:
			errs = append(errs, result)
		}
	}

	if len(errs) > 0 {
		msg := ""
		for _, err := range errs {
			msg += err.Error() + "\n"
		}

		return nil, errors.E(op, fmt.Errorf(msg))
	}

	assets, err := db.GetAssetsByDetails(h, creators, criteria.Prefixes, criteria.Assets, criteria.ExcludedAssets, db.Tiny)
	if err != nil {
		return nil, errors.E(op, err)
	}

	return assets, nil
}
