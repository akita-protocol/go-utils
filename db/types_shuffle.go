package db

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/akita-protocol/go-utils/errors"
)

// type ShuffleType string

// const (
// 	ShuffleTypeSlate ShuffleType = "slate" //blank arc19 slates that will get updated as they're purchased
// 	ShuffleTypeToken ShuffleType = "token" //token purchase for NFT redemption
// )

type Shuffle struct {
	ID    string  `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	NFDID *uint64 `structs:"nfd_id,omitempty" db:"nfd_id" json:"nfd_id,omitempty"`
	// Type              ShuffleType `structs:"type,omitempty" db:"type" json:"type,omitempty"`
	Title             string    `structs:"title,omitempty" db:"title" json:"title,omitempty"`
	About             string    `structs:"about,omitempty" db:"about" json:"about,omitempty"`
	Avatar            *string   `structs:"avatar,omitempty" db:"avatar" json:"avatar,omitempty"`
	AvatarID          *uint64   `structs:"avatar_id,omitempty" db:"avatar_id" json:"avatar_id,omitempty"`
	Banner            *string   `structs:"banner,omitempty" db:"banner" json:"banner,omitempty"`
	BannerID          *uint64   `structs:"banner_id,omitempty" db:"banner_id" json:"banner_id,omitempty"`
	MaxPerWallet      *uint64   `structs:"max_per_wallet,omitempty" db:"max_per_wallet" json:"max_per_wallet,omitempty"`
	MaxSold           *uint64   `structs:"max_sold,omitempty" db:"max_sold" json:"max_sold,omitempty"`
	ShuffleBeginsDate time.Time `structs:"shuffle_begins_date,omitempty" db:"shuffle_begins_date" json:"shuffle_begins_date,omitempty"`
	Creator           string    `structs:"creator,omitempty" db:"creator" json:"creator,omitempty"`
	CreatedDate       time.Time `structs:"created_date,omitempty" db:"created_date" json:"created_date,omitempty"`
	Address           string    `structs:"address,omitempty" db:"address" json:"address,omitempty"`
}

func ShuffleTableKeys() []string {
	return []string{
		"id",
		"nfd_id",
		"title",
		"about",
		"avatar",
		"avatar_id",
		"banner",
		"banner_id",
		"max_per_wallet",
		"max_sold",
		"shuffle_begins_date",
		"creator",
		"created_date",
		"address",
	}
}

func GetShuffle[H Handle](h H, id string) (*Shuffle, error) {
	const op errors.Op = "GetShuffle"
	query := fmt.Sprintf("select %v from %s.shuffle where id = ?", strings.Join(ShuffleTableKeys(), ", "), ShuffleDatabase())
	var shuffle Shuffle

	err := h.Get(&shuffle, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Shuffle Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &shuffle, err
}

func GetShufflesByNFDID[H Handle](h H, nfdID uint64) (*[]Shuffle, error) {
	const op errors.Op = "GetShufflesByNFDID"
	query := fmt.Sprintf("select %v from %s.shuffle where nfd_id = ?", strings.Join(ShuffleTableKeys(), ", "), ShuffleDatabase())
	var shuffles []Shuffle

	err := h.Select(&shuffles, query, nfdID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Shuffle Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &shuffles, err
}
