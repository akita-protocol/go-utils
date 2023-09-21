package compound

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/algorand/go-algorand-sdk/v2/crypto"
	"gitlab.algotoolchain.com/keelakay/utils/db"
	"gitlab.algotoolchain.com/keelakay/utils/errors"
	"gitlab.algotoolchain.com/keelakay/utils/misc"
)

type NFDWallet struct {
	*db.NFDWallet
	Vault bool `json:"vault,omitempty"`
}

func GetAllNFDWallets[H db.DBStruct](h H) (*[]NFDWallet, error) {
	const op errors.Op = "GetAllNFDWallets"
	query := fmt.Sprintf("select %s from %s.nfd_wallet", strings.Join(db.NFDWalletTableKeys(), ","), db.NFDDatabase())
	var list []db.NFDWallet

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

	idMap := map[uint64]struct{}{}
	nlist := []NFDWallet{}
	for _, wallet := range list {
		_, exists := idMap[wallet.ID]
		if !exists {
			idMap[wallet.ID] = struct{}{}
			addr := crypto.GetApplicationAddress(wallet.ID)
			vault := NFDWallet{
				NFDWallet: &db.NFDWallet{
					ID:       wallet.ID,
					Address:  addr.String(),
					Verified: misc.PointerBool(true),
					Deposit:  misc.PointerBool(false),
					Owner:    misc.PointerBool(false),
				},
				Vault: true,
			}
			nlist = append(nlist, vault)
		}
		nlist = append(nlist, NFDWallet{NFDWallet: &wallet})
	}

	return &nlist, nil
}

func GetNFDWallets[H db.DBStruct](h H, id uint64) (*[]NFDWallet, error) {
	const op errors.Op = "GetNFDWallets"
	query := fmt.Sprintf("select %s from %s.nfd_wallet where id = ?", strings.Join(db.NFDWalletTableKeys(), ","), db.NFDDatabase())
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

	addr := crypto.GetApplicationAddress(id)
	vault := NFDWallet{
		NFDWallet: &db.NFDWallet{
			ID:       id,
			Address:  addr.String(),
			Verified: misc.PointerBool(true),
			Deposit:  misc.PointerBool(false),
			Owner:    misc.PointerBool(false),
		},
		Vault: true,
	}

	nlist := append([]NFDWallet{vault}, list...)

	return &nlist, nil
}

func GetNFDVerifiedWallets[H db.DBStruct](h H, id uint64) (*[]NFDWallet, error) {
	const op errors.Op = "GetNFDWallets"
	query := fmt.Sprintf("select %s from %s.nfd_wallet where id = ? and verified = 1 order by deposit desc", strings.Join(db.NFDWalletTableKeys(), ","), db.NFDDatabase())
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

	addr := crypto.GetApplicationAddress(id)
	vault := NFDWallet{
		NFDWallet: &db.NFDWallet{
			ID:       id,
			Address:  addr.String(),
			Verified: misc.PointerBool(true),
			Deposit:  misc.PointerBool(false),
			Owner:    misc.PointerBool(false),
		},
		Vault: true,
	}

	nlist := append([]NFDWallet{vault}, list...)

	return &nlist, nil
}

func GetNFDVerifiedWalletAddressesByAdjacentAddresses[H db.DBStruct](h H, addresses []string) (*[]string, error) {
	const op errors.Op = "GetNFDVerifiedWalletAddressesByAdjacentAddresses"
	query := fmt.Sprintf("select %v from %s.nfd_wallet where verified = 1 and id in (select id from %v.nfd_wallet where address in (%s))", strings.Join(db.NFDWalletTableKeys(), ", "), db.NFDDatabase(), db.NFDDatabase(), strings.Repeat("?, ", len(addresses))[0:(len(addresses)*3)-2])
	var list []db.NFDWallet

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

	idMap := map[uint64]struct{}{}
	nlist := []string{}
	for _, wallet := range list {
		_, exists := idMap[wallet.ID]
		if !exists {
			idMap[wallet.ID] = struct{}{}
			addr := crypto.GetApplicationAddress(wallet.ID)
			nlist = append(nlist, addr.String())
		}
		nlist = append(nlist, wallet.Address)
	}
	nlist = misc.UniqueSlice(nlist)

	return &nlist, nil
}

func GetNFDVerifiedOrOwnerWallets[H db.DBStruct](h H, id uint64) (*[]NFDWallet, error) {
	const op errors.Op = "GetNFDWallets"
	query := fmt.Sprintf("select %s from %s.nfd_wallet where id = ? order by deposit desc", strings.Join(db.NFDWalletTableKeys(), ","), db.NFDDatabase())
	var list []db.NFDWallet
	nlist := []NFDWallet{}

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
		idMap := map[uint64]struct{}{}
		for _, wallet := range list {
			_, exists := idMap[wallet.ID]
			if !exists {
				idMap[wallet.ID] = struct{}{}
				addr := crypto.GetApplicationAddress(wallet.ID)
				vault := NFDWallet{
					NFDWallet: &db.NFDWallet{
						ID:       wallet.ID,
						Address:  addr.String(),
						Verified: misc.PointerBool(true),
						Deposit:  misc.PointerBool(false),
						Owner:    misc.PointerBool(false),
					},
					Vault: true,
				}
				nlist = append(nlist, vault)
			}

			if wallet.Verified != nil && *wallet.Verified {
				nlist = append(nlist, NFDWallet{NFDWallet: &wallet})
			}
		}
	} else {
		return &[]NFDWallet{{NFDWallet: &list[0]}}, nil
	}

	return &nlist, nil
}

func GetNFDWalletsByDiscordID[H db.DBStruct](h H, discordID string) (*[]NFDWallet, error) {
	const op errors.Op = "GetNFDWalletsByUserID"
	query := fmt.Sprintf("select %s from %s.nfd_wallet where id = (select id from %s.nfd where discord_id = ? order by last_changed_block desc limit 1) and (verified = 1 or owner = 1)", strings.Join(db.NFDWalletTableKeys(), ","), db.NFDDatabase(), db.NFDDatabase())
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

	addr := crypto.GetApplicationAddress(list[0].ID)
	vault := NFDWallet{
		NFDWallet: &db.NFDWallet{
			ID:       list[0].ID,
			Address:  addr.String(),
			Verified: misc.PointerBool(true),
			Deposit:  misc.PointerBool(false),
			Owner:    misc.PointerBool(false),
		},
		Vault: true,
	}

	nlist := append([]NFDWallet{vault}, list...)

	return &nlist, nil
}
