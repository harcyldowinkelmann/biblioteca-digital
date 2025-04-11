package service

import (
	"biblioteca-digital/internal/models/public"
	"biblioteca-digital/internal/repositories/public"
	"biblioteca-digital/pkg/pagination"

	"github.com/labstack/echo/v4"
)

type PermissaoService interface {
	BuscarPermissaoPorId(id int) (*models.Permissao, error)
	BuscarPermissoes(c echo.Context) (*pagination.Pagination, error)
	Cadastrar(permissao *models.Permissao) error
	Atualizar(id int, permissao *models.Permissao) (*models.Permissao, error)
	Ativar(id int) (error)
}

type permissaoService struct {
	repository repositories.PermissaoRepository
}

func NovoPermissaoService(permissaoRepo repositories.PermissaoRepository) PermissaoService {
	return &permissaoService{repository: permissaoRepo}
}

func (s *permissaoService) BuscarPermissaoPorId(id int) (*models.Permissao, error) {
	return s.repository.BuscarPermissaoPorId(id)
}

func (s *permissaoService) BuscarPermissoes(c echo.Context) (*pagination.Pagination, error) {
	return s.repository.BuscarPermissoes(c)
}

func (s *permissaoService) Cadastrar(permissao *models.Permissao) error {
	return s.repository.Cadastrar(permissao)
}

func (s *permissaoService) Atualizar(id int, updatedPermissao *models.Permissao) (*models.Permissao, error) {
	return s.repository.Atualizar(id, updatedPermissao)
}

func (s *permissaoService) Ativar(id int) error {
	return s.repository.Ativar(id)
}
