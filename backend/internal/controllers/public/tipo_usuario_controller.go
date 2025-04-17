package controllers

import (
	"biblioteca-digital/internal/config"
	models "biblioteca-digital/internal/models/public"
	services "biblioteca-digital/internal/services/public"
	"biblioteca-digital/pkg/mapear/constants"
	"biblioteca-digital/pkg/mapear/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TipoUsuarioController struct {
	service services.TipoUsuarioService
}

func NovoTipoUsuarioController(tipoUsuarioService services.TipoUsuarioService) *TipoUsuarioController {
	return &TipoUsuarioController{
		service: tipoUsuarioService,
	}
}

func (controller *TipoUsuarioController) BuscarTipoUsuarioPorId(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.ID_NAO_INFORMADO, Description: err.Error()})
	}

	tipoUsuario, err := controller.service.BuscarTipoUsuarioPorId(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Error{Message: constants.REGISTRO_NAO_ENCONTRADO, Description: err.Error()})
	}

	return c.JSON(http.StatusOK, tipoUsuario)
}

func (controller *TipoUsuarioController) BuscarUsuarios(c echo.Context) error {
	acoes, err := controller.service.BuscarUsuarios(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Message: constants.ERRO_LISTAGEM_REGISTRO, Description: err.Error()})
	}

	return c.JSON(http.StatusOK, acoes)
}

func (controller *TipoUsuarioController) Cadastrar(c echo.Context) error {
	var tipoUsuario models.TipoUsuario
	if err := c.Bind(&tipoUsuario); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.BODY_FALHA, Description: err.Error()})
	}

	if err := c.Validate(tipoUsuario); err != nil {
		return config.ValidationErrors(c, err)
	}

	if err := controller.service.Cadastrar(&tipoUsuario); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Message: constants.CADASTRO_FALHA_INSERCAO, Description: err.Error()})
	}

	return c.JSON(http.StatusCreated, response.Success{Message: constants.CADASTRO_REALIZADO})
}

func (controller *TipoUsuarioController) Atualizar(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.ID_NAO_INFORMADO, Description: err.Error()})
	}

	var tipoUsuarioAtualizar models.TipoUsuario
	if err := c.Bind(&tipoUsuarioAtualizar); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.BODY_FALHA, Description: err.Error()})
	}

	if err := c.Validate(tipoUsuarioAtualizar); err != nil {
		return config.ValidationErrors(c, err)
	}

	tipoUsuario, err := controller.service.Atualizar(id, &tipoUsuarioAtualizar)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Error{Message: constants.REGISTRO_NAO_ENCONTRADO, Description: err.Error()})
	}

	return c.JSON(http.StatusOK, response.SuccessBody{Message: constants.CADASTRO_ALTERADO, Body: tipoUsuario})
}

func (controller *TipoUsuarioController) Ativar(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.ID_NAO_INFORMADO, Description: err.Error()})
	}

	if err := controller.service.Ativar(id); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Message: constants.CADASTRO_FALHA_EXCLUSAO, Description: err.Error()})
	}

	return c.JSON(http.StatusOK, response.Success{Message: constants.CADASTRO_EXCLUIDO})
}
