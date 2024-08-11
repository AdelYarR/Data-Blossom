package main

import (
	"log"

	"github.com/AdelYarR/Data-Blossom/internal/api"
	"github.com/AdelYarR/Data-Blossom/pkg/repository"
	"github.com/AdelYarR/Data-Blossom/pkg/config"
)


func main() {
	apiconfig, err := api.NewConfig("./configs/config.toml")
	if err != nil {
		log.Fatal(err)
	}

	dbconfig, err := config.NewConfig("./configs/config.toml")
	if err != nil {
		log.Fatal(err)
	}

	db, err := repository.New(dbconfig)
	if err != nil {
		log.Fatal(err)
	}

	apiserver := api.New(apiconfig, db)

	if err = apiserver.Start(); err != nil {
		log.Fatal(err)
	}
}
