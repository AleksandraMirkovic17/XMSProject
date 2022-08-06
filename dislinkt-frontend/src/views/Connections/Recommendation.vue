<template>
<div v-if="loggedUser" class="recommendations-main">
  <div class="title">
    <h6 style="color: white; font-weight: bolder; margin: 2%">Add recommended friends({{recommendations.length}})</h6>
  </div>
  <div class="recommendations" style="margin-top: 2%;">
    <div v-for="(user,index) in recommendations" :key="index">
      <div class="profile-recommendation" style="margin-top: 2%; width: 100%">
        <div class="profile-view" v-on:click="redirectToProfile(user)" style="cursor: pointer">
          <div class="profile-icon" style="width: 10%">
            <svg xmlns="http://www.w3.org/2000/svg" width="30" height="30" fill="currentColor" class="bi bi-person-fill" viewBox="0 0 16 16">
              <path d="M3 14s-1 0-1-1 1-4 6-4 6 3 6 4-1 1-1 1H3zm5-6a3 3 0 1 0 0-6 3 3 0 0 0 0 6z"/>
            </svg>
          </div>
          <div class="info" style="display: flex; flex-direction: row; margin-left: 2%; margin-top: 2%">
            <p style="margin-left: 1%; font-size: 130%">{{user.name}}  </p>
            <p style="margin-left: 1%; font-size: 130%">{{user.surname}}</p>
            <p style="margin-left: 1%; font-size: 130%">({{user.username}})</p>
          </div>

        </div>
        <div class="mutual-info">
          {{user.isMutual}}
          <div v-if="user.isMutualFriends">
            <p class="info-mutual">{{user.MutualFriends}} mutual friends</p>
          </div>
          <div v-else>
            <p class="info-mutual">Followed by {{user.MutualFriends}}</p>
          </div>

        </div>
        <div class="profile-respond">
          <button v-if="user.connectionStatus=='NO_RELATION'" type="button" class="btn btn-light" style="right: 10%" v-on:click="follow(user)">+ Follow</button>
          <button v-if="user.connectionStatus=='PENDING'" type="button" class="btn btn-light" style="right: 10%" v-on:click="unsendRequest(user)">✔ Friend request sent</button>
          <div v-if="user.connectionStatus=='ACCEPT'">
            <button  type="button" class="btn btn-light" style=" right: 10%; border: 1pt black solid;" v-on:click="follow(user)"> ✔ Accept request</button>
            <button  type="button" class="btn btn-light" style=" right: 10%; margin-left: 0.5%; border: 1pt black solid;" v-on:click="RemoveFriendRequest(user)"> ✖ Decline request</button>
          </div>
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
  name: "Recommendation",
  data(){
    return{
      loggedUser: "",
      loggedUserDetails: "",
      recommendations: new Array(),

    }
  },
  mounted: function () {
    this.loggedUser = localStorage.getItem('user')
    UserService.getUserByUsername(JSON.parse(this.loggedUser).username).then(response =>{
      this.loggedUserDetails = response.data.User
      ConnectionService.GetRecommenadations(this.loggedUserDetails.id).then(response1 =>{
        console.log("recommendations:", response1.data.users)
        var friendNodes =  response1.data.users
        for (var f of friendNodes){
          this.pushNewRecommendation(f)
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
  methods: {
    pushNewRecommendation(f){
      UserService.getUserById(f.userID)
          .then(response2 =>{
            console.log(response2.data)
            let recommendedUser = {id: f.userID, name: response2.data.User.name, surname: response2.data.User.surname,
              username: response2.data.User.username,
              isMutualFriends: f.isMutual, MutualFriends: f.Mutual, connectionStatus: 'NO_RELATION', Public: !f.isPrivate}
            ConnectionService.GetConnectionDetail(this.loggedUserDetails.id, response2.data.User.id)
            .then(response3 =>{
              console.log("Connection details between recommended and logged user:", response3.data)
              recommendedUser.connectionStatus = response3.data.relation
              console.log("recommended user",recommendedUser)
              this.recommendations.push(recommendedUser)
            })
                .catch(err2 =>{
                  console.log(err2)
                  alert("It is impossible to get connection details between logged user and recommended user!")
                })


          })
          .catch(err =>{
            console.log("User with id: ", f.userID, " does not exists in UserService, just in ConnectionService!", err)
          })

    },
    redirectToProfile(user){
      this.$router.push("/profile/"+user.username)
    },
    follow(user){
      console.log(user, this.loggedUserDetails)
      if(user.Public || user.connectionStatus == 'ACCEPT'){
        ConnectionService.Connect(this.loggedUserDetails.id, user.id)
            .then( response => {
              console.log("connecting:", response);
              this.getNewRecommendations()
            })
            .catch(err => {
                  console.log(err)
                  alert("Error creating connection!")
                }
            )
      } else{
        ConnectionService.SendFriendRequest(this.loggedUserDetails.id, user.id)
            .then( response => {
              console.log("connecting:", response);
              this.getNewRecommendations()
            })
            .catch(err => {
                  console.log(err)
                  alert("Error sending request!")
                }
            )
      }


    },
    unsendRequest(user){
      ConnectionService.UnsendFriendRequest(this.loggedUserDetails.id, user.id)
          .then( response => {
            console.log("connecting:", response);
            this.getNewRecommendations()
          })
          .catch(err => {
                console.log(err)
                alert("Error unsending request!")
              }
          )

    },
    RemoveFriendRequest(user){
      ConnectionService.DeclineFriendRequest(this.loggedUserDetails.id, user.id)
          .then(response => {
            console.log("declining friend request: ", response);
            this.getNewRecommendations()
          })
          .catch(err => {
            console.log(err)
            alert("Error declining request!")
          })

    },
    getNewRecommendations() {
      this.recommendations = new Array()
      ConnectionService.GetRecommenadations(this.loggedUserDetails.id).then(response1 =>{
        console.log("recommendations:", response1.data.users)
        var friendNodes =  response1.data.users
        for (var f of friendNodes){
          this.pushNewRecommendation(f)
        }

      })
          .catch(err1 =>{
            alert("It is impossible to load friends!")
            console.log(err1)
          })
    }
  }
}
</script>

<style scoped>

.profile-view{
  margin-top: 2%;
  backround-color: white;
  padding: 1%;
  color: #e5e5e5;
  display: flex;
  flex-direction: row;
  horiz-align: center;
  vertical-align: center;
}

.recommendations{
  background-color: #333333;
  padding: 3%;
  border-radius: 15px;
}

.profile-recommendation{
  padding: 3%;
  padding-top: 0%;
  border-style: solid;
  border-width: 1pt;
  border-color: #e5e5e5;
}

.profile-view{
  padding: 3%;
}

.recommendations-main{
  margin-right: 0;
  width: 100%;
}

.info-mutual{
  color: white;
}

</style>