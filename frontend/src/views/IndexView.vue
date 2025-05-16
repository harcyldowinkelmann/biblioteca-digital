<template>
  <div class="home">
    <div class="main-div">
      <!-- Primeiro Material -->
      <div class="box">
        <!-- Primeira imagem -->
        <div class="sub-box">
          <v-img
            alt="Biblioteca Digital"
            :src="bookImage"
            transition="scale-transition"
            class="img-menu-bar"
          ></v-img>
        </div>

        <!-- Informações Material -->
        <div class="sub-box">
          <div class="list-infos-book">
            <div class="pa-6">
              <ul>
                <li>Titulo: Tempos Digitais</li>
                <li>Publicado em: 08/10/2022</li>
                <li>Páginas: 200 | Capítulos: 5 | Volume: 1 | Edição: 1</li>
                <li>Organizadores: Hélio Lemes Costa Jr</li>
              </ul>
            </div>
            <div class="saiba-mais">
              <ul title="Mais Informações">
                <li><span class="saiba-mais-link"><a href="#">🔗 - Download em PDF</a></span></li>
              </ul>
            </div>
          </div>
        </div>
      </div>

      <!-- Segundo Material -->
      <div class="box">
        <!-- Segunda imagem -->
        <div class="sub-box">
          <v-img
            alt="Biblioteca Digital"
            :src="bookImage2"
            transition="scale-transition"
            class="img-menu-bar"
          ></v-img>
        </div>

        <!-- Informações Material -->
        <div class="sub-box">
          <div class="list-infos-book">
            <div class="pa-6">
              <ul>
                <li>Titulo: Introdução à Tecnologia de Mostradores (Displays ou Telas)</li>
                <li>Publicado em: 12/07/2019</li>
                <li>Páginas: 126 | Capítulos: 7 | Volume: 2 | Edição: 2</li>
                <li>Organizadores: Andrei Frederico López e Jorge Tomioka</li>
              </ul>
            </div>
            <div class="saiba-mais">
              <ul title="Mais Informações">
                <li><span class="saiba-mais-link"><a href="#">🔗 - Download em PDF</a></span></li>
              </ul>
            </div>
          </div>
        </div>
      </div>

      <!-- Terceiro Material -->
      <div class="box">
        <!-- Terceira imagem -->
        <div class="sub-box">
          <v-img
            alt="Biblioteca Digital"
            :src="bookImage3"
            transition="scale-transition"
            class="img-menu-bar"
          ></v-img>
        </div>

        <!-- Informações Material -->
        <div class="sub-box">
          <div class="list-infos-book">
            <div>
              <ul>
                <li>Titulo: Tecnologia Vs Humanidade</li>
                <li>Publicado em: 04/04/2019</li>
                <li>Páginas: 200 | Capítulos: 5 | Volume: 1 | Edição: 1</li>
                <li>Organizadores: Gerd Leonhard</li>
              </ul>
            </div>
            <div class="saiba-mais">
              <ul title="Mais Informações" class="mais-informacoes">
                <li><span class="saiba-mais-link"><a href="#">🔗 - Download em PDF</a></span></li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>

    <FooterMain />
  </div>
</template>

<script>
import FooterMain from '@/components/FooterMain.vue';
import BookImage from '@/assets/images/content/book-image.jpg'
import BookImage2 from '@/assets/images/content/book-image-2.jpg'
import BookImage3 from '@/assets/images/content/book-image-3.jpg'

export default {
  name: 'IndexView',
  components: {FooterMain},
  data: () => ({
    bookImage: BookImage,
    bookImage2: BookImage2,
    bookImage3: BookImage3
  }),

  created() {
    const {cpf, senha} = this.$route.query;
    console.log('CPF recebido:', cpf);
    console.log('Senha recebida:', senha);
  }
}
</script>

<style scoped>

.home {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100vh;
  width: 100vw;
}

.main-div {
  width: 97vw;
  height: 50vh;
  background-color: rgba(36, 36, 36, 0.45); /* Define a cor com opacidade */
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-radius: 45px;
  margin-top: 70px;
  padding: 5px;
}

.box {
  width: 33%;
  height: 100%;
  border-radius: 45px;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 24px;
  font-weight: bold;
  color: black;
}

/* 📌 Faz a primeira box se tornar um flex container */
.box:first-child {
  display: flex;
  justify-content: space-between; /* Espaço entre as divs internas */
  align-items: center;
  padding: 2px;
}

.sub-box:last-child {
  border-radius: 0 45px 45px 0; /* Arredonda apenas o lado direito */
}

/* 📌 As duas sub-divs dentro da primeira box */
.sub-box {
  width: 49%; /* Ocupa quase metade da box */
  height: 100%;
  border-radius: 10px;
  display: flex;
  justify-content: center;
  align-items: center;
  color: black;
  border-radius: 45px;
}

.img-menu-bar {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 45px;
}

li {
  list-style: none;
  font-family: 'Times New Roman', Times, serif;
  font-weight: bold;
  font-size: small;
  color: white;
  text-align: left;
  margin: 15px 5px;
}

.mais-informacoes {
  list-style: disc;
}

.saiba-mais-link a {
  text-decoration: none;
  color: white;
  font-size: small;
}

.saiba-mais-link a:hover {
  text-decoration: underline;
}

</style>
