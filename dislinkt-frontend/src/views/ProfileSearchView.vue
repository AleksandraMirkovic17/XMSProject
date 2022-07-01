<template>
  <div>
    <div v-for="(user,index) in users" :key="index">
      <div class="profile-container">
        <div class="profile-icon">
          <svg xmlns="http://www.w3.org/2000/svg" width="30" height="30" fill="currentColor" class="bi bi-person-fill" viewBox="0 0 16 16">
            <path d="M3 14s-1 0-1-1 1-4 6-4 6 3 6 4-1 1-1 1H3zm5-6a3 3 0 1 0 0-6 3 3 0 0 0 0 6z"/>
          </svg>
        </div>
        <div class="info">
          <h4>{{user.name}} {{user.surname}}</h4>
          <h3>{{user.username}}</h3>
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
      users: []
    }
    },
  mounted() {
    //var searchParams = this.$route.params.search
    console.log(UserService.getUsers())
    UserService.getUsers().then(res => {
      this.users = res.data.users
    });


  },
  methods:{

  }
}
</script>

<style scoped>
.profile-container{
  border-radius: 20px;
  backround-color: white;
  color: #e5e5e5;
  display: flex;
  flex-direction: row;
  horiz-align: center;
  vertical-align: center;
}

</style>