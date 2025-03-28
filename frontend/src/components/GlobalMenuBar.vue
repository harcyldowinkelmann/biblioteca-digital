<template>
    <!--<nav>
        <router-link to="/">Home</router-link> |
        <router-link to="/about">About</router-link>
    </nav>-->
    <div class="global-menu-bar">
        <v-row align="center">
            <v-col cols="2">
                <v-img
                        alt="Biblioteca Digital"
                        :src="ImgMenuBar"
                        transition="scale-transition"
                        width="75"
                        class="img-menu-bar"
                ></v-img>
            </v-col>

            <v-col cols="8" align-self="auto">
                <span class="text-menu-bar">BIBLIOTECA DIGITAL {{ items[0] }}</span>
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
import ImgMenuBar from '@/assets/images/login/book-menu-img.png'

export default {
    name: "GlobalMenuView",
    data: () => ({
        ImgMenuBar: ImgMenuBar,
        items: [
            { title: 'Click Me' },
            { title: 'Click Me' },
            { title: 'Click Me' },
            { title: 'Click Me 2' },
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
    background-color: rgb(4, 4, 29);
    height: 75px;
    width: 100%;
    position: absolute;
    margin-top: 35px;
    border-radius: 50ch;
}

.img-menu-bar {
    margin-left: 50px;
}

.text-menu-bar {
    color: white;
    font-size: large;
    font-weight: bolder;
    align-items: start;
}

</style>
