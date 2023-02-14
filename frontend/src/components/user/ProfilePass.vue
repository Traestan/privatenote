<template>
   <div class="container">
        <div class="col-md-8 form-wrapper">
            <h2> Смена пароля </h2>
            <form id="create-user-form" @submit.prevent="createUser">
                <div div class="alert alert-danger" role="alert" v-if="errors.length">
                    <b>Пожалуйста исправьте указанные ошибки:</b>
                    <ul>
                        <li v-for="error in errors">{{ error }}</li>
                    </ul>
                </div>
                <div class="form-group col-md-12">
                    <label for="password"> Пароль </label>
                    <input type="text" id="password" v-model="password" name="password" class="form-control" placeholder="Enter password">
                </div>

                <div class="form-group col-md-12">
                    <label for="newpass"> Новый пароль </label>
                    <input type="text" id="newpass" v-model="newpass" name="newpass" class="form-control" placeholder="Enter newpass">
                </div>

                <div class="form-group col-md-12">
                    <label for="newpasscompare"> Повторите новый пароль </label>
                    <input type="text" id="newpasscompare" v-model="newpasscompare" name="newpasscompare" class="form-control" placeholder="Enter newpasscompare">
                </div>

                <div class="form-group col-md-4 pull-right">
                    <button class="btn btn-success" type="submit"> Изменить </button>
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
      newpass: "",
      password: "",
      newpasscompare:"",
      errors:[],
    };
  },
  created() {
    //this.date_posted = new Date().toLocaleDateString();
  },
  methods: {
    createUser() {
      let userData = {
        newpass: this.newpass,
        password: this.password,
        newpasscompare: this.newpasscompare,
      };
      this.__submitToServer(userData);
    },
    __submitToServer(data) {
      axios.post(`${server.baseURL}/user/changepass`, data).then(data => {
        if(data.data.error.Status==0){
          this.errors=[];
          router.push({ name: "home" });
        }else{
          this.errors=[];
          this.errors.push('newpass не существует или не валидный');
        }
      });
    }
  }
};
</script>