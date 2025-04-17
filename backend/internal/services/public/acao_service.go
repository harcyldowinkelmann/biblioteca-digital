package service

import (
	models "biblioteca-digital/internal/models/public"
	"biblioteca-digital/internal/repositories/public"
	"biblioteca-digital/pkg/pagination"

	"github.com/labstack/echo/v4"
)

type AcaoService interface {
	BuscarAcaoPorId(id int) (*models.Acao, error)
	BuscarAcoes(c echo.Context) (*pagination.Pagination, error)
	Cadastrar(acao *models.Acao) error
	Atualizar(id int, acao *models.Acao) (*models.Acao, error)
	Ativar(id int) (error)
}

type acaoService struct {
	repository repositories.AcaoRepository
}

func NovaAcaoService(acaoRepo repositories.AcaoRepository) AcaoService {
	return &acaoService{repository: acaoRepo}
}

func (s *acaoService) BuscarAcaoPorId(id int) (*models.Acao, error) {
	return s.repository.BuscarAcaoPorId(id)
}

func (s *acaoService) BuscarAcoes(c echo.Context) (*pagination.Pagination, error) {
	return s.repository.BuscarAcoes(c)
}

func (s *acaoService) Cadastrar(acao *models.Acao) error {
	return s.repository.Cadastrar(acao)
}

func (s *acaoService) Atualizar(id int, updatedAcao *models.Acao) (*models.Acao, error) {
	return s.repository.Atualizar(id, updatedAcao)
}

func (s *acaoService) Ativar(id int) error {
	return s.repository.Ativar(id)
}
