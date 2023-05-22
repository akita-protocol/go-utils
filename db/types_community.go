package db

import (
	"database/sql"
	"fmt"
	"strings"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type CommunityTab string

const (
	CommunityTabCollections CommunityTab = "collections"
	CommunityTabAbout       CommunityTab = "about"
	CommunityTabStaking     CommunityTab = "staking"
	CommunityTabSocial      CommunityTab = "social"
	CommunityTabRaffles     CommunityTab = "raffles"
	CommunityTabShuffles    CommunityTab = "shuffles"
	CommunityTabGallery     CommunityTab = "gallery"
)

type Community struct {
	ID   uint64       `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	AKTA *uint64      `structs:"akta,omitempty" db:"akta" json:"akta,omitempty"`
	Tab  CommunityTab `structs:"tab,omitempty" db:"tab" json:"tab,omitempty"`
}

func CommunityTableKeys() []string {
	return []string{"id", "akta", "tab"}
}

func GetCommunities[H Handle](h H, start, limit uint64) (*[]Community, error) {
	const op errors.Op = "GetCommunities"
	query := fmt.Sprintf("select %v from %v.community limit ?, ? order by akta desc", strings.Join(CommunityTableKeys(), ","), NFDDatabase())

	var communities *[]Community
	err := h.Select(&communities, query, start, limit)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Community Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return communities, nil
}

func GetCommunity[H Handle](h H, id uint64) (*Community, error) {
	const op errors.Op = "GetCommunity"
	query := fmt.Sprintf("select %v from %v.community where id = ?", strings.Join(CommunityTableKeys(), ","), NFDDatabase())

	var community Community
	err := h.Get(&community, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Community Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &community, nil
}
