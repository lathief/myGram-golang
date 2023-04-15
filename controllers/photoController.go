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

type PhotoController interface {
	GetAllPhotos(ctx *gin.Context)
	GetPhotoByID(ctx *gin.Context)
	CreatePhoto(ctx *gin.Context)
	UpdatePhoto(ctx *gin.Context)
	DeletePhoto(ctx *gin.Context)
}

func NewPhotoController(srv service.PhotoService) PhotoController {
	return &photoController{srv: srv}
}

type photoController struct {
	srv service.PhotoService
}

// Get All Photos
// @Tags Photo
// @Summary Get All Photos
// @Description Get All Photos
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Success 200 {object} helpers.BaseResponse{data=[]entities.ResponsePhoto}  "OK"
// @Failure 401 {object} helpers.BaseResponse "Unauthorization"
// @Failure 500 {object} helpers.BaseResponse "Internal Server Error"
// @Router /photo [GET]
func (c *photoController) GetAllPhotos(ctx *gin.Context) {
	response, err := c.srv.GetAllPhotos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helpers.NewResponse(http.StatusOK, response, "OK"))
}

// Get one photo by id
// @Tags Photo
// @Summary Get One photo by id
// @Description Get One photo by id
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Param photoID path int true "ID of the photo"
// @Success 200 {object} helpers.BaseResponse{data=entities.ResponsePhoto} "OK"
// @Failure 400 {object} helpers.BaseResponse "Bad Request"
// @Failure 401 {object} helpers.BaseResponse "Unauthorization"
// @Failure 500 {object} helpers.BaseResponse "Internal Server Error"
// @Router /photo/:photoId [GET]
func (c *photoController) GetPhotoByID(ctx *gin.Context) {
	photoParamID := ctx.Param("photoId")
	photoID, err := strconv.Atoi(photoParamID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	ID := uint(photoID)
	response, err := c.srv.GetPhotoByID(ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helpers.NewResponse(http.StatusOK, response, "OK"))
}

// Create new photo
// @Tags Photo
// @Summary Create new photo
// @Description Create new photo
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Param data body entities.RequestPhoto true "data"
// @Success 201 {object} helpers.BaseResponse "Created"
// @Failure 400 {object} helpers.BaseResponse "Bad Request"
// @Failure 401 {object} helpers.BaseResponse "Unauthorization"
// @Failure 500 {object} helpers.BaseResponse "Internal Server Error"
// @Router /photo [POST]
func (c *photoController) CreatePhoto(ctx *gin.Context) {
	dataReq := new(entities.RequestPhoto)
	data := new(entities.Photo)
	err := ctx.ShouldBindJSON(dataReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	data.UserID = uint(userData["id"].(float64))
	data.Title = dataReq.Title
	data.Caption = dataReq.Caption
	data.PhotoURL = dataReq.PhotoURL
	_, err = c.srv.CreatePhoto(*data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, helpers.NewResponse(http.StatusCreated, nil, "Created Photo Success"))
}

// Update a photo
// @Tags Photo
// @Summary Update a photo
// @Description Update a photo
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Param photoID path int true "ID of the photo"
// @Param data body entities.RequestPhoto true "data"
// @Success 200 {object} helpers.BaseResponse "OK"
// @Failure 400 {object} helpers.BaseResponse "Bad Request"
// @Failure 401 {object} helpers.BaseResponse "Unauthorization"
// @Failure 500 {object} helpers.BaseResponse "Internal Server Error"
// @Router /photo/:photoId [PUT]
func (c *photoController) UpdatePhoto(ctx *gin.Context) {
	dataReq := new(entities.RequestPhoto)
	data := new(entities.Photo)

	err := ctx.ShouldBindJSON(dataReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	photoParamID := ctx.Param("photoId")
	photoID, err := strconv.Atoi(photoParamID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	data.UserID = uint(userData["id"].(float64))
	data.Title = dataReq.Title
	data.Caption = dataReq.Caption
	data.PhotoURL = dataReq.PhotoURL
	ID := uint(photoID)
	_, err = c.srv.UpdatePhoto(*data, ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helpers.NewResponse(http.StatusOK, nil, "Updated Photo Success"))
}

// Delete a photo
// @Tags Photo
// @Summary Delete a photo
// @Description Delete a photo
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Param photoID path int true "ID of the photo"
// @Success 200 {object} helpers.BaseResponse "OK"
// @Failure 401 {object} helpers.BaseResponse "Unauthorization"
// @Failure 500 {object} helpers.BaseResponse "Internal Server Error"
// @Router /photo/:photoId [DELETE]
func (c *photoController) DeletePhoto(ctx *gin.Context) {
	paramKeyID := ctx.Param("photoId")
	photoID, _ := strconv.Atoi(paramKeyID)
	ID := uint(photoID)
	err := c.srv.DeletePhoto(ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helpers.NewResponse(http.StatusOK, nil, "Deleted Photo Success"))
}
