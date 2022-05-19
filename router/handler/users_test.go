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

func TestSignUpSuccessful(t *testing.T) {
	// given
	asserts := assert.New(t)

	usur := userSignUpRequest{}
	usur.User.Name = "homer"
	usur.User.Password = "password"
	postBody, _ := json.Marshal(usur)
	req, err := http.NewRequest("POST", "http://localhost:8008/api/users", bytes.NewBuffer(postBody))
	asserts.NoError(err)
	req.Header.Add("Content-Type", "application/json")

	// when
	resp, err := app.Test(req)

	// then
	asserts.NoError(err)
	body, err := ioutil.ReadAll(resp.Body)
	asserts.NoError(err)
	asserts.Equal(http.StatusCreated, resp.StatusCode)

	var respBody userSignUpResponse
	json.Unmarshal(body, &respBody)
	asserts.Equal("homer", respBody.User.Name)
	asserts.NotEmpty(respBody.Token)

	// finally
	err = resp.Body.Close()
	asserts.NoError(err)
	utils.ClearDB(db)
}

func TestSignUpInvalidPasswordLength(t *testing.T) {
	// given
	asserts := assert.New(t)

	usur := userSignUpRequest{}
	usur.User.Name = "homer"
	usur.User.Password = "pass"
	postBody, _ := json.Marshal(usur)
	req, err := http.NewRequest("POST", "http://localhost:8008/api/users", bytes.NewBuffer(postBody))
	asserts.NoError(err)
	req.Header.Add("Content-Type", "application/json")

	// when
	resp, err := app.Test(req)

	// then
	asserts.NoError(err)
	body, err := ioutil.ReadAll(resp.Body)
	asserts.NoError(err)
	asserts.Equal(http.StatusUnprocessableEntity, resp.StatusCode)

	var respBody errorResponse
	json.Unmarshal(body, &respBody)
	asserts.Equal("invalid password length", respBody.Errors["msg"])

	// finally
	err = resp.Body.Close()
	asserts.NoError(err)
	utils.ClearDB(db)
}

func TestSignUpEmptyUser(t *testing.T) {
	// given
	asserts := assert.New(t)
	
	usur := userSignUpRequest{}
	usur.User.Name = ""
	usur.User.Password = "password"
	postBody, _ := json.Marshal(usur)
	req, err := http.NewRequest("POST", "http://localhost:8008/api/users", bytes.NewBuffer(postBody))
	asserts.NoError(err)
	req.Header.Add("Content-Type", "application/json")

	// when
	resp, err := app.Test(req)

	// then
	asserts.NoError(err)
	body, err := ioutil.ReadAll(resp.Body)
	asserts.NoError(err)
	asserts.Equal(http.StatusUnprocessableEntity, resp.StatusCode)
	
	var respBody errorResponse
	json.Unmarshal(body, &respBody)
	asserts.Equal("invalid entity", respBody.Errors["msg"])
	
	// finally
	err = resp.Body.Close()
	asserts.NoError(err)
	utils.ClearDB(db)
}

func TestSignUpUsernameAlreadyInUse(t *testing.T) {
	// given
	asserts := assert.New(t)
	
	id := utils.RandomUserID()
	u := d.NewUser(id, "homer", "password")
	utils.PopulateUser(u, db)
	
	usur := userSignUpRequest{}
	usur.User.Name = "homer"
	usur.User.Password = "password"
	postBody, _ := json.Marshal(usur)
	req, err := http.NewRequest("POST", "http://localhost:8008/api/users", bytes.NewBuffer(postBody))
	asserts.NoError(err)
	req.Header.Add("Content-Type", "application/json")
	
	// when
	resp, err := app.Test(req)
	
	// then
	asserts.NoError(err)
	body, err := ioutil.ReadAll(resp.Body)
	asserts.NoError(err)
	asserts.Equal(http.StatusConflict, resp.StatusCode)
	
	var respBody errorResponse
	json.Unmarshal(body, &respBody) 
	asserts.Equal("already in use", respBody.Errors["msg"])
	
	// finally
	err = resp.Body.Close()
	asserts.NoError(err)
	utils.ClearDB(db)
}

func TestLoginSuccessful(t *testing.T) {
	// given
	asserts := assert.New(t)
	
	id := utils.RandomUserID()
	u := d.NewUser(id, "homer", "$2a$10$fEChGoAym287oEERp7XlQeFcnN7RIDmn70drD4liYvREHocfeySti")
	utils.PopulateUser(u, db)
	usur := userLoginRequest{}
	usur.Name = "homer"
	usur.Password = "password"
	postBody, _ := json.Marshal(usur)
	req, err := http.NewRequest("POST", "http://localhost:8008/api/users/login", bytes.NewBuffer(postBody))
	asserts.NoError(err)
	req.Header.Add("Content-Type", "application/json")
	
	// when
	resp, err := app.Test(req)
	
	// then
	asserts.NoError(err)
	body, err := ioutil.ReadAll(resp.Body)
	asserts.NoError(err)
	asserts.Equal(http.StatusOK, resp.StatusCode)
	
	var respBody userLoginResponse
	json.Unmarshal(body, &respBody)
	asserts.NotEmpty(respBody.Token)
	asserts.Equal("homer", respBody.User.Name)
	
	// finally
	err = resp.Body.Close()
	asserts.NoError(err)
	utils.ClearDB(db)
}

func TestLoginInvalidPassword(t *testing.T) {
	// given
	asserts := assert.New(t)
	
	id := utils.RandomUserID()
	u := d.NewUser(id, "homer", "password")
	utils.PopulateUser(u, db)
	usur := userLoginRequest{}
	usur.Name = "homer"
	usur.Password = "invalid"
	postBody, _ := json.Marshal(usur)
	req, err := http.NewRequest("POST", "http://localhost:8008/api/users/login", bytes.NewBuffer(postBody))
	asserts.NoError(err)
	req.Header.Add("Content-Type", "application/json")
	
	// when
	resp, err := app.Test(req)
	
	// then
	asserts.NoError(err)
	body, err := ioutil.ReadAll(resp.Body)
	asserts.NoError(err)
	asserts.Equal(http.StatusForbidden, resp.StatusCode)
	
	var respBody errorResponse
	json.Unmarshal(body, &respBody)
	asserts.Equal("invalid password", respBody.Errors["msg"])
	
	// finally
	err = resp.Body.Close()
	asserts.NoError(err)
	utils.ClearDB(db)
}

func TestLoginInvalidUsername(t *testing.T) {
	// given
	asserts := assert.New(t)
	
	id := utils.RandomUserID()
	u:= d.NewUser(id, "homer", "password")
	utils.PopulateUser(u, db)
	usur := userLoginRequest{}
	usur.Name = "marge"
	usur.Password = "password"
	postBody, _ := json.Marshal(usur)
	req, err := http.NewRequest("POST", "http://localhost:8008/api/users/login", bytes.NewBuffer(postBody))
	asserts.NoError(err)
	req.Header.Add("Content-Type", "application/json")
	
	// when
	resp, err := app.Test(req)

	// then
	asserts.NoError(err)
	body, err := ioutil.ReadAll(resp.Body)
	asserts.NoError(err)
	asserts.Equal(http.StatusNotFound, resp.StatusCode)
	
	var respBody errorResponse
	json.Unmarshal(body, &respBody)
	asserts.Equal("not found", respBody.Errors["msg"])
	
	// finally
	err = resp.Body.Close()
	asserts.NoError(err)
	utils.ClearDB(db)
}
