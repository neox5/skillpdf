package main

import (
	"fmt"
	"log"
	"os"

	"github.com/neox5/skillpdf/internal/cliapp"
	"github.com/neox5/skillpdf/pkg/pdf"
	"github.com/spf13/viper"
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

func loadConfig(configPath string) *pdf.Config {
	if configPath != "" {
		fmt.Println(configPath)
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			fmt.Printf("Warning: Config file %s does not exist. Using default configuration.\n", configPath)
			viper.SetConfigName("config")
			viper.AddConfigPath(".")
		} else {
			viper.SetConfigFile(configPath)
		}
	} else {
		fmt.Println("Warning: No config file specified. Using default configuration.")
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
	}

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Read config file error: %s", err)
	}

	config := &pdf.Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		log.Fatalf("Unmarshal config error: %s", err)
	}

	return config
}
