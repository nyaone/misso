package routers

import (
	"github.com/gin-gonic/gin"
)

func R(e *gin.Engine) {
	e.LoadHTMLGlob("templates/*.tmpl")

	rootGroup := e.Group("")

	// Public
	Public(rootGroup)

	// Login
	Login(rootGroup)

	// Consent
	Consent(rootGroup)

	// Logout
	Logout(rootGroup)

	// Userinfo
	User(rootGroup)
}
