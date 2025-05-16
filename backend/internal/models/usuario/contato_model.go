package usuario

import (
	"biblioteca-digital/internal/models"
	contato "biblioteca-digital/internal/models/public"
	"biblioteca-digital/pkg/mapear/constants"
	"strings"
	"errors"

	"gorm.io/gorm"
)

type Contato struct {
	models.BaseModel
	UsuarioId		uint	   				`json:"usuario_id" validate:"required"`
	Usuario			Usuario					`gorm:"foreignKey:UsuarioId;references:ID" json:"usuario,omitempty"`
	TipoContatoId	uint					`json:"tipo_contato_id" validate:"required"`
	TipoContato		contato.TipoContato		`gorm:"foreignKey:TipoContatoId;references:ID" json:"contato,omitempty"`
	Descricao 		string   				`json:"descricao" validate:"required"`
	Ativo     		bool 					`json:"ativo" validate:"required"`
}

func (c *Contato) BeforeCreate(tx *gorm.DB) (err error) {
	if c.Descricao == "" {
		return errors.New(constants.CAMPOS_OBRIGATORIOS)
	}

	c.Descricao = strings.ToUpper(c.Descricao)

	return nil
}

func (Contato) TableName() string {
	return "usuario.contato"
}
