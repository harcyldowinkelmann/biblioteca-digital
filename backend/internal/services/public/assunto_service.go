package service

import (
	"biblioteca-digital/internal/models/public"
	"biblioteca-digital/internal/repositories/public"
	"biblioteca-digital/pkg/pagination"

	"github.com/labstack/echo/v4"
)

type AssuntoService interface {
	BuscarAssuntoPorId(id int) (*models.Assunto, error)
	BuscarAssuntos(c echo.Context) (*pagination.Pagination, error)
	Cadastrar(assunto *models.Assunto) error
	Atualizar(id int, assunto *models.Assunto) (*models.Assunto, error)
	Ativar(id int) (error)
}

type assuntoService struct {
	repository repositories.AssuntoRepository
}

func NovoAssuntoService(assuntoRepo repositories.AssuntoRepository) AssuntoService {
	return &assuntoService{repository: assuntoRepo}
}

func (s *assuntoService) BuscarAssuntoPorId(id int) (*models.Assunto, error) {
	return s.repository.BuscarAssuntoPorId(id)
}

func (s *assuntoService) BuscarAssuntos(c echo.Context) (*pagination.Pagination, error) {
	return s.repository.BuscarAssuntos(c)
}

func (s *assuntoService) Cadastrar(assunto *models.Assunto) error {
	return s.repository.Cadastrar(assunto)
}

func (s *assuntoService) Atualizar(id int, updatedAssunto *models.Assunto) (*models.Assunto, error) {
	return s.repository.Atualizar(id, updatedAssunto)
}

func (s *assuntoService) Ativar(id int) error {
	return s.repository.Ativar(id)
}
