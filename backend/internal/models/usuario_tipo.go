package models

type UsuarioTipo struct {
	BaseModel
	Nome string `json:"nome" validate:"required"`
}

func (UsuarioTipo) TableName() string {
	return "usuario_tipo"
}
