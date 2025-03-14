package config

import (
	"ApiSup/pkg/mapear/constants"
	"ApiSup/pkg/mapear/response"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func ValidationErrors(c echo.Context, err error) error {
	validationErrors := err.(validator.ValidationErrors)

	errorMessages := make(map[string]string)

	for _, err := range validationErrors {
		fieldName := err.Field()

		if err.Tag() == "required" {
			errorMessages[fieldName] = constants.CAMPOS_OBRIGATORIOS

			continue
		}
		errorMessages[fieldName] = constants.CAMPOS_INVALIDOS
	}

	errorsResponse := response.ErrorResponseField{
		Message:     constants.PREENCHAS_TODOS_CAMPOS,
		Description: errorMessages,
	}

	return c.JSON(http.StatusBadRequest, errorsResponse)
}
