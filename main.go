package main

import (
	"log"
	"github.com/Bassiouni/money-mate-api/db"
	"github.com/Bassiouni/money-mate-api/router"
)

func init() {
	db.InitDB()

	router.BindRootRouter()
	router.BindUserRouter()
}

func main() {
	defer db.CloseConnection()

	if err := router.RootRouter.Run(":8080"); err != nil {
		log.Fatalf("Couldn't start the server: %v", err)
	}
}
