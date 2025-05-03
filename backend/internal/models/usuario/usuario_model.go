package usuario

import (
	"errors"
	"strings"

	"biblioteca-digital/internal/models"
	usuario "biblioteca-digital/internal/models/public"
	"biblioteca-digital/pkg/mapear/constants"

	"gorm.io/gorm"
	"time"
	"github.com/lib/pq"
)

type Usuario struct {
	TipoUsuarioId		uint	   				`json:"tipo_usuario_id" validate:"required"`
	TipoUsuario			usuario.TipoUsuario		`gorm:"foreignKey:TipoUsuarioId;references:ID" json:"tipo_usuario,omitempty"`
	ContatoId			uint 					`json:"contato_id" validate:"required"`
	Contato				string					`gorm:"foreignKey:ContatoId;references:ID" json:"contato,omitempty"`
	Nome				string 					`json:"nome" validate:"required"`
	Cpf					string					`json:"cpf" validate:"required"`
	DataNascimento		time.Time				`json:"data_nascimento" validate:"required"`
	Nacionalidade		string 					`json:"nacionalidade" validate:"required"`
	Senha				string 					`json:"senha" validate:"required"`
	Instituicao			string 					`json:"instituicao" validate:"required"`
	Interesses			pq.StringArray			`gorm:"type:varchar[]" json:"interesses" validate:"required,dive,required"`
	NivelEscolaridade	string 					`json:"nivel_escolaridade" validate:"required"`
	Ativo				bool 					`json:"ativo" validate:"required"`
	models.BaseModel
}

func (u *Usuario) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Nome == "" || u.Nacionalidade == "" || u.Senha == "" || u.Instituicao == "" || u.NivelEscolaridade == "" {
		return errors.New(constants.CAMPOS_OBRIGATORIOS)
	}

	u.Nome = strings.ToUpper(u.Nome)
	u.Nacionalidade = strings.ToUpper(u.Nacionalidade)
	u.Instituicao = strings.ToUpper(u.Instituicao)
	u.NivelEscolaridade = strings.ToUpper(u.NivelEscolaridade)

	return nil
}

func (Usuario) TableName() string {
	return "usuario.usuario"
}
