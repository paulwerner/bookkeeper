package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	d "github.com/paulwerner/bookkeeper/domain"
	"github.com/paulwerner/bookkeeper/utils"
	"github.com/stretchr/testify/assert"
)

func TestAppConfigGet(t *testing.T) {
	// given
	asserts := assert.New(t)

	id := utils.RandomUserID()
	u := d.NewUser(id, "homer", "password")
	utils.PopulateUser(u, db)
	req, err := http.NewRequest("GET", "http://localhost:8080/api/config", nil)
	asserts.NoError(err)
	req.Header.Add(createAuthHeader(id))

	// when
	resp, err := app.Test(req)

	// then
	asserts.NoError(err)
	body, err := ioutil.ReadAll(resp.Body)
	asserts.NoError(err)
	asserts.Equal(http.StatusOK, resp.StatusCode)

	var respBody appConfigResponse
	json.Unmarshal(body, &respBody)
	asserts.Equal(3, len(respBody.SupportedAccountTypes))
	asserts.Equal(d.CHECKING, respBody.SupportedAccountTypes[0])
	asserts.Equal(d.SAVINGS, respBody.SupportedAccountTypes[1])
	asserts.Equal(d.CREDIT_CARD, respBody.SupportedAccountTypes[2])

	// finally
	utils.ClearDB(db)
}

func TestAppConfigGetUnauthorizedFails(t *testing.T) {
	// given
	asserts := assert.New(t)

	id := utils.RandomUserID()
	u := d.NewUser(id, "homer", "password")
	utils.PopulateUser(u, db)
	req, err := http.NewRequest("GET", "http://localhost:8080/api/config", nil)
	asserts.NoError(err)

	// when
	resp, err := app.Test(req)

	// then
	asserts.NoError(err)
	body, err := ioutil.ReadAll(resp.Body)
	asserts.NoError(err)
	asserts.Equal(http.StatusForbidden, resp.StatusCode)

	var respBody errorResponse
	json.Unmarshal(body, &respBody)
	asserts.Equal("missing or malformed JWT", respBody.Errors["msg"])

	// finally
	utils.ClearDB(db)
}
