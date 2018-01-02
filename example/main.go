package main

import (
	"log"

	"github.com/soeyusuke/pego"
)

func main() {
	cmd, err := pego.New()
	if err != nil {
		log.Fatal(err)
	}
	_ = cmd
}
