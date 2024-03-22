package cliapp

import (
	"fmt"
	"os"

	"github.com/neox5/skillpdf/internal/config"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

var exampleCmd = &cli.Command{
	Name:      "example",
	Usage:     "creates example config; overwrite existing config",
	Args:      false,
	ArgsUsage: " ",
	Action:    exampleFunc,
}

var exampleFunc = func(cCtx *cli.Context) error {
	var exConfig = &config.Config{
		Columns: []config.SkillColumn{
			{
				Groups: []config.SkillGroup{
					{
						Name: "Programming Languages",
						Skills: []config.Skill{
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
						Skills: []config.Skill{
							{Name: "GitLab CI", Level: 10},
							{Name: "Azure Pipelines", Level: 7},
							{Name: "Keptn", Level: 8},
						},
					},
				},
			},
			{
				Groups: []config.SkillGroup{
					{
						Name: "Issue Tracking",
						Skills: []config.Skill{
							{Name: "Atlassian Jira", Level: 10},
							{Name: "GitLab Issues", Level: 7},
							{Name: "GitHub Issues", Level: 7},
						},
					},
				},
			},
		},
	}

	yamlData, err := yaml.Marshal(exConfig)
	if err != nil {
		return err
	}

	err = os.WriteFile("config.yaml", yamlData, 0644)
	if err != nil {
		return err
	}

	fmt.Println("successfully created example config.yaml")

	return nil
}
