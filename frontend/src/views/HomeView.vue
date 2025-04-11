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
            <div class="card">
                <div class="card-top">
                    <div class="card-left">
                        <v-img
                            alt="Biblioteca Digital"
                            :src="bookImage"
                            transition="scale-transition"
                            class="img-menu-bar"
                        ></v-img>
                    </div>
                    <div class="card-right">
                        <ul class="book-info">
                            <li><v-icon>{{ bookInfos[2].icon }}</v-icon><strong>{{ bookInfos[0].titulo }}</strong></li>
                            <li><v-icon>{{ bookInfos[0].icon }}</v-icon> {{ bookInfos[0].publicacao }}</li>
                            <li><v-icon>{{ bookInfos[0].icon }}</v-icon> Páginas: {{ bookInfos[0].paginas }}</li>
                            <li><v-icon>{{ bookInfos[0].icon }}</v-icon> Capítulos: {{ bookInfos[0].capitulos }}</li>
                            <li><v-icon>{{ bookInfos[0].icon }}</v-icon> Volume: {{ bookInfos[0].volume }}</li>
                            <li><v-icon>{{ bookInfos[0].icon }}</v-icon> Edição: {{ bookInfos[0].edicao }}</li>
                            <li><v-icon>{{ bookInfos[0].icon }}</v-icon> Organizadores: {{ bookInfos[0].organizadores }}</li>
                        </ul>
                    </div>
                </div>

                <div class="card-bottom">
                    <div class="bottom-half">Parte 1</div>
                    <div class="bottom-half">Parte 2</div>
                </div>
            </div>
  
            <div class="card">
              <div class="card-top">
                <div class="card-left">
                    <v-img
                        alt="Biblioteca Digital"
                        :src="bookImage2"
                        transition="scale-transition"
                        class="img-menu-bar"
                    ></v-img>
                </div>
                <div class="card-right">
                    <ul class="book-info">
                        <li><v-icon>{{ bookInfos[2].icon }}</v-icon> <strong>{{ bookInfos[1].titulo }}</strong></li>
                        <li><v-icon>{{ bookInfos[1].icon }}</v-icon> {{ bookInfos[1].publicacao }}</li>
                        <li><v-icon>{{ bookInfos[0].icon }}</v-icon> Páginas: {{ bookInfos[1].paginas }}</li>
                        <li><v-icon>{{ bookInfos[0].icon }}</v-icon> Capítulos: {{ bookInfos[1].capitulos }}</li>
                        <li><v-icon>{{ bookInfos[0].icon }}</v-icon> Volume: {{ bookInfos[1].volume }}</li>
                        <li><v-icon>{{ bookInfos[0].icon }}</v-icon> Edição: {{ bookInfos[1].edicao }}</li>
                        <li><v-icon>{{ bookInfos[0].icon }}</v-icon> Organizadores: {{ bookInfos[1].organizadores }}</li>
                    </ul>
                </div>
              </div>
              <div class="card-bottom">Inferior 2</div>
            </div>
  
            <div class="card">
              <div class="card-top">
                <div class="card-left">
                    <v-img
                        alt="Biblioteca Digital"
                        :src="bookImage3"
                        transition="scale-transition"
                        class="img-menu-bar"
                    ></v-img>
                </div>
                <div class="card-right">
                    <ul class="book-info">
                        <li><v-icon>{{ bookInfos[2].icon }}</v-icon> <strong>{{ bookInfos[2].titulo }}</strong></li>
                        <li><v-icon>{{ bookInfos[2].icon }}</v-icon> {{ bookInfos[2].publicacao }}</li>
                        <li><v-icon>{{ bookInfos[2].icon }}</v-icon> Páginas: {{ bookInfos[2].paginas }}</li>
                        <li><v-icon>{{ bookInfos[2].icon }}</v-icon> Capítulos: {{ bookInfos[2].capitulos }}</li>
                        <li><v-icon>{{ bookInfos[2].icon }}</v-icon> Volume: {{ bookInfos[2].volume }}</li>
                        <li><v-icon>{{ bookInfos[2].icon }}</v-icon> Edição: {{ bookInfos[2].edicao }}</li>
                        <li><v-icon>{{ bookInfos[2].icon }}</v-icon> Organizadores: {{ bookInfos[2].organizadores }}</li>
                    </ul>
                </div>
              </div>
              <div class="card-bottom">Inferior 3</div>
            </div>
          </div>
        </div>
  
        <v-divider vertical class="border-opacity-100" color="warning"></v-divider>
  
        <div class="right-panel">
          <!-- Conteúdo da direita -->
        </div>
      </div>
    </div>
