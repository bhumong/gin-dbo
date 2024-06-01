package middleware

import (
	"net/http"
	"strings"

	errors "github.com/bhumong/dbo-test/pkg/error"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
type RestError struct {
	Code    string `json:"code"`
	Field   string `json:"field"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}
type RestErrors struct {
	Errors []RestError `json:"errors"`
}

func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err != nil {
			switch e := err.Err.(type) {
			case errors.Error:
				c.JSON(e.HttpCode, gin.H{
					"error":       e.Message,
					"description": e.Description,
				})
			case validator.ValidationErrors:
				restErrors := RestErrors{
					Errors: make([]RestError, 0),
				}
				for _, fieldError := range e {
					restErrors.Errors = append(restErrors.Errors, RestError{
						Code:    strings.ToUpper("validation_" + fieldError.Tag()),
						Field:   fieldError.Field(),
						Message: fieldError.Error(),
						Detail:  "",
					})
				}
				c.JSON(http.StatusBadRequest, restErrors)
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
			}

			c.Abort()
		}
	}
}
