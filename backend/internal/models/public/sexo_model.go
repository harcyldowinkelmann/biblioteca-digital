package models

import (
	"errors"
	"strings"

	"biblioteca-digital/internal/models"
	"biblioteca-digital/pkg/mapear/constants"

	"gorm.io/gorm"
)

type Sexo struct {
	Descricao	string 		`json:"descricao" validate:"required"`
	Ativo		bool		`json:"ativo" validade:"required"`
	models.BaseModel
}

func (s *Sexo) BeforeCreate(tx *gorm.DB) (err error) {
	if s.Descricao == "" {
		return errors.New(constants.CAMPOS_OBRIGATORIOS)
	}

	s.Descricao = strings.ToUpper(s.Descricao)

	return nil
}

func (Sexo) TableName() string {
	return "public.sexo"
}
