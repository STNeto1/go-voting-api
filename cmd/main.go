package main

import (
	"fmt"
	"os"
	"voting/pkg/auth"
	"voting/pkg/common/config"
	"voting/pkg/common/db"
	"voting/pkg/user"
	"voting/pkg/voting"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	production := os.Getenv("RAILWAY_ENVIRONMENT") == "production"

	if !production {
		viper.SetConfigFile("./pkg/common/envs/.env")
		err := viper.ReadInConfig()
		if err != nil { // Handle errors reading the config file
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	config, err := config.LoadConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error loading config: %w", err))
	}

	r := gin.Default()
	h := db.Init(config.DBUrl)

	auth.RegisterRoutes(r, h)
	user.RegisterRoutes(r, h)
	voting.RegisterRoutes(r, h, config)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Gin API!",
		})
	})

	r.Run()
}
