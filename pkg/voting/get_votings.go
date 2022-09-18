package voting

import (
	"net/http"

	"voting/pkg/common/exceptions"
	"voting/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) GetVotings(c *gin.Context) {
	var votings []models.Voting

	if result := h.DB.Find(&votings); result.Error != nil {
		c.JSON(http.StatusBadRequest, exceptions.BadRequest("Error fetching votings"))
		return
	}

	c.JSON(http.StatusOK, &votings)
}
