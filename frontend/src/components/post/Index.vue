<template>
    <div>
      <div class="text-center">
        <h1>Список созданных заметок</h1>

        <div v-if="shorturls.length === 0">
          <h2> No shorturls found at the moment </h2>
        </div>
        
      </div>
      
      <div class="container"> 
        <div class="row">
          <div class="col-md-4" v-for="item in shorturls" >
            <div class="card mb-4 shadow-sm">
              <div class="card-body">
                <h2 class="card-img-top">{{item.title}}</h2>
                <router-link :to="{name: 'Edit', params: {id: item.number}}" class="btn btn-sm btn-outline-secondary">Изменить </router-link>
                <p class="card-text">Ttl: {{item.ttl}}</p>
                
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
</template>

<script>
// @ is an alias to /src
import { server } from "@/utils/helper";
import axios from "axios";

export default {
  data() {
    return {
      shorturls: [],
      baseURL:server.baseURL,
    };
  },
  created() {
    this.fetchShorturls();
  },
  methods: {
    fetchShorturls() {
      axios
        .get(`${server.baseURL}/note/list`)
        .then(data => {
            this.shorturls = data.data.data.note});
    },
    goto(url){
      window.location.href=this.baseURL+"/"+url;
    }
  }
};
</script>