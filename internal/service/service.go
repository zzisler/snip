package service

import (
	"crypto/rand"
	"log"

	"github.com/yihleego/base62"
)

func GenCode() string {

	bytes := make([]byte, 6)

	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	token := base62.StdEncoding.EncodeToString(bytes)

	return token

}
