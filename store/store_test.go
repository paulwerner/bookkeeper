package store_test

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/paulwerner/bookkeeper/infra"
)

var db *sql.DB

func TestMain(m *testing.M) {
	ctx := context.Background()
	container, database, err := infra.CreatePostgresTestContainer(ctx, "testdb")
	if err != nil {
		log.Fatal(err)
	}
	db = database
	errCode := m.Run()
	err = database.Close()
	if err != nil {
		log.Printf("failed to close db connection: %s", err)
	}
	err = container.Terminate(ctx)
	if err != nil {
		log.Printf("failed to terminate the tet container: %s", err)
	}
	os.Exit(errCode)
}
