package db

import (
	"database/sql"
	"fmt"
	"strings"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type StakingMerchant struct {
	ID          string  `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	Name        string  `structs:"name,omitempty" db:"name" json:"name,omitempty"`
	Description string  `structs:"description,omitempty" db:"description" json:"description,omitempty"`
	Logo        string  `structs:"logo,omitempty" db:"logo" json:"logo,omitempty"`
	Website     string  `structs:"website,omitempty" db:"website" json:"website,omitempty"`
	Email       *string `structs:"email,omitempty" db:"email" json:"email,omitempty"`
	Phone       *string `structs:"phone,omitempty" db:"phone" json:"phone,omitempty"`
}

func StakingMerchantTableKeys() []string {
	return []string{"id", "name", "description", "logo", "website", "email", "phone"}
}

func GetStakingMerchantByID[H Handle](h H, id string) (*StakingMerchant, error) {
	const op errors.Op = "GetMerchantByID"
	query := fmt.Sprintf("select %s from %s.merchant where id = ?", strings.Join(StakingMerchantTableKeys(), ","), StakingDatabase())
	var merchant StakingMerchant
	err := h.Get(&merchant, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "Staking Merchant Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}
	return &merchant, nil
}
