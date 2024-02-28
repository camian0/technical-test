<template>
  <div class="search my-3">
    <input
      class="rounded-md border-0 px-3.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 mx-2"
      type="text"
      v-model="model.search"
      autocomplete="given-name"
      placeholder="Ingresa un término para buscar"
    />

    <button
      type="submit"
      class="rounded-md bg-blue-400 mx-2 my-2.5 px-6 py-4 text-sm right-0 font-semibold text-white shadow-sm hover:bg-sky-600 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
      v-on:click="buscar"
    >
      Buscar
    </button>
  </div>
  {{ model.search }}
</template>

<script>
export default {
  name: "SearchComponent",

  data() {
    return {
      model: {
        search: "",
      },
      people: [
        {
          name: "Leslie Alexander",
          role: "Co-Founder / CEO",
          imageUrl:
            "https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80",
        },
        {
          name: "Juan ",
          role: "Co-Founder / CEO",
          imageUrl:
            "https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80",
        },
        {
          name: "Pepito",
          role: "Co-Founder / CEO",
          imageUrl:
            "https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80",
        },
      ],
    };
  },
  methods: {
    buscar() {
      console.log("dio click en el boton");
    },
  },
};
</script>

<style scoped>
.search {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}

.search input {
  min-width: 90%;
}
</style>
