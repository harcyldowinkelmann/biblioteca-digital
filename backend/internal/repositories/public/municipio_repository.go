package repositories

import (
	"biblioteca-digital/internal/config"
	models "biblioteca-digital/internal/models/public"
	"biblioteca-digital/pkg/pagination"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type MunicipioRepository interface {
	BuscarMunicipioPorId(id int) (*models.Municipio, error)
	BuscarMunicipios(c echo.Context) (*pagination.Pagination, error)
	Cadastrar(municipio *models.Municipio) error
	Atualizar(id int, municipio *models.Municipio) (*models.Municipio, error)
	Ativar(id int) error
}

type municipioRepository struct {
	db *gorm.DB
}

func NovoMunicipioRepository() MunicipioRepository {
	return &municipioRepository{db: config.DB}
}

func (r *municipioRepository) BuscarMunicipioPorId(id int) (*models.Municipio, error) {
	var municipio models.Municipio
	if err := r.db.Where("id = ?", id).First(&municipio).Error; err != nil {
		return nil, err
	}

	return &municipio, nil
}

func (r *municipioRepository) BuscarMunicipios(c echo.Context) (*pagination.Pagination, error) {
	var municipio []models.Municipio

	paginations, err := pagination.Paginate(c, r.db, &municipio, nil, "Municípios")

	if err != nil {
		return nil, err
	}

	return paginations, nil
}

func (r *municipioRepository) Cadastrar(municipio *models.Municipio) error {
	return r.db.Save(municipio).Error
}

func (r *municipioRepository) Atualizar(id int, updatedMunicipio *models.Municipio) (*models.Municipio, error) {
	municipio := new(models.Municipio)
	if err := r.db.First(municipio, id).Error; err != nil {
		return nil, err
	}
	if err := r.db.Model(municipio).Updates(updatedMunicipio).Error; err != nil {
		return nil, err
	}
	return municipio, nil
}

func (r *municipioRepository) Ativar(id int) error {
	municipio := new(models.Municipio)
	if err := r.db.First(municipio, id).Error; err != nil {
		return err
	}

	municipio.Ativo = !municipio.Ativo

	if err := r.db.Model(municipio).Update("ativo", municipio.Ativo).Error; err != nil {
		return err
	}
	return nil
}
