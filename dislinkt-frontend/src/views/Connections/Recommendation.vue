<template>
<div v-if="loggedUser" class="recommendations-main container">
    <h6>People you might know({{recommendations.length}})</h6>
  <hr>
  <div class="recommendations" style="margin-top: 2%;">
    <div v-for="(user,index) in recommendations" :key="index">
      <div class="panel profile-card-small" style="">
        <div class="profile-view" v-on:click="redirectToProfile(user)" style="cursor: pointer; display: flex; flex-direction: row" >
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
        <div class="mutual-info">
          {{user.isMutual}}
          <div v-if="user.isMutualFriends">
            <p class="info-mutual">{{user.MutualFriends}} mutual friends</p>
          </div>
          <div v-else>
            <p class="info-mutual">Followed by {{user.MutualFriends}}</p>
          </div>

        </div>
        <div class="profile-respond" style="margin-left: 30%;">
          <button v-if="user.connectionStatus=='NO_RELATION'" type="button" class="btn btn-round" style="right: 10%" v-on:click="follow(user)">+ Follow</button>
          <button v-if="user.connectionStatus=='PENDING'" type="button" class="btn btn-round" style="right: 10%" v-on:click="unsendRequest(user)">✔ Friend request sent</button>
          <div v-if="user.connectionStatus=='ACCEPT'">
            <button  type="button" class="btn btn-round" style=" right: 10%; border: 1pt black solid;" v-on:click="follow(user)"> ✔ Accept request</button>
            <button  type="button" class="btn btn-round" style=" right: 10%; margin-left: 0.5%; border: 1pt black solid;" v-on:click="RemoveFriendRequest(user)"> ✖ Decline request</button>
          </div>
        </div>

      </div>

    </div>

  </div>
  <h6 style="margin-top: 10%">Jobs for you({{jobs4u.length}})</h6>
  <hr>
  <div class="recommendations" style="margin-top: 2%;">
    <div v-for="(job,index) in jobs4u" :key="index">
      <div class="panel profile-card-small" style="">
        <div class="profile-view" v-on:click="redirectToProfileJob(job)" style="cursor: pointer; display: flex; flex-direction: row" >
          <div class="profile-icon" style="width: 20%">
            <div class="photo-container-small" style="text-align: center;">
              <p style="position: relative; align-content: center; margin: 13%; font-weight: bold;" v-if="job.position">
                {{job.position.charAt(0).toUpperCase()}}
              </p>
            </div>
          </div>
          <div class="info">
            <div  style="display: flex; flex-direction: row;  margin-top: 2%">
              <p style="margin-left: 1%; font-size: 130%;">{{job.position}}  </p>
            </div>
            <p class="username">@{{job.companyName}}</p>
          </div>

        </div>
        <div class="mutual-info">
          <badge v-for="(skill,index) in job.requiredSkills" :key="index" type="info" style="margin: 1%">
            {{skill}}
          </badge>
        </div>

      </div>

    </div>

  </div>

</div>
</template>

<script>
import UserService from "../../services/UserService";
import ConnectionService from "../../services/ConnectionService";
import JobService from "../../services/JobService";
import {Badge} from "../../components";


export default {
  name: "Recommendation",
  data(){
    return{
      loggedUser: "",
      loggedUserDetails: "",
      recommendations: new Array(),
      jobs4u: new Array()

    }
  },
  components:{
    Badge,
  },
  mounted: function () {
    this.loggedUser = localStorage.getItem('user')
    UserService.getUserByUsername(JSON.parse(this.loggedUser).username).then(response =>{
      this.loggedUserDetails = response.data.User
      ConnectionService.GetRecommenadations(this.loggedUserDetails.id).then(response1 =>{
        //console.log("recommendations:", response1.data.users)
        var friendNodes =  response1.data.users
        for (var f of friendNodes){
          this.pushNewRecommendation(f)
        }

      })
          .catch(err1 =>{
            alert("It is impossible to load friend recommendation!")
            console.log(err1)
          })
      JobService.GetJobRecommendations(this.loggedUserDetails.id).then( response2 =>{
        var jobNodes = response2.data.jobs
        for (var j of jobNodes){
          this.pushNewJobRecommendation(j)
        }
      }).catch(err1 =>{
            //alert("It is impossible to load job recommendation!")
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
           // console.log(response2.data)
            let recommendedUser = {id: f.userID, name: response2.data.User.name, surname: response2.data.User.surname,
              username: response2.data.User.username,
              isMutualFriends: f.isMutual, MutualFriends: f.Mutual, connectionStatus: 'NO_RELATION', Public: !f.isPrivate}
            ConnectionService.GetConnectionDetail(this.loggedUserDetails.id, response2.data.User.id)
            .then(response3 =>{
             // console.log("Connection details between recommended and logged user:", response3.data)
              recommendedUser.connectionStatus = response3.data.relation
             // console.log("recommended user",recommendedUser)
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
    pushNewJobRecommendation(f){
      JobService.GetJobById(f.jobID)
          .then(response2 =>{
            console.log(response2.data)
            let recommendedJob = {jobID: f.jobID,
              publisherId: response2.data.job.publisherId,
              datePosted: response2.data.job.datePosted,
              dateValid: response2.data.job.dateValid,
              companyName: f.companyName,
              position: f.position,
              jobDescription: f.jobDescription,
              requiredSkills: f.requiredSkills}
            this.jobs4u.push(recommendedJob)
          })
          .catch(err =>{
            console.log("It is impossible to get job by id "+f.jobID+"."+ err)
            alert("It is impossible to get job by id "+f.jobID+".")
          })

    },
    redirectToProfile(user){
      this.$router.push("/profile/"+user.username)
    },
    redirectToProfileJob(job){
      console.log("Redirecting "+job)
    },
    follow(user){
     // console.log(user, this.loggedUserDetails)
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
       // console.log("recommendations:", response1.data.users)
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
  padding: 1%;
  display: flex;
  flex-direction: row;
  horiz-align: center;
  vertical-align: center;
}

.recommendations-main{

}

</style>