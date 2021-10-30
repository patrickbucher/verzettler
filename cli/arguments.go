package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type FlashCardParams struct {
	Rows       int
	Cols       int
	FontSize   int
	FontPath   string
	InputPath  string
	OutputPath string
}

func ParseCommandLineArguments(cmdName string) FlashCardParams {
	const (
		rowsUsage       = "number or rows"
		colsUsage       = "number of columns"
		fontSizeUsage   = "font size in pt"
		fontPathUsage   = "path to font file (TTF)"
		outputPathUsage = "path for output file (PDF)"
	)
	var (
		rows       int
		cols       int
		fontSize   int
		fontPath   string
		outputPath string
	)
	flags := flag.NewFlagSet(cmdName, flag.ExitOnError)
	flags.IntVar(&rows, "rows", 4, rowsUsage)
	flags.IntVar(&rows, "r", 4, rowsUsage)
	flags.IntVar(&cols, "cols", 2, colsUsage)
	flags.IntVar(&cols, "c", 2, colsUsage)
	flags.IntVar(&fontSize, "size", 12, fontSizeUsage)
	flags.IntVar(&fontSize, "s", 12, fontSizeUsage)
	flags.StringVar(&fontPath, "font", "font.ttf", fontPathUsage)
	flags.StringVar(&fontPath, "f", "font.ttf", fontPathUsage)
	flags.StringVar(&outputPath, "out", "flashcards.pdf", fontPathUsage)
	flags.StringVar(&outputPath, "o", "flashcards.pdf", fontPathUsage)

	if len(os.Args) < 2 {
		Fail("usage: %s [input]", os.Args[0])
	}
	err := flags.Parse(os.Args[1:])
	if err != nil {
		Fail("parsing flags: %v", err)
	}
	if rows < 1 {
		Fail("rows/r must be at least 1, was %d", rows)
	}
	if cols < 1 {
		Fail("cols/c must be at least 1, was %d", cols)
	}
	if fontSize < 1 {
		Fail("size/s must be at least 1, was %d", fontSize)
	}

	args := flags.Args()
	if len(args) > 1 {
		Fail("excess arguments: %v", args[1:])
	}
	inputPath := args[0]

	return FlashCardParams{
		Rows:       rows,
		Cols:       cols,
		FontSize:   fontSize,
		FontPath:   fontPath,
		InputPath:  inputPath,
		OutputPath: outputPath,
	}
}

func Fail(msg string, args ...interface{}) {
	msg = strings.TrimSpace(msg) + "\n"
	fmt.Fprintf(os.Stderr, msg, args...)
	os.Exit(1)
}
