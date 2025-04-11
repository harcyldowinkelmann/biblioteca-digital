package repositories

import (
	"biblioteca-digital/internal/config"
	models "biblioteca-digital/internal/models/public"
	"biblioteca-digital/pkg/pagination"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type SexoRepository interface {
	BuscarSexoPorId(id int) (*models.Sexo, error)
	BuscarSexos(c echo.Context) (*pagination.Pagination, error)
	Cadastrar(sexo *models.Sexo) error
	Atualizar(id int, sexo *models.Sexo) (*models.Sexo, error)
	Ativar(id int) error
}

type sexoRepository struct {
	db *gorm.DB
}

func NovoSexoRepository() SexoRepository {
	return &sexoRepository{db: config.DB}
}

func (r *sexoRepository) BuscarSexoPorId(id int) (*models.Sexo, error) {
	var sexo models.Sexo
	if err := r.db.Where("id = ?", id).First(&sexo).Error; err != nil {
		return nil, err
	}

	return &sexo, nil
}

func (r *sexoRepository) BuscarSexos(c echo.Context) (*pagination.Pagination, error) {
	var sexo []models.Sexo

	paginations, err := pagination.Paginate(c, r.db, &sexo, nil, "Ações")

	if err != nil {
		return nil, err
	}

	return paginations, nil
}

func (r *sexoRepository) Cadastrar(sexo *models.Sexo) error {
	return r.db.Save(sexo).Error
}

func (r *sexoRepository) Atualizar(id int, updatedSexo *models.Sexo) (*models.Sexo, error) {
	sexo := new(models.Sexo)
	if err := r.db.First(sexo, id).Error; err != nil {
		return nil, err
	}
	if err := r.db.Model(sexo).Updates(updatedSexo).Error; err != nil {
		return nil, err
	}
	return sexo, nil
}

func (r *sexoRepository) Ativar(id int) error {
	sexo := new(models.Sexo)
	if err := r.db.First(sexo, id).Error; err != nil {
		return err
	}

	sexo.Ativo = !sexo.Ativo

	if err := r.db.Model(sexo).Update("ativo", sexo.Ativo).Error; err != nil {
		return err
	}
	return nil
}
