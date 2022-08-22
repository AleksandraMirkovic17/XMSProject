<template>
  <div class="friends" v-if="loggedUser">
    <div>
      <h4 class="title">Blocked users ({{blockeds.length}})</h4>
    </div>
    <div class="friends">
      <div v-for="(user,index) in blockeds" :key="index">
        <div class="panel profile-card-small" v-on:click="redirectToProfile(user)"  style="cursor: pointer; display: flex; flex-direction: row">
          <div class="profile-icon" style="width: 20%">
            <div class="photo-container-small">
              <p style="position: relative; align-content: center; margin: 13%; font-weight: bold;" v-if="user && user.name && user.surname">
                {{user.name.charAt(0).toUpperCase()}}{{user.surname.charAt(0).toUpperCase()}}
              </p>
            </div>
          </div>
          <div class="info">
            <div  style="display: flex; flex-direction: row;  margin-top: 2%; ">
              <p style="margin-left: 1%; font-size: 130%;">{{user.name}}  </p>
              <p style="margin-left: 1%; font-size: 130%">{{user.surname}}</p>
            </div>
            <p class="username">@{{user.username}}</p>
          </div>
        </div>
        <!--<div class="profile-container" v-on:click="redirectToProfile(user)" style="cursor: pointer">
          <div class="profile-icon" style="width: 10%">
            <svg xmlns="http://www.w3.org/2000/svg" width="30" height="30" fill="currentColor" class="bi bi-person-fill" viewBox="0 0 16 16">
              <path d="M3 14s-1 0-1-1 1-4 6-4 6 3 6 4-1 1-1 1H3zm5-6a3 3 0 1 0 0-6 3 3 0 0 0 0 6z"/>
            </svg>
          </div>
          <div class="info" style="display: flex; flex-direction: row;width: 90%">
            <h3 style="margin-left: 1%">{{user.name}} {{user.surname}} </h3>
            <h3>({{user.username}})</h3>
          </div>
        </div>-->

      </div>
    </div>
  </div>
</template>

<script>
import UserService from "../../services/UserService";
import ConnectionService from "../../services/ConnectionService";
export default {
  name: "Blockeds",
  data(){
    return{
      loggedUser: "",
      loggedUserDetails: "",
      blockeds: new Array()
    }
  },
  mounted: function () {
    this.loggedUser = localStorage.getItem('user')
    UserService.getUserByUsername(JSON.parse(this.loggedUser).username).then(response =>{
      this.loggedUserDetails = response.data.User
      ConnectionService.GetBlockeds(this.loggedUserDetails.id).then(response1 =>{
        console.log("blockeds", response1.data.users)
        var friendNodes =  response1.data.users
        for (var f of friendNodes){
          UserService.getUserById(f.userID)
              .then(response2 =>{
                console.log(response2.data)
                this.blockeds.push(response2.data.User)
              })
              .catch(err =>{
                console.log("User with id: ", f.userID, " does not exists in UserService, just in ConnectionService!", err)
              })
        }

      })
          .catch(err1 =>{
            alert("It is impossible to load blocked users!")
            console.log(err1)
          })
    })
        .catch(err =>{
          alert("It s impossible to get logged user!")
          console.log(err)
        })

  },
  methods:{
    redirectToProfile(user){
      this.$router.push("/profile/"+user.username)
    }
  }
}
</script>

<style scoped>

.profile-container{
  border-radius: 20px;
  border-style: solid;
  border-width: 1pt;
  border-color: #e5e5e5;
  backround-color: white;
  padding: 1%;
  color: #e5e5e5;
  display: flex;
  flex-direction: row;
  horiz-align: center;
  vertical-align: center;
}

</style>