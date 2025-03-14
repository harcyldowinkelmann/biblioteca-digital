package constants

const (
	CAMPOS_OBRIGATORIOS     string = "campo obrigatório"
	CAMPOS_INVALIDOS        string = "campo inválido"
	PREENCHAS_TODOS_CAMPOS  string = "preencha todos os campos corretamente"
	CADASTRO_ALTERADO       string = "cadastro foi alterado"
	CADASTRO_FALHA_INSERCAO string = "não foi possível realizar cadastro"
	CADASTRO_REALIZADO      string = "cadastro realizado"
	CADASTRO_EXCLUIDO       string = "cadastro excluído"
	CADASTRO_FALHA_EXCLUSAO string = "não foi possível excluir registro"
	ID_NAO_INFORMADO        string = "id não foi informado"
	ERROR_ATIVACAO_CONTA    string = "não foi possível realizar ativação da sua conta"
	BODY_FALHA              string = "falha no body"
	USUARIO_ENCONTRADO      string = "usuário encontrado"
	SENHA_INVALIDA          string = "senha inválida"
	REGISTRO_NAO_ENCONTRADO string = "registro não encontrado"
	FALHA_AUTENTICAO        string = "falha na autenticação"
	ACESSO_NAO_AUTORIZADO   string = "acesso não autorizado"
	ERRO_LISTAGEM_REGISTRO  string = "não foi possível listar registros"
	CPF_INVALIDO            string = "cpf informado já está vinculado a uma conta existente"
	ERROR_GERAR_TOKEN       string = "não foi possível gerar token de ativação de conta: %w"
	ERROR_CARREGAMENTO_ENV  string = "erro ao carregar o arquivo .env"
)
