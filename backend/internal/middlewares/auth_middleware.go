package middlewares

import (
	"biblioteca-digital/pkg/hashing"
	"biblioteca-digital/pkg/mapear/constants"
	"biblioteca-digital/pkg/mapear/responses"
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func VerifyTokenHandler() echo.MiddlewareFunc {
	jwtMiddleware := echojwt.JWT([]byte(hashing.Key))

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := jwtMiddleware(func(c echo.Context) error {
				return nil
			})(c); err != nil {
				return c.JSON(http.StatusUnauthorized, response.Error{
					Message:     constants.ACESSO_NAO_AUTORIZADO,
					Description: err,
				})
			}
			return next(c)
		}
	}
}
