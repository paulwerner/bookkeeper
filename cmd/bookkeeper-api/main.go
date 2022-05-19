package main

import (
	"fmt"
	"log"

	"github.com/paulwerner/bookkeeper/infra"
	"github.com/paulwerner/bookkeeper/ops"
	"github.com/paulwerner/bookkeeper/router/handler"
	"github.com/paulwerner/bookkeeper/security"
	"github.com/paulwerner/bookkeeper/store"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "pass"
	dbname   = "bookkeeperdb"
	sslmode  = "disable"

	serverHost = "http://localhost"
	serverPort = "8080"
)

func main() {
	db := infra.SetupPostgreSQLDatabase(
		host,
		port,
		user,
		password,
		dbname,
		sslmode,
	)
	defer db.Close()
	if err := ops.RunMigrations(db); err != nil {
		panic(err)
	}

	aH := security.NewAuthHandler()
	us := store.NewUserStore(db)
	as := store.NewAccountStore(db)
	txs := store.NewTransactionStore(db)
	h := handler.NewHandler(aH, us, as, txs)

	server := infra.NewFiberServer()
	h.Register(server)

	addr := fmt.Sprintf("%s:%s", serverHost, serverPort)
	log.Fatal(server.Listen(addr))
}
