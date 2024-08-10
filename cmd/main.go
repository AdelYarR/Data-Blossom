package main

import (
	"log"

	"github.com/AdelYarR/Data-Blossom/internal/api"
	"github.com/AdelYarR/Data-Blossom/pkg/repository"
	"github.com/AdelYarR/Data-Blossom/pkg/config"
)


func main() {
	apiconfig, err := api.NewConfig("C:/Users/user/Desktop/datablossom/configs/config.toml")
	if err != nil {
		log.Fatal(err)
	}

	dbconfig, err := config.NewConfig("C:/Users/user/Desktop/datablossom/configs/config.toml")
	if err != nil {
		log.Fatal(err)
	}

	db, err := repository.New(dbconfig)
	if err != nil {
		log.Fatal(err)
	}

	apiserver := api.New(apiconfig, db)

	err = apiserver.Start()
	if err != nil {
		log.Fatal(err)
	}
}
