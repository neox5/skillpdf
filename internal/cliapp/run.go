package cliapp

import (
	"fmt"
	"log"
	"os"

	"github.com/neox5/skillpdf/internal/config"
	"github.com/neox5/skillpdf/internal/pdf"
	"github.com/urfave/cli/v2"
)

func Run() {
	app := &cli.App{
		Name:    "skills",
		Usage:   "cli tool for generating a pdf listing your skills",
		Version: "1.0.0",
		Flags: []cli.Flag{
			configFlag,
			outputFlag,
		},
		Commands: []*cli.Command{
			exampleCmd,
		},
		Action: generatePdfFunc,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

var configFlag = &cli.StringFlag{
	Name:    "config",
	Value:   "config.yaml",
	Aliases: []string{"c"},
	Usage:   "Load configuration from `FILE`",
	Action: func(cCtx *cli.Context, v string) error {
		fileInfo, err := os.Stat(v)
		if err != nil {
			if os.IsNotExist(err) {
				return fmt.Errorf("config file does not exist: %v", err)
			}
			return err
		}
		if fileInfo.IsDir() {
			return fmt.Errorf("config filepath is a directory")
		}
		return nil
	},
}

var outputFlag = &cli.StringFlag{Name: "output", Value: "skills.pdf", Aliases: []string{"o"}, Usage: "Output `FILE`"}

var generatePdfFunc = func(cCtx *cli.Context) error {
	configPath := cCtx.String("config")
	outputPath := cCtx.String("output")

	cfg := config.Load(configPath)

	p := pdf.GenerateSkillsPdf(cfg)

	err := p.OutputFileAndClose(outputPath)
	if err != nil {
		log.Fatalf("PDF generation error: %v", err)
	}

	fmt.Printf("PDF successfully generated: %v\n", outputPath)

	return nil
}
