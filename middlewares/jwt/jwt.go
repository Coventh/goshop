package jwt

import (
	"goshop/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code string

		code = utils.Success
		token := c.GetHeader("Authorization")

		_, err := utils.ValidateToken(token)
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				code = utils.ErrorAuthCheckTokenTimeout
			default:
				code = utils.ErrorAuthCheckTokenFail
			}
		}

		if code != utils.Success {
			c.JSON(http.StatusUnauthorized, utils.PrepareResponse(nil, "Unauthorized", code))

			c.Abort()
			return
		}

		c.Next()
	}
}
