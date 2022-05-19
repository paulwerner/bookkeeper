package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	d "github.com/paulwerner/bookkeeper/domain"
	"github.com/paulwerner/bookkeeper/utils"
	"github.com/stretchr/testify/assert"
)

func TestTransactionsGetNotFound(t *testing.T) {
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
