package repositories

import (
	"biblioteca-digital/internal/config"
	"biblioteca-digital/internal/models/public"
	"biblioteca-digital/pkg/pagination"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AcaoRepository interface {
	BuscarAcaoPorId(id int) (*models.Acao, error)
	BuscarAcoes(c echo.Context) (*pagination.Pagination, error)
	Cadastrar(acao *models.Acao) error
	Atualizar(id int, acao *models.Acao) (*models.Acao, error)
	Ativar(id int) error
}

type acaoRepository struct {
	db *gorm.DB
}

func NovaAcaoRepository() AcaoRepository {
	return &acaoRepository{db: config.DB}
}

func (r *acaoRepository) BuscarAcaoPorId(id int) (*models.Acao, error) {
	var acao models.Acao
	if err := r.db.Where("id = ?", id).First(&acao).Error; err != nil {
		return nil, err
	}

	return &acao, nil
}

func (r *acaoRepository) BuscarAcoes(c echo.Context) (*pagination.Pagination, error) {
	var acao []models.Acao

	paginations, err := pagination.Paginate(c, r.db, &acao, nil, "Ações")

	if err != nil {
		return nil, err
	}

	return paginations, nil
}

func (r *acaoRepository) Cadastrar(acao *models.Acao) error {
	return r.db.Save(acao).Error
}

func (r *acaoRepository) Atualizar(id int, updatedAcao *models.Acao) (*models.Acao, error) {
	acao := new(models.Acao)
	if err := r.db.First(acao, id).Error; err != nil {
		return nil, err
	}
	if err := r.db.Model(acao).Updates(updatedAcao).Error; err != nil {
		return nil, err
	}
	return acao, nil
}

func (r *acaoRepository) Ativar(id int) error {
	acao := new(models.Acao)
	if err := r.db.First(acao, id).Error; err != nil {
		return err
	}

	acao.Ativo = !acao.Ativo

	if err := r.db.Model(acao).Update("ativo", acao.Ativo).Error; err != nil {
		return err
	}
	return nil
}
