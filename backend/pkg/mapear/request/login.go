package request

type Login struct {
	CPF   string `json:"cpf" validate:"required"`
	Senha string `json:"senha" validate:"required"`
}
