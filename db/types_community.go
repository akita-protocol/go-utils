package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
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
	ID      uint64       `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	Version string       `structs:"version,omitempty" db:"version" json:"version,omitempty"`
	AKTA    *uint64      `structs:"akta,omitempty" db:"akta" json:"akta,omitempty"`
	Tab     CommunityTab `structs:"tab,omitempty" db:"tab" json:"tab,omitempty"`
}

func CommunityTableKeys() []string {
	return []string{"id", "version", "akta", "coalesce(tab, 'about') as tab"}
}

func IsCommunity[H Handle](h H, id uint64) (bool, error) {
	const op errors.Op = "IsCommunity"
	query := fmt.Sprintf("select exists(select id from %v.community where id = ?)", NFDDatabase())

	var exists bool
	err := h.Get(&exists, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, errors.E(pkg, op, err)
	}

	return exists, nil
}

func GetCommunities[H Handle](h H, start, limit uint64) (*[]Community, error) {
	const op errors.Op = "GetCommunities"
	query := fmt.Sprintf("select %v from %v.community order by akta desc limit ?, ?", strings.Join(CommunityTableKeys(), ","), NFDDatabase())

	var communities []Community
	err := h.Select(&communities, query, start, limit)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Community Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &communities, nil
}

func GetAllCommunityVerifiedWallets[H Handle](h H) ([]string, error) {
	const op errors.Op = "GetAllCommunityVerifiedWallets"
	query := fmt.Sprintf("select address from %v.nfd_wallet where id in (select id from %v.community)", NFDDatabase(), NFDDatabase())

	var wallets []string
	err := h.Select(&wallets, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Community Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return wallets, nil
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

func DeleteCommunity[H Handle](h H, id uint64) error {
	const op errors.Op = "DeleteCommunity"
	query := fmt.Sprintf("delete from %v.community where id = ?", NFDDatabase())

	switch h := any(h).(type) {
	case *sqlx.Tx:
		stmt, err := h.Prepare(query)
		if err != nil {
			return errors.E(pkg, op, errors.Database, err, "Failed to Prepare Query")
		}

		_, err = stmt.Exec(id)
		if err != nil {
			return errors.E(pkg, op, errors.Database, err, "Failed to Execute Query")
		}
	case *sqlx.DB:
		_, err := h.Exec(query, id)
		if err != nil {
			return errors.E(pkg, op, errors.Database, err, "Failed to Execute Query")
		}
	}

	return nil
}

// *********************************************************************************************************************

func GetCommunityPower[H Handle](h H, id uint64) (uint64, error) {
	const op errors.Op = "GetCommunityPower"

	// this sql query is going to suck to write
	// i need to query -> the addresses that hold > 1 of the assets belonging to the community
	// both NFTs and ASAs are considered, and then expand the addresses to any addresses verified with an NFD that the wallets in the initial query hold
	// then i need to sum the holdings of akita of those addresses

	return 0, nil
}
