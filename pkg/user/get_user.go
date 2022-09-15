package user

import (
	"net/http"

	"ginn/pkg/common/error"
	"ginn/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) GetUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if result := h.DB.First(&user, "id = ?", id); result.Error != nil {
		c.JSON(http.StatusNotFound, error.NotFound("User was not found"))
		return
	}

	c.JSON(http.StatusOK, &user)
}
