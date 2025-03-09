package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hugo-berendi/ir-api/src/data"
	"github.com/hugo-berendi/ir-api/src/handlers/skills"
	"log"
)

func main() {
	data, err := data_loader.LoadDataFromFile()
	if err != nil {
		log.Fatalf("Error loading data: %v", err)
	}

	r := gin.Default()

	r.GET("/skills", func(c *gin.Context) { skills.GetSkills(c, data) })
	r.GET("/skills/:id/:level", func(c *gin.Context) { skills.GetSkillByIdWithLevel(c, data) })

	r.Run(":8080")
}
