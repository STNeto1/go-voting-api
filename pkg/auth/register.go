package auth

import (
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"voting/pkg/common/error"
	"voting/pkg/common/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RegisterRequestBody struct {
	Name     string `json:"name" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h handler) Register(c *gin.Context) {
	body := RegisterRequestBody{}

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
