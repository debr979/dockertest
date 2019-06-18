package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", Helloworld)

	if err := r.Run(":8081"); err != nil {
		log.Printf("%s", err)
	}
}

func Helloworld(c *gin.Context) {
	c.JSON(200, gin.H{"STATUS": "SUCCESS"})
	log.Printf("%s", "Hello World")
}
