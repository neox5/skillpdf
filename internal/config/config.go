package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

// Config defines the config structure
type Config struct {
	Settings PdfSettings   `yaml:"settings"`
	Columns  []SkillColumn `yaml:"columns"`
}

type PdfSettings struct {
	// IsLandscape sets the orientation from portrait (default) to landscape
	IsLandScape    bool    `yaml:"isLandscape"`
	LeftMargin     float64 `yaml:"leftMargin"`
	ColumnWidth    float64 `yaml:"columnWidth"`
	ColumnGap      float64 `yaml:"columnGap"`
	HeaderFontSize float64 `yaml:"headerFontSize"`
	SkillFontSize  float64 `yaml:"skillFontSize"`
	LineHeight     float64 `yaml:"lineHeight"`
	NameWidth      float64 `yaml:"nameWidth"`
	LevelWidth     float64 `yaml:"levelWidth"`
	Radius         float64 `yaml:"radius"`
	LevelDistance  float64 `yaml:"levelDistance"`
	FontHeader     string  `yaml:"fontHeader"`
	FontSkill      string  `yaml:"fontSkill"`
}

type SkillColumn struct {
	Groups []SkillGroup `yaml:"groups"`
}

// SkillGroup bundles one or more skills
type SkillGroup struct {
	Name   string  `yaml:"name"`
	Skills []Skill `yaml:"skills"`
}

// Skill combines a name with a level
type Skill struct {
	Name string `yaml:"name"`
	// Level from 0 to 10 or -1 to remove the graphical representation
	Level int `yaml:"level"`
}

func CreateDefaultConfig() *Config {
	return &Config{
		Settings: PdfSettings{
			IsLandScape:    false,
			LeftMargin:     10,
			ColumnWidth:    56.5,
			ColumnGap:      10,
			HeaderFontSize: 7,
			SkillFontSize:  8,
			LineHeight:     3.5,
			NameWidth:      34.5,
			LevelWidth:     22,
			Radius:         0.8,
			LevelDistance:  2.3,
			FontHeader:     "Arial",
			FontSkill:      "Arial",
		},
		Columns: []SkillColumn{
			{
				Groups: []SkillGroup{
					{
						Name: "Programming Languages",
						Skills: []Skill{
							{Name: "Javascirpt", Level: 9},
							{Name: "TypeScript", Level: 8},
							{Name: "Golang", Level: 8},
							{Name: "C", Level: 7},
							{Name: "C++", Level: 6},
							{Name: "C#", Level: 6},
							{Name: "HTML5", Level: 10},
							{Name: "CSS3", Level: 10},
						},
					},
					{
						Name: "DevOps",
						Skills: []Skill{
							{Name: "GitLab CI", Level: 10},
							{Name: "Azure Pipelines", Level: 7},
							{Name: "Keptn", Level: 8},
						},
					},
				},
			},
			{
				Groups: []SkillGroup{
					{
						Name: "Issue Tracking",
						Skills: []Skill{
							{Name: "Atlassian Jira", Level: 10},
							{Name: "GitLab Issues", Level: 7},
							{Name: "GitHub Issues", Level: 7},
						},
					},
				},
			},
		},
	}
}

func Load(configPath string) *Config {
	if configPath != "" {
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

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		log.Fatalf("Unmarshal config error: %s", err)
	}

	return config
}
