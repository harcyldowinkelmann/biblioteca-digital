package usuario

import (
	"errors"
	"strings"

	"biblioteca-digital/internal/models"
	"biblioteca-digital/pkg/mapear/constants"

	"gorm.io/gorm"
)

type Assunto struct {
	Descricao	string 		`json:"descricao" validate:"required"`
	Endpoint	string 		`json:"endpoint" validate:"required"`
	Ativo		bool		`json:"ativo" validade:"required"`
	models.BaseModel
}

func (a *Assunto) BeforeCreate(tx *gorm.DB) (err error) {
	if a.Descricao == "" {
		return errors.New(constants.CAMPOS_OBRIGATORIOS)
	}

	a.Descricao = strings.ToUpper(a.Descricao)

	return nil
}

func (Assunto) TableName() string {
	return "public.assunto"
}
