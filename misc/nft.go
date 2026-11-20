package misc

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/akita-protocol/go-utils/errors"
	"github.com/akita-protocol/go-utils/ipfs"
	"github.com/algorand/go-algorand-sdk/v2/client/v2/indexer"
	"github.com/algorand/go-algorand-sdk/v2/types"
)

const pkg errors.Pkg = "misc"

type NFTPreview struct {
	AsaID        uint64 `json:"asa_id"`
	Creator      string `json:"creator"`
	Name         string `json:"name"`
	Mime         string `json:"mime"`
	UnitName     string `json:"unit_name"`
	URL          string `json:"url"`
	ImageURL     string `json:"image_url"`
	TotalCreated uint64 `json:"total_created"`
	Decimals     uint64 `json:"decimals"`
}

func GetAsaPreviews(client *indexer.Client, asaIDs []uint64, includeNFDs bool) (*[]NFTPreview, error) {
	const op errors.Op = "GetAsaPreviews"
	var nfts []NFTPreview
	var fetchChan = make(chan interface{}, len(asaIDs))
	var actual int = 0

	for _, id := range asaIDs {
		actual++
		go func(id uint64) {
			_, asa, err := client.LookupAssetByID(id).Do(context.Background())
			if err != nil {
				err = errors.E(pkg, op, err, "failed to fetch asa")
				fetchChan <- err
				return
			}

			if asa.Deleted || !IsNFT(asa.Params.Url) || (!includeNFDs && asa.Params.UnitName == "NFD") {
				fetchChan <- nil
				return
			}

			nft := NFTPreview{
				AsaID:        asa.Index,
				Creator:      asa.Params.Creator,
				Name:         asa.Params.Name,
				UnitName:     asa.Params.UnitName,
				URL:          asa.Params.Url,
				ImageURL:     asa.Params.Url,
				TotalCreated: asa.Params.Total,
				Decimals:     asa.Params.Decimals,
			}

			if strings.Contains(asa.Params.Url, "template-") && asa.Params.UnitName != "NFD" {
				reserve, err := types.DecodeAddress(asa.Params.Reserve)
				if err != nil {
					err = errors.E(pkg, op, err, "failed to decode reserve address")
					fetchChan <- err
					return
				}

				asa.Params.Url, err = ipfs.ParseASAUrl(asa.Params.Url, reserve)
				if err != nil {
					err = errors.E(pkg, op, err, "failed to parse asa url")
					fetchChan <- err
					return
				}

				nft.ImageURL = asa.Params.Url
			} else if strings.Contains(asa.Params.Url, "template-") && asa.Params.UnitName == "NFD" {
				reserve, err := types.DecodeAddress(asa.Params.Reserve)
				if err != nil {
					err = errors.E(pkg, op, err, "failed to decode reserve address")
					fetchChan <- err
					return
				}

				nfdUrl, err := ipfs.ParseASAUrl(asa.Params.Url, reserve)
				if err != nil {
					err = errors.E(pkg, op, err, "failed to parse asa url")
					fetchChan <- err
					return
				}

				nfdData, err := GetNFDData(nfdUrl)
				if err != nil {
					err = errors.E(pkg, op, err, "failed to get nfd data")
					fetchChan <- err
					return
				}

				nft.ImageURL = nfdData.Properties.Verified.Avatar
			}

			if strings.HasSuffix(asa.Params.Url, "#arc3") {

				var arcData struct {
					Image string `json:"image"`
				}

				arc3resp, err := http.Get(GetUrl(asa.Params.Url))
				if err != nil {
					err = errors.E(pkg, op, err, "http.Get from cdn")
					fetchChan <- err
					return
				}
				defer arc3resp.Body.Close()

				if arc3resp.StatusCode != 200 {
					arc3resp, err = http.Get(GetUrl(asa.Params.Url))
					if err != nil {
						err = errors.E(pkg, op, err, "http.Get from ipfs.algonft.tools")
						fetchChan <- err
						return
					}

					if arc3resp.StatusCode != 200 {
						fmt.Println(GetUrl(asa.Params.Url))
						err = errors.E(pkg, op, "http.Get from ipfs.algonft.tools")
						fetchChan <- err
						return
					}
				}

				ipfsData, err := ioutil.ReadAll(arc3resp.Body)
				if err != nil {
					err = errors.E(pkg, op, err, "ioutil.ReadAll failed")
					fetchChan <- err
					return
				}

				err = json.Unmarshal(ipfsData, &arcData)
				if err != nil {
					err = errors.E(pkg, op, err, "json.Unmarshal failed")
					fetchChan <- err
					return
				}

				nft.ImageURL = GetUrl(arcData.Image)
			} else {
				nft.ImageURL = GetUrl(asa.Params.Url)
			}

			fetchChan <- &nft
		}(id)
	}

	for i := 0; i < actual; i++ {
		data := <-fetchChan
		switch result := data.(type) {
		case error:
			return nil, result
		case *NFTPreview:
			nfts = append(nfts, *result)
		default:
			// skipped
		}
	}

	return &nfts, nil
}

