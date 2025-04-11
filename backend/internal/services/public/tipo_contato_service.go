package service

import (
	"biblioteca-digital/internal/models/public"
	"biblioteca-digital/internal/repositories/public"
	"biblioteca-digital/pkg/pagination"

	"github.com/labstack/echo/v4"
)

type TipoContatoService interface {
	BuscarTipoContatoPorId(id int) (*models.TipoContato, error)
	BuscarTiposContatos(c echo.Context) (*pagination.Pagination, error)
	Cadastrar(tipoContato *models.TipoContato) error
	Atualizar(id int, tipoContato *models.TipoContato) (*models.TipoContato, error)
	Ativar(id int) (error)
}

type tipoContatoService struct {
	repository repositories.TipoContatoRepository
}

func NovoTipoContatoService(tipoContatoRepo repositories.TipoContatoRepository) TipoContatoService {
	return &tipoContatoService{repository: tipoContatoRepo}
}

func (s *tipoContatoService) BuscarTipoContatoPorId(id int) (*models.TipoContato, error) {
	return s.repository.BuscarTipoContatoPorId(id)
}

func (s *tipoContatoService) BuscarTiposContatos(c echo.Context) (*pagination.Pagination, error) {
	return s.repository.BuscarTiposContatos(c)
}

func (s *tipoContatoService) Cadastrar(tipoContato *models.TipoContato) error {
	return s.repository.Cadastrar(tipoContato)
}

func (s *tipoContatoService) Atualizar(id int, updatedTipoContato *models.TipoContato) (*models.TipoContato, error) {
	return s.repository.Atualizar(id, updatedTipoContato)
}

func (s *tipoContatoService) Ativar(id int) error {
	return s.repository.Ativar(id)
}
