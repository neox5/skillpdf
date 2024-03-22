package pdf

import (
	"github.com/jung-kurt/gofpdf"
)

// New creates a new gofpdf.Fpdf
func New(isLandscape bool) *gofpdf.Fpdf {
	orientation := "P"
	if isLandscape {
		orientation = "L"
	}
	pdf := gofpdf.New(orientation, "mm", "A4", "internal/fonts")
	pdf.AddFont("Montserrat", "", "Montserrat-Regular.json")
	pdf.AddFont("Ovo", "", "Ovo-Regular.json")
	return pdf
}
