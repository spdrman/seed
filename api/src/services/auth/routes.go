package auth

import "github.com/gin-gonic/gin"

// Routes for importing in main server router by calling the auth.Routes(router) function
func Routes(r *gin.RouterGroup) {
	r.GET("/auth/register", registerHandler)

	r.GET("/auth/verifyemail", verifyEmailHandler)

	r.GET("/auth/login", loginHandler)

	r.GET("/auth/forgotpassword", forgotPasswordHandler)

	r.GET("/aut/logout", logoutHandler)
}
