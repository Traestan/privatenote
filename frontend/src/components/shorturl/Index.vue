<template>
    <div>
      <div class="text-center">
        <h1>Список созданных ссылок</h1>

        <div v-if="shorturls.length === 0">
          <h2> No shorturls found at the moment </h2>
        </div>
        
      </div>
      
      <div class="container"> 
        <div class="row">
          <div class="col-md-4" v-for="shorturl in shorturls" :key="shorturl.ShortUrl">
            <div class="card mb-4 shadow-sm">
              <div class="card-body">
                <h2 class="card-img-top">{{ shorturl.ShortUrl}}</h2>
                <router-link :to="{name: 'EditShorturl', params: {shorturl: shorturl.ShortUrl}}" class="btn btn-sm btn-outline-secondary">Изменить </router-link>
                <!--
                <router-link :to="{name: 'goto', params: {shorturl: shorturl.ShortUrl}}" class="btn btn-sm btn-outline-success">Просмотр </router-link>
                -->
                
                <b-link @click='goto(shorturl.ShortUrl)' class="btn btn-sm btn-outline-success">Просмотр</b-link>

                <p class="card-text">{{ shorturl.SourceUrl}}</p>
                <p class="card-text">{{ shorturl.Date_Add}}</p>
                <p class="card-text"><router-link :to="{name: 'StatUrl', params: {shorturl: shorturl.ShortUrl}}">Аналитика </router-link></p>
                <p class="card-text"><router-link :to="{name: 'DelUrl', params: {shorturl: shorturl.ShortUrl}}">Удалить</router-link></p>
                <!--
                <p class="card-text"><router-link :to="{name: 'Shorturl', params: {id: shorturl.id}}">Записи в категории </router-link></p>
                -->
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
        .get(`${server.baseURL}/urls`)
        .then(data => (this.shorturls = data.data.Result));
    },
    goto(url){
      window.location.href=this.baseURL+"/"+url;
    }
  }
};
</script>