package voting

import (
	"voting/pkg/common/config"
	"voting/pkg/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/pusher/pusher-http-go"
	"gorm.io/gorm"
)

type handler struct {
	DB     *gorm.DB
	Pusher *pusher.Client
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB, config config.Config) {
	h := &handler{
		DB: db,
		Pusher: &pusher.Client{
			AppID:   config.PusherAppID,
			Key:     config.PusherKey,
			Secret:  config.PusherSecret,
			Cluster: config.PusherCluster,
			Secure:  config.PusherSecure,
		},
	}

	routes := r.Group("/votings")
	routes.GET("/user", middlewares.AuthorizeJWT(), h.GetUserVotings)
	routes.GET("/", h.GetVotings)
	routes.POST("/", middlewares.AuthorizeJWT(), h.CreateVoting)
	routes.POST("/vote", h.VoteOption)
	routes.GET("/:id", h.GetVoting)
	routes.DELETE("/:id", middlewares.AuthorizeJWT(), h.DeleteVoting)

}
