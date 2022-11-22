package main

import (
	"log"

	"github.com/krls08/go-web-app-sessions/cmd/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