func IsNFT(url string) bool {
	matches := []string{
		"ipfs/",
		"ipfs://",
		".png",
		".jpg",
		".jpeg",
		".jpeg",
		".webp",
		".gif",
		".svg",
		".mp4",
		".mp3",
		".wav",
		".webm",
		".mov",
		".m4a",
		".m4v",
		".json",
	}

	for _, match := range matches {
		if strings.Contains(url, match) {
			return true
		}
	}

	return false
}

func GetUrl(url string) string {
	cid, err := ipfs.GetCID(url)
	if err != nil {
		return url
	}
	return fmt.Sprintf("https://ipfs.algonft.tools/ipfs/%s", cid)
}

func GetNFDData(ipfsUrl string) (*NFD, error) {
	const op errors.Op = "GetNFDData"

	data, err := ipfs.GetIPFSData(ipfsUrl)
	if err != nil {
		return nil, errors.E(op, err)
	}

	var nfd NFD
	err = json.Unmarshal(data, &nfd)
	if err != nil {
		return nil, errors.E(op, err)
	}

	return &nfd, nil
}

type NFD struct {
	Name             string     `json:"name"`
	AppID            int64      `json:"appID"`
	State            string     `json:"state"`
	TimeCreated      string     `json:"timeCreated"`
	TimeChanged      string     `json:"timeChanged"`
	TimePurchased    string     `json:"timePurchased"`
	CurrentAsOfBlock int64      `json:"currentAsOfBlock"`
	AsaID            int64      `json:"asaID"`
	NfdAccount       string     `json:"nfdAccount"`
	Owner            string     `json:"owner"`
	Seller           string     `json:"seller"`
	MetaTags         []string   `json:"metaTags"`
	Category         string     `json:"category"`
	SaleType         string     `json:"saleType"`
	CAAlgo           []string   `json:"caAlgo"`
	UnverifiedCAAlgo []string   `json:"unverifiedCaAlgo"`
	DepositAccount   string     `json:"depositAccount"`
	Properties       Properties `json:"properties"`
	Asa              []ASA      `json:"ASA"`
	Image            string     `json:"image"`
}

type ASA struct {
	Index          int64  `json:"index"`
	Params         Params `json:"Params"`
	CachedImageKey string `json:"cachedImageKey"`
}

type Params struct {
	Creator     string `json:"creator"`
	Decimals    int64  `json:"decimals"`
	Manager     string `json:"manager"`
	Name        string `json:"name"`
	NameB64     string `json:"name-b64"`
	Reserve     string `json:"reserve"`
	Total       int64  `json:"total"`
	UnitName    string `json:"unit-name"`
	UnitNameB64 string `json:"unit-name-b64"`
	URL         string `json:"url"`
	URLB64      string `json:"url-b64"`
}

type Properties struct {
	UserDefined UserDefined `json:"userDefined"`
	Verified    Verified    `json:"verified"`
}

type UserDefined struct {
	Bio    string `json:"bio"`
	Caalgo string `json:"caalgo"`
	URL    string `json:"url"`
}

type Verified struct {
	Avatar      string `json:"avatar"`
	AvatarAsaID string `json:"avatarasaid"`
	Banner      string `json:"banner"`
	BannerAsaID string `json:"bannerasaid"`
	CAAlgo      string `json:"caAlgo"`
	Discord     string `json:"discord"`
	Twitter     string `json:"twitter"`
}
