package handler

import (
	"context"
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/paulwerner/bookkeeper/infra"
	"github.com/paulwerner/bookkeeper/ops"
	"github.com/paulwerner/bookkeeper/security"
	"github.com/paulwerner/bookkeeper/store"
)

var app *fiber.App
var db *sql.DB

func setWorkingDir() {
	wd, _ := os.Getwd()
	// go 2 directories up for migrations to be found
	os.Chdir(filepath.Dir(filepath.Dir(wd)))
}

func TestMain(m *testing.M) {
	setWorkingDir()
	ctx := context.Background()
	container, database, err := infra.CreatePostgresTestContainer(ctx, "testdb")
	if err != nil {
		panic(err)
	}

	err = ops.RunMigrations(database)
	if err != nil {
		panic(err)
	}

	aH := security.NewAuthHandler()
	uRW := store.NewUserStore(database)
	aRW := store.NewAccountStore(database)
	txRW := store.NewTransactionStore(database)

	h := NewHandler(aH, uRW, aRW, txRW)
	server := infra.NewFiberServer()

	h.Register(server)

	app = server
	db = database

	errCode := m.Run()

	err = database.Close()
	if err != nil {
		log.Printf("failed to close db connection: %s", err)
	}

	err = container.Terminate(ctx)
	if err != nil {
		log.Printf("failed to terminate the test container: %s", err)
	}

	os.Exit(errCode)
}