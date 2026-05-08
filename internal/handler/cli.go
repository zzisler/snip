package handler

import (
	"log"
	"os"
	"snip/internal/service"
	"snip/internal/storage"
)

var code string

func Cli() {

	switch os.Args[1] {
	case "add":
		code = service.GenCode()
		result := storage.Save(code, os.Args[2])
		if result != nil {
			log.Fatal(result)
		}
		log.Print(code)
	case "get":
		result, err := storage.Get(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		log.Print(result.URL)
	default:
		log.Fatal("unknown command")
	}

}
