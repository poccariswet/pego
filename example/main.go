package main

import (
	"log"

	"github.com/soeyusuke/pego"
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
	peg.Dir("YOUR_DIR")
	peg.ArgsSet("-c", "copy")

	filename := "sample.aac"
	peg.Run(filename)
}
