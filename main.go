package main

import (
	"log"

	"github.com/EmmanoelDan/importador/config"
	"github.com/EmmanoelDan/importador/router"
	_ "github.com/lib/pq"
)

func main() {
	_, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	router.Init()
}
