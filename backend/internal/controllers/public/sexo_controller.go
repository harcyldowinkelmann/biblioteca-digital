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

type SexoController struct {
	service services.SexoService
}

func NovoSexoController(sexoService services.SexoService) *SexoController {
	return &SexoController{
		service: sexoService,
	}
}

func (controller *SexoController) BuscarSexoPorId(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.ID_NAO_INFORMADO, Description: err.Error()})
	}

	sexo, err := controller.service.BuscarSexoPorId(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Error{Message: constants.REGISTRO_NAO_ENCONTRADO, Description: err.Error()})
	}

	return c.JSON(http.StatusOK, sexo)
}

func (controller *SexoController) BuscarSexos(c echo.Context) error {
	sexo, err := controller.service.BuscarSexos(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Message: constants.ERRO_LISTAGEM_REGISTRO, Description: err.Error()})
	}

	return c.JSON(http.StatusOK, sexo)
}

func (controller *SexoController) Cadastrar(c echo.Context) error {
	var sexo models.Sexo
	if err := c.Bind(&sexo); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.BODY_FALHA, Description: err.Error()})
	}

	if err := c.Validate(sexo); err != nil {
		return config.ValidationErrors(c, err)
	}

	if err := controller.service.Cadastrar(&sexo); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Message: constants.CADASTRO_FALHA_INSERCAO, Description: err.Error()})
	}

	return c.JSON(http.StatusCreated, response.Success{Message: constants.CADASTRO_REALIZADO})
}

func (controller *SexoController) Atualizar(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.ID_NAO_INFORMADO, Description: err.Error()})
	}

	var sexoAtualizar models.Sexo
	if err := c.Bind(&sexoAtualizar); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.BODY_FALHA, Description: err.Error()})
	}

	if err := c.Validate(sexoAtualizar); err != nil {
		return config.ValidationErrors(c, err)
	}

	sexo, err := controller.service.Atualizar(id, &sexoAtualizar)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Error{Message: constants.REGISTRO_NAO_ENCONTRADO, Description: err.Error()})
	}

	return c.JSON(http.StatusOK, response.SuccessBody{Message: constants.CADASTRO_ALTERADO, Body: sexo})
}

func (controller *SexoController) Ativar(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.ID_NAO_INFORMADO, Description: err.Error()})
	}

	if err := controller.service.Ativar(id); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Message: constants.CADASTRO_FALHA_EXCLUSAO, Description: err.Error()})
	}

	return c.JSON(http.StatusOK, response.Success{Message: constants.CADASTRO_EXCLUIDO})
}
