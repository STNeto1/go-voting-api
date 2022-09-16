package voting

import (
	"net/http"

	"voting/pkg/common/error"
	"voting/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) GetVoting(c *gin.Context) {
	id := c.Param("id")

	var voting models.Voting
	if err := h.DB.First(&voting, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, error.NotFound("Voting not found"))
		return
	}

	c.JSON(http.StatusOK, voting)
}
