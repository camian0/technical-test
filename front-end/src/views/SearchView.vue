<template>
  <h1>Bienvenido a la página de buscar</h1>
  <SearchComponent />
</template>

<script>
import SearchComponent from "../components/Search/SearchComponent.vue";
export default {
  name: "SearchView",
  data() {
    return {};
  },
  components: {
    SearchComponent,
  },
};
</script>
