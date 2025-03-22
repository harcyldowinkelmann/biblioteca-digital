package repositories

import (
	"biblioteca-digital/internal/config"
	"biblioteca-digital/internal/models/public"
	"biblioteca-digital/pkg/pagination"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TipoUsuarioRepository interface {
	BuscarTipoUsuarioPorId(id int) (*models.TipoUsuario, error)
	BuscarUsuarios(c echo.Context) (*pagination.Pagination, error)
	Cadastrar(tipoUsuario *models.TipoUsuario) error
	Atualizar(id int, tipoUsuario *models.TipoUsuario) (*models.TipoUsuario, error)
	Ativar(id int) error
}

type tipoUsuarioRepository struct {
	db *gorm.DB
}

func NovoTipoUsuarioRepository() TipoUsuarioRepository {
	return &tipoUsuarioRepository{db: config.DB}
}

func (r *tipoUsuarioRepository) BuscarTipoUsuarioPorId(id int) (*models.TipoUsuario, error) {
	var tipoUsuario models.TipoUsuario
	if err := r.db.Where("id = ?", id).First(&tipoUsuario).Error; err != nil {
		return nil, err
	}

	return &tipoUsuario, nil
}

func (r *tipoUsuarioRepository) BuscarUsuarios(c echo.Context) (*pagination.Pagination, error) {
	var tipoUsuario []models.TipoUsuario

	paginations, err := pagination.Paginate(c, r.db, &tipoUsuario, nil, "Usuarios")

	if err != nil {
		return nil, err
	}

	return paginations, nil
}

func (r *tipoUsuarioRepository) Cadastrar(tipoUsuario *models.TipoUsuario) error {
	return r.db.Save(tipoUsuario).Error
}

func (r *tipoUsuarioRepository) Atualizar(id int, updatedTipoUsuario *models.TipoUsuario) (*models.TipoUsuario, error) {
	tipoUsuario := new(models.TipoUsuario)
	if err := r.db.First(tipoUsuario, id).Error; err != nil {
		return nil, err
	}
	if err := r.db.Model(tipoUsuario).Updates(updatedTipoUsuario).Error; err != nil {
		return nil, err
	}
	return tipoUsuario, nil
}

func (r *tipoUsuarioRepository) Ativar(id int) error {
	tipoUsuario := new(models.TipoUsuario)
	if err := r.db.First(tipoUsuario, id).Error; err != nil {
		return err
	}

	tipoUsuario.Ativo = !tipoUsuario.Ativo

	if err := r.db.Model(tipoUsuario).Update("ativo", tipoUsuario.Ativo).Error; err != nil {
		return err
	}
	return nil
}
