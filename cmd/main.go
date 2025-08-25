package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	_ = router

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
