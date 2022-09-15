package user

import (
	"fmt"
	"net/http"

	"ginn/pkg/common/authorization"
	"ginn/pkg/common/error"
	"ginn/pkg/common/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UpdateUserRequestBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h handler) UpdateUser(c *gin.Context) {
	user, err := authorization.ExtractUser(c)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, error.Unauthorized("Unauthorized"))
		return
	}

	body := UpdateUserRequestBody{}

	// // getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, error.BadRequest("Invalid request body"))
		return
	}

	if body.Name != "" {
		user.Name = body.Name
	}

	if body.Email != "" && body.Email != user.Email {
		existingUser := models.User{}
		h.DB.First(&existingUser, "email = ?", body.Email)
		if existingUser.ID != "" {
			c.JSON(http.StatusBadRequest, error.BadRequest("Email already in use"))
			return
		}

		user.Email = body.Email
	}

	if body.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, error.InternalServerError("Internal error"))
			return
		}
		user.Password = string(hashedPassword)
	}

	if result := h.DB.Save(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, error.InternalServerError("Internal error"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
