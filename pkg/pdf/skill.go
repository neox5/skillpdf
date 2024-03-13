package pdf

type SkillColumn struct {
	Groups []SkillGroup `yaml:"groups"`
}

// SkillGroup bundles one or more skills
type SkillGroup struct {
	Name   string  `yaml:"name"`
	Skills []Skill `yaml:"skills"`
}

// Skill combines a name with a level from 0 to 10
type Skill struct {
	Name  string `yaml:"name"`
	Level int    `yaml:"level"`
}
