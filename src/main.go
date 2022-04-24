package main

import (
	"github.com/madrid-recicla/server/src/dbconnector"
	"github.com/madrid-recicla/server/src/router"
)

func main() {
	go dbconnector.InitMongoClient()
	router.InitRouter()
}
