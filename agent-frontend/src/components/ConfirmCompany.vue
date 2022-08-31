<template>
    <div  class="company-reservations">
  <table class="table table-striped table-bordered">
            <thead>
                <tr>
                    <th>Company name</th>
                    <th>Username</th>
                    <th>Request</th>
                </tr>
            </thead>
            <tbody v-for="c in companies" :key="c.id">
                <tr v-if="c.validated!=true">
                    <td>{{c.name}}</td>
                    <td>{{c.owner.username}}</td>
                    <td>          <button type="button" class="btn btn-secondary"  v-on:click="Request(c)">Accept</button>
</td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script>

import axios from 'axios'
import {devServer} from "../../vue.config";

export default{
    data(){
        return{
           companies: null,
           loggedUser: null,
           token: ''
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
      axios.get(devServer.proxy+'api/company')
      .then(response1=>{
        this.companies=response1.data;
         console.log("Ovo su sve kompanije:", this.companies)
      })
    })
    },
    methods:
    {
      Request(company)
      {
           let token1 = localStorage.getItem('token').substring(1, localStorage.getItem('token').length-1);

        console.log(company.id)
        let id=company.id
      axios.put(devServer.proxy+'api/company/validate/'+id, {
      headers: {
        'Authorization' : 'Bearer ' + token1,
      }
    })
      .then(response=>
      {
        console.log(response.data)
      })
      }
    }

}
</script>


<style scoped>
.company-reservations{
  margin-bottom: 20%;
  margin-left: 15%;
  margin-right: 10%;
  margin-top: 10%;
}

</style>