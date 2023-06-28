package misc

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/algorand/go-algorand-sdk/encoding/msgpack"
	"github.com/algorand/go-algorand-sdk/types"
	"github.com/algorand/go-algorand-sdk/v2/client/v2/common/models"
	"github.com/algorand/go-algorand-sdk/v2/client/v2/indexer"
	cerrors "gitlab.algotoolchain.com/keelakay/utils/errors"
)

func AddrToED25519PublicKey(a types.Address) (pk ed25519.PublicKey) {
	pk = make([]byte, len(a))
	copy(pk, a[:])
	return
}

func RawTransactionBytesToSign(tx types.Transaction) []byte {
	// Encode the transaction as msgpack
	encodedTx := msgpack.Encode(tx)

	// Prepend the hashable prefix
	msgParts := [][]byte{[]byte("TX"), encodedTx}
	return bytes.Join(msgParts, nil)
}

func VerifySignature(tx types.Transaction, pk ed25519.PublicKey, sig types.Signature) bool {
	toBeSigned := RawTransactionBytesToSign(tx)
	return ed25519.Verify(pk, toBeSigned, sig[:])
}

// CheckSignature verifies that stx is either a single signature
func CheckSignature(stx types.SignedTxn) error {
	if stx.Sig == (types.Signature{}) {
		return errors.New("msig/lsig not supported")
	}

	// ensure other signature fields are empty
	if len(stx.Msig.Subsigs) != 0 || stx.Msig.Version != 0 || stx.Msig.Threshold != 0 {
		return errors.New("tx has both a sig and msig")
	}

	if !stx.Lsig.Blank() {
		return errors.New("tx has both a sig and lsig")
	}

	fmt.Println("Auth Address: ", stx.AuthAddr.String())

	var pk ed25519.PublicKey
	if stx.AuthAddr.String() != types.ZeroAddress.String() {
		pk = AddrToED25519PublicKey(stx.AuthAddr)
	} else {
		pk = AddrToED25519PublicKey(stx.Txn.Sender)
	}

	if !VerifySignature(stx.Txn, pk, stx.Sig) {
		return errors.New("signature is invalid")
	}

	return nil
}

func GetLatestAssetConfigProperties(client *indexer.Client, assetID uint64) (map[string]string, error) {
	const op cerrors.Op = "GetLatestAssetConfig"
	var data map[string]interface{}
	properties := map[string]string{}

	nextToken := ""
	loop := true
	var lastResponse models.TransactionsResponse

	for loop {
		loop = false
		nextPage := nextToken
		resp, err := client.LookupAssetTransactions(assetID).NextToken(nextPage).TxType("acfg").Do(context.Background())
		if err != nil {
			return properties, cerrors.E(op, err)
		}

		if resp.NextToken != "" {
			lastResponse = resp
			nextToken = resp.NextToken
			loop = true
			continue
		}

		if len(resp.Transactions) == 0 {
			resp = lastResponse
		}

		txn := resp.Transactions[len(resp.Transactions)-1]

		err = json.Unmarshal(txn.Note, &data)
		if err != nil {
			return properties, cerrors.E(op, err)
		}

		props, ok := data["properties"]
		if !ok {
			return properties, cerrors.E(op, "properties not found")
		}

		interumProps, ok := props.(map[string]interface{})
		if !ok {
			return properties, cerrors.E(op, "properties is not a map of string to interface")
		}

		for key, value := range interumProps {
			properties[key] = value.(string)
		}

		break
	}

	return properties, nil
}

func Arc69Check(client *indexer.Client, assetID uint64) (bool, map[string]string, error) {
	const op cerrors.Op = "Arc69Check"
	var data map[string]interface{}
	isArc69 := false
	properties := map[string]string{}

	nextToken := ""
	loop := true
	var lastResponse models.TransactionsResponse

	for loop {
		loop = false
		nextPage := nextToken
		resp, err := client.LookupAssetTransactions(assetID).NextToken(nextPage).TxType("acfg").Do(context.Background())
		if err != nil {
			return isArc69, properties, cerrors.E(op, err)
		}

		if resp.NextToken != "" {
			lastResponse = resp
			nextToken = resp.NextToken
			loop = true
			continue
		}

		if len(resp.Transactions) == 0 {
			resp = lastResponse
		}

		txn := resp.Transactions[len(resp.Transactions)-1]

		err = json.Unmarshal(txn.Note, &data)
		if err != nil {
			return isArc69, properties, cerrors.E(op, err)
		}

		props, ok := data["properties"]
		if !ok {
			return isArc69, properties, cerrors.E(op, "properties not found")
		}

		interumProps, ok := props.(map[string]interface{})
		if !ok {
			return isArc69, properties, cerrors.E(op, "properties is not a map of string to interface")
		}

		for key, value := range interumProps {
			properties[key] = value.(string)
		}

		interfaceStandard, ok := data["standard"]
		if !ok {
			return isArc69, properties, cerrors.E(op, "standard not found")
		}

		standard, ok := interfaceStandard.(string)
		if !ok {
			return isArc69, properties, cerrors.E(op, "interface standard is not a string")
		}

		if strings.ToLower(standard) == "arc69" {
			isArc69 = true
		}

		break
	}

	return isArc69, properties, nil
}
