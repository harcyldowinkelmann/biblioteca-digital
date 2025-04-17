package service

import (
	models "biblioteca-digital/internal/models/public"
	"biblioteca-digital/internal/repositories/public"
	"biblioteca-digital/pkg/pagination"

	"github.com/labstack/echo/v4"
)

type MunicipioService interface {
	BuscarMunicipioPorId(id int) (*models.Municipio, error)
	BuscarMunicipios(c echo.Context) (*pagination.Pagination, error)
	Cadastrar(municipio *models.Municipio) error
	Atualizar(id int, municipio *models.Municipio) (*models.Municipio, error)
	Ativar(id int) (error)
}

type municipioService struct {
	repository repositories.MunicipioRepository
}

func NovoMunicipioService(municipioRepo repositories.MunicipioRepository) MunicipioService {
	return &municipioService{repository: municipioRepo}
}

func (s *municipioService) BuscarMunicipioPorId(id int) (*models.Municipio, error) {
	return s.repository.BuscarMunicipioPorId(id)
}

func (s *municipioService) BuscarMunicipios(c echo.Context) (*pagination.Pagination, error) {
	return s.repository.BuscarMunicipios(c)
}

func (s *municipioService) Cadastrar(municipio *models.Municipio) error {
	return s.repository.Cadastrar(municipio)
}

func (s *municipioService) Atualizar(id int, updatedMunicipio *models.Municipio) (*models.Municipio, error) {
	return s.repository.Atualizar(id, updatedMunicipio)
}

func (s *municipioService) Ativar(id int) error {
	return s.repository.Ativar(id)
}
