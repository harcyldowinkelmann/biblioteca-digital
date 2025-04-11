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

type MunicipioController struct {
	service services.MunicipioService
}

func NovoMunicipioController(municipioService services.MunicipioService) *MunicipioController {
	return &MunicipioController{
		service: municipioService,
	}
}

func (controller *MunicipioController) BuscarMunicipioPorId(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.ID_NAO_INFORMADO, Description: err.Error()})
	}

	municipio, err := controller.service.BuscarMunicipioPorId(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Error{Message: constants.REGISTRO_NAO_ENCONTRADO, Description: err.Error()})
	}

	return c.JSON(http.StatusOK, municipio)
}

func (controller *MunicipioController) BuscarMunicipios(c echo.Context) error {
	municipio, err := controller.service.BuscarMunicipios(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Message: constants.ERRO_LISTAGEM_REGISTRO, Description: err.Error()})
	}

	return c.JSON(http.StatusOK, municipio)
}

func (controller *MunicipioController) Cadastrar(c echo.Context) error {
	var municipio models.Municipio
	if err := c.Bind(&municipio); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.BODY_FALHA, Description: err.Error()})
	}

	if err := c.Validate(municipio); err != nil {
		return config.ValidationErrors(c, err)
	}

	if err := controller.service.Cadastrar(&municipio); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Message: constants.CADASTRO_FALHA_INSERCAO, Description: err.Error()})
	}

	return c.JSON(http.StatusCreated, response.Success{Message: constants.CADASTRO_REALIZADO})
}

func (controller *MunicipioController) Atualizar(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.ID_NAO_INFORMADO, Description: err.Error()})
	}

	var municipioAtualizar models.Municipio
	if err := c.Bind(&municipioAtualizar); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.BODY_FALHA, Description: err.Error()})
	}

	if err := c.Validate(municipioAtualizar); err != nil {
		return config.ValidationErrors(c, err)
	}

	municipio, err := controller.service.Atualizar(id, &municipioAtualizar)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Error{Message: constants.REGISTRO_NAO_ENCONTRADO, Description: err.Error()})
	}

	return c.JSON(http.StatusOK, response.SuccessBody{Message: constants.CADASTRO_ALTERADO, Body: municipio})
}

func (controller *MunicipioController) Ativar(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: constants.ID_NAO_INFORMADO, Description: err.Error()})
	}

	if err := controller.service.Ativar(id); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Message: constants.CADASTRO_FALHA_EXCLUSAO, Description: err.Error()})
	}

	return c.JSON(http.StatusOK, response.Success{Message: constants.CADASTRO_EXCLUIDO})
}
