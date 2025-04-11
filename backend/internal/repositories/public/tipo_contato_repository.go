package repositories

import (
	"biblioteca-digital/internal/config"
	models "biblioteca-digital/internal/models/public"
	"biblioteca-digital/pkg/pagination"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TipoContatoRepository interface {
	BuscarTipoContatoPorId(id int) (*models.TipoContato, error)
	BuscarTiposContatos(c echo.Context) (*pagination.Pagination, error)
	Cadastrar(tipoContato *models.TipoContato) error
	Atualizar(id int, tipoContato *models.TipoContato) (*models.TipoContato, error)
	Ativar(id int) error
}

type tipoContatoRepository struct {
	db *gorm.DB
}

func NovoTipoContatoRepository() TipoContatoRepository {
	return &tipoContatoRepository{db: config.DB}
}

func (r *tipoContatoRepository) BuscarTipoContatoPorId(id int) (*models.TipoContato, error) {
	var tipoContato models.TipoContato
	if err := r.db.Where("id = ?", id).First(&tipoContato).Error; err != nil {
		return nil, err
	}

	return &tipoContato, nil
}

func (r *tipoContatoRepository) BuscarTiposContatos(c echo.Context) (*pagination.Pagination, error) {
	var tipoContato []models.TipoContato

	paginations, err := pagination.Paginate(c, r.db, &tipoContato, nil, "Ações")

	if err != nil {
		return nil, err
	}

	return paginations, nil
}

func (r *tipoContatoRepository) Cadastrar(tipoContato *models.TipoContato) error {
	return r.db.Save(tipoContato).Error
}

func (r *tipoContatoRepository) Atualizar(id int, updatedTipoContato *models.TipoContato) (*models.TipoContato, error) {
	tipoContato := new(models.TipoContato)
	if err := r.db.First(tipoContato, id).Error; err != nil {
		return nil, err
	}
	if err := r.db.Model(tipoContato).Updates(updatedTipoContato).Error; err != nil {
		return nil, err
	}
	return tipoContato, nil
}

func (r *tipoContatoRepository) Ativar(id int) error {
	tipoContato := new(models.TipoContato)
	if err := r.db.First(tipoContato, id).Error; err != nil {
		return err
	}

	tipoContato.Ativo = !tipoContato.Ativo

	if err := r.db.Model(tipoContato).Update("ativo", tipoContato.Ativo).Error; err != nil {
		return err
	}
	return nil
}
