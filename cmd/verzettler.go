package main

import (
	"log"

	"github.com/patrickbucher/verzettler"
)

func main() {
	params := verzettler.ParseCommandLineArguments("verzettler")
	if err := verzettler.BuildFlashCardsPDF(params); err != nil {
		log.Fatal(err)
	}
}
