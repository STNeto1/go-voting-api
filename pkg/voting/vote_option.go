package voting

import (
	"errors"
	"net/http"
	"time"
	"voting/pkg/common/models"

	"voting/pkg/common/error"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateVoteRequestBody struct {
	Voting string `json:"voting" binding:"required"`
	Option string `json:"option" binding:"required"`
}

func (h handler) VoteOption(c *gin.Context) {

	body := CreateVoteRequestBody{}

	// getting request's body
	if err := c.ShouldBindJSON(&body); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]error.ApiError, len(ve))
			for i, fe := range ve {
				out[i] = error.ApiError{Param: fe.Field(), Message: error.MsgForTag(fe)}
			}

			c.JSON(http.StatusBadRequest, error.BadValidation(out))
			return
		}

	}

	var voting models.Voting
	if err := h.DB.First(&voting, "id = ?", body.Voting).Error; err != nil {
		c.JSON(http.StatusNotFound, error.NotFound("Voting not found"))
		return
	}

	if time.Now().Before(voting.Start) {
		c.JSON(http.StatusBadRequest, error.BadRequest("Voting has not started"))
		return
	}

	if time.Now().After(voting.End) {
		c.JSON(http.StatusBadRequest, error.BadRequest("Voting has ended"))
		return
	}

	var option models.VotingOption
	if err := h.DB.First(&option, "id = ? AND voting_id = ?", body.Option, body.Voting).Error; err != nil {
		c.JSON(http.StatusNotFound, error.NotFound("Option not found"))
		return
	}

	option.Votes++
	h.DB.Model(&option).Update("votes", option.Votes)

	c.JSON(http.StatusCreated, gin.H{"message": "Voted successfully"})
}
