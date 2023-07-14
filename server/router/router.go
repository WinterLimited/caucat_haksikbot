package router

import "github.com/gin-gonic/gin"

func Load(r *gin.Engine) {
	// Load haksik routes
	loadHaksikRoutes(r)

	// Load user routes
	loadUserRoutes(r)

	// Load menu routes
	loadMenuRoutes(r)
}
