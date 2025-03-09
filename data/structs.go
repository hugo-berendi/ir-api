package data

// Skill structs
type SkillParam struct {
	Name        string `json:"name"`
	BaseValue   int    `json:"baseValue"`
	ScalingRate int    `json:"scalingRate"`
}

type SkillRune struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	SkillParams []SkillParam `json:"skillParams"`
}

type Skill struct {
	FirstRune  SkillRune `json:"firstRune"`
	SecondRune SkillRune `json:"secondRune"`
}

// Main data struct
type Data struct {
	Skills []Skill `json:"skills"`
}
