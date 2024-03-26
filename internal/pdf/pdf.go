package pdf

import (
	"strings"

	"github.com/jung-kurt/gofpdf"
	"github.com/neox5/skillpdf/internal/config"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("settings.leftMargin", 10)
	viper.SetDefault("settings.columnWidth", 56.5)
	viper.SetDefault("settings.columnGap", 10)
	viper.SetDefault("settings.headerFontSize", 7)
	viper.SetDefault("settings.skillFontSize", 8)
	viper.SetDefault("settings.lineHeight", 3.5)
	viper.SetDefault("settings.nameWidth", 34.5)
	viper.SetDefault("settings.levelWidth", 22)
	viper.SetDefault("settings.radius", 0.8)
	viper.SetDefault("settings.levelDistance", 2.3)
	viper.SetDefault("settings.fontHeader", "Arial")
	viper.SetDefault("settings.fontSkill", "Arial")
}

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

func GenerateSkillsPdf(cfg *config.Config) *gofpdf.Fpdf {
	s := cfg.Settings
	p := New(s.IsLandScape)
	p.AddPage()

	p.SetLeftMargin(s.LeftMargin)

	WriteSectionHeader(&s, p, "Skills")

	originY := p.GetY()
	for i, c := range cfg.Columns {
		p.SetY(originY)
		p.SetLeftMargin(s.LeftMargin + float64(i)*(s.ColumnWidth+s.ColumnGap))
		for _, g := range c.Groups {
			WriteSkillGroup(&s, p, g)
		}
	}

	return p
}

func WriteSectionHeader(cfg *config.PdfSettings, p *gofpdf.Fpdf, headerName string) {
	p.SetFont(cfg.FontHeader, "", 12)
	p.SetCellMargin(0)
	p.CellFormat(cfg.ColumnWidth, 16, headerName, "", 2, "", false, 0, "")
}

// WriteSkillGroup writes a complete SkillGroup to a pdf page
func WriteSkillGroup(cfg *config.PdfSettings, p *gofpdf.Fpdf, g config.SkillGroup) {
	writeGroupHeader(cfg, p, g)
	p.Ln(1.5)
	for _, s := range g.Skills {
		writeSkill(cfg, p, s)
		p.Ln(1.3 * cfg.LineHeight)
	}
}

func writeGroupHeader(cfg *config.PdfSettings, p *gofpdf.Fpdf, g config.SkillGroup) {
	p.SetFont(cfg.FontHeader, "", cfg.HeaderFontSize)
	p.SetCellMargin(0)
	p.CellFormat(cfg.ColumnWidth, cfg.LineHeight, strings.ToUpper(g.Name), "", 2, "", false, 0, "")
	x, y := p.GetXY()
	setDrawBlue(p)
	p.Line(x, y, x+cfg.ColumnWidth, y)
}

func writeSkill(cfg *config.PdfSettings, p *gofpdf.Fpdf, s config.Skill) {
	p.SetFont(cfg.FontSkill, "", cfg.SkillFontSize)
	p.SetCellMargin(3)
	p.CellFormat(cfg.NameWidth, cfg.LineHeight, s.Name, "", 0, "", false, 0, "")
	writeLevel(cfg, p, s.Level)
}

func writeLevel(cfg *config.PdfSettings, p *gofpdf.Fpdf, lvl int) {
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
		p.Circle(x, y+cfg.Radius+1, cfg.Radius, style)

		x = x + cfg.LevelDistance
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
