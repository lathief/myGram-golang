package middleware

import (
	"myGram/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.NewResponse(http.StatusUnauthorized, nil, err.Error()))
			return
		}
		c.Set("userData", verifyToken)
		c.Next()
	}
}
