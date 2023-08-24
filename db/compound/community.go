package compound

import (
	"fmt"
	"strings"

	"gitlab.algotoolchain.com/keelakay/utils/db"
	"gitlab.algotoolchain.com/keelakay/utils/errors"
	"gitlab.algotoolchain.com/keelakay/utils/misc"
)

type Community struct {
	*db.Community
	Tokens      []db.CommunityToken     `json:"tokens,omitempty"`
	Associates  []db.CommunityAssociate `json:"associates,omitempty"`
	Collections []Collection            `json:"collections,omitempty"`
	Faq         []db.CommunityFaq       `json:"faq,omitempty"`
	Extras      []db.CommunityExtras    `json:"extras,omitempty"`
	// Staking     []db.Staking         `json:"staking,omitempty"`
}

type CommunityGetExclude string

const (
	CommunityGetExcludeTokens      CommunityGetExclude = "tokens"
	CommunityGetExcludeCollections CommunityGetExclude = "collections"
	CommunityGetExcludeFaq         CommunityGetExclude = "faq"
	// CommunityGetExcludeStaking     CommunityGetExclude = "staking"
	CommunityGetExcludeExtras CommunityGetExclude = "extras"
)

func GetCommunity[H db.Handle](h H, nfdID uint64, exclude ...CommunityGetExclude) (*Community, error) {
	const op errors.Op = "GetCommunitiesByNFDID"
	var community Community
	buffer := (5 - len(exclude))
	rChan := make(chan interface{}, buffer)
	defer close(rChan)

	go func() {
		comm, err := db.GetCommunity(h, nfdID)
		if err != nil {
			rChan <- err
			return
		}
		rChan <- comm
	}()

	if !misc.InSlice(CommunityGetExcludeTokens, exclude) {
		go func() {
			tokens, err := db.GetCommunityTokens(h, nfdID)
			if err != nil {
				rChan <- err
				return
			}
			rChan <- tokens
		}()
	}

	if !misc.InSlice(CommunityGetExcludeCollections, exclude) {
		go func() {
			collections, err := GetCollectionsByNFDID(h, nfdID)
			if err != nil {
				rChan <- err
				return
			}
			rChan <- collections
		}()
	}

	if !misc.InSlice(CommunityGetExcludeFaq, exclude) {
		go func() {
			faq, err := db.GetCommunityFaq(h, nfdID, 0, 10)
			if err != nil {
				rChan <- err
				return
			}
			rChan <- faq
		}()
	}

	if !misc.InSlice(CommunityGetExcludeExtras, exclude) {
		go func() {
			extras, err := db.GetCommunityExtras(h, nfdID)
			if err != nil {
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
		case *db.Community:
			community.Community = result
		case *[]db.CommunityToken:
			community.Tokens = *result
		case *[]Collection:
			community.Collections = *result
		case *[]db.CommunityFaq:
			community.Faq = *result
		case *[]db.CommunityExtras:
			community.Extras = *result
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
			if db.ErrNoRows(err) && strings.Contains(err.Error(), "operation: GetCommunity") {
				return nil, err
			}

			msg += err.Error() + "\n"
		}

		return nil, errors.E(op, fmt.Errorf(msg))
	}

	return &community, nil
}

func DeleteCommunity[H db.Handle](h H, nfdID uint64) error {
	const op errors.Op = "DeleteCommunity"
	var err error

	err = db.DeleteCommunity(h, nfdID)
	if err != nil {
		return errors.E(op, err)
	}

	// community_json
	err = db.DeleteCommunityJson(h, nfdID)
	if err != nil {
		return errors.E(op, err)
	}

	// community_tokens
	err = db.DeleteCommunityTokens(h, nfdID)
	if err != nil {
		return errors.E(op, err)
	}

	// community_faq
	err = db.DeleteCommunityFaq(h, nfdID)
	if err != nil {
		return errors.E(op, err)
	}

	// community_extras
	err = db.DeleteCommunityExtras(h, nfdID)
	if err != nil {
		return errors.E(op, err)
	}

	// collection
	collections, err := db.GetCollectionsByNFDID(h, nfdID)
	if err != nil {
		return errors.E(op, err)
	}

	for _, collection := range *collections {

		err = db.DeleteCollectionPrefixes(h, collection.ID)
		if err != nil {
			return errors.E(op, err)
		}

		err = db.DeleteCollectionArtists(h, collection.ID)
		if err != nil {
			return errors.E(op, err)
		}

		err = db.DeleteCollectionAssets(h, collection.ID)
		if err != nil {
			return errors.E(op, err)
		}

		err = db.DeleteCollectionExcludedAssets(h, collection.ID)
		if err != nil {
			return errors.E(op, err)
		}

		err = db.DeleteCollectionExtras(h, collection.ID)
		if err != nil {
			return errors.E(op, err)
		}

		properties, err := db.GetProperties(h, collection.ID)
		if err != nil {
			return errors.E(op, err)
		}

		for i := range *properties {
			property := (*properties)[i]

			err = db.DeletePropertyValues(h, property.ID)
			if err != nil {
				return errors.E(op, err)
			}

			err = db.DeletePropertyValueExtras(h, property.ID)
			if err != nil {
				return errors.E(op, err)
			}
		}

		err = db.DeleteCollectionProperties(h, collection.ID)
		if err != nil {
			return errors.E(op, err)
		}
	}

	err = db.DeleteCollectionsByNFDID(h, nfdID)
	if err != nil {
		return errors.E(op, err)
	}

	return nil
}

// *************************************************************

func GetCommunityAssetIDs[H db.Handle](h H, nfdID uint64) ([]uint64, error) {
	const op errors.Op = "GetCommunityAssetIDs"

	buffer := 2
	rChan := make(chan interface{}, buffer)
	defer close(rChan)

	go func() {
		var assets = []uint64{}
		cols, err := db.GetCollectionsByNFDID(h, nfdID)
		if err != nil {
			rChan <- err
			return
		}

		for i := range *cols {
			col := (*cols)[i]

			colAssets, err := GetCollectionAssets(h, nfdID, col.ID)
			if err != nil {
				rChan <- err
				return
			}

			for i := range *colAssets {
				asset := (*colAssets)[i]
				assets = append(assets, asset.ID)
			}
		}

		rChan <- assets
	}()

	go func() {
		tokens, err := db.GetCommunityTokens(h, nfdID)
		if err != nil {
			rChan <- err
			return
		}
		rChan <- tokens
	}()

	var errs []error
	var ids = []uint64{}
	for i := 0; i < buffer; i++ {
		data := <-rChan
		switch result := data.(type) {
		case *[]uint64:
			ids = append(ids, *result...)
		case *[]db.CommunityToken:
			for _, token := range *result {
				ids = append(ids, token.AssetID)
			}
		}
	}

	if len(errs) > 0 {
		msg := ""
		for _, err := range errs {
			if db.ErrNoRows(err) && strings.Contains(err.Error(), "operation: GetCommunity") {
				return nil, err
			}

			msg += err.Error() + "\n"
		}

		return nil, errors.E(op, fmt.Errorf(msg))
	}

	ids = misc.UniqueSlice(ids)
	return ids, nil
}
