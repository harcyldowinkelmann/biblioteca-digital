# API Golang com Echo Framework (PostgreSQL) - Biblioteca Digital

Este guia detalha como configurar o ambiente para desenvolvimento da API do projeto Biblioteca Digital utilizando [Go](https://golang.org/) e [Air](https://github.com/cosmtrek/air), uma ferramenta de live reload para aplicações Go.

## Pré-requisitos

- Sistema operacional: Windows, macOS ou Linux
- Terminal ou prompt de comando com permissões administrativas

## Instalação do Go

1. Acesse a página oficial de downloads do Go: [https://golang.org/dl/](https://golang.org/dl/).
2. Baixe o instalador correspondente ao seu sistema operacional.
3. Siga as instruções do instalador para completar a instalação.
4. Verifique se o Go foi instalado corretamente:
   ```bash
   go version
   ```

## Instalação do Air

1. Instale o Air usando o comando `go install`:
   ```bash
   go install github.com/cosmtrek/air@latest
   ```
2. Verifique se o Air foi instalado corretamente:
   ```bash
   air -v
   ```
   O comando acima deve retornar a versão instalada do Air.

## Executando o Air no Projeto

1. Certifique-se de estar no diretório raiz do projeto.
2. Navegue até a pasta `cmd/server` (ou o diretório principal do servidor Go no projeto):
   ```bash
   cd cmd/server
   ```
3. Inicie o Air:
   ```bash
   air
   ```

O servidor será iniciado com suporte a live reload. Qualquer alteração no código-fonte reiniciará o servidor automaticamente.

---

# Documentação Paginação

O recurso fornece uma implementação eficiente de paginação para aplicações utilizando **Echo** e **GORM**, permitindo consultas otimizadas com filtros, preload de relacionamentos e metadados.

---

## Exemplo de Uso

### Estrutura do Repositório

No exemplo abaixo, utilizamos o padrão **Repository** para separar a lógica de acesso ao banco de dados.

```go
package repositories

import (
	"meuprojeto/models"
	"meuprojeto/pagination"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// userRepository representa o repositório de usuários
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository cria uma nova instância de userRepository
func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

// Listagem retorna a lista paginada de usuários
func (r *userRepository) Listagem(c echo.Context) (*pagination.Pagination, error) {
	var users []models.UsuarioView

	// Chamada da função de paginação
	paginations, err := pagination.Paginate(c, r.db, &users, nil, "UsuarioTipo")

	if err != nil {
		return nil, err
	}

	return paginations, nil
}
```

---

### Exemplo no Handler

No handler, chamamos o método `Listagem` do repositório para obter os dados paginados.

```go
package handlers

import (
	"net/http"
	"meuprojeto/repositories"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userRepo *repositories.UserRepository
}

// NewUserHandler cria uma nova instância de UserHandler
func NewUserHandler(userRepo *repositories.UserRepository) *UserHandler {
	return &UserHandler{userRepo: userRepo}
}

// ListUsers é o endpoint para listar usuários com paginação
func (h *UserHandler) ListUsers(c echo.Context) error {
	// Chamada ao repositório
	paginations, err := h.userRepo.Listagem(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Retorna os dados paginados
	return c.JSON(http.StatusOK, paginations)
}
```

---

### Exemplo de Resposta

Dada a configuração acima, ao fazer uma requisição para o endpoint, o resultado será semelhante a:

**Requisição:**
```bash
GET /users?page=1&limit=10
```

**Resposta:**
```json
{
    "total_records": 3,
    "total_pages": 1,
    "items": [
        {
            "id": 1,
            "nome": "Harcyldo Leonardo Viana Winkelmann",
            "cpf": "060.532.471-97",
            "usuario_tipo_id": 1,
            "tipo": {
                "ID": 1,
                "CreatedAt": "2024-11-18T12:46:49.003212Z",
                "UpdatedAt": "2024-11-18T12:46:49.003212Z",
                "DeletedAt": null,
                "nome": "Comum"
            }
        }
    ]
}
```

## Vantagens do Padrão

- **Modularidade**: Lógica de acesso ao banco isolada no repositório.
- **Reutilização**: `Paginate` pode ser usado em diferentes repositórios.
- **Extensibilidade**: Fácil inclusão de filtros e preload de relacionamentos conforme necessário.

---

# Subindo a Aplicação com Docker

Para iniciar a aplicação com PostgreSQL, siga os passos abaixo:

1. **Subir o PostgreSQL**

Abra um terminal e navegue até o diretório raiz do projeto. Primeiro, suba o contêiner PostgreSQL:

```bash
docker-compose  up --build postgres -d
```

Isso iniciará o serviço PostgreSQL. O contêiner estará acessível na porta `5432`.

2. **Subir os demais serviços**

Depois que o PostgreSQL estiver rodando, suba o restante dos serviços da aplicação (Go + Air + PGAdmin):

```bash
docker-compose up -d
```

O servidor Go estará acessível na porta `8080`.

3. **Verificar se tudo está funcionando**

Verifique os logs dos contêineres para garantir que tudo está rodando corretamente:

```bash
docker-compose logs -f
```

---

# Acesse a aplicação

Após iniciar os serviços, você pode acessar a aplicação através do seguinte URL no seu navegador:

- Acesse a API: [http://localhost:8080](http://localhost:8080)

---

# Parando a Aplicação

Para parar a aplicação e desligar os contêineres Docker, siga os passos abaixo:

1. Abra um terminal e navegue até o diretório raiz do projeto.
2. Execute o seguinte comando para parar e remover os contêineres, redes e volumes criados:

```bash
docker-compose down
```

---

Agora sua aplicação está configurada para rodar com Docker, incluindo o **PostgreSQL**, o servidor Go com suporte a live reload utilizando o Air.
