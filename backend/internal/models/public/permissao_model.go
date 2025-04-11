package models

import (
	"errors"
	"strings"

	"biblioteca-digital/internal/models"
	"biblioteca-digital/pkg/mapear/constants"

	"gorm.io/gorm"
)

type Permissao struct {
	Descricao string   `json:"descricao" validate:"required"`
	AcaoId    uint     `json:"acao_id" validate:"required" gorm:"not null"`
	Acao      Acao     `gorm:"foreignKey:AcaoId;references:ID" json:"acao,omitempty"`
	Ativo     bool     `json:"ativo" validate:"required"`
	models.BaseModel
}

func (p *Permissao) BeforeCreate(tx *gorm.DB) (err error) {
	if p.Descricao == "" {
		return errors.New(constants.CAMPOS_OBRIGATORIOS)
	}

	p.Descricao = strings.ToUpper(p.Descricao)

	return nil
}

func (Permissao) TableName() string {
	return "public.permissao"
}
