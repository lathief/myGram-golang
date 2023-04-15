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

type CommentController interface {
	GetAllComments(ctx *gin.Context)
	GetCommentByID(ctx *gin.Context)
	CreateComment(ctx *gin.Context)
	UpdateComment(ctx *gin.Context)
	DeleteComment(ctx *gin.Context)
}

func NewCommentController(srv service.CommentService) CommentController {
	return &commentController{srv: srv}
}

type commentController struct {
	srv service.CommentService
}

// Get All comments
// @Tags Comment
// @Summary Get comments
// @Description Get comments
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Success 200 {object} helpers.BaseResponse{data=[]entities.ResponseComment}  "OK"
// @Failure 401 {object} helpers.BaseResponse "Unauthorization"
// @Failure 500 {object} helpers.BaseResponse "Internal Server Error"
// @Router /comment [GET]
func (c *commentController) GetAllComments(ctx *gin.Context) {
	response, err := c.srv.GetAllComments()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helpers.NewResponse(http.StatusOK, response, "OK"))
}

// Get one comment by id
// @Tags Comment
// @Summary Get one comment by id
// @Description Get one comment by id
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Param photoID path int true "ID of the Comment"
// @Success 200 {object} helpers.BaseResponse{data=entities.ResponseComment} "OK"
// @Failure 400 {object} helpers.BaseResponse "Bad Request"
// @Failure 401 {object} helpers.BaseResponse "Unauthorization"
// @Failure 500 {object} helpers.BaseResponse "Internal Server Error"
// @Router /comment/:commentID [GET]
func (c *commentController) GetCommentByID(ctx *gin.Context) {
	commentParamID := ctx.Param("commentId")
	commentID, err := strconv.Atoi(commentParamID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	ID := uint(commentID)
	response, err := c.srv.GetCommentByID(ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helpers.NewResponse(http.StatusOK, response, "OK"))
}

// Create a comment
// @Tags Comment
// @Summary Create a comment
// @Description Create a comment
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Param commentID path int true "ID of the comment"
// @Param data body entities.RequestComment true "data"
// @Success 201 {object} helpers.BaseResponse "Created"
// @Failure 400 {object} helpers.BaseResponse "Bad Request"
// @Failure 401 {object} helpers.BaseResponse "Unauthorization"
// @Failure 500 {object} helpers.BaseResponse "Internal Server Error"
// @Router /comment [POST]
func (c *commentController) CreateComment(ctx *gin.Context) {
	data := new(entities.Comment)

	err := ctx.ShouldBindJSON(data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	data.UserID = uint(userData["id"].(float64))

	_, err = c.srv.CreateComment(*data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, helpers.NewResponse(http.StatusCreated, nil, "Created Comment Success"))
}

// Update a comment
// @Tags Comment
// @Summary Update a comment
// @Description Update a comment
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Param commentID path int true "ID of the comment"
// @Param data body entities.RequestComment true "data"
// @Success 200 {object} helpers.BaseResponse "OK"
// @Failure 400 {object} helpers.BaseResponse "Bad Request"
// @Failure 401 {object} helpers.BaseResponse "Unauthorization"
// @Failure 500 {object} helpers.BaseResponse "Internal Server Error"
// @Router /comment/:commentID [PUT]
func (c *commentController) UpdateComment(ctx *gin.Context) {
	data := new(entities.Comment)

	err := ctx.ShouldBindJSON(data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	commentParamID := ctx.Param("commentId")
	commentID, err := strconv.Atoi(commentParamID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	data.UserID = uint(userData["id"].(float64))

	ID := uint(commentID)
	_, err = c.srv.UpdateComment(*data, ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helpers.NewResponse(http.StatusOK, nil, "Updated Comment Success"))
}

// Delete a comment
// @Tags Comment
// @Summary Delete a comment
// @Description Delete a comment
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Param commentID path int true "ID of the comment"
// @Success 200 {object} helpers.BaseResponse "OK"
// @Failure 400 {object} helpers.BaseResponse "Bad Request"
// @Failure 401 {object} helpers.BaseResponse "Unauthorization"
// @Failure 500 {object} helpers.BaseResponse "Internal Server Error"
// @Router /comment/:commentID [DELETE]
func (c *commentController) DeleteComment(ctx *gin.Context) {
	paramKeyID := ctx.Param("commentId")
	commentID, _ := strconv.Atoi(paramKeyID)
	ID := uint(commentID)
	err := c.srv.DeleteComment(ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helpers.NewResponse(http.StatusOK, nil, "Deleted Comment Success"))
}
