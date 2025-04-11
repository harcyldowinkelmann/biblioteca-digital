package models

import (
	"errors"
	"strings"

	"biblioteca-digital/internal/models"
	"biblioteca-digital/pkg/mapear/constants"

	"gorm.io/gorm"
)

type TipoContato struct {
	Descricao	string 		`json:"descricao" validate:"required"`
	Ativo		bool		`json:"ativo" validade:"required"`
	models.BaseModel
}

func (tc *TipoContato) BeforeCreate(tx *gorm.DB) (err error) {
	if tc.Descricao == "" {
		return errors.New(constants.CAMPOS_OBRIGATORIOS)
	}

	tc.Descricao = strings.ToUpper(tc.Descricao)

	return nil
}

func (TipoContato) TableName() string {
	return "public.tipo_contato"
}
