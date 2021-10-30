package paging

type Page struct {
	Items []string
	Rows  int
	Cols  int
}

func NewPage(rows, cols int) Page {
	n := rows * cols
	items := make([]string, n)
	return Page{items, rows, cols}
}

type Sheet struct {
	Front Page
	Back  Page
}

// Distribute distributes...
func Distribute(pairs map[string]string, rows, cols int) []Sheet {
	perPage := rows * cols
	sheets := make([]Sheet, 0)
	frontSeq := buildFrontPageIndexSequence(rows, cols)
	backSeq := buildBackPageIndexSequence(rows, cols)
	i := 0
	front := NewPage(rows, cols)
	back := NewPage(rows, cols)
	for key, value := range pairs {
		front.Items[frontSeq[i]] = key
		back.Items[backSeq[i]] = value
		i++
		if i == perPage {
			sheets = append(sheets, Sheet{front, back})
			front = NewPage(rows, cols)
			back = NewPage(rows, cols)
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
