package models

import (
	"ApiSup/pkg/mapear/constants"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Usuario struct {
	BaseModel
	Nome          string `json:"nome" validate:"required"`
	CPF           string `json:"cpf" validate:"required"`
	Email         string `json:"email" validate:"required"`
	Senha         string `json:"senha" validate:"required"`
	UsuarioTipoID uint   `json:"usuario_tipo_id"  validate:"required"`
}

type UsuarioView struct {
	ID            uint        `json:"id"`
	Nome          string      `json:"nome"`
	CPF           string      `json:"cpf"`
	UsuarioTipoID uint        `json:"usuario_tipo_id"  validate:"required"`
	UsuarioTipo   UsuarioTipo `json:"tipo" gorm:"foreignKey:UsuarioTipoID" validate:"-"`
}

func UniqueCPF(cpf string, tx *gorm.DB, model interface{}) (err error) {
	var count int64

	if err := tx.Model(&model).Where("cpf = ?", cpf).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New(constants.CPF_INVALIDO)
	}

	return nil
}

func Criptografia(senha string) (password string, err error) {
	if senha != "" {
		hashedSenha, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
		if err != nil {
			return senha, err
		}
		senha = string(hashedSenha)
	}

	return senha, nil
}

func (u *Usuario) BeforeCreate(tx *gorm.DB) (err error) {
	err = UniqueCPF(u.CPF, tx, u)
	if err != nil {
		return err
	}

	senha, err := Criptografia(u.Senha)
	if err != nil {
		return err
	}

	u.Senha = senha

	return nil
}

func (Usuario) TableName() string {
	return "usuario"
}

func (UsuarioView) TableName() string {
	return "usuario"
}
