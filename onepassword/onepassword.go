package onepassword

import (
	"crypto/sha512"
	"encoding/base64"
	baseErrors "errors"
	"fmt"
	"strings"

	"github.com/1Password/connect-sdk-go/connect"
	"github.com/1Password/connect-sdk-go/onepassword"
	"github.com/algorand/go-algorand-sdk/v2/crypto"
	"github.com/algorand/go-algorand-sdk/v2/mnemonic"
	"gitlab.algotoolchain.com/keelakay/utils/errors"
	"gitlab.algotoolchain.com/keelakay/utils/misc"
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

func NewAccountInOnePassword(client connect.Client) (*crypto.Account, error) {
	const op errors.Op = "GetAccountFromOnePassword"

	account := crypto.GenerateAccount()

	mnemonics, err := mnemonic.FromPrivateKey(account.PrivateKey)
	if err != nil {
		return nil, errors.E(op, err, "failed to convert private key to mnemonic")
	}

	mnemonic := strings.Split(mnemonics, " ")

	item := &onepassword.Item{
		Category: onepassword.Login,
		Title:    fmt.Sprintf("Algorand Wallet - %s", account.Address.String()),
		Fields: []*onepassword.ItemField{
			{
				Label: "1",
				Value: string(mnemonic[0]),
			},
			{
				Label: "2",
				Value: string(mnemonic[1]),
			},
			{
				Label: "3",
				Value: string(mnemonic[2]),
			},
			{
				Label: "4",
				Value: string(mnemonic[3]),
			},
			{
				Label: "5",
				Value: string(mnemonic[4]),
			},
			{
				Label: "6",
				Value: string(mnemonic[5]),
			},
			{
				Label: "7",
				Value: string(mnemonic[6]),
			},
			{
				Label: "8",
				Value: string(mnemonic[7]),
			},
			{
				Label: "9",
				Value: string(mnemonic[8]),
			},
			{
				Label: "10",
				Value: string(mnemonic[9]),
			},
			{
				Label: "11",
				Value: string(mnemonic[10]),
			},
			{
				Label: "12",
				Value: string(mnemonic[11]),
			},
			{
				Label: "13",
				Value: string(mnemonic[12]),
			},
			{
				Label: "14",
				Value: string(mnemonic[13]),
			},
			{
				Label: "15",
				Value: string(mnemonic[14]),
			},
			{
				Label: "16",
				Value: string(mnemonic[15]),
			},
			{
				Label: "17",
				Value: string(mnemonic[16]),
			},
			{
				Label: "18",
				Value: string(mnemonic[17]),
			},
			{
				Label: "19",
				Value: string(mnemonic[18]),
			},
			{
				Label: "20",
				Value: string(mnemonic[19]),
			},
			{
				Label: "21",
				Value: string(mnemonic[20]),
			},
			{
				Label: "22",
				Value: string(mnemonic[21]),
			},
			{
				Label: "23",
				Value: string(mnemonic[22]),
			},
			{
				Label: "24",
				Value: string(mnemonic[23]),
			},
			{
				Label: "25",
				Value: string(mnemonic[24]),
			},
		},
	}

	_, err = client.CreateItem(item, OnePasswordVaultID)
	if err != nil {
		return nil, errors.E(op, err, "failed to convert private key to mnemonic")
	}

	return &account, nil
}
