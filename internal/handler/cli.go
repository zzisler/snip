package handler

import (
	"fmt"
	"os"
	"snip/internal/service"
	"snip/internal/storage"
	"snip/internal/web"
)

var code string

func Cli() {

	if len(os.Args) < 2 {
		fmt.Println("usage: snip <command>. Use 'snip help' for more info.")
		return
	}

	switch os.Args[1] {
	case "help":

		fmt.Println("Usage: <command> [options] [arguments]\n" +
			"  Available commands:\n" +
			"   - add <url>\n" +
			"      Add a new URL to shorten and receive a short code.\n" +
			"   - get <code>\n" +
			"      Retrieve the original URL by using the provided short code.")

	case "add":

		if len(os.Args) < 3 {
			fmt.Println("usage: snip add <url>")
			return
		}

		code = service.GenCode()

		err := storage.Save(code, os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(code)

	case "get":

		if len(os.Args) < 3 {
			fmt.Println("usage: snip get <code>")
			return
		}

		result, err := storage.Get(os.Args[2])
		if err != nil {
			fmt.Println(result)
			return
		}

		fmt.Println(result.URL)

	case "serve":

		web.Start()

	default:
		fmt.Println("unknown command")
		return
	}

}
