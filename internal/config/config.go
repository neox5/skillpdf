package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

// Config defines the config structure
type Config struct {
	Columns []SkillColumn `yaml:"columns"`
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
