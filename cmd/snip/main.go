package main

import (
	"snip/internal/handler"
	"snip/internal/storage"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	storage.Connect()

	handler.Cli()

}
