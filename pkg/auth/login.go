package auth

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"voting/pkg/common/authorization"
	"voting/pkg/common/error"
	"voting/pkg/common/models"

	"github.com/gin-gonic/gin"
)

type LoginRequestBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h handler) Login(c *gin.Context) {
	body := LoginRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, error.BadRequest("Invalid request body"))
		return
	}

	user := models.User{}
	h.DB.First(&user, "email = ?", body.Email)
	if user.ID == "" {
		c.JSON(http.StatusBadRequest, error.BadRequest("Invalid credentials"))
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, error.BadRequest("Invalid credentials"))
		return
	}

	token := authorization.GenerateToken(user.ID)

	c.JSON(http.StatusOK, gin.H{"access_token": token})
}
