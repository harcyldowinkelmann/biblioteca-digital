package models

import (
	"errors"
	"strings"

	"biblioteca-digital/internal/models"
	"biblioteca-digital/pkg/mapear/constants"

	"gorm.io/gorm"
)

type Acao struct {
	Descricao  string       `json:"descricao" validate:"required"`
	Endpoint   string       `json:"endpoint" validate:"required"`
	Ativo      bool         `json:"ativo" validate:"required"`
	Permissoes []Permissao  `gorm:"foreignKey:AcaoId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"permissoes,omitempty"`
	models.BaseModel
}

func (a *Acao) BeforeCreate(tx *gorm.DB) (err error) {
	if a.Descricao == "" {
		return errors.New(constants.CAMPOS_OBRIGATORIOS)
	}

	a.Descricao = strings.ToUpper(a.Descricao)
	return nil
}

func (Acao) TableName() string {
	return "public.acao"
}
