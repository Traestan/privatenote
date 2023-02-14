<template>
 <div class="container">
  <h4 class="text-center mt-20">
   <small>
    <button class="btn btn-success" v-on:click="navigate()"> Ссылки </button>
   </small>
  </h4>
  <div class="col-md-8 form-wrapper">
   <h2> Редактирование ссылки </h2>
   <form id="edit-shortslug-form" @submit.prevent="editShortslug">
    <div class="form-group col-md-12">
     <label for="slug"> Slug </label>
     <input type="text" id="slug" v-model="category.SourceUrl" name="slug" class="form-control">
    </div>
    <div class="form-group col-md-4 pull-right">
     <button class="btn btn-success" type="submit"> Изменить </button>
    </div>
   </form>
  </div>
 </div>
</template>

<script>
import { server } from "../../utils/helper";
import axios from "axios";
import router from "../../router";
export default {
  data() {
    return {
      id: 0,
      category: {}
    };
  },
  created() {
    this.id = this.$route.params.shorturl;
    this.getShortslug();
  },
  methods: {
    editShortslug() {
      let shortslugData = {
        title: this.category.title,
        description: this.category.description,
        body: this.category.body,
        slug: this.category.SourceUrl
      };
      axios
        .put(`${server.baseURL}/edit/${this.id}`, shortslugData)
        .then(data => {
          router.push({ name: "Shorturls" });
        });
    },
    getShortslug() {
      axios
        .get(`${server.baseURL}/get/${this.id}`)
        .then(data => (this.category = data.data.Result));
    },
    navigate() {
      router.go(-1);
    }
  }
};
</script>