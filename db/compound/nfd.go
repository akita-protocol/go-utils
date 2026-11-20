package compound

import (
	"fmt"

	"github.com/akita-protocol/go-utils/db"
	"github.com/akita-protocol/go-utils/errors"
	"github.com/akita-protocol/go-utils/misc"
)

type NFD struct {
	*db.NFD
	BannerURL  *string `json:"banner_url,omitempty"`
	BannerMime *string `json:"banner_mime,omitempty"`
	AvatarURL  *string `json:"avatar_url,omitempty"`
	AvatarMime *string `json:"avatar_mime,omitempty"`
}

type NFDGetExclude string

const (
	NFDGetExcludeBanner NFDGetExclude = "banner"
	NFDGetExcludeAvatar NFDGetExclude = "avatar"
)

func GetNFD[H db.Handle](h H, nfdID uint64, st db.TableKeySlice, exclude ...NFDGetExclude) (*NFD, error) {
	const op errors.Op = "GetNFD"

	nfd, err := db.GetNFD(h, nfdID, st)
	if err != nil {
		return nil, errors.E(op, err)
	}

	if nfd == nil {
		return nil, errors.E(op, errors.DatabaseResultNotFound, fmt.Errorf("NFD Not Found"))
	}

	compNFD := &NFD{
		NFD: nfd,
	}
	buffer := (2 - len(exclude))
	if buffer == 0 {
		return compNFD, nil
	}

	rChan := make(chan interface{}, buffer)
	defer close(rChan)

	if !misc.InSlice(NFDGetExcludeBanner, exclude) {
		if nfd.BannerID != nil {
			go func() {
				bannerAsset, err := db.GetAsset(h, *nfd.BannerID, db.TKSTiny)
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

	if !misc.InSlice(NFDGetExcludeAvatar, exclude) {
		if nfd.AvatarID != nil {
			go func() {
				avatarAsset, err := db.GetAsset(h, *nfd.AvatarID, db.TKSTiny)
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

	var errs []error
	for i := 0; i < buffer; i++ {
		data := <-rChan
		switch result := data.(type) {
		case LabeledAsset:
			switch result.Label {
			case "banner":
				compNFD.BannerMime = result.Asset.Mime
				if result.Asset.ImageURL != nil && *result.Asset.ImageURL != "" {
					compNFD.BannerURL = result.Asset.ImageURL
				} else {
					compNFD.BannerURL = result.Asset.URL
				}
			case "avatar":
				compNFD.AvatarMime = result.Asset.Mime
				if result.Asset.ImageURL != nil && *result.Asset.ImageURL != "" {
					compNFD.AvatarURL = result.Asset.ImageURL
				} else {
					compNFD.AvatarURL = result.Asset.URL
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

	return compNFD, nil
}

func GetNFDByAddress[H db.Handle](h H, address string, st db.TableKeySlice, exclude ...NFDGetExclude) (*NFD, error) {
	const op errors.Op = "GetNFDByAddress"

	nfds, err := db.GetNFDsByAddress(h, address, st)
	if err != nil {
		return nil, err
	}

	if nfds == nil || len(*nfds) == 0 {
		return nil, errors.E(op, errors.DatabaseResultNotFound, fmt.Errorf("NFD Not Found"))
	}

	nfd := &(*nfds)[0]

	compNFD := &NFD{
		NFD: nfd,
	}
	buffer := (2 - len(exclude))
	if buffer == 0 {
		return compNFD, nil
	}

	rChan := make(chan interface{}, buffer)
	defer close(rChan)

	if !misc.InSlice(NFDGetExcludeBanner, exclude) {
		if nfd.BannerID != nil {
			go func() {
				bannerAsset, err := db.GetAsset(h, *nfd.BannerID, db.TKSTiny)
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

	if !misc.InSlice(NFDGetExcludeAvatar, exclude) {
		if nfd.AvatarID != nil {
			go func() {
				avatarAsset, err := db.GetAsset(h, *nfd.AvatarID, db.TKSTiny)
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

	var errs []error
	for i := 0; i < buffer; i++ {
		data := <-rChan
		switch result := data.(type) {
		case LabeledAsset:
			switch result.Label {
			case "banner":
				compNFD.BannerMime = result.Asset.Mime
				if result.Asset.ImageURL != nil && *result.Asset.ImageURL != "" {
					compNFD.BannerURL = result.Asset.ImageURL
				} else {
					compNFD.BannerURL = result.Asset.URL
				}
			case "avatar":
				compNFD.AvatarMime = result.Asset.Mime
				if result.Asset.ImageURL != nil && *result.Asset.ImageURL != "" {
					compNFD.AvatarURL = result.Asset.ImageURL
				} else {
					compNFD.AvatarURL = result.Asset.URL
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

	return compNFD, nil
}

func GetNFDByName[H db.Handle](h H, name string, st db.TableKeySlice, exclude ...NFDGetExclude) (*NFD, error) {
	const op errors.Op = "GetNFD"

	nfd, err := db.GetNFDByName(h, name, st)
	if err != nil {
		return nil, errors.E(op, err)
	}

	if nfd == nil {
		return nil, errors.E(op, errors.DatabaseResultNotFound, fmt.Errorf("NFD Not Found"))
	}

	compNFD := &NFD{
		NFD: nfd,
	}
	buffer := (2 - len(exclude))
	if buffer == 0 {
		return compNFD, nil
	}

	rChan := make(chan interface{}, buffer)
	defer close(rChan)

	if !misc.InSlice(NFDGetExcludeBanner, exclude) {
		if nfd.BannerID != nil {
			go func() {
				bannerAsset, err := db.GetAsset(h, *nfd.BannerID, db.TKSTiny)
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

	if !misc.InSlice(NFDGetExcludeAvatar, exclude) {
		if nfd.AvatarID != nil {
			go func() {
				avatarAsset, err := db.GetAsset(h, *nfd.AvatarID, db.TKSTiny)
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

	var errs []error
	for i := 0; i < buffer; i++ {
		data := <-rChan
		switch result := data.(type) {
		case LabeledAsset:
			switch result.Label {
			case "banner":
				compNFD.BannerMime = result.Asset.Mime
				if result.Asset.ImageURL != nil && *result.Asset.ImageURL != "" {
					compNFD.BannerURL = result.Asset.ImageURL
				} else {
					compNFD.BannerURL = result.Asset.URL
				}
			case "avatar":
				compNFD.AvatarMime = result.Asset.Mime
				if result.Asset.ImageURL != nil && *result.Asset.ImageURL != "" {
					compNFD.AvatarURL = result.Asset.ImageURL
				} else {
					compNFD.AvatarURL = result.Asset.URL
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

	return compNFD, nil
}
