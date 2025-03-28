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
                        width="250"
                        class="mx-auto"
                    >
                    </v-img>
                </div>
                <v-card class="pa-4 mx-auto box" elevation="4">
                    <v-form
                        method="post"
                        ref="form"
                    >
                        <div class="text-subtitle-1 text-medium-emphasis d-flex align-center justify-space-between">
                            CPF
                        </div>
        
                        <v-text-field
                            rounded="xl"
                            density="compact"
                            placeholder="Entre com seu CPF..."
                            prepend-inner-icon="mdi-account-key"
                            variant="outlined"
                            v-mask="['###.###.###-##']"
                        ></v-text-field>
                        <div class="text-subtitle-1 text-medium-emphasis d-flex align-center justify-space-between">
                            Senha
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
                        <v-btn
                            block
                            rounded="xl"
                            class="mb-8"
                            color="#007744"
                            size="large"
                            variant="flat"
                            @click="logar()"
                        >
                            LOGIN
                        </v-btn>
                    </v-form>
                </v-card>
            </v-main>
            <div class="versaoSistema">
                <span> Biblioteca Digital 2025 - versão: <a class="link">{{ pack.version }}</a></span>
            </div>
        </div>
    </v-app>
  </template>


<script>
import personImg from '@/assets/images/login/person-img.png'

export default {
    name: "LoginView",
    data: () => ({
        pack: require('../../package.json'),
        personImg: personImg
    })
}
</script>

<style scoped lang="css">
    @media screen and (min-width: 640px) {
        .form-login {
            width: 30%;
        }
    }

    .main-screen {
        background-color: #3a6391;
    }

    .versaoSistema {
        text-align: center;
        margin-top: 2rem;
        font-weight: bold;
    }

    .link {
        color: #007744;
        text-decoration: none;
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
        background-color: #3a6391;
        display: flex;
        justify-content: center; /* Centraliza horizontalmente */
        align-items: center; /* Centraliza verticalmente */
        flex-direction: column; /* Alinha as divs uma abaixo da outra */
    }

    .box {
        width: 500px;
    }

    .text-field-rounded {
        border-radius: 100px !important;
    }
</style>
