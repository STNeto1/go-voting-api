package auth

import (
	"voting/pkg/common/authorization"

	"github.com/gin-gonic/gin"
)

func (h handler) Profile(c *gin.Context) {
	user := authorization.ExtractUser(c)

	c.JSON(200, user)
}
