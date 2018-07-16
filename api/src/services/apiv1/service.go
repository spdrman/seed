package auth

import (
	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "register here!",
	})
}

func login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login here!",
	})
}

func forgotPassword(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Forgot password here!",
	})
}

func verifyemail(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Verify email here!",
	})
}
