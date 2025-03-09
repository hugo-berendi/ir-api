package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hugo-berendi/ir-api/data"
	"github.com/hugo-berendi/ir-api/handlers/skills"
	"log"
)

func main() {
	data, err := data.LoadDataFromFile()
	if err != nil {
		log.Fatalf("Error loading data: %v", err)
	}

	r := gin.Default()

	r.GET("/skills", func(c *gin.Context) { skills.GetSkills(c, data) })
	r.GET("/skills/:id/:level", func(c *gin.Context) { skills.GetSkillByIdWithLevel(c, data) })

	r.Run(":8080")
}
