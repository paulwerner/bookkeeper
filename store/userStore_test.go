package store

import (
	"testing"

	d "github.com/paulwerner/bookkeeper/pkg/domain"
	"github.com/paulwerner/bookkeeper/utils"
	"github.com/stretchr/testify/assert"
)

func TestUserCreate(t *testing.T) {
	// given
	asserts := assert.New(t)

	cut := NewUserStore(db)
	id := utils.RandomUserID()
	name := "homer"
	password := "password"

	// when
	result, err := cut.Create(id, name, password)

	// then
	asserts.NoError(err)
	asserts.NotNil(result)
	asserts.Equal(id, result.ID)
	asserts.Equal(name, result.Name)
	asserts.Equal(password, result.Password)

	// finally
	utils.ClearDB(db)
}

func TestUserFindByID(t *testing.T) {
	// given
	asserts := assert.New(t)

	cut := NewUserStore(db)
	id := utils.RandomUserID()
	name := "homer"
	password := "password"
	u := d.NewUser(id, name, password)
	utils.PopulateUser(u, db)

	// when
	result, err := cut.FindByID(id)

	// then
	asserts.NoError(err)
	asserts.NotNil(result)
	asserts.Equal(id, result.ID)
	asserts.Equal(name, result.Name)
	asserts.Equal(password, result.Password)

	// finally
	utils.ClearDB(db)
}

func TestUserFindByIDNotFound(t *testing.T) {
	// given
	asserts := assert.New(t)

	cut := NewUserStore(db)
	id := utils.RandomUserID()

	// when
	result, err := cut.FindByID(id)

	// then
	asserts.Equal(err, d.ErrNotFound)
	asserts.Nil(result)
}

func TestUserFindByName(t *testing.T) {
	// given
	asserts := assert.New(t)

	id := utils.RandomUserID()
	name := "homer"
	password := "password"
	u := d.NewUser(id, name, password)
	utils.PopulateUser(u, db)

	cut := NewUserStore(db)
	// when
	result, err := cut.FindByName(name)

	// then
	asserts.NoError(err)
	asserts.NotNil(result)
	asserts.Equal(id, result.ID)
	asserts.Equal(name, result.Name)
	asserts.Equal(password, result.Password)

	// finally
	utils.ClearDB(db)
}
func TestUserFindByNameNotFound(t *testing.T) {
	// given
	asserts := assert.New(t)

	cut := NewUserStore(db)
	name := "homer"

	// when
	result, err := cut.FindByName(name)

	// then
	asserts.Equal(err, d.ErrNotFound)
	asserts.Nil(result)
}
