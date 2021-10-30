package main

import "github.com/patrickbucher/verzettler"

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
	verzettler.BuildFlashCardsPDF(words, "example.pdf", "font.ttf")
}
