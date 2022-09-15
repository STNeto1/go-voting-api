package auth

import (
	"fmt"
	"net/http"

	"ginn/pkg/common/authorization"
	"ginn/pkg/common/error"

	"github.com/gin-gonic/gin"
)

func (h handler) Profile(c *gin.Context) {
	user, err := authorization.ExtractUser(c)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, error.Unauthorized("Unauthorized"))
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}
