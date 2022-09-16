package voting

import (
	"errors"
	"net/http"
	"time"
	"voting/pkg/common/models"

	"voting/pkg/common/authorization"
	"voting/pkg/common/error"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateVotingRequestBody struct {
	Title string `json:"title" binding:"required,min=3,max=255"`
	Start string `json:"start" binding:"required,datetime=2006-01-02"`
	End   string `json:"end" binding:"required,datetime=2006-01-02"`
}

func (h handler) CreateVoting(c *gin.Context) {

	body := CreateVotingRequestBody{}

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
