package service

import (
	"biblioteca-digital/internal/models/public"
	"biblioteca-digital/internal/repositories/public"
	"biblioteca-digital/pkg/pagination"

	"github.com/labstack/echo/v4"
)

type SexoService interface {
	BuscarSexoPorId(id int) (*models.Sexo, error)
	BuscarSexos(c echo.Context) (*pagination.Pagination, error)
	Cadastrar(sexo *models.Sexo) error
	Atualizar(id int, sexo *models.Sexo) (*models.Sexo, error)
	Ativar(id int) (error)
}

type sexoService struct {
	repository repositories.SexoRepository
}

func NovoSexoService(sexoRepo repositories.SexoRepository) SexoService {
	return &sexoService{repository: sexoRepo}
}

func (s *sexoService) BuscarSexoPorId(id int) (*models.Sexo, error) {
	return s.repository.BuscarSexoPorId(id)
}

func (s *sexoService) BuscarSexos(c echo.Context) (*pagination.Pagination, error) {
	return s.repository.BuscarSexos(c)
}

func (s *sexoService) Cadastrar(sexo *models.Sexo) error {
	return s.repository.Cadastrar(sexo)
}

func (s *sexoService) Atualizar(id int, updatedSexo *models.Sexo) (*models.Sexo, error) {
	return s.repository.Atualizar(id, updatedSexo)
}

func (s *sexoService) Ativar(id int) error {
	return s.repository.Ativar(id)
}
