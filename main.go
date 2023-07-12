package main

import (
	"log"
	"money-mate/db"
	"money-mate/router"
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
