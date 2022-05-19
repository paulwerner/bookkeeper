package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	d "github.com/paulwerner/bookkeeper/domain"
	"github.com/paulwerner/bookkeeper/utils"
	"github.com/stretchr/testify/assert"
)

func TestAccountCreateWithDescription(t *testing.T) {
	// given
	defer utils.ClearDB(db)
	asserts := assert.New(t)

	uID := utils.RandomUserID()
	u := d.NewUser(uID, "homer", "password")
	utils.PopulateUser(u, db)

	acr := accountCreateRequest{}
	acr.Account.Name = "Main Account"
	description := "some description"
	acr.Account.Description = &description
	acr.Account.Type = d.CHECKING
	acr.Account.CurrentBalance.Value = int64(2342)
	acr.Account.CurrentBalance.Currency = "EUR"
	postBody, err := json.Marshal(acr)
	asserts.NoError(err)

	req, err := http.NewRequest("POST", "http://localhost:8080/api/accounts", bytes.NewBuffer(postBody))
	asserts.NoError(err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(createAuthHeader(uID))

	// when
	resp, err := app.Test(req)

	// then
	asserts.NoError(err)
	var respBody accountCreateResponse
	body, err := ioutil.ReadAll(resp.Body)
	asserts.NoError(err)
	asserts.Equal(http.StatusCreated, resp.StatusCode)

	json.Unmarshal(body, &respBody)
	asserts.NotNil(respBody.Account.ID)
	asserts.Equal(acr.Account.Name, respBody.Account.Name)
	asserts.Equal(acr.Account.Description, respBody.Account.Description)
	asserts.Equal(acr.Account.Description, respBody.Account.Description)
	asserts.Equal(acr.Account.Type, respBody.Account.Type)
	asserts.Equal("€23.42", respBody.Account.BalanceFormatted)
}

func TestAccountCreateWithoutDescription(t *testing.T) {
	// given
	defer utils.ClearDB(db)
	asserts := assert.New(t)

	uID := utils.RandomUserID()
	u := d.NewUser(uID, "homer", "password")
	utils.PopulateUser(u, db)

	acr := accountCreateRequest{}
	acr.Account.Name = "Main Account"
	acr.Account.Type = d.CHECKING
	acr.Account.CurrentBalance.Value = int64(2342)
	acr.Account.CurrentBalance.Currency = "EUR"
	postBody, err := json.Marshal(acr)
	asserts.NoError(err)

	req, err := http.NewRequest("POST", "http://localhost:8080/api/accounts", bytes.NewBuffer(postBody))
	asserts.NoError(err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(createAuthHeader(uID))

	// when
	resp, err := app.Test(req)

	// then
	asserts.NoError(err)
	var respBody accountCreateResponse
	body, err := ioutil.ReadAll(resp.Body)
	asserts.NoError(err)
	asserts.Equal(http.StatusCreated, resp.StatusCode)

	json.Unmarshal(body, &respBody)
	asserts.NotNil(respBody.Account.ID)
	asserts.Equal(acr.Account.Name, respBody.Account.Name)
	asserts.Equal(acr.Account.Description, respBody.Account.Description)
	asserts.Equal(acr.Account.Type, respBody.Account.Type)
	asserts.Equal("€23.42", respBody.Account.BalanceFormatted)
}

func TestAccountCreateNameAlreadyExists(t *testing.T) {
	// given
	defer utils.ClearDB(db)
	asserts := assert.New(t)

	uID := utils.RandomUserID()
	u := d.NewUser(uID, "homer", "password")
	utils.PopulateUser(u, db)

	id := utils.RandomAccountID()
	a := d.NewAccount(id, *u, "Main Account", nil, d.CHECKING, 2342, "EUR")
	utils.PopulateAccount(a, db)

	acr := accountCreateRequest{}
	acr.Account.Name = "Main Account"
	acr.Account.Type = d.CHECKING
	acr.Account.CurrentBalance.Value = int64(2342)
	acr.Account.CurrentBalance.Currency = "EUR"
	postBody, err := json.Marshal(acr)
	asserts.NoError(err)

	req, err := http.NewRequest("POST", "http://localhost:8080/api/accounts", bytes.NewBuffer(postBody))
	asserts.NoError(err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(createAuthHeader(uID))

	// when
	resp, err := app.Test(req)

	// then
	asserts.NoError(err)
	var respBody errorResponse
	body, err := ioutil.ReadAll(resp.Body)
	asserts.NoError(err)
	asserts.Equal(http.StatusConflict, resp.StatusCode)

	json.Unmarshal(body, &respBody)
	asserts.Equal("already in use", respBody.Errors["msg"])
}

func TestAccountsGetReturnsEmptyAccounts(t *testing.T) {
	// given
	defer utils.ClearDB(db)
	asserts := assert.New(t)

	uID := utils.RandomUserID()
	u := d.NewUser(uID, "homer", "password")
	utils.PopulateUser(u, db)

	req, err := http.NewRequest("GET", "http://localhost:8080/api/accounts", nil)
	asserts.NoError(err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(createAuthHeader(uID))

	// when
	resp, err := app.Test(req)

	// then
	asserts.NoError(err)
	var respBody accountsGetResponse
	body, err := ioutil.ReadAll(resp.Body)
	asserts.NoError(err)
	asserts.Equal(http.StatusOK, resp.StatusCode)

	json.Unmarshal(body, &respBody)
	asserts.Empty(respBody.Accounts)
}

func TestAccountsGetReturnsAccounts(t *testing.T) {
	// given
	defer utils.ClearDB(db)
	asserts := assert.New(t)

	uID := utils.RandomUserID()
	u := d.NewUser(uID, "homer", "password")
	utils.PopulateUser(u, db)

	aID1 := utils.RandomAccountID()
	a1 := d.NewAccount(aID1, *u, "Main Account", nil, d.CHECKING, 2342, "EUR")
	utils.PopulateAccount(a1, db)

	aID2 := utils.RandomAccountID()
	description2 := "my second account"
	a2 := d.NewAccount(aID2, *u, "Second Account", &description2, d.SAVINGS, 4223, "EUR")
	utils.PopulateAccount(a2, db)

	aID3 := utils.RandomAccountID()
	description3 := "my credit card"
	a3 := d.NewAccount(aID3, *u, "CC Account", &description3, d.CREDIT_CARD, -10050, "EUR")
	utils.PopulateAccount(a3, db)

	req, err := http.NewRequest("GET", "http://localhost:8080/api/accounts", nil)
	asserts.NoError(err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(createAuthHeader(uID))

	// when
	resp, err := app.Test(req)

	// then
	asserts.NoError(err)
	var respBody accountsGetResponse
	body, err := ioutil.ReadAll(resp.Body)
	asserts.NoError(err)
	asserts.Equal(http.StatusOK, resp.StatusCode)

	json.Unmarshal(body, &respBody)
	asserts.Equal(3, len(respBody.Accounts))

	asserts.Equal(aID1, respBody.Accounts[0].ID)
	asserts.Equal("Main Account", respBody.Accounts[0].Name)
	asserts.Nil(respBody.Accounts[0].Description)
	asserts.Equal(d.CHECKING, respBody.Accounts[0].Type)
	asserts.Equal("€23.42", respBody.Accounts[0].BalanceFormatted)

	asserts.Equal(aID2, respBody.Accounts[1].ID)
	asserts.Equal("Second Account", respBody.Accounts[1].Name)
	asserts.Equal(&description2, respBody.Accounts[1].Description)
	asserts.Equal(d.SAVINGS, respBody.Accounts[1].Type)
	asserts.Equal("€42.23", respBody.Accounts[1].BalanceFormatted)

	asserts.Equal(aID3, respBody.Accounts[2].ID)
	asserts.Equal("CC Account", respBody.Accounts[2].Name)
	asserts.Equal(&description3, respBody.Accounts[2].Description)
	asserts.Equal(d.CREDIT_CARD, respBody.Accounts[2].Type)
	asserts.Equal("-€100.50", respBody.Accounts[2].BalanceFormatted)
}
