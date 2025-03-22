package routes

import (
	controller "biblioteca-digital/internal/controllers"
	controllerPublic "biblioteca-digital/internal/controllers/public"
	"biblioteca-digital/internal/middlewares"
	"biblioteca-digital/internal/repositories"
	repo "biblioteca-digital/internal/repositories/public"
	"biblioteca-digital/internal/services"
	servicesPublic "biblioteca-digital/internal/services/public"

	"github.com/labstack/echo/v4"
)

func InitializeRoutes(e *echo.Echo) {
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	auth := e.Group("/auth")
	auth.POST("/login", userController.Login)

	protected := e.Group("/api")
	protected.Use(middlewares.VerifyTokenHandler())

	protected.POST("/usuarios", userController.Created)
	protected.GET("/usuarios", userController.Listagem)
	protected.GET("/usuarios/:id", userController.Get)
	protected.PUT("/usuarios/:id", userController.Updated)
	protected.DELETE("/usuarios/:id", userController.Deleted)

	/**
	 * Rotas para Gerenciamento dos tipos de Acoes cadastradas no sistema
	 */
	 acaoRepo := repo.NovaAcaoRepository()
	 acaoService := servicesPublic.NovaAcaoService(acaoRepo)
	 acaoController := controllerPublic.NovaAcaoController(acaoService)
 
	 acao := e.Group("/acao")
	 acao.Use(middlewares.VerifyTokenHandler())
 
	 acao.POST("/", acaoController.Cadastrar)
	 acao.GET("/", acaoController.BuscarAcoes)
	 acao.GET("/acao/:id", acaoController.BuscarAcaoPorId)
	 acao.PUT("/acao/:id", acaoController.Atualizar)
	 acao.PATCH("/acao/:id", acaoController.Ativar)
	 /// ---------------------------------------------------

	 /**
	 * Rotas para Gerenciamento dos tipos de Usuários cadastrados no sistema
	 */
	 tipoUsuarioRepo := repo.NovoTipoUsuarioRepository()
	 tipoUsuarioService := servicesPublic.NovoTipoUsuarioService(tipoUsuarioRepo)
	 tipoUsuarioController := controllerPublic.NovoTipoUsuarioController(tipoUsuarioService)
 
	 tipoUsuario := e.Group("/tipo-usuario")
	 tipoUsuario.Use(middlewares.VerifyTokenHandler())
 
	 tipoUsuario.POST("/", tipoUsuarioController.Cadastrar)
	 tipoUsuario.GET("/", tipoUsuarioController.BuscarUsuarios)
	 tipoUsuario.GET("/tipo-usuario/:id", tipoUsuarioController.BuscarTipoUsuarioPorId)
	 tipoUsuario.PUT("/tipo-usuario/:id", tipoUsuarioController.Atualizar)
	 tipoUsuario.PATCH("/tipo-usuario/:id", tipoUsuarioController.Ativar)
	 /// ---------------------------------------------------
}
