package voting

import (
	"net/http"

	"voting/pkg/common/authorization"
	"voting/pkg/common/error"
	"voting/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) DeleteVoting(c *gin.Context) {

	user := authorization.ExtractUser(c)

	id := c.Param("id")

	var voting models.Voting
	if err := h.DB.First(&voting, "id = ? AND user_id = ?", id, user.ID).Error; err != nil {
		c.JSON(http.StatusNotFound, error.NotFound("Voting not found"))
		return
	}

	if result := h.DB.Delete(&voting); result.Error != nil {
		c.JSON(http.StatusInternalServerError, error.InternalServerError("Internal error"))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
