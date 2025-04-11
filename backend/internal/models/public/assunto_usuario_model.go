package models

import (
	"biblioteca-digital/internal/models"
)

type AssuntoUsuario struct {
	models.BaseModel
	UsuarioId uint `json:"usuario_id" validate:"required" gorm:"not null;uniqueIndex:idx_usuario_assunto"`
	AssuntoId uint `json:"assunto_id" validate:"required" gorm:"not null;uniqueIndex:idx_usuario_assunto"`
	Ativo     bool `json:"ativo" validate:"required"`
}

func (AssuntoUsuario) TableName() string {
	return "public.assunto_usuario"
}
