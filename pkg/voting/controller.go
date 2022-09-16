package voting

import (
	"voting/pkg/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/votings")
	routes.GET("/user", middlewares.AuthorizeJWT(), h.GetUserVotings)
	routes.GET("/", h.GetVotings)
	routes.POST("/", middlewares.AuthorizeJWT(), h.CreateVoting)
	routes.POST("/vote", h.VoteOption)
	routes.GET("/:id", h.GetVoting)
	routes.DELETE("/:id", middlewares.AuthorizeJWT(), h.DeleteVoting)

}
