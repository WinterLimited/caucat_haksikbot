package router

import "github.com/gin-gonic/gin"

func Load(r *gin.Engine) {
	// Load home routes
	loadHaksikRoutes(r)
}
