<template>
  <div class="text-center">
    <div class="col-sm-12">
      <h4 style="margin-top: 30px;">
        <button class="btn btn-success" v-on:click="navigate()"> Все ссылки</button>
        </h4>
      <hr>
      <h2> {{ shorturl.Shorturl.SourceUrl }} </h2>
      <span class="glyphicon glyphicon-time"></span> {{shorturl.Shorturl.Date_Add}}.
      <b-table striped hover :items="shorturl.Statist"></b-table>
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
      shorturl: {}
    };
  },
  created() {
    this.id = this.$route.params.shorturl;
    this.getPost();
    //console.log(this.post);
  },
  methods: {
    getPost() {
      axios
        .get(`${server.baseURL}/stat/${this.id}`)
        .then(data => (this.shorturl = data.data));
    },
    navigate() {
      router.go(-1);
    }
  }
};
</script>