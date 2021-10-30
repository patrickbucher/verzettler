package rendering

import (
	"github.com/patrickbucher/verzettler/paging"
	"github.com/signintech/gopdf"
)

const (
	yMargin = 25.0
	xMargin = 20.0
)

var (
	a4width  = gopdf.PageSizeA4.W
	a4height = gopdf.PageSizeA4.H
)

func DistributeWords(pdf *gopdf.GoPdf, page paging.Page, rows, cols int) {
	yOffset := 0.0
	xOffset := 0.0
	for i, item := range page {
		if i%cols == 0 {
			pdf.SetX(xMargin)
			xOffset = 0.0
			if i > 0 {
				yOffset += a4height / float64(rows)
				pdf.SetY(yOffset + yMargin)
			} else {
				pdf.SetY(yMargin)
			}
		} else {
			xOffset += a4width / float64(cols)
			pdf.SetX(xOffset + xMargin)
		}
		pdf.Cell(nil, item)
	}
}

func DrawGrid(pdf *gopdf.GoPdf, rows, cols int) {
	pdf.SetLineWidth(1)
	pdf.SetLineType("dashed")
	colWidth := a4width / float64(cols)
	xpos := 0.0
	for col := 0; col < (cols - 1); col++ {
		xpos += colWidth
		pdf.Line(xpos, 0, xpos, a4height)
	}
	rowHeight := a4height / float64(rows)
	ypos := 0.0
	for row := 0; row < (rows - 1); row++ {
		ypos += rowHeight
		pdf.Line(0, ypos, a4width, ypos)
	}
}
