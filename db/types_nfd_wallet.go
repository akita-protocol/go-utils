package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"gitlab.algotoolchain.com/keelakay/utils/errors"
	"gitlab.algotoolchain.com/keelakay/utils/misc"
)

type NFDWallet struct {
	// ID is the app id of the NFD
	ID       uint64 `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	Address  string `structs:"address,omitempty" db:"address" json:"address,omitempty"`
	Verified *bool  `structs:"verified,omitempty" db:"verified" json:"verified,omitempty"`
	Deposit  *bool  `structs:"deposit,omitempty" db:"deposit" json:"deposit,omitempty"`
}

func NFDWalletTableKeys() []string {
	return []string{"id", "address", "verified", "deposit"}
}

func GetAllNFDWallets[H DBStruct](h H) (*[]NFDWallet, error) {
	const op errors.Op = "GetAllNFDWallets"
	query := fmt.Sprintf("select %s from %s.nfd_wallet", strings.Join(NFDWalletTableKeys(), ","), NFDDatabase())
	var list []NFDWallet

	err := h.Select(&list, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "wallets not found")
		}
		return nil, errors.E(pkg, op, err)
	}

	if !(len(list) > 0) {
		return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, fmt.Errorf("wallets not found"))
	}

	return &list, nil
}

func GetNFDWallets[H DBStruct](h H, id uint64) (*[]NFDWallet, error) {
	const op errors.Op = "GetNFDWallets"
	query := fmt.Sprintf("select %s from %s.nfd_wallet where id = ?", strings.Join(NFDWalletTableKeys(), ","), NFDDatabase())
	var list []NFDWallet

	err := h.Select(&list, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "wallets not found")
		}
		return nil, errors.E(pkg, op, err)
	}

	if !(len(list) > 0) {
		return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, fmt.Errorf("wallets not found"))
	}

	return &list, nil
}

func GetNFDVerifiedWallets[H DBStruct](h H, id uint64) (*[]NFDWallet, error) {
	const op errors.Op = "GetNFDWallets"
	query := fmt.Sprintf("select %s from %s.nfd_wallet where id = ? and verified = 1 order by deposit desc", strings.Join(NFDWalletTableKeys(), ","), NFDDatabase())
	var list []NFDWallet

	err := h.Select(&list, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "wallets not found")
		}
		return nil, errors.E(pkg, op, err)
	}

	if !(len(list) > 0) {
		return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, fmt.Errorf("wallets not found"))
	}

	return &list, nil
}

func GetNFDVerifiedWalletAddressesByAdjacentAddresses[H DBStruct](h H, addresses []string) (*[]string, error) {
	const op errors.Op = "GetNFDVerifiedWalletAddressesByAdjacentAddresses"
	query := fmt.Sprintf("select distinct(address) from %s.nfd_wallet where verified = 1 and address in (%s)", NFDDatabase(), strings.Repeat("?, ", len(addresses))[0:(len(addresses)*3)-2])
	var list []string

	err := h.Select(&list, query, misc.ToInterfaceSlice(addresses)...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "wallets not found")
		}
		return nil, errors.E(pkg, op, err)
	}

	if !(len(list) > 0) {
		return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, fmt.Errorf("wallets not found"))
	}

	return &list, nil
}

func GetNFDVerifiedOrOwnerWallets[H DBStruct](h H, id uint64) (*[]NFDWallet, error) {
	const op errors.Op = "GetNFDWallets"
	query := fmt.Sprintf("select %s from %s.nfd_wallet where id = ? order by deposit desc", strings.Join(NFDWalletTableKeys(), ","), NFDDatabase())
	var list []NFDWallet
	var nlist []NFDWallet

	err := h.Select(&list, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "wallets not found")
		}
		return nil, errors.E(pkg, op, err)
	}

	if !(len(list) > 0) {
		return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, fmt.Errorf("wallets not found"))
	}

	if len(list) > 1 {
		for i := 1; i < len(list); i++ {
			if list[i].Verified != nil && *list[i].Verified {
				nlist = append(nlist, list[i])
			}
		}
	} else {
		return &list, nil
	}

	return &nlist, nil
}

func GetNFDWalletsByDiscordID[H DBStruct](h H, discordID string) (*[]NFDWallet, error) {
	const op errors.Op = "GetNFDWalletsByUserID"
	query := fmt.Sprintf("select %s from %s.nfd_wallet where id = (select id from %s.nfd where discord_id = ? limit 1)", strings.Join(NFDWalletTableKeys(), ","), NFDDatabase(), NFDDatabase())
	var list []NFDWallet

	err := h.Select(&list, query, discordID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "nfdwallets not found")
		}
		return nil, errors.E(pkg, op, err)
	}

	if !(len(list) > 0) {
		return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, fmt.Errorf("nfdwallets not found"))
	}

	return &list, nil
}

func DeleteNFDWalletsByIDAndAddress[H DBStruct](h H, id uint64, address string) error {
	const op errors.Op = "DeleteNFDWalletsByIDAndAddress"
	query := fmt.Sprintf("delete from %s.nfd_wallet where id = ? and address = ?", NFDDatabase())
	var err error

	switch h := any(h).(type) {
	case *sqlx.Tx:
		stmt, err := h.Prepare(query)
		if err != nil {
			return errors.E(pkg, op, errors.Database, err, "Failed to Prepare Query")
		}

		_, err = stmt.Exec(id, address)
		if err != nil {
			return errors.E(pkg, op, errors.Database, err, "Failed to Execute Query")
		}
	case *sqlx.DB:
		_, err = h.Exec(query, id, address)
		if err != nil {
			return errors.E(pkg, op, errors.Database, err, "Failed to Execute Query")
		}
	}

	return nil
}

// DeleteNFDWalletNotIn deletes collections that are not included in a list for a given NFD
func DeleteNFDWalletNotIn[H Handle](h H, id uint64, addresses ...string) error {
	const op errors.Op = "DeleteNFDWalletNotIn"
	var err error
	data := append([]interface{}{id}, misc.ToInterfaceSlice(addresses)...)
	fmt.Println(data)
	query := fmt.Sprintf("delete from %s.nfd_wallet where id = ?", NFDDatabase())

	if len(addresses) > 0 {
		qMarks := []rune(strings.Repeat("?, ", len(addresses)))
		query += fmt.Sprintf(" and address not in (%s)", string(qMarks[0:len(qMarks)-2]))
	}

	switch h := any(h).(type) {
	case *sqlx.Tx:
		stmt, err := h.Prepare(query)
		if err != nil {
			return errors.E(pkg, op, errors.Database, err, "Failed to Prepare Query")
		}

		_, err = stmt.Exec(data...)
		if err != nil {
			return errors.E(pkg, op, errors.Database, err, "Failed to Execute Query")
		}
	case *sqlx.DB:
		_, err = h.Exec(query, data...)
		if err != nil {
			return errors.E(pkg, op, errors.Database, err, "Failed to Execute Query")
		}
	}

	return nil
}
