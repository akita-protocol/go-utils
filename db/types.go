package db

import (
	"github.com/jmoiron/sqlx"
)

type DBStruct interface {
	*sqlx.DB | *sqlx.Tx
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

type DBObject interface {
	// global types go here
}

// type AkcEntry struct {
// 	ID   string  `structs:"id,omitempty" db:"id" json:"id,omitempty"`
// 	Addr string  `structs:"addr,omitempty" db:"addr" json:"addr,omitempty"`
// 	Txn  *string `structs:"txn,omitempty" db:"txn" json:"txn,omitempty"`
// }

// func GetAkcEntry[H DBStruct](h H, id string) (*AkcEntry, error) {
// 	const op errors.Op = "GetAkcEntry"
// 	const query = "select id, addr, txn from akc_entry where id = ?"
// 	var r AkcEntry

// 	err := h.Get(&r, query, id)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "AkcEntry Not Found")
// 		}
// 		return nil, errors.E(pkg, op, err)
// 	}
// 	return &r, nil
// }

// type EntryAsset struct {
// 	ID      string `structs:"id,omitempty" db:"id" json:"id,omitempty"`
// 	EntryID string `structs:"entry_id,omitempty" db:"entry_id" json:"entry_id,omitempty"`
// 	AssetID uint64 `structs:"asset_id,omitempty" db:"asset_id" json:"asset_id,omitempty"`
// }

// func GetEntryAssets[H DBStruct](h H, entryID string) (*[]EntryAsset, error) {
// 	const op errors.Op = "GetEntryAssets"
// 	const query = "select id, entry_id, asset_id from entry_asset where entry_id = ?"
// 	var list []EntryAsset

// 	err := h.Select(&list, query, entryID)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "EntryAsset Not Found")
// 		}
// 		return nil, errors.E(pkg, op, err)
// 	}
// 	return &list, nil
// }

// type EntryWithAssets struct {
// 	Entry  *AkcEntry
// 	Assets *[]EntryAsset
// }

// func GetAkcEntryWithAssets[H DBStruct](h H, id string) (*EntryWithAssets, error) {
// 	const op errors.Op = "GetAkcEntryWithAssets"
// 	rChan := make(chan interface{}, 2)
// 	defer close(rChan)
// 	var r EntryWithAssets

// 	go func() {
// 		entry, err := GetAkcEntry(h, id)
// 		if err != nil {
// 			rChan <- err
// 			return
// 		}
// 		rChan <- entry
// 	}()

// 	go func() {
// 		assets, err := GetEntryAssets(h, id)
// 		if err != nil {
// 			rChan <- err
// 			return
// 		}
// 		rChan <- assets
// 	}()

// 	for i := 0; i < 2; i++ {
// 		data := <-rChan
// 		switch result := data.(type) {
// 		case *AkcEntry:
// 			r.Entry = result
// 		case *[]EntryAsset:
// 			r.Assets = result
// 		case error:
// 			return nil, errors.E(op, result)
// 		}
// 	}

// 	return &r, nil
// }
