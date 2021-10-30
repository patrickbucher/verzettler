package main

import (
	"github.com/patrickbucher/verzettler"
)

func main() {
	params := verzettler.ParseCommandLineArguments("verzettler")
	verzettler.BuildFlashCardsPDF(params)
}
