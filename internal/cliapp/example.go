package cliapp

import (
	"fmt"
	"os"

	"github.com/neox5/skillpdf/internal/config"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

const (
	defaultConfigPath = "./config.yaml"
)

var exampleCmd = &cli.Command{
	Name:      "example",
	Usage:     "creates example config; overwrite existing config",
	Args:      false,
	ArgsUsage: " ",
	Action:    exampleFunc,
}

var exampleFunc = func(cCtx *cli.Context) error {
	var exConfig = config.CreateDefaultConfig()

	yamlData, err := yaml.Marshal(exConfig)
	if err != nil {
		return err
	}

	err = os.WriteFile(defaultConfigPath, yamlData, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("successfully created example config: %v\n", defaultConfigPath)
	return nil
}
