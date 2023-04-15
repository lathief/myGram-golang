package middleware

import (
	"myGram/database"
	"myGram/entities"
	"myGram/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		photoId, err := strconv.Atoi(c.Param("photoId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse(http.StatusBadRequest, nil, "Invalid Path Variable"))
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Photo := entities.Photo{}

		err = db.Select("user_id").First(&Photo, uint(photoId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, helpers.NewResponse(http.StatusNotFound, nil, "Data doesn't exist"))
			return
		}

		if Photo.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.NewResponse(http.StatusUnauthorized, nil, "You are not allowed to access this data"))
			return
		}
		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		commentId, err := strconv.Atoi(c.Param("commentId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, helpers.NewResponse(http.StatusBadRequest, nil, "Invalid Path Variable"))
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Comment := entities.Comment{}

		err = db.Select("user_id").First(&Comment, uint(commentId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, helpers.NewResponse(http.StatusNotFound, nil, "Data doesn't exist"))
			return
		}

		if Comment.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.NewResponse(http.StatusUnauthorized, nil, "You are not allowed to access this data"))
			return
		}

		c.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid Parameter",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		SocialMedia := entities.SocialMedia{}

		err = db.Select("user_id").First(&SocialMedia, uint(socialMediaId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Data Not Found",
				"message": "Data doesn't exist",
			})
			return
		}

		if SocialMedia.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}
