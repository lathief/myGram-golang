package controllers

import (
	"myGram/entities"
	"myGram/helpers"
	"myGram/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	srv service.UserService
}

func NewUserController(srv service.UserService) AuthController {
	return &authController{srv}
}

// Register new user
// @Tags User
// @Summary Register new user
// @Description Register new user
// @Accept  json
// @Produce  json
// @Param data body entities.RequestRegister true "data"
// @Success 201 {object} helpers.BaseResponse "Created"
// @Failure 400 {object} helpers.BaseResponse "Bad Request"
// @Failure 500 {object} helpers.BaseResponse "Internal Server Error"
// @Router /user/register [POST]
func (c *authController) Register(ctx *gin.Context) {
	data := new(entities.User)
	dataReq := new(entities.RequestRegister)
	if err := ctx.ShouldBindJSON(dataReq); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	data.Email = dataReq.Email
	data.Password = dataReq.Password
	data.Username = dataReq.Username
	data.Age = dataReq.Age
	err := c.srv.Create(data)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, helpers.NewResponse(http.StatusCreated, nil, "Register Successfully"))
}

// Login new user
// @Tags User
// @Summary Login new user
// @Description Login new user
// @Accept  json
// @Produce  json
// @Param data body entities.RequestLogin true "data"
// @Success 200 {object} helpers.BaseResponse "OK"
// @Failure 400 {object} helpers.BaseResponse "Bad Request"
// @Failure 500 {object} helpers.BaseResponse "Internal Server Error"
// @Router /user/login [POST]
func (c *authController) Login(ctx *gin.Context) {
	data := new(entities.RequestLogin)

	if err := ctx.ShouldBindJSON(data); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}

	response, err := c.srv.Login(data)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helpers.NewResponse(http.StatusOK, response, "Login Successfully"))
}