</template>

<script>
import BookImage from '@/assets/images/content/book-image.jpg'
import BookImage2 from '@/assets/images/content/book-image-2.jpg'
import BookImage3 from '@/assets/images/content/book-image-3.jpg'

export default {
  name: 'HomeView',
  data: () => ({
    bookImage: BookImage,
    bookImage2: BookImage2,
    bookImage3: BookImage3,
    bookInfos: [
        {
            titulo: 'Tempos Digitais',
            publicacao: '08/10/2022',
            paginas: '200',
            capitulos: '5',
            volume: '1',
            edicao: '1',
            organizadores: 'Hélio Lemes Costa Jr',
            icon: 'mdi-clock'
        },
        {
            titulo: 'Introdução à Tecnologia de Mostradores (Displays ou Telas)',
            publicacao: '12/07/2019',
            paginas: '126',
            capitulos: '7',
            volume: '2',
            edicao: '1',
            organizadores: 'Andrei Frederico López e Jorge Tomioka',
            icon: 'mdi-clock'
        },
        {
            titulo: 'Tecnologia Vs Humanidade',
            publicacao: '04/04/2019',
            paginas: '200',
            capitulos: '5',
            volume: '1',
            edicao: '1',
            organizadores: 'Gerd Leonhard',
            icon: 'mdi-clock'
        },
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
  margin-top: 50px;
  width: 95%;
  height: 550px;
  background-color: lightblue;
  display: flex;
  flex-direction: row;
}

.left-panel {
  width: 70%;
  background-color: #e0f7fa;
  padding: 1rem;
  border-top-left-radius: 45px;
  border-bottom-left-radius: 45px;
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
}

.right-panel {
  width: 30%;
  background-color: #ffe0b2;
  padding: 1rem;
  border-top-right-radius: 45px;
  border-bottom-right-radius: 45px;
  box-sizing: border-box;
}

.top-section {
  height: 10%;
  display: flex;
  width: 100%;
  margin-bottom: 1rem;
}

.top-left {
  width: 60%;
  background-color: #bbdefb;
  display: flex;
  align-items: center;
  padding-left: 1rem;
}

.top-right {
  width: 40%;
  background-color: #c8e6c9;
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
  background-color: #f5f5f5;
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
  background-color: #e1bee7;
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
  background-color: #fff9c4;
  display: flex;
  padding: 0.5rem;
  box-sizing: border-box;
  border: 1px black solid;
}

.card-bottom {
  height: 25%;
  width: 100%;
  background-color: #ffe0b2;
  display: flex;
  flex-direction: column; /* Divide verticalmente */
  padding: 0;
  margin: 0;
  border: 1px black solid;
}

.bottom-half {
  width: 100%;
  height: 50%; /* 50% da altura do card-bottom */
  display: flex;
  justify-content: center;
  align-items: center;
  border-top: 1px solid rgba(0, 0, 0, 0.1); /* opcional: separação visual */
}

.book-info {
  font-size: 0.75rem;
  color: black;
  text-align: left;
  padding: 0.5rem;
  list-style-type: disc;      /* ativa as bolinhas */
  padding-left: 1.2rem;       /* espaço para as bolinhas */
}

.book-info li {
  margin-bottom: 0.2rem;
}
</style>
