<template>
    <v-container
      class="fill-height"
      fluid
    >
      <v-row align="center" justify="center">
        <v-col cols="12" md="10">
          <div class="div-cadastro">
            <v-form method="post" ref="form">
              <v-row>
                <v-col cols="12">
                  <p style="font-weight: bolder; font-size: 1.5em;">CADASTRE-SE</p>
                </v-col>

                <v-col cols="6">
                  <div class="text-subtitle-1 text-medium-emphasis d-flex align-center justify-space-between">
                    <span>Nome</span>
                  </div>
                  <v-text-field
                    rounded="xl"
                    density="compact"
                    placeholder="Entre com seu Nome..."
                    prepend-inner-icon="mdi-account-key"
                    variant="outlined"
                    color="primary"
                  ></v-text-field>
                </v-col>

                <v-col cols="6">
                  <div class="text-subtitle-1 text-medium-emphasis d-flex align-center justify-space-between">
                    <span>Instituição de ensino</span>
                  </div>
                  <v-text-field
                    rounded="xl"
                    density="compact"
                    placeholder="Instituição..."
                    prepend-inner-icon="mdi-account-key"
                    variant="outlined"
                    color="primary"
                  ></v-text-field>
                </v-col>
              </v-row>

              <v-row>
                <v-col cols="6">
                  <div class="text-subtitle-1 text-medium-emphasis d-flex align-center justify-space-between">
                    <span>E-mail</span>
                  </div>
                  <v-text-field
                    rounded="xl"
                    density="compact"
                    placeholder="Entre com seu e-mail..."
                    prepend-inner-icon="mdi-account-key"
                    variant="outlined"
                    color="primary"
                  ></v-text-field>
                </v-col>

                <v-col cols="6">
                  <div class="text-subtitle-1 text-medium-emphasis d-flex align-center justify-space-between">
                    <span>Senha</span>
                  </div>
                  <v-text-field
                    rounded="xl"
                    density="compact"
                    placeholder="Entre com uma senha..."
                    prepend-inner-icon="mdi-account-key"
                    variant="outlined"
                    color="primary"
                    type="password"
                  ></v-text-field>
                </v-col>
              </v-row>

              <v-row>
                <v-col cols="6">
                  <div class="text-subtitle-1 text-medium-emphasis d-flex align-center justify-space-between">
                    <span>Eu sou:</span>
                  </div>
                  <v-autocomplete
                      rounded="xl"
                      variant="outlined"
                      density="compact"
                      :items="tipo_usuario"
                      item-value="id"
                      item-title="tipo"
                      auto-select-first
                      chips
                      clearable
                      deletable-chips
                  ></v-autocomplete>
                </v-col>

                <v-col cols="6" class="d-flex align-center">
                  <v-btn
                      rounded="xl"
                      class="mb-0"
                      color="indigo"
                      size="large"
                      variant="flat"
                      @click="cadastrar()"
                  >
                      Cadastrar-me
                  </v-btn>
                </v-col>
              </v-row>
            </v-form>
          </div>
        </v-col>
      </v-row>
    </v-container>
</template>

<script>
export default {
  name: 'CadastroView',
  data: () => ({
    tipo_usuario: [
      {
        "id": 1,
        "tipo": "Aluno"
      },
      {
        "id": 2,
        "tipo": "Professor"
      },
      {
        "id": 3,
        "tipo": "Curador Especializado"
      },
      {
        "id": 4,
        "tipo": "Suporte"
      }
    ]
  })
}
</script>

<style>
html, body, #app {
    height: 100%;
    margin: 0;
}

.div-cadastro {
    border: 1px solid black;
    padding: 24px;
    border-radius: 8px;
    background-color: white;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
    width: 100%;
}

span {
  font-weight: bolder;
}
</style>
