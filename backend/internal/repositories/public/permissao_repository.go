package repositories

import (
	"biblioteca-digital/internal/config"
	models "biblioteca-digital/internal/models/public"
	"biblioteca-digital/pkg/pagination"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PermissaoRepository interface {
	BuscarPermissaoPorId(id int) (*models.Permissao, error)
	BuscarPermissoes(c echo.Context) (*pagination.Pagination, error)
	Cadastrar(permissao *models.Permissao) error
	Atualizar(id int, permissao *models.Permissao) (*models.Permissao, error)
	Ativar(id int) error
}

type permissaoRepository struct {
	db *gorm.DB
}

func NovoPermissaoRepository() PermissaoRepository {
	return &permissaoRepository{db: config.DB}
}

func (r *permissaoRepository) BuscarPermissaoPorId(id int) (*models.Permissao, error) {
	var permissao models.Permissao
	if err := r.db.Where("id = ?", id).First(&permissao).Error; err != nil {
		return nil, err
	}

	return &permissao, nil
}

func (r *permissaoRepository) BuscarPermissoes(c echo.Context) (*pagination.Pagination, error) {
	var permissao []models.Permissao

	paginations, err := pagination.Paginate(c, r.db, &permissao, nil, "Permissões")

	if err != nil {
		return nil, err
	}

	return paginations, nil
}

func (r *permissaoRepository) Cadastrar(permissao *models.Permissao) error {
	return r.db.Save(permissao).Error
}

func (r *permissaoRepository) Atualizar(id int, updatedPermissao *models.Permissao) (*models.Permissao, error) {
	permissao := new(models.Permissao)
	if err := r.db.First(permissao, id).Error; err != nil {
		return nil, err
	}
	if err := r.db.Model(permissao).Updates(updatedPermissao).Error; err != nil {
		return nil, err
	}
	return permissao, nil
}

func (r *permissaoRepository) Ativar(id int) error {
	permissao := new(models.Permissao)
	if err := r.db.First(permissao, id).Error; err != nil {
		return err
	}

	permissao.Ativo = !permissao.Ativo

	if err := r.db.Model(permissao).Update("ativo", permissao.Ativo).Error; err != nil {
		return err
	}
	return nil
}
