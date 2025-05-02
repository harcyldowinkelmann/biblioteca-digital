package public

import (
	"errors"
	"strings"

	"biblioteca-digital/internal/models"
	"biblioteca-digital/pkg/mapear/constants"

	"gorm.io/gorm"
)

type Permissao struct {
	TipoUsuarioId 	uint	   		`json:"tipo_usuario_id" validate:"required"`
	TipoUsuario		TipoUsuario		`gorm:"foreignKey:TipoUsuarioId;references:ID" json:"tipo_usuario,omitempty"`
	Descricao 		string   		`json:"descricao" validate:"required"`
	AcaoId    		uint     		`json:"acao_id" validate:"required" gorm:"not null"`
	Acao      		Acao     		`gorm:"foreignKey:AcaoId;references:ID" json:"acao,omitempty"`
	Ativo     		bool     		`json:"ativo" validate:"required"`
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
