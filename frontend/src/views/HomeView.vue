<template>
    <div class="container">
      <div class="content">
        <div class="left-panel">
          <div class="top-section">
            <div class="top-left">
              <h2>Recentes</h2>
            </div>
            <div class="top-right">
              <div class="filter-container">
                <span>Filtrar</span>
                <v-icon>mdi-filter</v-icon>
              </div>
            </div>
          </div>
  
          <div class="bottom-section">
                <div class="card" v-for="(book, index) in books" :key="index">
                    <div class="card-top">
                        <div class="card-left">
                            <v-img
                              alt="Capa do livro"
                              :src="book.image"
                              transition="scale-transition"
                              class="img-menu-bar"
                            />
                        </div>

                        <div class="card-right">
                            <ul class="book-info">
                                <li><v-icon>{{ book.icon }}</v-icon> <strong>{{ book.titulo }}</strong></li>
                                <li><v-icon>{{ book.icon }}</v-icon> {{ book.publicacao }}</li>
                                <li><v-icon>{{ book.icon }}</v-icon> Páginas: {{ book.paginas }}</li>
                                <li><v-icon>{{ book.icon }}</v-icon> Capítulos: {{ book.capitulos }}</li>
                                <li><v-icon>{{ book.icon }}</v-icon> Volume: {{ book.volume }}</li>
                                <li><v-icon>{{ book.icon }}</v-icon> Edição: {{ book.edicao }}</li>
                                <li><v-icon>{{ book.icon }}</v-icon> Organizadores: {{ book.organizadores }}</li>
                            </ul>
                        </div>
                    </div>

                    <div class="card-bottom">
                        <div class="bottom-half">
                            <span class="rating-material">Avaliar:</span>
                            <v-rating hover :length="5" size="x-small" :model-value="3" />
                        </div>
                        <div class="bottom-half">
                            <div class="imagem-container">
                                <img :src="share" alt="Imagem">
                            </div>
                            <div class="texto-container">
                                <p>Compartilhar.</p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
  
        <v-divider vertical class="border-opacity-100" color="white"></v-divider>
  
        <div class="right-panel">
            <div class="right-top-section">
                <div class="text-card-right-panel">
                  <h3>Favoritos</h3>
                </div>
            </div>
            <div class="right-bottom-section">
                <div class="right-bottom-item">
                    <div class="bottom-item-left">
                        <v-img
                            alt="Capa do livro"
                            :src="books[0].image"
                            transition="scale-transition"
                            class="img-menu-bar"
                        />
                    </div>
                    <div class="bottom-item-right">
                        <ul class="book-info-right-card">
                            <li><v-icon>mdi-search</v-icon> <strong>{{ books[0].titulo }}</strong></li>
                            <li><v-icon>mdi-search</v-icon> {{ books[0].publicacao }}</li>
                            <li><v-icon>mdi-search</v-icon> Páginas: {{ books[0].paginas }} | Capítulos: {{ books[0].capitulos }} | Volume: {{ books[0].volume }} | Edição: {{ books[0].edicao }}</li>
                            <li><v-icon>mdi-search</v-icon> Organizadores: {{ books[0].organizadores }}</li>
                        </ul>
                    </div>
                </div>
                <div class="right-bottom-item">
                    <div class="bottom-item-left">
                        <v-img
                            alt="Capa do livro"
                            :src="books[1].image"
                            transition="scale-transition"
                            class="img-menu-bar"
                        />
                    </div>
                    <div class="bottom-item-right">
                        <ul class="book-info-right-card">
                            <li><v-icon>mdi-search</v-icon> <strong>{{ books[1].titulo }}</strong></li>
                            <li><v-icon>mdi-search</v-icon> {{ books[1].publicacao }}</li>
                            <li><v-icon>mdi-search</v-icon> Páginas: {{ books[1].paginas }} | Capítulos: {{ books[1].capitulos }} | Volume: {{ books[1].volume }} | Edição: {{ books[1].edicao }}</li>
                            <li><v-icon>mdi-search</v-icon> Organizadores: {{ books[1].organizadores }}</li>
                        </ul>
                    </div>
                </div>
                <div class="right-bottom-item">
                    <div class="bottom-item-left">
                        <v-img
                            alt="Capa do livro"
                            :src="books[2].image"
                            transition="scale-transition"
                            class="img-menu-bar"
                        />
                    </div>
                    <div class="bottom-item-right">
                        <ul class="book-info-right-card">
                            <li><v-icon>mdi-search</v-icon> <strong>{{ books[2].titulo }}</strong></li>
                            <li><v-icon>mdi-search</v-icon> {{ books[2].publicacao }}</li>
                            <li><v-icon>mdi-search</v-icon> Páginas: {{ books[2].paginas }} | Capítulos: {{ books[2].capitulos }} | Volume: {{ books[2].volume }} | Edição: {{ books[2].edicao }}</li>
                            <li><v-icon>mdi-search</v-icon> Organizadores: {{ books[2].organizadores }}</li>
                        </ul>
                    </div>
                </div>
            </div>
        </div>
      </div>
    </div>
</template>

