package store

import (
	"context"
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/paulwerner/bookkeeper/infra"
	"github.com/paulwerner/bookkeeper/ops"
)

var db *sql.DB

func setWorkingDir() {
	wd, _ := os.Getwd()
	// go 1 directory up for migrations to be found
	os.Chdir(filepath.Dir(wd))
}

func TestMain(m *testing.M) {
	setWorkingDir()
	ctx := context.Background()
	container, database, err := infra.CreatePostgresTestContainer(ctx, "testdb")
	if err != nil {
		log.Fatal(err)
	}

	err = ops.RunMigrations(database)
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
