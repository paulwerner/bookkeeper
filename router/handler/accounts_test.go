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

	// finally
	utils.ClearDB(db)
}

func TestAccountCreateWithoutDescription(t *testing.T) {
	// given
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

	// finally
	utils.ClearDB(db)
}

func TestAccountCreateNameAlreadyExists(t *testing.T) {
	// given
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

	// finally
	utils.ClearDB(db)
}
