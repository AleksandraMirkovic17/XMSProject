<template>
  <div class="requests" v-if="loggedUser" style=" width: 100%;">

    <div>
      <h4 class="title">Friend requests ({{requests.length}})</h4>
    </div>
    <div class="friends">
      <div v-for="(user,index) in requests" :key="index">
        <div class="profile-container" >
          <div class="profile-informations" v-on:click="redirectToProfile(user)" style="cursor: pointer">
            <div class="profile-icon" style="width: 10%">
              <svg xmlns="http://www.w3.org/2000/svg" width="30" height="30" fill="currentColor" class="bi bi-person-fill" viewBox="0 0 16 16">
                <path d="M3 14s-1 0-1-1 1-4 6-4 6 3 6 4-1 1-1 1H3zm5-6a3 3 0 1 0 0-6 3 3 0 0 0 0 6z"/>
              </svg>
            </div>
            <div class="info" style="display: flex; flex-direction: row;width: 90%">
              <h3 style="margin-left: 1%">{{user.name}} {{user.surname}} </h3>
              <h3>({{user.username}})</h3>
            </div>
          </div>
          <div  style="display: flex; flex-direction: row;  margin: auto; margin-left: 20%; margin-top: 2%">
            <button  type="button" class="btn btn-light" style=" right: 10%; border: 1pt black solid;" v-on:click="Connect(user, index)"> ✔ Accept request</button>
            <button  type="button" class="btn btn-light" style=" right: 10%; margin-left: 0.5%; border: 1pt black solid;" v-on:click="RemoveFriendRequest(user, index)"> ✖ Decline request</button>
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
  name: "FriendReqests",
  data(){
    return{
      loggedUser: "",
      loggedUserDetails: "",
      requests: new Array()
    }
  },
  mounted: function () {
    this.loggedUser = localStorage.getItem('user')
    UserService.getUserByUsername(JSON.parse(this.loggedUser).username).then(response =>{
      this.loggedUserDetails = response.data.User
      ConnectionService.GetFriendRequests(this.loggedUserDetails.id).then(response1 =>{
        console.log("requests", response1.data.users)
        var friendNodes =  response1.data.users
        for (var f of friendNodes){
          UserService.getUserById(f.userID)
              .then(response2 =>{
                console.log(response2.data)
                this.requests.push(response2.data.User)
              })
              .catch(err =>{
                console.log("User with id: ", f.userID, " does not exists in UserService, just in ConnectionService!", err)
              })
        }

      })
          .catch(err1 =>{
            alert("It is impossible to load requests!")
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
    },
    RemoveFriendRequest(user, index){
      ConnectionService.DeclineFriendRequest(this.loggedUserDetails.id, user.id)
      .then(response => {
        console.log("removing friend request", response)
        this.requests.splice(index,1)
      })
          .catch(err => {
                console.log(err)
                alert("Error declining request!")
              }
          )

    },
    Connect(user, index){
      ConnectionService.Connect(this.loggedUserDetails.id, user.id)
          .then( response => {
            console.log("connecting:", response);
            this.requests.splice(index,1)
          })
          .catch(err => {
                console.log(err)
                alert("Error creating connection!")
              }
          )
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

}

.profile-informations{
  display: flex;
  flex-direction: row;
  horiz-align: center;
  vertical-align: center;

}

</style>