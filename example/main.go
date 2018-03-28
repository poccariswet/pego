package main

import (
	"log"
	"os"

	"github.com/soeyusuke/pego"
)

var (
	home_path = os.Getenv("HOME")
)

func main() {
	peg, err := pego.New()
	if err != nil {
		log.Fatal(err)
	}

	files := []string{
		"a1.mp3",
		"a2.mp3",
	}

	peg.Concat(files)
	peg.Dir(home_path)
	peg.ArgsSet("-c", "copy")

	filename := "sample.aac"
	peg.run(filename)
}
