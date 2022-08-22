<template>
<div class="friends" v-if="loggedUser" style=" width: 100%;">
  <div>
    <h4 class="title">Friends ({{friends.length}})</h4>
  </div>
  <div class="friends">
    <div v-for="(user,index) in friends" :key="index">
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

    </div>
  </div>
</div>
</template>

<script>
import UserService from "../../services/UserService";
import ConnectionService from "../../services/ConnectionService";
export default {
  name: "Friends",
  data(){
    return{
      loggedUser: "",
      loggedUserDetails: "",
      friends: new Array()
    }
  },
  mounted: function () {
    this.loggedUser = localStorage.getItem('user')
    UserService.getUserByUsername(JSON.parse(this.loggedUser).username).then(response =>{
      this.loggedUserDetails = response.data.User
      ConnectionService.GetFriends(this.loggedUserDetails.id).then(response1 =>{
          console.log("friends", response1.data.users)
        var friendNodes =  response1.data.users
        for (var f of friendNodes){
          UserService.getUserById(f.userID)
          .then(response2 =>{
            console.log(response2.data)
            this.friends.push(response2.data.User)
          })
          .catch(err =>{
            console.log("User with id: ", f.userID, " does not exists in UserService, just in ConnectionService!", err)
          })
        }

      })
      .catch(err1 =>{
        alert("It is impossible to load friends!")
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

.profiles-view{
  margin: 2%;
}

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