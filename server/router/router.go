package router

import (
	"github.com/gin-gonic/gin"
	"server/store"
)

func Load(r *gin.Engine, s *store.Store) {
	// Load haksik routes
	loadHaksikRoutes(r, s)

	// Load user routes
	loadUserRoutes(r, s)

	// Load menu routes
	loadMenuRoutes(r, s)
}
