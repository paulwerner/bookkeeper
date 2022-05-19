package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	d "github.com/paulwerner/bookkeeper/pkg/domain"
	"github.com/paulwerner/bookkeeper/utils"
	"github.com/stretchr/testify/assert"
)

func TestTransactionGetNotFound(t *testing.T) {
	// given
	defer utils.ClearDB(db)
	asserts := assert.New(t)

	uID := utils.RandomUserID()
	u := d.NewUser(uID, "homer", "password")
	utils.PopulateUser(u, db)

	aID := utils.RandomAccountID()
	a := d.NewAccount(aID, *u, "Main Account", nil, d.CHECKING, 2342, "EUR")
	utils.PopulateAccount(a, db)

	url := fmt.Sprintf("http://localhost:8080/api/accounts/%s/transactions/%s", aID, "invalid")
	req, err := http.NewRequest("GET", url, nil)
	asserts.NoError(err)

	req.Header.Set(createAuthHeader(uID))

	// when
	resp, err := app.Test(req)

	// then
	asserts.NoError(err)
	asserts.NotNil(resp)
	asserts.Equal(http.StatusNotFound, resp.StatusCode)
	var respBody errorResponse
	body, err := ioutil.ReadAll(resp.Body)
	asserts.NoError(err)

	json.Unmarshal(body, &respBody)
	asserts.Equal("not found", respBody.Errors["msg"])
}

func TestTransactionGet(t *testing.T) {
	// given
	defer utils.ClearDB(db)
	asserts := assert.New(t)

	uID := utils.RandomUserID()
	u := d.NewUser(uID, "homer", "password")
	utils.PopulateUser(u, db)

	aID := utils.RandomAccountID()
	a := d.NewAccount(aID, *u, "Main Account", nil, d.CHECKING, 2342, "EUR")
	utils.PopulateAccount(a, db)

	txID := utils.RandomTransactionID()
	tx := d.NewTransaction(txID, *a, nil, 2342, "EUR")
	utils.PopulateTransaction(tx, db)
	url := fmt.Sprintf("http://localhost:8080/api/accounts/%s/transactions/%s", aID, txID)
	req, err := http.NewRequest("GET", url, nil)
	asserts.NoError(err)

	req.Header.Set(createAuthHeader(uID))

	// when
	resp, err := app.Test(req)

	// then
	asserts.NoError(err)
	asserts.NotNil(resp)
	asserts.Equal(http.StatusOK, resp.StatusCode)

	var respBody transactionGetResponse
	body, err := ioutil.ReadAll(resp.Body)
	asserts.NoError(err)

	json.Unmarshal(body, &respBody)
	asserts.Equal(txID, respBody.Transaction.ID)
	asserts.Nil(respBody.Transaction.Description)
	asserts.Equal("€23.42", respBody.Transaction.AmountFormatted)
}

func TestTransactionsGet(t *testing.T) {
	// given
	defer utils.ClearDB(db)
	asserts := assert.New(t)

	uID := utils.RandomUserID()
	u := d.NewUser(uID, "homer", "password")
	utils.PopulateUser(u, db)

	aID := utils.RandomAccountID()
	a := d.NewAccount(aID, *u, "Main Account", nil, d.CHECKING, 2342, "EUR")
	utils.PopulateAccount(a, db)

	txID1 := utils.RandomTransactionID()
	tx1 := d.NewTransaction(txID1, *a, nil, 2342, "EUR")
	utils.PopulateTransaction(tx1, db)

	txID2 := utils.RandomTransactionID()
	description2 := "some description"
	tx2 := d.NewTransaction(txID2, *a, &description2, 4223, "EUR")
	utils.PopulateTransaction(tx2, db)

	url := fmt.Sprintf("http://localhost:8080/api/accounts/%s/transactions", aID)
	req, err := http.NewRequest("GET", url, nil)
	asserts.NoError(err)

	req.Header.Set(createAuthHeader(uID))

	// when
	resp, err := app.Test(req)

	// then
	asserts.NoError(err)
	asserts.NotNil(resp)
	asserts.Equal(http.StatusOK, resp.StatusCode)

	var respBody transactionsGetResponse
	body, err := ioutil.ReadAll(resp.Body)
	asserts.NoError(err)

	json.Unmarshal(body, &respBody)
	asserts.Equal(2, len(respBody.Transactions))

	asserts.NotNil(respBody.Transactions[0].ID)
	asserts.Nil(respBody.Transactions[0].Description)
	asserts.Equal("€23.42", respBody.Transactions[0].AmountFormatted)

	asserts.NotNil(respBody.Transactions[1].ID)
	asserts.Equal(&description2, respBody.Transactions[1].Description)
	asserts.Equal("€42.23", respBody.Transactions[1].AmountFormatted)
}

func TestTransactionsCreateSuccessful(t *testing.T) {
	// given
	defer utils.ClearDB(db)
	asserts := assert.New(t)

	uID := utils.RandomUserID()
	u := d.NewUser(uID, "homer", "password")
	utils.PopulateUser(u, db)

	aID := utils.RandomAccountID()
	a := d.NewAccount(aID, *u, "Main Account", nil, d.CHECKING, 2342, "EUR")
	utils.PopulateAccount(a, db)

	tcr := transactionCreateRequest{}
	tcr.Transaction.Amount = int64(2342)
	tcr.Transaction.Currency = "EUR"
	reqBody, err := json.Marshal(tcr)
	asserts.NoError(err)

	url := fmt.Sprintf("http://localhost:8080/api/accounts/%s/transactions", aID)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	asserts.NoError(err)

	req.Header.Set(createAuthHeader(uID))
	req.Header.Set("Content-Type", "application/json")

	// when
	resp, err := app.Test(req)

	// then
	asserts.NoError(err)
	asserts.NotNil(resp)
	asserts.Equal(http.StatusCreated, resp.StatusCode)

	var respBody transactionGetResponse
	body, err := ioutil.ReadAll(resp.Body)
	asserts.NoError(err)

	json.Unmarshal(body, &respBody)
	asserts.NotNil(respBody.Transaction.ID)
	asserts.Nil(respBody.Transaction.Description)
	asserts.Equal("€23.42", respBody.Transaction.AmountFormatted)
}

func TestTransactionsCreateAccountNotFound(t *testing.T) {
	// given
	defer utils.ClearDB(db)
	asserts := assert.New(t)

	uID := utils.RandomUserID()
	u := d.NewUser(uID, "homer", "password")
	utils.PopulateUser(u, db)

	tcr := transactionCreateRequest{}
	tcr.Transaction.Amount = int64(2342)
	tcr.Transaction.Currency = "EUR"
	reqBody, err := json.Marshal(tcr)
	asserts.NoError(err)

	url := fmt.Sprintf("http://localhost:8080/api/accounts/%s/transactions", "invalid")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	asserts.NoError(err)

	req.Header.Set(createAuthHeader(uID))
	req.Header.Set("Content-Type", "application/json")

	// when
	resp, err := app.Test(req)

	// then
	asserts.NoError(err)
	asserts.NotNil(resp)
	asserts.Equal(http.StatusNotFound, resp.StatusCode)

	var respBody errorResponse
	body, err := ioutil.ReadAll(resp.Body)
	asserts.NoError(err)

	json.Unmarshal(body, &respBody)
	asserts.Equal("not found", respBody.Errors["msg"])
}
