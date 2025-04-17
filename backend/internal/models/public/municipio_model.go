package public

import (
	"errors"
	"strings"

	"biblioteca-digital/internal/models"
	"biblioteca-digital/pkg/mapear/constants"

	"gorm.io/gorm"
)

type Municipio struct {
	Descricao	string 		`json:"descricao" validate:"required"`
	Ativo		bool		`json:"ativo" validade:"required"`
	models.BaseModel
}

func (m *Municipio) BeforeCreate(tx *gorm.DB) (err error) {
	if m.Descricao == "" {
		return errors.New(constants.CAMPOS_OBRIGATORIOS)
	}

	m.Descricao = strings.ToUpper(m.Descricao)

	return nil
}

func (Municipio) TableName() string {
	return "public.municipio"
}
