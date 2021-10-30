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

// DistributeWords renders the given page's words on a PDF.
func DistributeWords(pdf *gopdf.GoPdf, page paging.Page) {
	yOffset := 0.0
	xOffset := 0.0
	for i, item := range page.Items {
		if i%page.Cols == 0 {
			pdf.SetX(xMargin)
			xOffset = 0.0
			if i > 0 {
				yOffset += a4height / float64(page.Rows)
				pdf.SetY(yOffset + yMargin)
			} else {
				pdf.SetY(yMargin)
			}
		} else {
			xOffset += a4width / float64(page.Cols)
			pdf.SetX(xOffset + xMargin)
		}
		pdf.Cell(nil, item)
	}
}

// DrawGrid renders a grid on the PDF according to the page's rows and cols.
func DrawGrid(pdf *gopdf.GoPdf, page paging.Page) {
	pdf.SetLineWidth(1)
	pdf.SetLineType("dashed")
	colWidth := a4width / float64(page.Cols)
	xpos := 0.0
	for col := 0; col < (page.Cols - 1); col++ {
		xpos += colWidth
		pdf.Line(xpos, 0, xpos, a4height)
	}
	rowHeight := a4height / float64(page.Rows)
	ypos := 0.0
	for row := 0; row < (page.Rows - 1); row++ {
		ypos += rowHeight
		pdf.Line(0, ypos, a4width, ypos)
	}
}
