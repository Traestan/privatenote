<template>
 <div class="container">
  <h4 class="text-center mt-20">
   <small>
    <button class="btn btn-success" v-on:click="navigate()"> View All Posts </button>
   </small>
  </h4>
  <div class="col-md-8 form-wrapper">
   <h2> Изменить запись </h2>
   <form id="edit-post-form" @submit.prevent="editPost">
    <div class="form-group col-md-12">
     <label for="title"> Title </label>
     <input type="text" id="title" v-model="post.title" name="title" class="form-control" placeholder="Enter title">
    </div>
    <div class="form-group col-md-12">
     <label for="description"> Description </label>
     <input type="text" id="description" v-model="post.description" name="description" class="form-control" placeholder="Enter Description">
    </div>
    <div class="form-group col-md-12">
     <label for="text"> Текст </label>
     <textarea id="text" cols="30" rows="5" v-model="post.text" class="form-control"></textarea>
    </div>
    <div class="form-group col-md-12">
     <label for="ttl"> Время жизни </label>
     <!-- <input type="text" id="ttl" v-model="post.ttl" name="ttl" class="form-control"> -->
     <select v-model="post.ttl" class="form-control" >
        <option v-for="(park,index) in categories" v-bind:key="park" v-bind:value="index" selected="key == post.ttl">
          {{park}}
        </option>
      </select>
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
      post: {},
      ttls:this.getTtls()
    };
  },
  created() {
    this.id = this.$route.params.id;
    this.getPost();
  },
  methods: {
    editPost() {
      let postData = {
        title: this.post.title,
        description: this.post.description,
        text: this.post.text,
        ttl: this.post.ttl,
      };
      axios
        .put(`${server.baseURL}/note/edit/${this.id}`, postData)
        .then(data => {
          router.push({ name: "home" });
        });
    },
    getPost() {
      axios
        .get(`${server.baseURL}/note/get/${this.id}`)
        .then(data => (this.post = data.data.data.note));
    },
    getTtls() {
      axios
        .get(`${server.baseURL}/service/ttl`)
        .then(data => {
          this.categories = data.data.data;

        });
    },
    navigate() {
      router.go(-1);
    }
  }
};
</script>