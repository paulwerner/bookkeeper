package store

import (
	"testing"

	d "github.com/paulwerner/bookkeeper/domain"
	"github.com/paulwerner/bookkeeper/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	// given
	cut := NewUserStore(db)
	id := d.UserID("some-user-id")
	name := "homer"
	password := "password"

	// when
	result, err := cut.Create(id, name, password)

	// then
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, id, result.ID)
	assert.Equal(t, name, result.Name)
	assert.Equal(t, password, result.Password)

	// finally
	utils.ClearDB(db)
}

func TestFindByID(t *testing.T) {
	// given
	cut := NewUserStore(db)
	id := d.UserID("some-user-id")
	name := "homer"
	password := "password"
	u := d.NewUser(id, name, password)
	utils.PopulateUser(u, db)

	// when
	result, err := cut.FindByID(id)

	// then
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, id, result.ID)
	assert.Equal(t, name, result.Name)
	assert.Equal(t, password, result.Password)

	// finally
	utils.ClearDB(db)
}

func TestFindByIDNotFound(t *testing.T) {
	// given
	cut := NewUserStore(db)
	id := d.UserID("some-user-id")

	// when
	result, err := cut.FindByID(id)

	// then
	assert.Equal(t, err, d.ErrNotFound)
	assert.Nil(t, result)
}

func TestFindByName(t *testing.T) {
	// given
	cut := NewUserStore(db)
	id := d.UserID("some-user-id")
	name := "homer"
	password := "password"
	u := d.NewUser(id, name, password)
	utils.PopulateUser(u, db)

	// when
	result, err := cut.FindByName(name)

	// then
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, id, result.ID)
	assert.Equal(t, name, result.Name)
	assert.Equal(t, password, result.Password)

	// finally
	utils.ClearDB(db)
}
func TestFindByNameNotFound(t *testing.T) {
	// given
	cut := NewUserStore(db)
	name := "homer"

	// when
	result, err := cut.FindByName(name)

	// then
	assert.Equal(t, err, d.ErrNotFound)
	assert.Nil(t, result)
}
