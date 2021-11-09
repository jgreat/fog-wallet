package api

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tyler-smith/go-bip39"
	"github.com/tyler-smith/go-bip39/wordlists"
)

// GetAccounts - get all accounts
func (a *Api) GetAccounts(ctx echo.Context) error {
	// var result []Account
	var accounts []Account

	dbResult := a.DB.Find(&accounts)
	if dbResult.Error != nil {
		return sendApiError(ctx, http.StatusInternalServerError, "DB error")
	}

	return ctx.JSONPretty(http.StatusOK, accounts, "  ")
}

// PostAccounts - Create an account
func (a *Api) PostAccounts(ctx echo.Context) error {
	var newAccount NewAccount

	// This should make sure the content matches struct.
	err := ctx.Bind(&newAccount)
	if err != nil {
		return sendApiError(ctx, http.StatusBadRequest, "Invalid format for Account")
	}

	// generate the account id - this should be fun
	var accountID string = generateAccountID()

	var account Account
	account.Name = newAccount.Name
	account.AccountID = &accountID

	dbResult := a.DB.Create(&account)
	if dbResult.Error != nil {
		return sendApiError(ctx, http.StatusInternalServerError, "Failed to create account")
	}

	return ctx.JSONPretty(http.StatusOK, account, "  ")
}

func generateAccountID() string {

	// Create 24 word mnemonic
	bip39.SetWordList(wordlists.English)
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	log.Println("Mnemonic: ", mnemonic)

	// We need to check that there isn't a colliding account?
	// Scan from current block index - TODO figure out how to get block index

	// account_key is ? slip-0010 (is this b58?)

	return mnemonic
}
