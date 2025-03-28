<template>
    <v-app class="main-screen">
        <div class="form-login">
            <LoadLogin v-if="carregando" :message="message" />
            <v-main v-else class="container">
                <div class="box">
                    <v-img
                        alt="Biblioteca Digital"
                        :src="personImg"
                        transition="scale-transition"
                        width="150"
                        class="mx-auto"
                    >
                    </v-img>
                </div>
                <v-card variant="primary" class="pa-4 mx-auto box" elevation="0">
                    <v-form
                        method="post"
                        ref="form"
                    >
                        <div class="text-subtitle-1 text-medium-emphasis d-flex align-center justify-space-between">
                            <span>CPF</span>
                        </div>
                        <v-text-field
                            rounded="xl"
                            density="compact"
                            placeholder="Entre com seu CPF..."
                            prepend-inner-icon="mdi-account-key"
                            variant="outlined"
                            v-mask="['###.###.###-##']"
                            color="primary"
                        ></v-text-field>

                        <div class="text-subtitle-1 text-medium-emphasis d-flex align-center justify-space-between">
                            <span>Senha</span>
                        </div>
                        <v-text-field
                            rounded="xl"
                            :append-inner-icon="visible ? 'mdi-eye-off' : 'mdi-eye'"
                            :type="visible ? 'text' : 'password'"
                            density="compact"
                            placeholder="Entre com sua senha..."
                            prepend-inner-icon="mdi-lock-outline"
                            variant="outlined"
                        ></v-text-field>

                        <div class="esqueci-minha-senha">
                            <span><a href="#">Esqueci minha senha</a></span>
                        </div>

                        <v-btn
                            rounded="xl"
                            class="mb-8"
                            color="indigo"
                            size="large"
                            variant="flat"
                            @click="logar()"
                        >
                            LOGIN
                        </v-btn>
                    </v-form>

                    <div class="nao-tenho-cadastro">
                        <a href="#">
                            <span>🔗 - Não tenho cadastro</span>
                        </a>
                    </div>
                </v-card>
            </v-main>
        </div>
        <div class="versaoSistema">
            <span> Biblioteca Digital 2025 - Versão: <a class="link">{{ pack.version }} </a> | Projeto de Sistemas Computacionais II - UFGD | Copyright 2024 Todos os direitos reservados &copy;</span>
        </div>
    </v-app>
  </template>


<script>
import personImg from '@/assets/images/login/person-img.png'
import linkExterno from '@/assets/images/login/link-externo.jpg'

export default {
    name: "LoginView",
    data: () => ({
        pack: require('../../package.json'),
        personImg: personImg,
        linkExterno: linkExterno
    })
}
</script>

<style scoped lang="css">

    * {
        color: white;
        font-weight: bolder;
    }

    .main-screen {
        background-color: #3a6391;
        display: flex;
        justify-content: center; /* Centraliza horizontalmente */
        align-items: center; /* Centraliza verticalmente */
        height: 100vh; /* Ocupa toda a altura da tela */
    }

    @media screen and (min-width: 640px) {
        .form-login {
            width: 30%;
            background: #3a6391;
            padding: 20px;
            border-radius: 10px;
            margin-top: 175px;
        }
    }

    #border {
        border: 5px solid rgb(194, 60, 60); /* Cor e espessura da borda */
        border-radius: 10px; /* Borda arredondada opcional */
        padding: 10px; /* Para evitar que o conteúdo encoste na borda */
    }

    .versaoSistema {
        text-align: center;
        margin-top: 2rem;
        font-weight: bold;
        font-size: small;
    }

    .link {
        text-decoration: underline;
        font-weight: bold;
    }

    #login {
        height: 100%;
        width: 100%;
        position: absolute;
        top: 0;
        left: 0;
        content: "";
        z-index: 0;
    }

    .form-login {
        padding: 25px 25px;
        width: 100%;
    }

    .container {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
    }

    .box {
        width: 500px;
    }

    .esqueci-minha-senha {
        font-weight: lighter !important;
        text-align: left;
        font-size: 14px;
        margin-top: 5px;
        margin-left: 12px; /* Ajusta o alinhamento para combinar com os campos */
    }

    .esqueci-minha-senha a {
        color: white; /* Ajuste conforme necessário */
        text-decoration: underline;
        font-weight: normal;
    }

    .nao-tenho-cadastro {
        text-align: left;
        font-size: 14px;
        margin-top: 5px;
        margin-left: 12px; /* Ajuste conforme necessário */
        display: flex;
        align-items: center; /* Alinha o texto com a imagem */
        gap: 8px; /* Espaço entre o ícone e o texto */
    }

    .nao-tenho-cadastro a {
        color: white; /* Ajuste conforme necessário */
        text-decoration: none;
        font-weight: normal;
        display: flex;
        align-items: center;
    }

    .nao-tenho-cadastro a:hover {
        text-decoration: underline;
    }

</style>
