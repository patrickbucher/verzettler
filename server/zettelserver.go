package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/patrickbucher/verzettler"
)

const (
	defaultRows     = 4
	defaultCols     = 2
	defaultFontSize = 12
	fontPath        = "font.ttf"
)

func main() {
	if fontFile, err := os.Open(fontPath); err != nil {
		log.Fatalf("unable to open font file '%s': %v", fontPath, err)
	} else {
		fontFile.Close()
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "server/index.html")
	})
	http.HandleFunc("/zettel", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		defer r.Body.Close()

		pairsFile, _, err := r.FormFile("pairs")
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		} else {
			defer pairsFile.Close()
		}

		inputFile, err := os.CreateTemp("./", "input")
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer os.Remove(inputFile.Name())

		if _, err := io.Copy(inputFile, pairsFile); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		rows := getIntFormFieldOr(r, "rows", defaultRows)
		cols := getIntFormFieldOr(r, "cols", defaultCols)
		fontSize := getIntFormFieldOr(r, "size", defaultFontSize)

		outputFile, err := os.CreateTemp("./", "output")
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		outputFile.Close()
		defer os.Remove(outputFile.Name())

		params := verzettler.FlashCardParams{
			Rows:       rows,
			Cols:       cols,
			FontSize:   fontSize,
			FontPath:   fontPath,
			InputPath:  inputFile.Name(),
			OutputPath: outputFile.Name(),
		}
		if err = verzettler.BuildFlashCardsPDF(params); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		pdfFile, err := os.Open(outputFile.Name())
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer pdfFile.Close()
		w.Header().Set("Content-Type", "application/pdf")
		http.ServeFile(w, r, pdfFile.Name())
	})
	log.Fatal(http.ListenAndServe("0.0.0.0:3771", nil))
}

func getIntFormFieldOr(r *http.Request, name string, fallback int) int {
	raw := r.FormValue(name)
	parsed, err := strconv.Atoi(raw)
	if err != nil {
		return fallback
	}
	return parsed
}
