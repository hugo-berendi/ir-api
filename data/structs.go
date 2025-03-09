package data

// Skill structs
type SkillParam struct {
	Name        string `json:"name"`
	BaseValue   int    `json:"baseValue"`
	ScalingRate int    `json:"scalingRate"`
}

type Skill struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	SkillParams []SkillParam `json:"skillParams"`
}

// Main data struct
type Data struct {
	Skills []Skill `json:"skills"`
}
