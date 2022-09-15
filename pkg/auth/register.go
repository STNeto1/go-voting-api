package auth

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"voting/pkg/common/error"
	"voting/pkg/common/models"

	"github.com/gin-gonic/gin"
)

type RegisterRequestBody struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h handler) Register(c *gin.Context) {
	body := RegisterRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, error.BadRequest("Invalid request body"))
		return
	}

	existingUser := models.User{}
	h.DB.First(&existingUser, "email = ?", body.Email)
	if existingUser.ID != "" {
		c.JSON(http.StatusBadRequest, error.BadRequest("Email already in use"))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, error.InternalServerError("Internal error"))
		return
	}

	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hashedPassword),
	}

	if result := h.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, error.InternalServerError("Internal error"))
		return
	}

	c.JSON(http.StatusCreated, gin.H{"statusCode": http.StatusCreated, "message": "User created"})
}
