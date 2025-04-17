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

type AcaoController struct {
	service services.AcaoService
}

func NovaAcaoController(acaoService services.AcaoService) *AcaoController {
	return &AcaoController{
		service: acaoService,
	}
}

func (controller *AcaoController) BuscarAcaoPorId(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.ID_NAO_INFORMADO, Description: err.Error()})
	}

	acao, err := controller.service.BuscarAcaoPorId(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Error{Message: constants.REGISTRO_NAO_ENCONTRADO, Description: err.Error()})
	}

	return c.JSON(http.StatusOK, acao)
}

func (controller *AcaoController) BuscarAcoes(c echo.Context) error {
	acoes, err := controller.service.BuscarAcoes(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Message: constants.ERRO_LISTAGEM_REGISTRO, Description: err.Error()})
	}

	return c.JSON(http.StatusOK, acoes)
}

func (controller *AcaoController) Cadastrar(c echo.Context) error {
	var acao models.Acao
	if err := c.Bind(&acao); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.BODY_FALHA, Description: err.Error()})
	}

	if err := c.Validate(acao); err != nil {
		return config.ValidationErrors(c, err)
	}

	if err := controller.service.Cadastrar(&acao); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Message: constants.CADASTRO_FALHA_INSERCAO, Description: err.Error()})
	}

	return c.JSON(http.StatusCreated, response.Success{Message: constants.CADASTRO_REALIZADO})
}

func (controller *AcaoController) Atualizar(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.ID_NAO_INFORMADO, Description: err.Error()})
	}

	var acaoAtualizar models.Acao
	if err := c.Bind(&acaoAtualizar); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.BODY_FALHA, Description: err.Error()})
	}

	if err := c.Validate(acaoAtualizar); err != nil {
		return config.ValidationErrors(c, err)
	}

	acao, err := controller.service.Atualizar(id, &acaoAtualizar)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Error{Message: constants.REGISTRO_NAO_ENCONTRADO, Description: err.Error()})
	}

	return c.JSON(http.StatusOK, response.SuccessBody{Message: constants.CADASTRO_ALTERADO, Body: acao})
}

func (controller *AcaoController) Ativar(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.ID_NAO_INFORMADO, Description: err.Error()})
	}

	if err := controller.service.Ativar(id); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Message: constants.CADASTRO_FALHA_EXCLUSAO, Description: err.Error()})
	}

	return c.JSON(http.StatusOK, response.Success{Message: constants.CADASTRO_EXCLUIDO})
}
