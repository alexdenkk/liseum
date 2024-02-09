package main

import (
	"alexdenkk/liseum/model"
	"alexdenkk/liseum/pkg/db"
	"alexdenkk/liseum/server"
	"flag"
	"log"
)

var (
	addr   string
	jwt    string
	dbName string
	dbPort string
	dbUser string
	dbPswd string
)

func main() {
	parseFlags()

	// connecting to db
	serverDB, err := db.Connect(dbName)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("db connected")

	// migration
	model.Migrate(serverDB)
	log.Println("migration complete")

	// create app
	booksApp := server.New(serverDB, []byte(jwt), addr)
	log.Println("server initialized")

	// run
	log.Fatal(booksApp.Run())
}

func parseFlags() {
	flag.StringVar(&addr, "address", ":8080", "address and(or) port for app")
	flag.StringVar(&jwt, "jwt", "fuck", "jwt sign key for user tokens")
	flag.StringVar(&dbName, "dbname", "liseum.db", "database name")
	flag.Parse()
}
