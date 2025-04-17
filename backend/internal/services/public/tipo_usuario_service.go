package service

import (
	models "biblioteca-digital/internal/models/public"
	"biblioteca-digital/internal/repositories/public"
	"biblioteca-digital/pkg/pagination"

	"github.com/labstack/echo/v4"
)

type TipoUsuarioService interface {
	BuscarTipoUsuarioPorId(id int) (*models.TipoUsuario, error)
	BuscarUsuarios(c echo.Context) (*pagination.Pagination, error)
	Cadastrar(tipoUsuario *models.TipoUsuario) error
	Atualizar(id int, tipoUsuario *models.TipoUsuario) (*models.TipoUsuario, error)
	Ativar(id int) (error)
}

type tipoUsuarioService struct {
	repository repositories.TipoUsuarioRepository
}

func NovoTipoUsuarioService(tipoUsuarioRepo repositories.TipoUsuarioRepository) TipoUsuarioService {
	return &tipoUsuarioService{repository: tipoUsuarioRepo}
}

func (s *tipoUsuarioService) BuscarTipoUsuarioPorId(id int) (*models.TipoUsuario, error) {
	return s.repository.BuscarTipoUsuarioPorId(id)
}

func (s *tipoUsuarioService) BuscarUsuarios(c echo.Context) (*pagination.Pagination, error) {
	return s.repository.BuscarUsuarios(c)
}

func (s *tipoUsuarioService) Cadastrar(tipoUsuario *models.TipoUsuario) error {
	return s.repository.Cadastrar(tipoUsuario)
}

func (s *tipoUsuarioService) Atualizar(id int, updatedTipoUsuario *models.TipoUsuario) (*models.TipoUsuario, error) {
	return s.repository.Atualizar(id, updatedTipoUsuario)
}

func (s *tipoUsuarioService) Ativar(id int) error {
	return s.repository.Ativar(id)
}
