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
			for _, skillParam := range skill.SkillParams {
				param := skillParam.BaseValue * (skillParam.ScalingRate * level)
				skill.Description = strings.ReplaceAll(skill.Description, fmt.Sprintf("{%s}", skillParam.Name), strconv.Itoa(param))
			}
			c.JSON(http.StatusOK, gin.H{"name": skill.Name, "description": skill.Description})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Skill not found"})
}
