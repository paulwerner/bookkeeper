package store

import (
	"testing"

	d "github.com/paulwerner/bookkeeper/pkg/domain"
	"github.com/paulwerner/bookkeeper/utils"
	"github.com/stretchr/testify/assert"
)

func TestAccountCreateAccountWithoutDescription(t *testing.T) {
	// given
	defer utils.ClearDB(db)
	asserts := assert.New(t)

	uID := utils.RandomUserID()
	u := d.NewUser(uID, "homer", "password")
	utils.PopulateUser(u, db)

	id := utils.RandomAccountID()
	a := d.NewAccount(id, *u, "Main Account", nil, d.CHECKING, int64(23), "EUR")

	cut := NewAccountStore(db)

	// when
	result, err := cut.Create(*a)

	// then
	asserts.NoError(err)
	asserts.NotNil(result)
	asserts.Equal(id, result.ID)
	asserts.Equal(a.Name, result.Name)
	asserts.Nil(result.Description)
	asserts.Equal(a.Type, result.Type)
	asserts.Equal(a.BalanceValue, result.BalanceValue)
	asserts.Equal(a.BalanceCurrency, result.BalanceCurrency)
}

func TestAccountCreateAccountWithDescription(t *testing.T) {
	// given
	defer utils.ClearDB(db)
	asserts := assert.New(t)

	id := utils.RandomAccountID()
	uID := utils.RandomUserID()
	u := d.NewUser(uID, "homer", "password")
	utils.PopulateUser(u, db)
	name := "Main Account"
	description := "some description"
	accountType := d.CHECKING
	balanceValue := int64(23)
	balanceCurrency := "EUR"
	a := d.NewAccount(id, *u, name, &description, accountType, balanceValue, balanceCurrency)

	cut := NewAccountStore(db)

	// when
	result, err := cut.Create(*a)

	// then
	asserts.NoError(err)
	asserts.NotNil(result)
	asserts.Equal(id, result.ID)
	asserts.Equal(name, result.Name)
	asserts.Equal(&description, result.Description)
	asserts.Equal(accountType, result.Type)
	asserts.Equal(balanceValue, result.BalanceValue)
	asserts.Equal(balanceCurrency, result.BalanceCurrency)
}

func TestAccountFindByIDAndUser(t *testing.T) {
	// given
	defer utils.ClearDB(db)
	asserts := assert.New(t)

	id := utils.RandomAccountID()
	uID := utils.RandomUserID()
	u := d.NewUser(uID, "homer", "password")
	utils.PopulateUser(u, db)
	name := "Main Account"
	description := "some description"
	accountType := d.CHECKING
	balanceValue := int64(23)
	balanceCurrency := "EUR"

	account := d.NewAccount(id, *u, name, &description, accountType, balanceValue, balanceCurrency)
	utils.PopulateAccount(account, db)

	cut := NewAccountStore(db)

	// when
	result, err := cut.FindByIDAndUser(id, uID)

	// then
	asserts.NoError(err)
	asserts.NotNil(result)
	asserts.Equal(id, result.ID)
	asserts.Equal(name, result.Name)
	asserts.Equal(&description, result.Description)
	asserts.Equal(accountType, result.Type)
	asserts.Equal(balanceValue, result.BalanceValue)
	asserts.Equal(balanceCurrency, result.BalanceCurrency)
}

func TestAccountFindByIDAndUserNotFound(t *testing.T) {
	// given
	defer utils.ClearDB(db)
	asserts := assert.New(t)

	id := utils.RandomAccountID()
	uID := utils.RandomUserID()

	cut := NewAccountStore(db)

	// when
	result, err := cut.FindByIDAndUser(id, uID)

	// then
	asserts.Equal(d.ErrNotFound, err)
	asserts.Nil(result)
}

func TestAccountFindByUserAndName(t *testing.T) {
	// given
	defer utils.ClearDB(db)
	asserts := assert.New(t)

	id := utils.RandomAccountID()
	uID := utils.RandomUserID()
	u := d.NewUser(uID, "homer", "password")
	utils.PopulateUser(u, db)
	name := "Main Account"
	description := "some description"
	accountType := d.CHECKING
	balanceValue := int64(23)
	balanceCurrency := "EUR"

	account := d.NewAccount(id, *u, name, &description, accountType, balanceValue, balanceCurrency)
	utils.PopulateAccount(account, db)

	cut := NewAccountStore(db)

	// when
	result, err := cut.FindByUserAndName(uID, name)

	// then
	asserts.NoError(err)
	asserts.NotNil(result)
	asserts.Equal(id, result.ID)
	asserts.Equal(name, result.Name)
	asserts.Equal(&description, result.Description)
	asserts.Equal(accountType, result.Type)
	asserts.Equal(balanceValue, result.BalanceValue)
	asserts.Equal(balanceCurrency, result.BalanceCurrency)
}

func TestAccountFindByUserAndNameNotFound(t *testing.T) {
	// given
	asserts := assert.New(t)

	uID := utils.RandomUserID()
	name := "Main Account"

	cut := NewAccountStore(db)

	// when
	result, err := cut.FindByUserAndName(uID, name)

	// then
	asserts.Equal(d.ErrNotFound, err)
	asserts.Nil(result)
}
