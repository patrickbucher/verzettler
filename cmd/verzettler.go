package main

import (
	"fmt"
	"strings"

	"github.com/signintech/gopdf"
)

const (
	cols    = 3
	rows    = 5
	yMargin = 25.0
	xMargin = 20.0
)

var (
	a4width  = gopdf.PageSizeA4.W
	a4height = gopdf.PageSizeA4.H
)

func main() {
	words := map[string]string{
		"Sehenswürdigkeit": "достопримечательность",
		"Haus":             "дом",
		"Katze":            "кот",
		"Hund":             "собака",
		"Mädchen":          "девочка",
		"Bär":              "медведь",
		"Russland":         "Россия",
		"Junge":            "мальчик",
	}
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	err := pdf.AddTTFFont("font", "font.ttf")
	if err != nil {
		panic(err)
	}
	err = pdf.SetFont("font", "", 14)
	if err != nil {
		panic(err)
	}

	sheets := Distribute(words, rows, cols)
	for _, sheet := range sheets {
		pdf.AddPage()
		drawGrid(&pdf)
		fmt.Println(strings.Join(sheet.Front, ","))
		distributeWords(&pdf, sheet.Front)

		pdf.AddPage()
		drawGrid(&pdf)
		fmt.Println(strings.Join(sheet.Back, ","))
		distributeWords(&pdf, sheet.Back)
	}
	pdf.WritePdf("example.pdf")
}

func distributeWords(pdf *gopdf.GoPdf, page Page) {
	yOffset := 0.0
	xOffset := 0.0
	for i, item := range page {
		if i%cols == 0 {
			pdf.SetX(xMargin)
			xOffset = 0.0
			if i > 0 {
				yOffset += a4height / rows
				pdf.SetY(yOffset + yMargin)
			} else {
				pdf.SetY(yMargin)
			}
		} else {
			xOffset += a4width / cols
			pdf.SetX(xOffset + xMargin)
		}
		pdf.Cell(nil, item)
	}
}

func drawGrid(pdf *gopdf.GoPdf) {
	pdf.SetLineWidth(1)
	pdf.SetLineType("dashed")
	colWidth := a4width / cols
	xpos := 0.0
	for col := 0; col < (cols - 1); col++ {
		xpos += colWidth
		pdf.Line(xpos, 0, xpos, a4height)
	}
	rowHeight := a4height / rows
	ypos := 0.0
	for row := 0; row < (rows - 1); row++ {
		ypos += rowHeight
		pdf.Line(0, ypos, a4width, ypos)
	}
}

type Page []string

type Sheet struct {
	Front Page
	Back  Page
}

func Distribute(pairs map[string]string, rows, cols int) []Sheet {
	perPage := rows * cols
	sheets := make([]Sheet, 0)
	frontSeq := buildFrontPageIndexSequence(rows, cols)
	backSeq := buildBackPageIndexSequence(rows, cols)
	i := 0
	front := make(Page, perPage)
	back := make(Page, perPage)
	for key, value := range pairs {
		front[frontSeq[i]] = key
		back[backSeq[i]] = value
		i++
		if i == perPage {
			sheets = append(sheets, Sheet{front, back})
			front = make(Page, perPage)
			back = make(Page, perPage)
			i = 0
		}
	}
	if i > 0 {
		sheets = append(sheets, Sheet{front, back})
	}
	return sheets
}

// front: enumerate by column, then by row
// |-------|
// | 0 | 1 |
// |---|---|
// | 2 | 3 |
// |---|---|
// | 4 | 5 |
// |---|---|
// | 6 | 7 |
// |---|---|
func buildFrontPageIndexSequence(rows, cols int) []int {
	cells := rows * cols
	indexSequence := make([]int, cells)
	for i := 0; i < cells; i++ {
		indexSequence[i] = i
	}
	return indexSequence
}

// back: flip on long edge
// |-------|
// | 1 | 0 |
// |---|---|
// | 3 | 2 |
// |---|---|
// | 5 | 4 |
// |---|---|
// | 7 | 6 |
// |---|---|
func buildBackPageIndexSequence(rows, cols int) []int {
	cells := rows * cols
	indexSequence := make([]int, cells)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			i := cols - 1 - c + r*cols
			indexSequence[r*cols+c] = i
		}
	}
	return indexSequence
}
