<template>
   <div class="container">
        <div class="col-md-8 form-wrapper">
          <h2> Регистрация </h2>
          <form id="create-user-form" @submit.prevent="createUser">
              <div div class="alert alert-danger" role="alert" v-if="errors.length">
                <b>Пожалуйста исправьте указанные ошибки:</b>
                <ul>
                  <li v-for="error in errors">{{ error }}</li>
                </ul>
              </div>
               <div class="form-group col-md-12">
                <label for="email"> Email </label>
                <input type="text" id="email" v-model="email" name="email" class="form-control" placeholder="Enter email">
               </div>
              <div class="form-group col-md-12">
                  <label for="password"> Пароль </label>
                  <input type="text" id="password" v-model="password" name="password" class="form-control" placeholder="Enter password">
              </div>
              <div class="form-group col-md-4 pull-right">
                  <button class="btn btn-success" type="submit"> Зарегистрироваться </button>
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
  created() {
    //this.date_posted = new Date().toLocaleDateString();
  },
  methods: {
    createUser() {
      let userData = {
        email: this.email,
        password: this.password,
      };
      this.__submitToServer(userData);
    },
    __submitToServer(data) {
      //const str = JSON.stringify(this.postBody);
      axios.post(`${server.baseURL}/user/register`, data).then(data => {
        //console.log(data);
        //console.log(data.data.error);
        /*if(data.data.error!==null){
          this.errors=[];
          this.errors.push('Email существует или не валидный');
        }else{
          this.errors=[];
          router.push({ name: "home" });
        }*/
        if (data.data.error===undefined){
          this.errors=[];
          router.push({ name: "home" });
        }else{
          this.errors=[];
          this.errors.push('Email не существует или не валидный');
        }
      });
    }
  }
};
</script>