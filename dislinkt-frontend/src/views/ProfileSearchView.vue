<template>
  <div>
    <div v-for="user in users" :key="user.name">
      <p>nene</p>

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

</style>