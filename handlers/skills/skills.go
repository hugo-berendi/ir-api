package skills

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hugo-berendi/ir-api/data"
	"net/http"
	"strconv"
	"strings"
)

func GetSkills(c *gin.Context, data data.Data) {
	c.JSON(http.StatusOK, data.Skills)
}

func GetSkillByIdWithLevel(c *gin.Context, data data.Data) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid skill ID"})
		return
	}

	level, err := strconv.Atoi(c.Param("level"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid level"})
		return
	}

	for idx, skill := range data.Skills {
		if idx == id {
			c.JSON(http.StatusOK, calcSkillRunes(skill, level))
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Skill not found"})
}

func calcSkillRunes(skill data.Skill, level int) data.Skill {
	for _, skillParam := range skill.FirstRune.SkillParams {
		param := skillParam.BaseValue * (skillParam.ScalingRate * level)
		skill.FirstRune.Description = strings.ReplaceAll(skill.FirstRune.Description, fmt.Sprintf("{%s}", skillParam.Name), strconv.Itoa(param))
	}
	for _, skillParam := range skill.SecondRune.SkillParams {
		param := skillParam.BaseValue * (skillParam.ScalingRate * level)
		skill.SecondRune.Description = strings.ReplaceAll(skill.SecondRune.Description, fmt.Sprintf("{%s}", skillParam.Name), strconv.Itoa(param))
	}

	return skill
}
