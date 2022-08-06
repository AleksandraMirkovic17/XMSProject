<template>
<div class="connections" style="display: flex; flex-direction: row; ">
  <div class="b-sidebar" style="width: 20%; height: 80vw; margin-top: 3%; border-right: #8d4141 2pt solid;">
    <div class="d-flex flex-column flex-shrink-0 p-3 text-bg-dark" style="width: 100%; height: 95%; margin-bottom: 2%">
      <a href="/" class="d-flex align-items-center mb-3 mb-md-0 me-md-auto text-white text-decoration-none">
        <svg class="bi pe-none me-2" width="40" height="32"><use xlink:href="#bootstrap"/></svg>
        <span class="fs-4">Connections</span>
      </a>
      <hr>
      <ul class="nav nav-pills flex-column mb-auto">
        <li class="nav-item">
          <div :class="{customactive:display=='friends'}">

          <a href="#" class="nav-link text-white" aria-current="page" v-on:click="DisplayFriends">
              <p style="font-size: 150%">Friends</p>

          </a>
    </div>

    </li>
        <li>
          <div :class="{customactive:display=='requests'}">

          <a href="#" class="nav-link text-white" v-on:click="DisplayRequests">
            <p style="font-size: 150%">Friend requests</p>
          </a>
          </div>
        </li>
        <li>
          <div :class="{customactive:display=='blockeds'}">

          <a href="#" class="nav-link text-white" v-on:click="DisplayBlockeds">
            <p style="font-size: 150%">Blocked</p>
          </a>
            </div>
        </li>
      </ul>
      <hr>
    </div>
  </div>
  <div class="view" style="width: 55%; margin-left: 20%; height: 100%; border-right: #e85a4f 1pt solid; padding: 1%;">
    <div class="router-elem">
      <div v-if="display=='friends'">
        <Friends></Friends>
      </div>
      <div v-if="display=='requests'">
        <FriendReqests></FriendReqests>
      </div>
      <div v-if="display=='blockeds'">
        <Blockeds></Blockeds>
      </div>


    </div>
  </div>
  <div style="padding: 1%; border-left: #e85a4f 1pt solid; width: 25%">
    <Recommendation></Recommendation>
  </div>

</div>
</template>

<script>
import Friends from "./Connections/Friends";
import Blockeds from "./Connections/Blockeds";
import FriendReqests from "./Connections/FriendReqests";
import Recommendation from "./Connections/Recommendation";
export default {
  name: "Connections",
  components: {Friends, Blockeds, FriendReqests, Recommendation },
  data: function (){
    return{
      display: 'friends'
    }
  },
  mounted() {
  },
  methods:{
    DisplayFriends(){
      this.display = 'friends'
    },
    DisplayRequests(){
      this.display = 'requests'
    },
    DisplayBlockeds(){
      this.display='blockeds'
    }
  }
}
</script>

<style scoped>

.customactive{
  background-color: #e85a4f;
  color: white;
  font-weight: bold;
  border-radius: 20px;
}

</style>