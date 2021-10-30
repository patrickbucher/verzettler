package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/patrickbucher/verzettler/paging"
	"github.com/patrickbucher/verzettler/rendering"
	"github.com/signintech/gopdf"
)

const (
	rows       = 4
	cols       = 3
	outputPath = "example.pdf"
	fontPath   = "font.ttf"
	inputPath  = "pairs.json"
)

func main() {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		panic(fmt.Sprintf("open file '%s': %v", inputPath, err))
	}
	defer inputFile.Close()

	buf := bytes.NewBufferString("")
	if _, err = io.Copy(buf, inputFile); err != nil {
		panic(fmt.Sprintf("read from file '%s': %v", inputPath, err))
	}

	words := make(map[string]string)
	if err = json.Unmarshal(buf.Bytes(), &words); err != nil {
		panic(fmt.Sprintf("unmarshal JSON from file '%s': %v", inputPath, err))
	}
	BuildFlashCardsPDF(words, outputPath, fontPath)
}

// BuildFlashCardsPDF takes a map of word pairs, a path to store the resulting
// PDF, the path to a font file (TTF), and renders a PDF of flash cards.
func BuildFlashCardsPDF(pairs map[string]string, pdfOutputPath, fontPath string) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	err := pdf.AddTTFFont("font", fontPath)
	if err != nil {
		panic(err)
	}
	err = pdf.SetFont("font", "", 14)
	if err != nil {
		panic(err)
	}

	sheets := paging.Distribute(pairs, rows, cols)
	for _, sheet := range sheets {
		pdf.AddPage()
		rendering.DrawGrid(&pdf, sheet.Front)
		rendering.DistributeWords(&pdf, sheet.Front)

		pdf.AddPage()
		rendering.DrawGrid(&pdf, sheet.Back)
		rendering.DistributeWords(&pdf, sheet.Back)
	}
	pdf.WritePdf("example.pdf")
}
