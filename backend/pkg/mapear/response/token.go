package response

import "biblioteca-digital/internal/models"

type Token struct {
	Token   string         `json:"token"`
	Usuario models.Usuario `json:"usuario"`
}
