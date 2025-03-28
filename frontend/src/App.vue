<template>
  <GlobalMenuBar />
  <router-view />
</template>

<script>
import GlobalMenuBar from './components/GlobalMenuBar.vue';

export default {
  name: 'App',
  components: {GlobalMenuBar}
}
</script>

<style>

* {
  margin: 0px;
  padding: 0px;
}

#app {
  background-color: #3a6391;
  min-height: 100vh;
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
