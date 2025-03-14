package repositories

import (
	"ApiSup/internal/config"
	"ApiSup/internal/models"
	"ApiSup/pkg/pagination"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByCPF(cpf string) (*models.Usuario, error)
	Detalhar(id int) (*models.Usuario, error)
	Listagem(c echo.Context) (*pagination.Pagination, error)
	Novo(user *models.Usuario) error
	Editar(id int, updatedUser *models.Usuario) (*models.Usuario, error)
	Deletar(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{db: config.DB}
}

func (r *userRepository) GetUserByCPF(cpf string) (*models.Usuario, error) {
	var user models.Usuario
	if err := r.db.Where("cpf = ?", cpf).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Detalhar(id int) (*models.Usuario, error) {
	user := new(models.Usuario)
	if err := r.db.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Listagem(c echo.Context) (*pagination.Pagination, error) {
	var users []models.UsuarioView

	paginations, err := pagination.Paginate(c, r.db, &users, nil, "UsuarioTipo")

	if err != nil {
		return nil, err
	}

	return paginations, nil
}

func (r *userRepository) Novo(user *models.Usuario) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Editar(id int, updatedUser *models.Usuario) (*models.Usuario, error) {
	user := new(models.Usuario)
	if err := r.db.First(user, id).Error; err != nil {
		return nil, err
	}
	if err := r.db.Model(user).Updates(updatedUser).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Deletar(id int) error {
	return r.db.Delete(&models.Usuario{}, id).Error
}
