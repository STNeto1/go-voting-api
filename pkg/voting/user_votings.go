package voting

import (
	"net/http"

	"voting/pkg/common/authorization"
	"voting/pkg/common/error"
	"voting/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) GetUserVotings(c *gin.Context) {
	user := authorization.ExtractUser(c)

	var votings []models.Voting

	if result := h.DB.Find(&votings).Where("user_id = ?", user.ID); result.Error != nil {
		c.JSON(http.StatusBadRequest, error.BadRequest("Error fetching votings"))
		return
	}

	c.JSON(http.StatusOK, &votings)
}
