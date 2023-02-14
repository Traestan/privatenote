<template>
   <div class="container">
        <div class="col-md-8 form-wrapper">
            <h2> Авторизация пользователя </h2>
            <form id="login-form" @submit.prevent="loginUser">
                <div div class="alert alert-danger" role="alert" v-if="errors.length">
                    <b>Пожалуйста исправьте указанные ошибки:</b>
                    <ul>
                        <li v-for="error in errors">{{ error }}</li>
                    </ul>
                </div>
                <div class="form-group col-md-12">
                    <label for="email"> Email </label>
                    <input type="text" id="email" v-model="email" name="email" class="form-control" placeholder="example@example">
               </div>
                <div class="form-group col-md-12">
                    <label for="password"> Password </label>
                    <input type="text" id="password" v-model="password" name="password" class="form-control" placeholder="password">
                </div>
                <div class="form-group col-md-12">
                  <div class="row">
                    <div class="form-group col-md-2 pull-right">
                        <button class="btn btn-info" type="submit"> Войти </button>
                    </div> 
                    <div class="form-group col-md-4 pull-right">
                      <b-link :to="{ path: 'forgot'}">Забыли пароль?</b-link>
                    </div>
                  </div> 
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
        email: "",
        password: "",
        errors:[],
    };
  },
  methods: {
    loginUser() {
      let userData = {
        email: this.email,
        password: this.password,
      };
      this.__submitToServer(userData);
    },
    __submitToServer(data) {
      
      axios.post(`${server.baseURL}/user/login`, data).then(data => {
        if (data.data.error.Status===0){
          this.errors=[];
          localStorage.setItem('user',data.data.data.user.email);
          localStorage.setItem('jwt',data.data.data.user.token);
          this.$cookie.set('token',data.data.data.user.token);

          if (localStorage.getItem('jwt') != null){
            this.$emit('loggedIn');
            if(this.$route.params.nextUrl != null){
              this.$router.push({ name: "home" });
            }else {
              this.$router.push({ name: "Notes"});
            }
          }
        }else{
          this.errors=[];
          this.errors.push('Не верный логин и/или пароль');
        }
      });
    }
  }
};
</script>