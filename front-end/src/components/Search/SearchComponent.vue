<template>
<!-- modal -->
<dialog id="favDialog">
    <div class="modal-content">
      <span class="title">Contenido del correo</span>
      <section class="headers">
        <p>
          <b>Desde:</b>{{ modal.from }}
        </p>
        <p>
          <b>Para:</b> {{ modal.to }}
        </p>
        <p>
          <b>Asunto:</b>{{ modal.subject }}
        </p>
        <p>
          <b>Fecha:</b> {{ modal.date }}
        </p>
      </section>
      <p>
        <b>Contenido del correo:</b>
        <pre>{{ modal.content }}</pre>
      </p>
      <section class="description"></section>
      <menu class="buttons">
        <button
          id="cancel"
          class="rounded-md bg-blue-400 mx-2 my-2.5 px-6 py-4 text-sm right-0 font-semibold text-white shadow-sm hover:bg-sky-600 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
          @click="closeModal"
        >
          Cerrar
        </button>
      </menu>
    </div>
  </dialog> 

  <!-- buscador -->
  <div class="search my-3">
    <div class="field-card">
      <select
        placeholder="Selecciona"
        id="field"
        class="select mt-3"
        v-model="model.fieldSearch"
      >
        <option selected>{{ this.model.fieldSearch }}</option>
        <option v-for="(v, k) in FIELDS" :value="k">{{ v }}</option>
      </select>
    </div>

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

  <!-- tabla -->
  <div class="table mx-2 my-5 w-full" v-if="tableData.length != 0">
    <table class="">
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
        <tr
          class="even:bg-gray-50 odd:bg-gray-300"
          v-for="item in tableData"
          v-on="item"
        >
          <td style="min-width: 205px" class="">{{ item.Date }}</td>
          <td class="">{{ item.From }}</td>
          <td class="">{{ item.To }}</td>
          <td style="width: 215px" class="">{{ item.Subject }}</td>
          <td class="">
            <menu>
              <button id="updateDetails" class="rounded-md bg-blue-400 mx-2 my-2.5 px-2 py-3 text-sm right-0 font-semibold text-white shadow-sm hover:bg-sky-600 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600" @click="openModal(item)">Ver detalles</button>
            </menu>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import { postData } from "../../Request/Request";
import ModalComponent from "../Modal.vue";
import { FIELDS } from "../../Config/Constants";
export default {
  name: "SearchComponent",
  setup() {
    return {
      FIELDS,
    };
  },  
  components: {
    ModalComponent,
  },
  data() {
    return {
      data: [],
      tableData: [],
      body: {
        search_type: "matchall",
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
        fieldSearch: "Escoje el campo de búsqueda",
      },
      indexedDB: "enron",
      modal: {
        from: "",
        to: "",
        date: "",
        subject: "",
        content: "",
      },
    };
  },
  methods: {
    openModal(email){
      let updateButton = document.getElementById("updateDetails");
      let cancelButton = document.getElementById("cancel");
      let favDialog = document.getElementById("favDialog");
      this.load(email);
      favDialog.showModal()
    },

    closeModal(){
      let updateButton = document.getElementById("updateDetails");
      let cancelButton = document.getElementById("cancel");
      let favDialog = document.getElementById("favDialog");
      this.clearModal();
      favDialog.close();
    },

    // modal() {
    //   var updateButton = document.getElementById("updateDetails");
    //   var cancelButton = document.getElementById("cancel");
    //   var favDialog = document.getElementById("favDialog");

    //   // Update button opens a modal dialog
    //   updateButton.addEventListener("click", function () {
    //     favDialog.showModal();
    //   });

    //   // Form cancel button closes the dialog box
    //   cancelButton.addEventListener("click", function () {
    //     favDialog.close();
    //   });
    // },

    async search() {
      this.body.query.field = this.model.fieldSearch;
      this.body.query.term = this.model.search;
      // let field = this.model.fieldSearch;
      // let search = this.model.search;

      // this.body.query_string = { query: `${field}:${search}` };
      // console.log("modelo", this.model);
      // console.log("body", this.body);
      await postData(`search`, this.body)
        .then((res) => {
          if (res != null) {
            this.data = res.data;
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

    load(email) {
      this.modal.to = email.To;
      this.modal.from = email.From;
      this.modal.subject = email.Subject;
      this.modal.date = email.Date;
      this.modal.content = email.Content;
    },

    clearModal(){
      this.modal= {
        from: "",
        to: "",
        date: "",
        subject: "",
        content: ""
      }
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
  min-width: 74%;
}

table {
  background-color: white;
  border-radius: 5px;
  width: 95%;
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
pre {
  white-space: pre-wrap; /* Since CSS 2.1 */
  white-space: -moz-pre-wrap; /* Mozilla, since 1999 */
  white-space: -pre-wrap; /* Opera 4-6 */
  white-space: -o-pre-wrap; /* Opera 7 */
  word-wrap: break-word; /* Internet Explorer 5.5+ */
}

#favDialog{
  width: 80%;
}

.field-card {
  display: flex;
  flex-direction: column;
}

.select {
  color: black;
  border-radius: 5px;
  height: 45px;
  min-width: 183px;
}

.modal-content {
  display: flex;
  flex-direction: column;
  padding: 10px;
  /* text-align: center; */
}

.headers {
  margin: 10px 0px;
}
.buttons {
  display: flex;
  justify-content: center;
}
.buttons button {
  margin: 5px 10px;
  padding: 10px 20px;
}

.title {
  text-align: center;
  font-size: 50px;
  font-weight: 600;
}
</style>
