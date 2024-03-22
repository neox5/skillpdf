package pdf

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/neox5/skillpdf/internal/config"
)

func GenerateSkillsPdf(cfg *config.Config) *gofpdf.Fpdf {
	p := New(cfg.Settings.IsLandScape)
	p.AddPage()

	originY := p.GetY()

	p.SetLeftMargin(LeftMargin)
	for i, c := range cfg.Columns {
		p.SetY(originY)
		p.SetLeftMargin(LeftMargin + float64(i)*(ColumnWidth+ColumnGap))
		for _, g := range c.Groups {
			WriteSkillGroup(p, g)
		}
	}

	return p
}
