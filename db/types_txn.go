package db

import "gitlab.keelakay.com/keelakay/utils/uuid"

type Transaction struct {
	ID          string  `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	TxID        *string `structs:"tx_id,omitempty" db:"tx_id" json:"tx_id,omitempty"`
	FromAddress *string `structs:"faddress,omitempty" db:"faddress" json:"faddress,omitempty"`
	ToAddress   string  `structs:"taddress,omitempty" db:"taddress" json:"taddress,omitempty"`
	Token       string  `structs:"token,omitempty" db:"token" json:"token,omitempty"`
	Amount      uint64  `structs:"amount,omitempty" db:"amount" json:"amount,omitempty"`
	Completed   *bool   `structs:"completed,omitempty" db:"completed" json:"completed,omitempty"`
}

func NewTransaction() *Transaction {
	return &Transaction{
		ID: uuid.New(uuid.Transaction),
	}
}
