package pdf

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
