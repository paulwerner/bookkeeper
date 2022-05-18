package store

import (
	"testing"

	d "github.com/paulwerner/bookkeeper/domain"
	"github.com/paulwerner/bookkeeper/utils"
	"github.com/stretchr/testify/assert"
)

func TestTransactionCreateWithoutDescription(t *testing.T) {
	// given
	asserts := assert.New(t)

	uID := utils.RandomUserID()
	u := d.NewUser(uID, "homer", "pasword")
	utils.PopulateUser(u, db)
	aID := utils.RandomAccountID()
	a := d.NewAccount(aID, *u, "Main Account", nil, d.CHECKING, int64(23), "EUR")
	utils.PopulateAccount(a, db)
	
	id := utils.RandomTransactionID()
	amount := int64(23)
	currency := "EUR"

	cut := NewtransactionStore(db)

	// when
	result, err := cut.Create(id, aID, nil, amount, currency)

	// then
	asserts.NoError(err)
	asserts.NotNil(result)
	asserts.Equal(id, result.ID)
	asserts.Nil(result.Description)
	asserts.Equal(amount, result.Amount)
	asserts.Equal(currency, result.Currency)

	// finally
	utils.ClearDB(db)
}

func TestTransactionCreateWithDescription(t *testing.T) {
	// given
	asserts := assert.New(t)

	uID := utils.RandomUserID()
	u := d.NewUser(uID, "homer", "pasword")
	utils.PopulateUser(u, db)
	aID := utils.RandomAccountID()
	a := d.NewAccount(aID, *u, "Main Account", nil, d.CHECKING, int64(23), "EUR")
	utils.PopulateAccount(a, db)

	id := utils.RandomTransactionID()
	amount := int64(23)
	currency := "EUR"

	cut := NewtransactionStore(db)

	// when
	result, err := cut.Create(id, aID, nil, amount, currency)

	// then
	asserts.NoError(err)
	asserts.NotNil(result)
	asserts.Equal(id, result.ID)
	asserts.Nil(result.Description)
	asserts.Equal(amount, result.Amount)
	asserts.Equal(currency, result.Currency)

	// finally
	utils.ClearDB(db)
}