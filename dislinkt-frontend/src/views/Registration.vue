<template>
<div class="registration">
  <form>
    <h1>
      Join Dislinkt today!
    </h1>
    <br>
    <div class="input-group mb-3">
      <span class="input-group-text" id="basic-addon1">Name</span>
      <input type="text" class="form-control" placeholder="Name" aria-label="Name" aria-describedby="basic-addon1" v-model="user.name">
    </div>
    <div class="input-group mb-3">
      <span class="input-group-text" id="basic-addon2">Surname</span>
      <input type="text" class="form-control" placeholder="Surname" aria-label="Surname" aria-describedby="basic-addon1" v-model="user.surname">
    </div>
    <div class="input-group mb-3">
      <span class="input-group-text" id="basic-addon3">Username</span>
      <input type="text" class="form-control" placeholder="Userrname" aria-label="Username" aria-describedby="basic-addon1" v-model="user.username">
    </div>
    <div class="input-group mb-3">
      <span class="input-group-text" id="basic-addon4">Email</span>
      <input type="text" class="form-control" placeholder="email@example.com" aria-label="Email" aria-describedby="basic-addon1" v-model="user.email">
    </div>
    <div class="input-group mb-3">
      <span class="input-group-text" id="basic-addon5">Phone number</span>
      <input type="text" class="form-control" placeholder="838942868109041" aria-label="Surname" aria-describedby="basic-addon1" v-model="user.contactPhone">
    </div>
    <div class="input-group mb-3">
      <span class="input-group-text" id="basic-addon6">Date of birth</span>
      <input type="date" class="form-control"  aria-describedby="basic-addon1" v-model="user.date">
    </div>
    <div class="input-group mb-3">
      <span class="input-group-text" id="basic-addon7">Gender</span>
      <select class="form-select" aria-label="Default select example" v-model="user.gender">
        <option selected>Select</option>
        <option value="male">Male</option>
        <option value="female">Female</option>
        <option value="other">Other</option>
      </select>
    </div>
    <div class="input-group mb-3">
      <span class="input-group-text" id="basic-addon8">Password</span>
      <input type="password" class="form-control" aria-label="Surname" aria-describedby="basic-addon1" v-model="user.password">
    </div>
    <div class="input-group mb-3">
      <span class="input-group-text" id="basic-addon9">Repeat password</span>
      <input type="password" class="form-control" aria-label="Surname" aria-describedby="basic-addon1">
    </div>
    <br>
    <button type="button" class="btn btn-dark" style="width: 100%; height: 60px" v-on:click="registerUser()">Register</button>
    <br>
    <p style="font-size: 12pt">
      Already have an account? <a href="/login">Log in.</a>
    </p>


  </form>

</div>
</template>

<script>

import UserService from '../services/UserService';

export default {
  name: "Registration",
  data(){
      return {
          user: {}
      }
  },
  methods: {
      registerUser(){
          UserService.registerUser(this.user);
      }
  }
}
</script>

<style scoped>

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