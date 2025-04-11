package controllers

import (
	"biblioteca-digital/internal/config"
	"biblioteca-digital/internal/models/public"
	services "biblioteca-digital/internal/services/public"
	"biblioteca-digital/pkg/mapear/constants"
	"biblioteca-digital/pkg/mapear/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TipoContatoController struct {
	service services.TipoContatoService
}

func NovoTipoContatoController(tipoContatoService services.TipoContatoService) *TipoContatoController {
	return &TipoContatoController{
		service: tipoContatoService,
	}
}

func (controller *TipoContatoController) BuscarTipoContatoPorId(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.ID_NAO_INFORMADO, Description: err.Error()})
	}

	tipoContato, err := controller.service.BuscarTipoContatoPorId(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Error{Message: constants.REGISTRO_NAO_ENCONTRADO, Description: err.Error()})
	}

	return c.JSON(http.StatusOK, tipoContato)
}

func (controller *TipoContatoController) BuscarTiposContatos(c echo.Context) error {
	assunto, err := controller.service.BuscarTiposContatos(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Message: constants.ERRO_LISTAGEM_REGISTRO, Description: err.Error()})
	}

	return c.JSON(http.StatusOK, assunto)
}

func (controller *TipoContatoController) Cadastrar(c echo.Context) error {
	var tipoContato models.TipoContato
	if err := c.Bind(&tipoContato); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.BODY_FALHA, Description: err.Error()})
	}

	if err := c.Validate(tipoContato); err != nil {
		return config.ValidationErrors(c, err)
	}

	if err := controller.service.Cadastrar(&tipoContato); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Message: constants.CADASTRO_FALHA_INSERCAO, Description: err.Error()})
	}

	return c.JSON(http.StatusCreated, response.Success{Message: constants.CADASTRO_REALIZADO})
}

func (controller *TipoContatoController) Atualizar(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.ID_NAO_INFORMADO, Description: err.Error()})
	}

	var tipoContatoAtualizar models.TipoContato
	if err := c.Bind(&tipoContatoAtualizar); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.BODY_FALHA, Description: err.Error()})
	}

	if err := c.Validate(tipoContatoAtualizar); err != nil {
		return config.ValidationErrors(c, err)
	}

	tipoContato, err := controller.service.Atualizar(id, &tipoContatoAtualizar)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Error{Message: constants.REGISTRO_NAO_ENCONTRADO, Description: err.Error()})
	}

	return c.JSON(http.StatusOK, response.SuccessBody{Message: constants.CADASTRO_ALTERADO, Body: tipoContato})
}

func (controller *TipoContatoController) Ativar(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.ID_NAO_INFORMADO, Description: err.Error()})
	}

	if err := controller.service.Ativar(id); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Message: constants.CADASTRO_FALHA_EXCLUSAO, Description: err.Error()})
	}

	return c.JSON(http.StatusOK, response.Success{Message: constants.CADASTRO_EXCLUIDO})
}
