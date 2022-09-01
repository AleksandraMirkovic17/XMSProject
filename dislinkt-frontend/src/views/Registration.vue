<template>
  <div class="root">
    <div class="registration">
      <form>
        <h1>
          Join Dislinkt today!
        </h1>
        <br>
        Name
        <div class="input-group mb-3">
          <input type="text" class="form-control" placeholder="Name" aria-label="Name" aria-describedby="basic-addon1" v-model="user.user.name">
        </div>
        Surname
        <div class="input-group mb-3">
          <input type="text" class="form-control" placeholder="Surname" aria-label="Surname" aria-describedby="basic-addon1" v-model="user.user.surname">
        </div>
        Username
        <div class="input-group mb-3">
          <input type="text" class="form-control" placeholder="Userrname" aria-label="Username" aria-describedby="basic-addon1" v-model="user.user.username">
        </div>
        Email
        <div class="input-group mb-3">
          <input type="text" class="form-control" placeholder="email@example.com" aria-label="Email" aria-describedby="basic-addon1" v-model="user.user.email">
        </div>
        Phone
        <div class="input-group mb-3">
          <input type="text" class="form-control" placeholder="838942868109041" aria-label="Surname" aria-describedby="basic-addon1" v-model="user.user.contactPhone">
        </div>
        Date of birth
        <div class="input-group mb-3">
          <input type="date" class="form-control"  aria-describedby="basic-addon1" v-model="user.user.date">
        </div>
        Gender
        <div class="input-group mb-3">
          <select class="form-select" aria-label="Default select example" v-model="user.user.gender">
            <option selected>Select</option>
            <option value=1>Male</option>
            <option value=2>Female</option>
            <option value=3>Other</option>
          </select>
        </div>
        Password
        <div class="input-group mb-3">
          <input type="password" class="form-control" aria-label="Surname" aria-describedby="basic-addon1" v-model="user.user.password">
        </div>
        Repeat password
        <div class="input-group mb-3">
          <input type="password" class="form-control" aria-label="Surname" aria-describedby="basic-addon1" v-model="user.user.repeatPassword">
        </div>
        <span class="switch">
          <input type="checkbox" class="switch" id="isPublic" v-model="user.user.public">
          <label for="isPublic" v-if="user.user.public">Public</label>
          <label for="isPublic" v-if="!user.user.public">Private</label>
        </span>
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
              "dateOfBirth": this.user.user.date,
              "password": this.user.user.password,
              "public": this.user.user.public
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
  background-color: #ffffff;
  border-radius: 20px;
  border: 1pt solid red;
  margin: 30%;
  margin-top: 2%;
  padding: 5%;
  align: center;
  vertical-align: center;
  align-content: center;
}

.registration form{
  width: 100%;
  align: center;
}

/* SWITCH */

.switch {
    font-size: 1rem;
    position: relative;
}
.switch input {
    position: absolute;
    height: 1px;
    width: 1px;
    background: none;
    border: 0;
    clip: rect(0 0 0 0);
    clip-path: inset(50%);
    overflow: hidden;
    padding: 0;
}
.switch input + label {
    position: relative;
    min-width: calc(calc(2.375rem * .8) * 2);
    border-radius: calc(2.375rem * .8);
    height: calc(2.375rem * .8);
    line-height: calc(2.375rem * .8);
    display: inline-block;
    cursor: pointer;
    outline: none;
    user-select: none;
    vertical-align: middle;
    text-indent: calc(calc(calc(2.375rem * .8) * 2) + .5rem);
}
.switch input + label::before, .switch input + label::after {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    width: calc(calc(2.375rem * .8) * 2);
    bottom: 0;
    display: block;
}
.switch input + label::before {
    right: 0;
    background-color: #dee2e6;
    border-radius: calc(2.375rem * .8);
    transition: 0.2s all;
}
.switch input + label::after {
    top: 2px;
    left: 2px;
    width: calc(calc(2.375rem * .8) - calc(2px * 2));
    height: calc(calc(2.375rem * .8) - calc(2px * 2));
    border-radius: 50%;
    background-color: white;
    transition: 0.2s all;
}
.switch input:checked + label::before {
    background-color: #e56830;
}
.switch input:checked + label::after {
    margin-left: calc(2.375rem * .8);
}
.switch input:focus + label::before {
    outline: none;
    box-shadow: 0 0 0 0.2rem #f34d0098;
}
.switch input:disabled + label {
    color: #868e96;
    cursor: not-allowed;
}
.switch input:disabled + label::before {
    background-color: #e9ecef;
}
.switch.switch-sm {
    font-size: 0.875rem;
}
.switch.switch-sm input + label {
    min-width: calc(calc(1.9375rem * .8) * 2);
    height: calc(1.9375rem * .8);
    line-height: calc(1.9375rem * .8);
    text-indent: calc(calc(calc(1.9375rem * .8) * 2) + .5rem);
}
.switch.switch-sm input + label::before {
    width: calc(calc(1.9375rem * .8) * 2);
}
.switch.switch-sm input + label::after {
    width: calc(calc(1.9375rem * .8) - calc(2px * 2));
    height: calc(calc(1.9375rem * .8) - calc(2px * 2));
}
.switch.switch-sm input:checked + label::after {
    margin-left: calc(1.9375rem * .8);
}
.switch.switch-lg {
    font-size: 1.25rem;
}
.switch.switch-lg input + label {
    min-width: calc(calc(3rem * .8) * 2);
    height: calc(3rem * .8);
    line-height: calc(3rem * .8);
    text-indent: calc(calc(calc(3rem * .8) * 2) + .5rem);
}
.switch.switch-lg input + label::before {
    width: calc(calc(3rem * .8) * 2);
}
.switch.switch-lg input + label::after {
    width: calc(calc(3rem * .8) - calc(2px * 2));
    height: calc(calc(3rem * .8) - calc(2px * 2));
}
.switch.switch-lg input:checked + label::after {
    margin-left: calc(3rem * .8);
}
.switch + .switch {
    margin-left: 1rem;
}


</style>