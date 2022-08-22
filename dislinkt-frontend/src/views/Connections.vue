<template>
<div class="connections" style="display: flex; flex-direction: row; width: 100%;   justify-content: center; ">

    <div class="router-elem" style="width: 70%; margin: 0">
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
</template>

<script>
import Friends from "./Connections/Friends";
import Blockeds from "./Connections/Blockeds";
import FriendReqests from "./Connections/FriendReqests";
export default {
  name: "Connections",
  components: {Friends, Blockeds, FriendReqests },
  props: {display: String},
  data: function (){
    return{
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