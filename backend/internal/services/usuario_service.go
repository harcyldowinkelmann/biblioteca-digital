package services

import (
	"ApiSup/internal/models"
	"ApiSup/internal/repositories"
	"ApiSup/pkg/mapear/constants"
	"ApiSup/pkg/mapear/request"
	"ApiSup/pkg/pagination"
	"errors"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Authenticate(body request.Login) (*models.Usuario, error)
	Detalhar(id int) (*models.Usuario, error)
	Listagem(c echo.Context) (*pagination.Pagination, error)
	Novo(user *models.Usuario) error
	Editar(id int, updatedUser *models.Usuario) (*models.Usuario, error)
	Deletar(id int) error
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{repository: userRepo}
}

func (s *userService) Authenticate(body request.Login) (*models.Usuario, error) {
	user, err := s.repository.GetUserByCPF(body.CPF)
	if err != nil {
		return nil, errors.New(constants.USUARIO_ENCONTRADO)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Senha), []byte(body.Senha))
	if err != nil {
		return nil, errors.New(constants.SENHA_INVALIDA)
	}

	return user, nil
}

func (s *userService) Detalhar(id int) (*models.Usuario, error) {
	return s.repository.Detalhar(id)
}

func (s *userService) Listagem(c echo.Context) (*pagination.Pagination, error) {
	return s.repository.Listagem(c)
}

func (s *userService) Novo(user *models.Usuario) error {
	return s.repository.Novo(user)
}

func (s *userService) Editar(id int, updated *models.Usuario) (*models.Usuario, error) {
	if updated.Senha != "" {
		hashedSenha, err := models.Criptografia(updated.Senha)
		if err != nil {
			return nil, err
		}

		updated.Senha = hashedSenha
	}

	return s.repository.Editar(id, updated)
}

func (s *userService) Deletar(id int) error {
	return s.repository.Deletar(id)
}
