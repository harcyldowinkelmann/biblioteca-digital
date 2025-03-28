<template>
    <div class="global-menu-bar">
        <v-row align="center">
            <v-col cols="1">
                <v-img
                    alt="Biblioteca Digital"
                    :src="ImgMenuBar"
                    transition="scale-transition"
                    width="70"
                    class="img-menu-bar"
                ></v-img>
            </v-col>

            <v-col cols="3" class="text-col">
                <span class="text-menu-bar">BIBLIOTECA DIGITAL</span>
            </v-col>

            <!-- Barra de Pesquisa -->
            <v-col cols="6">
                <v-text-field
                    type="text"
                    variant="outlined"
                    label="Pesquise algum material"
                    prepend-inner-icon="mdi-magnify"
                    hide-details
                    single-line
                    rounded="xl"
                    base-color="white"
                    color="#ffffff"
                ></v-text-field>
            </v-col>

            <v-col cols="2">
                <v-menu
                    transition="scale-transition"
                >
                    <template v-slot:activator="{ props }">
                        <v-btn
                            color="primary"
                            v-bind="props"
                        > 
                            Menu 
                        </v-btn>
                    </template>

                    <v-list>
                        <v-list-item
                            v-for="(item, index) in items"
                            :key="index"
                            :value="index"
                        >
                            <v-list-item-title>{{ item.title }}</v-list-item-title>
                        </v-list-item>
                    </v-list>
                </v-menu>
            </v-col>
        </v-row>
    </div>
</template>

<script>
import ImgMenuBar from '@/assets/images/login/img-logo-menu-bar.png'
import MenuImg from '@/assets/images/content/menu.png'

export default {
    name: "GlobalMenuView",
    data: () => ({
        ImgMenuBar: ImgMenuBar,
        MenuImg: MenuImg,
        items: [
            { title: 'Minhas Mídias' },
            { title: 'Suporte' }
        ]
    })
}
</script>

<style scoped>
nav {
    padding: 30px;
}

nav a {
    font-weight: bold;
    color: #2c3e50;
}

nav a.router-link-exact-active {
    color: #42b983;
}

.global-menu-bar {
    background-color: #242424;
    /*background-color: black;*/
    height: 70px;
    width: 100%;
    position: absolute;
    margin-top: 35px;
    border-radius: 50ch;
}

.img-menu-bar {
    margin-left: 50px;
}

.text-col {
    display: flex;
    align-items: center; /* Alinha o texto verticalmente */
}

.text-menu-bar {
    color: white;
    font-size: large;
    font-weight: bolder;
    margin-left: 10px; /* Espaço entre o texto e a imagem */
}

</style>
