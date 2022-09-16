package voting

import (
	"fmt"
	"net/http"

	"voting/pkg/common/authorization"
	"voting/pkg/common/error"
	"voting/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) GetUserVotings(c *gin.Context) {
	user, err := authorization.ExtractUser(c)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, error.Unauthorized("Unauthorized"))
		return
	}

	var votings []models.Voting

	if result := h.DB.Find(&votings).Where("user_id = ?", user.ID); result.Error != nil {
		c.JSON(http.StatusBadRequest, error.BadRequest("Error fetching votings"))
		return
	}

	c.JSON(http.StatusOK, &votings)
}
