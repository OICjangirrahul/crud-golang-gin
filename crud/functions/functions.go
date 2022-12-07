package functions

import (
	"github.com/go-playground/validator/v10"
	"crud/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	}
	return "Unknown error"
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func Errr(user *models.User, c *gin.Context) {
	if err := c.ShouldBindJSON(user); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
				//    a := []string{}
				//    a = append(a,out[i].Field)
				//    a = append(a,out[i].Message)
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}

		return
	}
}