<script>
import BookImage1 from '@/assets/images/content/book-image.jpg'
import BookImage2 from '@/assets/images/content/book-image-2.jpg'
import BookImage3 from '@/assets/images/content/book-image-3.jpg'
import Compartilhar from '@/assets/images/home/share-img.webp'

export default {
  data: () => ({
    share: Compartilhar,
    books: [
      {
        titulo: 'Tempos Digitais',
        publicacao: '08/10/2022',
        paginas: '200',
        capitulos: '5',
        volume: '1',
        edicao: '1',
        organizadores: 'Hélio Lemes Costa Jr',
        icon: 'mdi-clock',
        image: BookImage1
      },
      {
        titulo: 'Introdução à Tecnologia de Mostradores',
        publicacao: '12/07/2019',
        paginas: '126',
        capitulos: '7',
        volume: '2',
        edicao: '1',
        organizadores: 'Andrei Frederico López',
        icon: 'mdi-clock',
        image: BookImage2
      },
      {
        titulo: 'Tecnologia Vs Humanidade',
        publicacao: '04/04/2019',
        paginas: '200',
        capitulos: '5',
        volume: '1',
        edicao: '1',
        organizadores: 'Gerd Leonhard',
        icon: 'mdi-clock',
        image: BookImage3
      }
    ]
  })
}
</script>

<style scoped>
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.content {
  background-color: rgba(36, 36, 36, 0.45);
  border-radius: 45px;
  margin-top: 50px;
  width: 95%;
  height: 450px; /* No pc de casa, aumentar para 550px */
  display: flex;
  flex-direction: row;
}

.left-panel {
  width: 70%;
  padding: 1rem;
  border-top-left-radius: 45px;
  border-bottom-left-radius: 45px;
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
}

.top-section {
  height: 10%;
  display: flex;
  width: 100%;
  margin-bottom: 1rem;
}

.top-left {
  width: 60%;
  font-weight: bolder;
  color: white;
  display: flex;
  align-items: center;
  padding-left: 1rem;
}

.top-right {
  width: 40%;
  color: white;
  display: flex;
  justify-content: flex-end;
  align-items: center;
  padding-right: 1rem;
}

.filter-container {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.bottom-section {
  display: flex;
  gap: 1rem;
  height: 90%;
  width: 100%;
}

.card {
  flex: 1;
  padding: 0;
  border-radius: 12px;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.card-top {
  display: flex;
  height: 75%;
  width: 100%;
}

.card-left,
.card-right {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  padding: 0;
  margin: 0;
}

.card-left {
  width: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  overflow: hidden; /* Garante que nada ultrapasse a div */
}

.book-cover {
  max-height: 90%;
  max-width: 90%;
  object-fit: contain; /* Mantém proporção sem cortar */
  border-radius: 8px;
}

.card-right {
  width: 50%;
  display: flex;
  padding: 0.5rem;
  box-sizing: border-box;
}

.card-bottom {
  height: 25%;
  width: 100%;
  display: flex;
  flex-direction: column; /* Divide verticalmente */
  padding: 0;
  margin: 0;
}

.bottom-half {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
}

.book-info {
  font-size: 0.65rem; /* No pc de casa, aumentar para 0.75rem */
  color: black;
  text-align: left;
  padding: 0.5rem;
  list-style-type: disc;      /* ativa as bolinhas */
  padding-left: 1.2rem;       /* espaço para as bolinhas */
}

.book-info li {
  margin-bottom: 0.2rem;
  color: white;
}

.rating-material {
  font-size: 1rem;
  font-weight: bolder;
  color: white;
  padding-top: 3px;
}

.imagem-container {
  width: 20%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.imagem-container img {
  width: 75%;
  height: 75%;
  object-fit: cover;
}

.texto-container {
  width: 50%;
  height: 100%;
  font-weight: bolder;
  color: white;
  display: flex;
  justify-content: center;
  align-items: center;
}

/* Div panel da direita */

.right-panel {
  width: 30%;
  padding: 1rem;
  border-top-right-radius: 45px;
  border-bottom-right-radius: 45px;
  box-sizing: border-box;
}

.right-top-section {
  width: 100%;
  height: 10%;
}

.right-bottom-section {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  width: 100%;
  height: 90%;
}

.right-bottom-item {
  display: flex;
  width: 100%;
  height: 30%;
  background-color: rgba(36, 36, 36, 0.45);
}

.bottom-item-left {
  width: 25%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.bottom-item-right {
  width: 75%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: flex-start;
  padding-left: 10px;
}

.book-info-right-card {
  font-size: 0.6rem;
  margin: 0;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: flex-start;
  gap: 4px;
}

.book-info-right-card li {
  display: flex;
  align-items: center;
  gap: 1px;
  color: white;
}

.text-card-right-panel {
  display: flex;
  align-items: center;         /* Centraliza verticalmente */
  justify-content: flex-start; /* Alinha à esquerda horizontalmente */
  height: 100%;                /* Garante que a altura seja respeitada */
  padding-left: 1rem;          /* (opcional) afasta um pouco da borda esquerda */
  font-weight: bolder;
  color: white;
}
</style>
