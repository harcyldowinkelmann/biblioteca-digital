<template>
  <div class="home">
    <div class="main-div">
      <div class="box">
        <div class="sub-box">
          <v-img
            alt="Biblioteca Digital"
            :src="bookImage"
            transition="scale-transition"
            width="100"
            class="img-menu-bar"
          ></v-img>
        </div>
        <div class="sub-box">
          B
        </div>
      </div>
      <div class="box">2</div>
      <div class="box">3</div>
    </div>

    <FooterMain />
  </div>
</template>

<script>
import FooterMain from '@/components/FooterMain.vue';
import BookImage from '@/assets/images/content/book-image.png'

export default {
  name: 'HomeView',
  components: {FooterMain},
  data: () => ({
    bookImage: BookImage
  })
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
  border: 1px solid yellow;
}

.main-div {
  width: 97vw;
  height: 60vh;
  background-color: #242424;
  opacity: 45%;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-radius: 45px;
  margin-top: 100px;
  padding: 5px;
}

.box {
  width: 33%;
  height: 100%;
  background-color: white;
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
  padding: 5px;
}

/* 📌 As duas sub-divs dentro da primeira box */
.sub-box {
  width: 49%; /* Ocupa quase metade da box */
  height: 100%;
  background-color: lightblue;
  border-radius: 10px;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 18px;
  font-weight: bold;
  color: black;
  border-radius: 45px;
}

.img-menu-bar {
  width: 100%;
  height: 100%;
}

</style>
