<template>
  <div id="app">
    <div id="nav">
      <b-navbar toggleable="lg" type="dark" variant="info">
          <b-collapse id="nav-collapse" is-nav>
            <b-navbar-nav>
              <b-nav-item to="/">Home</b-nav-item>

              <b-nav-item v-if="auth!==true" to="/login">Вход</b-nav-item>
              <b-nav-item v-if="auth!==true" to="/register">Регистрация</b-nav-item>
              
              <b-nav-item v-if="auth==true" href="/create" >Добавить запись</b-nav-item>
              <b-nav-item v-if="auth==true" href="/list/url" >Список заметок</b-nav-item>
              <b-nav-item-dropdown v-if="auth==true">
                <template v-slot:button-content>
                  <em>Профиль</em>
                </template>
                <b-dropdown-item href="/user/changepass">Пароль</b-dropdown-item>
              </b-nav-item-dropdown>

              <b-nav-item v-if="auth==true" href="/logout" >Выход</b-nav-item>
              
            </b-navbar-nav>
          </b-collapse>
      </b-navbar>
    </div>
    <router-view/>
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
      auth:false,
    };
  },
  created() {
    this.auth=false;
    if(localStorage.getItem('jwt') == null){
      this.auth=false;
    }else{
      this.auth=true;
    }
  },
  methods: {
    logout() {
      let userData = {
        email: this.email,
        password: this.password,
      };
      this.__submitToServer(userData);
    }
  }
};
</script>