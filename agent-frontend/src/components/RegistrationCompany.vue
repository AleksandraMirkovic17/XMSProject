<template class="tem">
  <div class="adventure-registration">
    <form>
      <h3 class="mb-3">Company registration</h3>
      <div class="col-11">
        <label for="adventure-name" class="form-label">Company name</label>
        <input type="text" class="form-control" id="adventure-name" placeholder="E.g. Titanic" v-model = "name" required>
      </div>

      <div class="col-11">
        <label for="adventure-name" class="form-label">Company description</label>
        <input type="text" class="form-control" id="adventure-name" placeholder="E.g. Titanic" v-model = "description" required>
      </div>

      <div class="col-11">
        <label for="adventure-name" class="form-label">Company contact info</label>
        <input type="text" class="form-control" id="adventure-name" placeholder="E.g. Titanic" v-model = "contactInfo" required>
      </div>
      
<hr>
      <button class="w-100 btn btn-primary btn-lg" type="button" v-on:click="Registration">Register a company </button>
    </form>

  </div>


</template>
<script>

import axios from 'axios'
import {devServer} from "../../vue.config";

export default{
    data(){
        return{
          name: '',
          description: '',
          contactInfo: '',
          loggedUser: null
        }
    },
     mounted(){
    
   this.token = localStorage.getItem('token').substring(1, localStorage.getItem('token').length-1);
    axios.get(devServer.proxy+'api/auth/user_data', {
      headers: {
        'Authorization' : 'Bearer ' + this.token,
      }
    })
    .then(response => {
      this.loggedUser =response.data
      console.log("Ovaj user je ulogovan:", this.loggedUser.roles[0].name)
    })
    },
    methods:{
    Registration()
    {
      if(this.name=='' || this.contactInfo=='' || this.description=='')
      {
        alert('fill in all the fields')
        return;
      }
      axios.post(devServer.proxy+'api/company',
            {
              "name": this.name,
              "description": this.description,
              "contactInfo": this.contactInfo
            }
            )
            .then(response => {
              alert(response.data)
              
       });
    }
    }
}
</script>

<style scoped>
.adventure-registration{

  width: 60%;
  horiz-align: center;
  margin-left: 20%;
  margin-top: 2%;
  border-radius: 3%;
  padding: 3%;
  background-blend-mode: lighten;
}
.removeImageBtn{
    position: absolute;
}
.double-field{
  display: flex;
  flex-direction: row;
}

.double-field .form-check{
  margin-left: 0.5%;
}

.adventure-registration .form-label{
}



</style>