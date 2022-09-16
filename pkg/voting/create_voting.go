package voting

import (
	"net/http"
	"time"
	"voting/pkg/common/models"

	"voting/pkg/common/authorization"
	"voting/pkg/common/error"

	"github.com/gin-gonic/gin"
)

type CreateVotingRequestBody struct {
	Title string `json:"title" binding:"required"`
	Start string `json:"start" binding:"required"`
	End   string `json:"end" binding:"required"`
}

func (h handler) CreateVoting(c *gin.Context) {

	body := CreateVotingRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, error.BadRequest("Invalid request body"))
		return
	}

	user, err := authorization.ExtractUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, error.Unauthorized("Unauthorized"))
		return
	}

	start, _ := time.Parse("2006-01-02", body.Start)
	end, _ := time.Parse("2006-01-02", body.End)

	voting := models.Voting{
		Title:  body.Title,
		Start:  start,
		End:    end,
		UserID: user.ID,
	}

	if result := h.DB.Create(&voting); result.Error != nil {
		c.JSON(http.StatusInternalServerError, error.InternalServerError("Internal error"))
		return
	}

	c.JSON(http.StatusCreated, voting)
}
