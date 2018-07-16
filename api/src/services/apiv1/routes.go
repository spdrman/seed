package auth

import (
	"github.com/gin-gonic/gin"
)

// Routes for importing in main server router by calling the auth.Routes(router) function
func Routes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.GET("/register", register)
		//other routes here

		auth.GET("/verifyemail", verifyemail)

		auth.GET("/login", login)

		auth.GET("/forgotpassword", forgotPassword)

	}
}
