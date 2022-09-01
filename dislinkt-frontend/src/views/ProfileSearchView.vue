<template>
  <div>
    <div class="search">
      <div v-for="(user,index) in users" :key="index">
        <div class="panel profile-card-small" v-if="!loggedUser" v-on:click="redirectToProfile(user)"  style="cursor: pointer; display: flex; flex-direction: row">
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
        <div class="panel profile-card-small" v-if="loggedUser" v-on:click="redirectToProfile(user)"  style="cursor: pointer; display: flex; flex-direction: row">
          <div class="profile-icon" style="width: 20%">
            <div class="photo-container-small">
              <p style="position: relative; align-content: center; margin: 13%; font-weight: bold;" v-if="user && user.Name && user.Surname">
                {{user.Name.charAt(0).toUpperCase()}}{{user.Surname.charAt(0).toUpperCase()}}
              </p>
            </div>
          </div>
          <div class="info">
            <div  style="display: flex; flex-direction: row;  margin-top: 2%; ">
              <p style="margin-left: 1%; font-size: 130%;">{{user.Name}}  </p>
              <p style="margin-left: 1%; font-size: 130%">{{user.Surname}}</p>
            </div>
            <p class="username">@{{user.Username}}</p>
          </div>
        </div>

      </div>
    </div>
  </div>

</template>

<script>
import UserService from "../services/UserService";
export default {
  name: "ProfileSearchView",
  data(){
    return {
      users: new Array(),
      loggedUser: '',
      loggedUserDetails : {},
    }
    },
  mounted() {
    this.loggedUser = localStorage.getItem('user')
    var searchParams = this.$route.params.search
    if (this.loggedUser){
      UserService.getUserByUsername(JSON.parse(this.loggedUser).username).then(response2 => {
        this.loggedUserDetails = response2.data.User
        UserService.searchUsersLogged(this.loggedUserDetails.id, searchParams)
        .then(response3 =>{
          console.log(response3.data.Users)
          this.users = response3.data.Users
        })
      })

    }else {
      UserService.searchUsers(searchParams).then(res => {
        this.users = res.data.users
      });

    }


  },
  methods:{
    redirectToProfile(user){
      if (!this.loggedUser){
        this.$router.push("/profile/"+user.username)
      } else {
        this.$router.push("/profile/"+user.Username)
      }
    }

  }
}
</script>

<style scoped>

.search{
  margin-top: 10%;
  margin-right: 20%;
  margin-left: 20%;
}

.profile-card-small{
  background-color: #e1e1e1;
  padding: 3%;
  margin-top: 2%;
  width: 100%;
  box-shadow: 0px 10px 25px 0px rgba(0, 0, 0, 0.3);
}

.photo-container-small{
  background-color: #f6f6f6;
  width: 70px;
  height: 70px;
  border-radius: 50%;
  overflow: hidden;
  margin: 0 auto;
  box-shadow: 0px 10px 25px 0px rgba(0, 0, 0, 0.3);
  align-content: center;
  vertical-align: center;
  text-align: center;
  horiz-align: center;
}


</style>