package pdf

import (
	"strings"

	"github.com/jung-kurt/gofpdf"
	"github.com/neox5/skillpdf/internal/config"
)

const (
	LeftMargin     = 10
	ColumnWidth    = 56.5
	ColumnGap      = 10
	HeaderFontSize = 7
	SkillFontSize  = 8
	lineHeight     = 3.5
	nameWidth      = 34.5
	levelWidth     = 22
	r              = 0.8 // radius
	levelDistance  = 2.3

	fontHeader = "Arial"
	fontSkill  = "Arial"
)

func WriteSectionHeader(p *gofpdf.Fpdf, headerName string) {
	p.SetFont(fontHeader, "", 12)
	p.SetCellMargin(0)
	p.CellFormat(ColumnWidth, 16, headerName, "", 2, "", false, 0, "")
}

// WriteSkillGroup writes a complete SkillGroup to a pdf page
func WriteSkillGroup(p *gofpdf.Fpdf, g config.SkillGroup) {
	writeGroupHeader(p, g)
	p.Ln(1.5)
	for _, s := range g.Skills {
		writeSkill(p, s)
		p.Ln(1.3 * lineHeight)
	}
}

func writeGroupHeader(p *gofpdf.Fpdf, g config.SkillGroup) {
	p.SetFont(fontHeader, "", HeaderFontSize)
	p.SetCellMargin(0)
	p.CellFormat(ColumnWidth, lineHeight, strings.ToUpper(g.Name), "", 2, "", false, 0, "")
	x, y := p.GetXY()
	setDrawBlue(p)
	p.Line(x, y, x+ColumnWidth, y)
}

func writeSkill(p *gofpdf.Fpdf, s config.Skill) {
	p.SetFont(fontSkill, "", SkillFontSize)
	p.SetCellMargin(3)
	p.CellFormat(nameWidth, lineHeight, s.Name, "", 0, "", false, 0, "")
	writeLevel(p, s.Level)
}

func writeLevel(p *gofpdf.Fpdf, lvl int) {
	if lvl == -1 {
		return
	}

	setFillBlue(p)
	setDrawGray(p)
	x, y := p.GetXY()
	style := ""
	for i := 0; i < 10; i++ {
		if i < lvl {
			style = "F" // filled
		} else {
			style = "D" // outlined only
		}
		p.Circle(x, y+r+1, r, style)

		x = x + levelDistance
	}
}

func setDrawBlue(p *gofpdf.Fpdf) {
	p.SetDrawColor(11, 79, 143)
}

func setDrawGray(p *gofpdf.Fpdf) {
	p.SetDrawColor(218, 218, 218)
}

func setFillBlue(p *gofpdf.Fpdf) {
	p.SetFillColor(11, 79, 144)
}
