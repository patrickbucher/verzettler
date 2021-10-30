// verzettler converts a JSON file of key-value pairs into PDF flash cards.
package verzettler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/signintech/gopdf"
)

// BuildFlashCardsPDF takes a map of word pairs, a path to store the resulting
// PDF, the path to a font file (TTF), and renders a PDF of flash cards.
func BuildFlashCardsPDF(params FlashCardParams) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	err := pdf.AddTTFFont("font", params.FontPath)
	if err != nil {
		panic(err)
	}
	err = pdf.SetFont("font", "", params.FontSize)
	if err != nil {
		panic(err)
	}

	inputFile, err := os.Open(params.InputPath)
	if err != nil {
		Fail("open file '%s': %v", params.InputPath, err)
	}
	defer inputFile.Close()
	buf := bytes.NewBufferString("")
	if _, err = io.Copy(buf, inputFile); err != nil {
		Fail("read from file '%s': %v", params.InputPath, err)
	}
	words := make(map[string]string)
	if err = json.Unmarshal(buf.Bytes(), &words); err != nil {
		panic(fmt.Sprintf("unmarshal JSON from file '%s': %v", params.InputPath, err))
	}

	sheets := Distribute(words, params.Rows, params.Cols)
	for _, sheet := range sheets {
		pdf.AddPage()
		DrawGrid(&pdf, sheet.Front)
		DistributeWords(&pdf, sheet.Front)

		pdf.AddPage()
		DrawGrid(&pdf, sheet.Back)
		DistributeWords(&pdf, sheet.Back)
	}
	err = pdf.WritePdf(params.OutputPath)
	if err != nil {
		Fail("write PDF to %s: %v", params.OutputPath, err)
	}
}
