package auth

import (
	"github.com/gin-gonic/gin"
)

func registerHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "register here!",
	})
}

func loginHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login here!",
	})
}

func forgotPasswordHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Forgot password here!",
	})
}

func verifyEmailHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Verify email here!",
	})
}

func logoutHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "logout here!",
	})
}
