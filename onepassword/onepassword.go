package onepassword

import (
	"crypto/sha512"
	"encoding/base64"
	baseErrors "errors"
	"fmt"
	"strings"

	"github.com/1Password/connect-sdk-go/connect"
	"github.com/1Password/connect-sdk-go/onepassword"
	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/mnemonic"
	"gitlab.keelakay.com/keelakay/utils/errors"
	"gitlab.keelakay.com/keelakay/utils/misc"
)

var (
	OnePasswordVaultID        string = "7or5yifsmv4yuqwg6wlrcrng4e"
	OnePasswordVerifyWalletID string = "ruif77mdgirsthjaq55b7tmhfa"
)

type SecretsManager struct {
	secret string
}

func (s *SecretsManager) Load(client connect.Client) error {
	const op errors.Op = "Load"

	item, err := client.GetItemByTitle("Verify - Token Secret", OnePasswordVaultID)
	if err != nil {
		var opErr *onepassword.Error
		if baseErrors.As(err, &opErr) {
			fmt.Printf("message=%s, status code=%d\n",
				opErr.Message,
				opErr.StatusCode,
			)
		}

		return errors.E(op, errors.Network, err, "failed to get item from 1Password")
	}

	s.secret = item.Fields[0].Value

	return nil
}

func (s *SecretsManager) Sign(payload string) string {
	hasher := sha512.New()
	hasher.Write([]byte(fmt.Sprintf("%s+%s", s.secret, payload)))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func (s *SecretsManager) Verify(payload, signature string) bool {
	return s.Sign(payload) == signature
}

func GetAccountFromOnePassword(client connect.Client) (*crypto.Account, error) {
	const op errors.Op = "GetAccountFromOnePassword"

	item, err := client.GetItem(OnePasswordVerifyWalletID, OnePasswordVaultID)
	if err != nil {
		var opErr *onepassword.Error
		if baseErrors.As(err, &opErr) {
			fmt.Printf("message=%s, status code=%d\n",
				opErr.Message,
				opErr.StatusCode,
			)
		}

		return nil, errors.E(op, errors.Network, err, "failed to get item from 1Password")
	}

	fieldNames := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25"}
	sl := []string{}
	for _, field := range item.Fields {
		if misc.InSlice(field.Label, fieldNames) {
			sl = append(sl, field.Value)
		}
	}

	key, err := mnemonic.ToPrivateKey(strings.Join(sl, " "))
	if err != nil {
		return nil, errors.E(op, err, "failed to convert mnemonic to private key")
	}

	account, err := crypto.AccountFromPrivateKey(key)
	if err != nil {
		return nil, errors.E(op, err, "failed to convert private key to account")
	}

	return &account, nil
}
