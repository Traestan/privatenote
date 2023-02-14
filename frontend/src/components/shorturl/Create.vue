<template>
  <div class="container">
    <div class="col-md-8 form-wrapper">
      <h2> Добавить ссылку </h2>
      <form id="create-shorturls-form" @submit.prevent="createShorturls" enctype="multipart/form-data">
        <div class="form-group col-md-12">
          <label for="slug"> Slug </label>
          <input type="text" id="slug" v-model="slug" name="slug" class="form-control">
        </div>
        
        <div class="form-group col-md-4 pull-right">
          <button class="btn btn-success" type="submit"> Добавить </button>
        </div>          
      </form>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import { server } from "../../utils/helper";
import router from "../../router";
export default {
  data() {
    return {
      slug: "hello",
    };
  },
  methods: {
    createShorturls() {
      let shorturlsData = {
        title: this.title,
        description: this.description,
        body: this.body,
        slug: this.slug,
        //date_created: this.date_posted
      };
      this.__submitToServer(shorturlsData);
    },
    __submitToServer(data) {
      //console.log(data);
      axios.post(`${server.baseURL}/add`, data).then(data => {
        console.log(data);
        router.push({ name: "home" });
      });
    },
    onFileChange(e) {
      var files = e.target.files || e.dataTransfer.files;
      if (!files.length)
        return;
      //this.createImage(files[0]);
    }

  }
};
</script>