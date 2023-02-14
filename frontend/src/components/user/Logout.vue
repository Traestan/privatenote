<template>
    
</template>
 
<script>
import axios from "axios";
import { server } from "../../utils/helper";
import router from "../../router";
export default {
    data (){
      return {
      }
    },
    created:function(){
        localStorage.removeItem("user");
        localStorage.removeItem("jwt");
        this.$cookie.delete('token');
    }
};
</script>