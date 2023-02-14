<template>
   <div class="container">
        <div class="col-md-8 form-wrapper">
            <h2> Восстановление пароля </h2>
            <form id="login-form" @submit.prevent="loginUser">
              <div div class="alert alert-danger" role="alert" v-if="errors.length">
                <b>Пожалуйста исправьте указанные ошибки:</b>
                <ul>
                  <li v-for="error in errors">{{ error }}</li>
                </ul>
              </div>

              <div div class="alert alert-success" role="alert" v-if="send==1">
                <b>Новый пароль отправлен вам на email</b>
              </div>

                <div class="form-group col-md-12">
                    <label for="email"> Email </label>
                    <input type="text" id="email" v-model="email" name="email" class="form-control" placeholder="example@example">
                </div>
                
                <div class="form-group col-md-4 pull-right">
                    <button class="btn btn-info" type="submit"> Восстановить </button>
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
        errors:[],
        send:0,
    };
  },
  methods: {
    loginUser() {
      let userData = {
        email: this.email,
      };
      this.__submitToServer(userData);
    },
    __submitToServer(data) {
      
      axios.post(`${server.baseURL}/user/forgot`, data).then(data => {
        console.log(data);
        //router.push({ name: "home" });
        if (data.data.email===undefined){
          this.errors=[];
          this.errors.push('Email не существует или не валидный');
        }else{
          this.errors=[];
          this.send=1;
          console.log(data.data.email);
        }
        //this.$cookie.set('token',data.data.token);
        //console.log(data.headers);
      });
    }
  }
};
</script>