package main

import "github.com/gin-gonic/gin"
import "bitbucket.org/devshell/yolo/backend/src/services/auth"

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/api/v1")
	{
		auth.Routes(v1)
	}

	router.Run(":8080") // listen and serve on 127.0.0.1:8080
}
