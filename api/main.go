package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_postgres_s3/api/modules"
	"log"
)

func main() {

	router := gin.Default()
	gin.SetMode(gin.DebugMode)

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/url.insert", InsertUrl)

	}
	err := router.Run(":8080")
	if err != nil {
		log.Println(err)
	}

}

func InsertUrl(c *gin.Context) {

	id := modules.InsertDBurl("localhost.local")
	fmt.Println(id)
}
