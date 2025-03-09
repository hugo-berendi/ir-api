package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// structs
type SkillParam struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	BaseValue   int    `json:"baseValue"`
	ScalingRate int    `json:"scalingRate"`
}

type Skill struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	SkillParams []SkillParam `json:"skillParams"`
}

type Data struct {
	Skills []Skill `json:"skills"`
}

var data Data

func loadDataFromFile() {
	file, err := os.Open("data.json")
	if err != nil {
		panic("Failed to open data.json")
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic("Failed to read data.json")
	}

	if err := json.Unmarshal(bytes, &data); err != nil {
		panic("Failed to parse json")
	}
}

func getSkills(c *gin.Context) {
	c.JSON(http.StatusOK, data.Skills)
}

func getSkillByIdWithLevel(c *gin.Context) {
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

	for _, skill := range data.Skills {
		if skill.ID == id {
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

func main() {
	loadDataFromFile()
	r := gin.Default()

	r.GET("/skills", getSkills)
	r.GET("/skills/:id/:level", getSkillByIdWithLevel)

	r.Run(":8080")
}
