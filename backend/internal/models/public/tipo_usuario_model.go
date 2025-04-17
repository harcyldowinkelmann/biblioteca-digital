package public

import (
	"errors"
	"strings"

	"biblioteca-digital/internal/models"
	"biblioteca-digital/pkg/mapear/constants"

	"gorm.io/gorm"
)

type TipoUsuario struct {
	Descricao	string 		`json:"descricao" validate:"required"`
	Ativo		bool 		`json:"ativo" validate:"required"`
	models.BaseModel
}

func (tu *TipoUsuario) BeforeCreate(tx *gorm.DB) (err error) {
	if tu.Descricao == "" {
		return errors.New(constants.CAMPOS_OBRIGATORIOS)
	}

	tu.Descricao = strings.ToUpper(tu.Descricao)

	return nil
}

func (TipoUsuario) TableName() string {
	return "public.tipo_usuario"
}
