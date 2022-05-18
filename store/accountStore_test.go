package store

import (
	"testing"

	d "github.com/paulwerner/bookkeeper/domain"
	"github.com/paulwerner/bookkeeper/utils"
	"github.com/stretchr/testify/assert"
)

func TestAccountCreateAccountWithoutDescription(t *testing.T) {
	// given
	asserts := assert.New(t)

	id := utils.RandomAccountID()
	uID := utils.RandomUserID()
	user := d.NewUser(uID, "homer", "password")
	utils.PopulateUser(user, db)
	name := "Main Account"
	accountType := d.CHECKING
	balanceValue := int64(23)
	balanceCurrency := "EUR"

	cut := NewAccountStore(db)

	// when
	result, err := cut.Create(id, uID, name, nil, accountType, balanceValue, balanceCurrency)

	// then
	asserts.NoError(err)
	asserts.NotNil(result)
	asserts.Equal(id, result.ID)
	asserts.Equal(name, result.Name)
	asserts.Nil(result.Description)
	asserts.Equal(accountType, result.Type)
	asserts.Equal(balanceValue, result.BalanceValue)
	asserts.Equal(balanceCurrency, result.BalanceCurrency)

	// finally
	utils.ClearDB(db)
}

func TestAccountCreateAccountWithDescription(t *testing.T) {
	// given
	tx, _ := db.Begin()
	defer tx.Rollback()
	asserts := assert.New(t)

	id := utils.RandomAccountID()
	uID := utils.RandomUserID()
	user := d.NewUser(uID, "homer", "password")
	utils.PopulateUser(user, db)
	name := "Main Account"
	description := "some description"
	accountType := d.CHECKING
	balanceValue := int64(23)
	balanceCurrency := "EUR"

	cut := NewAccountStore(db)

	// when
	result, err := cut.Create(id, uID, name, &description, accountType, balanceValue, balanceCurrency)

	// then
	asserts.NoError(err)
	asserts.NotNil(result)
	asserts.Equal(id, result.ID)
	asserts.Equal(name, result.Name)
	asserts.Equal(&description, result.Description)
	asserts.Equal(accountType, result.Type)
	asserts.Equal(balanceValue, result.BalanceValue)
	asserts.Equal(balanceCurrency, result.BalanceCurrency)

	// finally
	utils.ClearDB(db)
}

func TestAccountFindByIDAndUser(t *testing.T) {
	// given
	tx, _ := db.Begin()
	defer tx.Rollback()
	asserts := assert.New(t)

	id := utils.RandomAccountID()
	uID := utils.RandomUserID()
	user := d.NewUser(uID, "homer", "password")
	utils.PopulateUser(user, db)
	name := "Main Account"
	description := "some description"
	accountType := d.CHECKING
	balanceValue := int64(23)
	balanceCurrency := "EUR"

	account := d.NewAccount(id, *user, name, &description, accountType, balanceValue, balanceCurrency)
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

	// finally
	utils.ClearDB(db)
}

func TestAccountFindByIDAndUserNotFound(t *testing.T) {
	// given
	tx, _ := db.Begin()
	defer tx.Rollback()
	asserts := assert.New(t)

	id := utils.RandomAccountID()
	uID := utils.RandomUserID()

	cut := NewAccountStore(db)

	// when
	result, err := cut.FindByIDAndUser(id, uID)

	// then
	asserts.Equal(d.ErrNotFound, err)
	asserts.Nil(result)

	// finally
	utils.ClearDB(db)
}

func TestAccountFindByUserAndName(t *testing.T) {
	// given
	tx, _ := db.Begin()
	defer tx.Rollback()
	asserts := assert.New(t)

	id := utils.RandomAccountID()
	uID := utils.RandomUserID()
	user := d.NewUser(uID, "homer", "password")
	utils.PopulateUser(user, db)
	name := "Main Account"
	description := "some description"
	accountType := d.CHECKING
	balanceValue := int64(23)
	balanceCurrency := "EUR"

	account := d.NewAccount(id, *user, name, &description, accountType, balanceValue, balanceCurrency)
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

	// finally
	utils.ClearDB(db)
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

	// finally
	utils.ClearDB(db)
}
