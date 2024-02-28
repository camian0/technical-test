<template>
  <NavBar />
  <main class="container mx-auto py-2">
    <RouterView />
  </main>

  <footer>
    <span class="content">Todos los derechos reservados&copy;</span>
  </footer>
</template>
<script setup>
import { RouterView } from "vue-router";
import NavBar from "./components/navBar/Navbar.vue";
</script>

<style>
body {
  background-color: rgb(231 231 231);
  font-family: "Helvetica", "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
}

main {
  min-height: 100vh;
}

footer {
  color: white;
  background-color: #1f2937;
  height: 60px;
  margin-top: auto;
  border-top-left-radius: 10px;
  border-top-right-radius: 10px;
  text-align: center;
}

.content {
  text-align: end;
  padding: 10px;
}
</style>
