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
      v-on:click="search"
    >
      Buscar
    </button>
  </div>

  <section
    class="card relative isolate overflow-hidden px-6 py-24 sm:py-32 lg:px-8 rounded-md mx-2 my-2"
    v-if="tableData.length == 0"
  >
    <div class="mx-auto max-w-2xl lg:max-w-4xl">
      <p class="flex justify-center">
        La tabla no contiene información para mostrar
      </p>
    </div>
  </section>

  <div class="table mx-2 my-5" v-if="tableData.length != 0">
    <table class="table-fixed border-collapse">
      <thead>
        <tr>
          <th class="">Fecha</th>
          <th class="">Desde</th>
          <th class="">Para</th>
          <th class="">Asunto</th>
          <th class="">Contenido</th>
        </tr>
      </thead>
      <tbody>
        <tr class="even:bg-gray-50 odd:bg-gray-300" v-for="item in tableData">
          <td style="min-width: 205px" class="">{{ item.Date }}</td>
          <td class="">{{ item.From }}</td>
          <td class="">{{ item.To }}</td>
          <td style="min-width: 215px" class="">{{ item.Subject }}</td>
          <td class="">
            <pre>
            {{ item.Content }}
            </pre>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import { postData } from "../../Request/Request";
export default {
  name: "SearchComponent",
  data() {
    return {
      data: [],
      tableData: [],
      body: {
        search_type: "match",
        query: {
          term: "",
          field: "",
        },
        sort_fields: ["-@timestamp"],
        from: 0,
        max_results: 20,
        _source: [],
      },
      model: {
        search: "",
      },
      indexedDB: "enron",
    };
  },
  methods: {
    async search() {
      this.body.query.field = "content";
      this.body.query.term = "fuck";
      await postData(`api/${this.indexedDB}/_search`, this.body)
        .then((res) => {
          if (res != null) {
            this.data = res;
            console.log("data recibida", this.data);
          }

          this.processData(this.data.hits.hits);
        })
        .catch((err) => {
          console.log("error", err);
        });
    },

    processData(emails) {
      let processed = [];
      emails.forEach((item) => {
        let object = {
          Date: item._source.Date,
          From: item._source.From,
          To: item._source.To,
          Subject: item._source.Subject,
          Content: item._source.content,
        };

        processed.push(object);
      });

      this.tableData = processed;
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

table {
  background-color: white;
  border-radius: 5px;
  overflow-y: visible;
}

th {
  padding: 5px;
}

td {
  padding: 3px;
}

.card {
  background-color: rgb(255 255 255 / 16%);
}
</style>
