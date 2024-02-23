<template>
  <nav class="bg-gray-800">
    <div class="content">
      <RouterLink to="/">Inicio</RouterLink>
      <RouterLink to="/search">Buscar</RouterLink>
    </div>
  </nav>
</template>

<script>
import { RouterLink } from "vue-router";
</script>

<style scoped>
nav {
  height: 55px;
  color: white;
  border-bottom-left-radius: 10px;
  border-bottom-right-radius: 10px;
}
.content {
  /* background-color: blueviolet; */
  text-align: right;
  width: 100%;
  padding: 15px;
}

a {
  margin: 10px;
  padding: 10px;
}

a:hover {
  background-color: #374151;
  padding: 15px 10px;
  border-radius: 10px;
}

.router-link-active {
  background-color: #4b5563;
  padding: 15px 10px;
  border-radius: 10px;
}
</style>
