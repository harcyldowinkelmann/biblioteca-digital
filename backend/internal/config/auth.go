package config

import (
	"ApiSup/internal/models"
	"ApiSup/pkg/hashing"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(usuario models.Usuario) (string, error) {
	claims := jwt.MapClaims{
		"id":   usuario.ID,
		"nome": usuario.Nome,
		"cpf":  usuario.CPF,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(hashing.Key))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
