package controllers

import (
	"myGram/entities"
	"myGram/helpers"
	"myGram/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type SocialMediaController interface {
	GetAllSosmed(ctx *gin.Context)
	GetSosmedByID(ctx *gin.Context)
	CreateSosmed(ctx *gin.Context)
	UpdateSosmed(ctx *gin.Context)
	DeleteSosmed(ctx *gin.Context)
}

func NewSosmedController(srv service.SocialMediaService) SocialMediaController {
	return &socialMediaController{srv: srv}
}

type socialMediaController struct {
	srv service.SocialMediaService
}

// Get All social media
// @Tags Social Media
// @Summary Get All social media
// @Description Get All social media
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Success 200 {object} helpers.BaseResponse{data=[]entities.ResponseSocialMedia}  "OK"
// @Failure 401 {object} helpers.BaseResponse "Unauthorization"
// @Failure 500 {object} helpers.BaseResponse "Internal Server Error"
// @Router /socialmedia [GET]
func (c *socialMediaController) GetAllSosmed(ctx *gin.Context) {
	response, err := c.srv.GetAllSosmeds()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helpers.NewResponse(http.StatusOK, response, "OK"))
}

// Get one social media by id
// @Tags Social Media
// @Summary Get One social media by id
// @Description Get One social media by id
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Param photoID path int true "ID of the social media"
// @Success 200 {object} helpers.BaseResponse{data=entities.ResponseSocialMedia} "OK"
// @Failure 400 {object} helpers.BaseResponse "Bad Request"
// @Failure 401 {object} helpers.BaseResponse "Unauthorization"
// @Failure 500 {object} helpers.BaseResponse "Internal Server Error"
// @Router /socialmedia/:socialmediaid [GET]
func (c *socialMediaController) GetSosmedByID(ctx *gin.Context) {
	sosmedParamID := ctx.Param("socialMediaId")
	sosmedID, err := strconv.Atoi(sosmedParamID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	ID := uint(sosmedID)
	response, err := c.srv.GetSosmedByID(ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helpers.NewResponse(http.StatusOK, response, "OK"))
}

// Create new social media
// @Tags Social Media
// @Summary Create new social media
// @Description Create social media
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Param data body entities.RequestSocialMedia true "data"
// @Success 201 {object} helpers.BaseResponse "Created"
// @Failure 400 {object} helpers.BaseResponse "Bad Request"
// @Failure 401 {object} helpers.BaseResponse "Unauthorization"
// @Failure 500 {object} helpers.BaseResponse "Internal Server Error"
// @Router /socialmedia [POST]
func (c *socialMediaController) CreateSosmed(ctx *gin.Context) {
	data := new(entities.SocialMedia)

	err := ctx.ShouldBindJSON(data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	data.UserID = uint(userData["id"].(float64))

	_, err = c.srv.CreateSosmed(*data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, helpers.NewResponse(http.StatusCreated, nil, "Created Social Media Success"))
}

// Update by id social media
// @Tags Social Media
// @Summary Update by id social media
// @Description Update by id social media
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Param socialmediaid path int true "ID of the social media"
// @Param data body entities.RequestSocialMedia true "data"
// @Success 200 {object} helpers.BaseResponse "OK"
// @Failure 400 {object} helpers.BaseResponse "Bad Request"
// @Failure 401 {object} helpers.BaseResponse "Unauthorization"
// @Failure 500 {object} helpers.BaseResponse "Internal Server Error"
// @Router /socialmedia/:socialmediaid [PUT]
func (c *socialMediaController) UpdateSosmed(ctx *gin.Context) {
	data := new(entities.SocialMedia)

	err := ctx.ShouldBindJSON(data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	sosmedParamID := ctx.Param("socialMediaId")
	sosmedID, err := strconv.Atoi(sosmedParamID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	data.UserID = uint(userData["id"].(float64))

	ID := uint(sosmedID)
	_, err = c.srv.UpdateSosmed(*data, ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helpers.NewResponse(http.StatusOK, nil, "Updated Social Media Success"))
}

// Delete by id social media
// @Tags Social Media
// @Summary Delete by id social media
// @Description Delete by id social media
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Param socialmediaid path int true "ID of the social media"
// @Success 200 {object} helpers.BaseResponse "OK"
// @Failure 400 {object} helpers.BaseResponse "Bad Request"
// @Failure 401 {object} helpers.BaseResponse "Unauthorization"
// @Failure 500 {object} helpers.BaseResponse "Internal Server Error"
// @Router /socialmedia/:socialmediaid [DELETE]
func (c *socialMediaController) DeleteSosmed(ctx *gin.Context) {
	paramKeyID := ctx.Param("socialMediaId")
	sosmedID, _ := strconv.Atoi(paramKeyID)
	ID := uint(sosmedID)
	err := c.srv.DeleteSosmed(ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helpers.NewResponse(http.StatusOK, nil, "Deleted Social Media Success"))
}
