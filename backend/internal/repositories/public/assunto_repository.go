package repositories

import (
	"biblioteca-digital/internal/config"
	models "biblioteca-digital/internal/models/public"
	"biblioteca-digital/pkg/pagination"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AssuntoRepository interface {
	BuscarAssuntoPorId(id int) (*models.Assunto, error)
	BuscarAssuntos(c echo.Context) (*pagination.Pagination, error)
	Cadastrar(assunto *models.Assunto) error
	Atualizar(id int, assunto *models.Assunto) (*models.Assunto, error)
	Ativar(id int) error
}

type assuntoRepository struct {
	db *gorm.DB
}

func NovoAssuntoRepository() AssuntoRepository {
	return &assuntoRepository{db: config.DB}
}

func (r *assuntoRepository) BuscarAssuntoPorId(id int) (*models.Assunto, error) {
	var assunto models.Assunto
	if err := r.db.Where("id = ?", id).First(&assunto).Error; err != nil {
		return nil, err
	}

	return &assunto, nil
}

func (r *assuntoRepository) BuscarAssuntos(c echo.Context) (*pagination.Pagination, error) {
	var assunto []models.Assunto

	paginations, err := pagination.Paginate(c, r.db, &assunto, nil, "Ações")

	if err != nil {
		return nil, err
	}

	return paginations, nil
}

func (r *assuntoRepository) Cadastrar(assunto *models.Assunto) error {
	return r.db.Save(assunto).Error
}

func (r *assuntoRepository) Atualizar(id int, updatedAssunto *models.Assunto) (*models.Assunto, error) {
	assunto := new(models.Assunto)
	if err := r.db.First(assunto, id).Error; err != nil {
		return nil, err
	}
	if err := r.db.Model(assunto).Updates(updatedAssunto).Error; err != nil {
		return nil, err
	}
	return assunto, nil
}

func (r *assuntoRepository) Ativar(id int) error {
	assunto := new(models.Assunto)
	if err := r.db.First(assunto, id).Error; err != nil {
		return err
	}

	assunto.Ativo = !assunto.Ativo

	if err := r.db.Model(assunto).Update("ativo", assunto.Ativo).Error; err != nil {
		return err
	}
	return nil
}
