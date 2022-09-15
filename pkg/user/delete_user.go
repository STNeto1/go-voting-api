package user

import (
	"fmt"
	"net/http"

	"voting/pkg/common/authorization"
	"voting/pkg/common/error"

	"github.com/gin-gonic/gin"
)

func (h handler) DeleteUser(c *gin.Context) {
	user, err := authorization.ExtractUser(c)
	fmt.Println(user)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, error.Unauthorized("Unauthorized"))
		return
	}

	h.DB.Delete(&user)

	c.JSON(http.StatusNoContent, nil)
}
