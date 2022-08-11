<template>
  <div class="root">
    <div class="registration">
      <form>
        <h1>
          Join Dislinkt today!
        </h1>
        <br>
        <div class="input-group mb-3">
          <span class="input-group-text" id="basic-addon1">Name</span>
          <input type="text" class="form-control" placeholder="Name" aria-label="Name" aria-describedby="basic-addon1" v-model="user.user.name">
        </div>
        <div class="input-group mb-3">
          <span class="input-group-text" id="basic-addon2">Surname</span>
          <input type="text" class="form-control" placeholder="Surname" aria-label="Surname" aria-describedby="basic-addon1" v-model="user.user.surname">
        </div>
        <div class="input-group mb-3">
          <span class="input-group-text" id="basic-addon3">Username</span>
          <input type="text" class="form-control" placeholder="Userrname" aria-label="Username" aria-describedby="basic-addon1" v-model="user.user.username">
        </div>
        <div class="input-group mb-3">
          <span class="input-group-text" id="basic-addon4">Email</span>
          <input type="text" class="form-control" placeholder="email@example.com" aria-label="Email" aria-describedby="basic-addon1" v-model="user.user.email">
        </div>
        <div class="input-group mb-3">
          <span class="input-group-text" id="basic-addon5">Phone number</span>
          <input type="text" class="form-control" placeholder="838942868109041" aria-label="Surname" aria-describedby="basic-addon1" v-model="user.user.contactPhone">
        </div>
        <div class="input-group mb-3">
          <span class="input-group-text" id="basic-addon6">Date of birth</span>
          <input type="date" class="form-control"  aria-describedby="basic-addon1" v-model="user.user.date">
        </div>
        <div class="input-group mb-3">
          <span class="input-group-text" id="basic-addon7">Gender</span>
          <select class="form-select" aria-label="Default select example" v-model="user.user.gender">
            <option selected>Select</option>
            <option value=1>Male</option>
            <option value=2>Female</option>
            <option value=3>Other</option>
          </select>
        </div>
        <div class="input-group mb-3">
          <span class="input-group-text" id="basic-addon8">Password</span>
          <input type="password" class="form-control" aria-label="Surname" aria-describedby="basic-addon1" v-model="user.user.password">
        </div>
        <div class="input-group mb-3">
          <span class="input-group-text" id="basic-addon9">Repeat password</span>
          <input type="password" class="form-control" aria-label="Surname" aria-describedby="basic-addon1" v-model="user.user.repeatPassword">
        </div>
        <div class="form-check form-switch" style="width: 180px">
          <input class="form-check-input"  style="width: 50px; height: 25px" type="checkbox" role="switch" id="flexSwitchCheckChecked" checked v-model="user.user.public">
          <label class="form-check-label" for="flexSwitchCheckChecked" style="font-weight: bolder; font-size: 18px">Public account</label>
        </div>
        <br>
        <p>If you select this option your account and posts will be visible to anyone! You can chenge at setting later!</p>
        <br>
        <button type="button" class="btn btn-dark" style="width: 100%; height: 60px" v-on:click="registerUser()">Register</button>
        <br>
        <p style="font-size: 12pt">
          Already have an account? <a href="/login">Log in.</a>
        </p>


      </form>

    </div>
  </div>

</template>

<script>

import UserService from '../services/UserService';

export default {
  name: "Registration",
  data(){
      return {
          user: {
            user : {
              /*{
    "user" : {
        "name": "Angelina",
        "surname": "Jolie",
        "username": "angelina",
        "email": "angelina@gmail.com",
        "contactPhone": "734848932894",
        "dateOfBirth": "2000-01-01T01:30:15.01Z",
        "password": "123"
    }
}*/
            }
          }
      }
  },
  methods: {
      registerUser(){
        if(this.user.user.password != this.user.user.repeatPassword){
          alert("Passwords are not same!")
          return
        }
        alert(this.user.user.name)
        alert(this.user.user.public)
          UserService.registerUser({
            "user" : {
              "name": this.user.user.name,
              "surname": this.user.user.surname,
              "username": this.user.user.username,
              "gender": 2,
              "email": this.user.user.email,
              "contactPhone": this.user.user.contactPhone,
              "dateOfBirth": "2000-01-01T01:30:15.01Z",
              "password": this.user.user.password,
              "Public": this.user.user.public
            }

            }
          );
      }
  }
}
</script>

<style scoped>

.root{
  margin: 0;

}

.registration{
  background-color: #e85a4f;
  border-radius: 20px;
  margin: 30%;
  margin-top: 2%;
  padding: 5%;
  horiz-align: center;
  vertical-align: center;
  align-content: center;
}

.registration form{
  width: 100%;
  horiz-align: center;
}

</style>