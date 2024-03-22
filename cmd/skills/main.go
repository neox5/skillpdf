package main

import (
	"github.com/neox5/skillpdf/internal/cliapp"
)

func main() {
	cliapp.Run()
	// var configPath string
	// outputFilename := "skills.pdf"
	// if len(os.Args) > 1 {
	// 	configPath = os.Args[1]
	// 	if len(os.Args) > 2 {
	// 		outputFilename = os.Args[2]
	// 	}
	// }

	// cfg := loadConfig(configPath)

	// p := pdf.New(true)
	// p.AddPage()

	// originY := p.GetY()

	// p.SetLeftMargin(pdf.LeftMargin)
	// for i, c := range cfg.Columns {
	// 	p.SetY(originY)
	// 	p.SetLeftMargin(pdf.LeftMargin + float64(i)*(pdf.ColumnWidth+pdf.ColumnGap))
	// 	for _, g := range c.Groups {
	// 		pdf.WriteSkillGroup(p, g)
	// 	}
	// }

	// err := p.OutputFileAndClose(outputFilename)
	// if err != nil {
	// 	log.Fatalf("PDF generation error: %v", err)
	// }

	// fmt.Println("PDF successfully generated")
}
