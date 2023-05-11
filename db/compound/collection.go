package compound

import (
	"fmt"

	"gitlab.algotoolchain.com/keelakay/utils/db"
	"gitlab.algotoolchain.com/keelakay/utils/errors"
	"gitlab.algotoolchain.com/keelakay/utils/misc"
)

type Collection struct {
	*db.Collection
	Prefixes       []string              `json:"prefixes,omitempty"`
	Assets         []uint64              `json:"assets,omitempty"`
	ExcludedAssets []uint64              `json:"excluded_assets,omitempty"`
	Artists        []string              `json:"artists,omitempty"`
	Properties     []Property            `json:"properties,omitempty"`
	Extras         []db.CollectionExtras `json:"extras,omitempty"`
}

type CollectionGetExclude string

const (
	CollectionGetExcludePrefixes       CollectionGetExclude = "prefixes"
	CollectionGetExcludeAssets         CollectionGetExclude = "assets"
	CollectionGetExcludeExcludedAssets CollectionGetExclude = "excluded_assets"
	CollectionGetExcludeArtists        CollectionGetExclude = "artists"
	CollectionGetExcludeProperties     CollectionGetExclude = "properties"
	CollectionGetExcludeExtras         CollectionGetExclude = "extras"
)

func GetCollectionsByNFDID[H db.Handle](h H, nfdID string, exclude ...CollectionGetExclude) (*[]Collection, error) {
	const op errors.Op = "GetCollectionsByNFDID"

	var collections []Collection
	cols, err := db.GetCollectionsByNFDID(h, nfdID)
	if err != nil {
		return nil, errors.E(op, err)
	}

	for _, col := range *cols {
		compCollection := Collection{
			Collection: &col,
		}
		buffer := (6 - len(exclude))
		rChan := make(chan interface{}, buffer)
		defer close(rChan)

		if !misc.InSlice(CollectionGetExcludePrefixes, exclude) {
			go func() {
				prefixes, err := db.GetCollectionPrefixes(h, col.ID)
				if err != nil && err.(*errors.Error).Kind != errors.DatabaseResultNotFound {
					rChan <- err
					return
				}
				rChan <- prefixes
			}()
		}

		if !misc.InSlice(CollectionGetExcludeAssets, exclude) {
			go func() {
				assets, err := db.GetCollectionAssets(h, col.ID)
				if err != nil && err.(*errors.Error).Kind != errors.DatabaseResultNotFound {
					rChan <- err
					return
				}
				rChan <- assets
			}()
		}

		if !misc.InSlice(CollectionGetExcludeExcludedAssets, exclude) {
			go func() {
				excludedAssets, err := db.GetCollectionExcludedAssets(h, col.ID)
				if err != nil && err.(*errors.Error).Kind != errors.DatabaseResultNotFound {
					rChan <- err
					return
				}
				rChan <- excludedAssets
			}()
		}

		if !misc.InSlice(CollectionGetExcludeArtists, exclude) {
			go func() {
				artists, err := db.GetCollectionArtistByCollection(h, col.ID)
				if err != nil && err.(*errors.Error).Kind != errors.DatabaseResultNotFound {
					rChan <- err
					return
				}
				rChan <- artists
			}()
		}

		if !misc.InSlice(CollectionGetExcludeProperties, exclude) {
			go func() {
				properties, err := GetProperties(h, col.ID)
				if err != nil && err.(*errors.Error).Kind != errors.DatabaseResultNotFound {
					rChan <- err
					return
				}
				rChan <- properties
			}()
		}

		if !misc.InSlice(CollectionGetExcludeExtras, exclude) {
			go func() {
				extras, err := db.GetCollectionExtras(h, col.ID)
				if err != nil && err.(*errors.Error).Kind != errors.DatabaseResultNotFound {
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
				for _, prefix := range *result {
					prefixes = append(prefixes, prefix.Prefix)
				}
				compCollection.Prefixes = prefixes
			case *[]db.CollectionAsset:
				assets := []uint64{}
				for _, asset := range *result {
					assets = append(assets, asset.AsaID)
				}
				compCollection.Assets = assets
			case *[]db.CollectionExcludedAsset:
				assets := []uint64{}
				for _, asset := range *result {
					assets = append(assets, asset.AsaID)
				}
				compCollection.ExcludedAssets = assets
			case *[]db.CollectionArtist:
				artists := []string{}
				for _, artist := range *result {
					artists = append(artists, artist.Address)
				}
				compCollection.Artists = artists
			case *[]Property:
				compCollection.Properties = *result
			case *[]db.CollectionExtras:
				compCollection.Extras = *result
			case error:
				errs = append(errs, err)
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
