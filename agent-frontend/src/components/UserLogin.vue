<template>
<section class="bg-dark text-light p-5 text-center text-sm-start">
<div id="loginForm">

    <br>
    <br>
    <br>
    <br>
    <br>
    <br>
    <br>
    <br>
   <h1>Login</h1>
   <br>
   <input placeholder="Email" class="inputKredencijali" type="text" v-model="email"/>
   <br>
   <input placeholder="Password" class="inputKredencijali" type="password" v-model="password"/>
   <br>
   <br>
   <br>
   <button class="buttonLogin"  v-on:click="Login">Login</button>
    <br>
    
    <br>
    <br>
    <br>
    <br>
    <br>
    <br>
    <br>
    <br>
</div>
</section>
</template>
<script>
import axios from 'axios'


//import {onBeforeMount} from "vue";
export default{
    data(){
        return{
            userType: '',
            email: '',
            password: '',
            fieldEmpty: false,
            show:false,
            loggedUser:null
        }
    },
    methods:{
    
      Login(){
        localStorage.removeItem('token');
        localStorage.clear();
  
        this.CheckIfEmpty()
        if(!this.fieldEmpty){
          axios
              .post('http://localhost:8080/auth/login',
              {
                "email": this.email,
                "password": this.password
              })
              .then(response => {
                if(!response.data){
                  alert('Bad username or password')
                  return
                }
                    localStorage.setItem('token', JSON.stringify(response.data.accessToken));
                    let token = localStorage.getItem('token').substring(1, localStorage.getItem('token').length-1);
                     axios.get('http://localhost:8080/auth/userData', {
                    headers: {
                    'Authorization' : 'Bearer ' + token,
                  }
                    })
               .then(response1 => {
                     this.loggedUser =response1.data
                     console.log(response1.data.userType)
              })
              })
          .catch((err)=>
          alert(err))
        }else alert('Wrong username or password!');

      },

      CheckIfEmpty(){
        if(this.email === '' || this.password === '') this.fieldEmpty = true;
      }
      
    }
}
</script>
<style>

  h1{
    text-align: center;
  }

  #loginForm{
      text-align:center;
  }

  #loginForm input{
    
  width: 300px;
	margin:5px 0;
	padding:10px;
	border-radius:20px;
	border: 2px solid #eee;
	box-shadow:0 0 15px 4px rgba(0,0,0,0.06);
  }

  input:focus{
	background: #FCFCFC;
	outline: none;
	
  }
  
 .buttonLogin {
	padding:10px;
	border:none;
	background-color:#2ECC71;
	color:#fff;
	font-weight:600;
	border-radius:20px;
	width:300px;
  }
  .buttonLogin:hover{
	background: #333;
  }
</style>