<template>
    <div>
      <div class="text-center">
        <h1>Short Service User Auth</h1>
        <p> ver 0.8 This is the description of the blog built with Golang, Vue.js and MongoDB</p>
       
      </div>
      <div class="container">
        <div class="row">
        </div>
      </div>
    </div>
</template>

<script>
// @ is an alias to /src
import { server } from "@/utils/helper";
import axios from "axios";

export default {
  metaInfo: {
    title: 'Vue App',
        meta: [
          { vmid: 'description', property: 'description', content: 'Vue App' },
          { vmid: 'og:title', property: 'og:title', content: 'Vue App' },
          { vmid: 'og:description', property: 'og:description', content: 'Vue App' },
        ],
  },
  data() {
    return {
      posts: [],
      
    };
  },
  created() {
    //this.fetchPosts();
  },
  methods: {}
};
</script>