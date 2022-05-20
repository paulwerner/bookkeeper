package main

import (
	"fmt"
	"log"
	"os"

	"github.com/paulwerner/bookkeeper/infra"
	"github.com/paulwerner/bookkeeper/ops"
	"github.com/paulwerner/bookkeeper/router/handler"
	"github.com/paulwerner/bookkeeper/security"
	"github.com/paulwerner/bookkeeper/store"
)

func main() {
	db := infra.SetupPostgreSQLDatabase(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"))
	defer db.Close()
	ops.RunMigrations(db)

	aH := security.NewAuthHandler()
	us := store.NewUserStore(db)
	as := store.NewAccountStore(db)
	txs := store.NewTransactionStore(db)
	h := handler.NewHandler(aH, us, as, txs)

	server := infra.NewFiberServer()
	h.Register(server)

	addr := fmt.Sprintf(
		"%s:%s",
		os.Getenv("SERVER_HOST"),
		os.Getenv("SERVER_PORT"))
	log.Fatal(server.Listen(addr))
}
