package controllers

import (
	"ApiSup/internal/config"
	"ApiSup/internal/models"
	"ApiSup/internal/services"
	"ApiSup/pkg/mapear/constants"
	"ApiSup/pkg/mapear/request"
	"ApiSup/pkg/mapear/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	service services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		service: userService,
	}
}

func (controller *UserController) Login(c echo.Context) error {
	var body request.Login
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.BODY_FALHA, Description: err.Error()})
	}

	if err := c.Validate(body); err != nil {
		return config.ValidationErrors(c, err)
	}

	var usuario *models.Usuario
	usuario, err := controller.service.Authenticate(body)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.Error{Message: constants.ACESSO_NAO_AUTORIZADO, Description: err.Error()})
	}

	tokenString, err := config.GenerateJWT(*usuario)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Message: constants.FALHA_AUTENTICAO, Description: err.Error()})
	}

	response := response.Token{
		Token:   tokenString,
		Usuario: *usuario,
	}

	return c.JSON(http.StatusOK, response)
}

func (controller *UserController) Listagem(c echo.Context) error {
	users, err := controller.service.Listagem(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Message: constants.ERRO_LISTAGEM_REGISTRO, Description: err.Error()})
	}

	return c.JSON(http.StatusOK, users)
}

func (controller *UserController) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.ID_NAO_INFORMADO, Description: err.Error()})
	}

	user, err := controller.service.Detalhar(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Error{Message: constants.REGISTRO_NAO_ENCONTRADO, Description: err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (controller *UserController) Created(c echo.Context) error {
	var user models.Usuario
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.BODY_FALHA, Description: err.Error()})
	}

	if err := c.Validate(user); err != nil {
		return config.ValidationErrors(c, err)
	}

	if err := controller.service.Novo(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Message: constants.CADASTRO_FALHA_INSERCAO, Description: err.Error()})
	}

	return c.JSON(http.StatusCreated, response.Success{Message: constants.CADASTRO_REALIZADO})
}

func (controller *UserController) Updated(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.ID_NAO_INFORMADO, Description: err.Error()})
	}

	var updatedUser models.Usuario
	if err := c.Bind(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.BODY_FALHA, Description: err.Error()})
	}

	if err := c.Validate(updatedUser); err != nil {
		return config.ValidationErrors(c, err)
	}

	user, err := controller.service.Editar(id, &updatedUser)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Error{Message: constants.REGISTRO_NAO_ENCONTRADO, Description: err.Error()})
	}

	return c.JSON(http.StatusOK, response.SuccessBody{Message: constants.CADASTRO_ALTERADO, Body: user})
}

func (controller *UserController) Deleted(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.ID_NAO_INFORMADO, Description: err.Error()})
	}

	if err := controller.service.Deletar(id); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Message: constants.CADASTRO_FALHA_EXCLUSAO, Description: err.Error()})
	}

	return c.JSON(http.StatusOK, response.Success{Message: constants.CADASTRO_EXCLUIDO})
}
