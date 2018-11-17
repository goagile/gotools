package main

import (
	"gofpdf"
)

func main() {
	var (
		orientation string = "P"  // "P" - Portrait, "L" - Landscape
		units       string = "mm" // "mm", "cm", "pt", "in"
		size        string = "A4" // "A3", "A4", "A5", "Letter", "Legal", "Tabloid"
		fontDir     string = "."  // An empty string is replaced with "."
	)

	pdf := gofpdf.New(orientation, units, size, fontDir)
}
