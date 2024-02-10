package main

import (
	"alexdenkk/liseum/model"
	"alexdenkk/liseum/pkg/db"
	"alexdenkk/liseum/pkg/hash"
	"alexdenkk/liseum/server"
	"flag"
	"log"
)

var (
	addr          string
	jwt           string
	dbName        string
	adminLogin    string
	adminPassword string
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

	// creating admin from server
	if adminLogin != "" && adminPassword != "" {
		result := serverDB.Create(&model.User{
			Login:    adminLogin,
			Password: hash.Hash(adminPassword),
		})

		if result.Error != nil {
			log.Fatal(err)
			return
		}

		log.Println("user created")
		return
	}

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
	flag.StringVar(&adminLogin, "admin-login", "", "admin login")
	flag.StringVar(&adminPassword, "admin-pass", "", "admin password")
	flag.Parse()
}
