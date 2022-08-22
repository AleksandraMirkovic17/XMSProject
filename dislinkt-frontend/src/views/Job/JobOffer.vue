<template>
<div>
  <div class="job-view profile-card-small" style="display: flex; flex-direction: column">
    <div style="display: flex; flex-direction: row">
      <div class="photo-container-small" style="text-align: center; ">
        <p style="position: relative; align-content: center; margin: 13%; margin-right: 2%; font-weight: bold;" v-if="position">
          {{position.charAt(0).toUpperCase()}}
        </p>
      </div>
      <div class="info" style="width: 100%">
        <div  style="display: flex; flex-direction: row; ">
          <p style="margin-left: 1%; font-size: 130%; font-weight: bolder">{{position}}  </p>
        </div>
        <p class="username" style="margin: 0; text-align: left">@{{companyName}}</p>
      </div>

    </div>
    <div style="margin-top: 2%; text-align: left">
      {{description}}
    </div>
    <div style="margin-top: 2%; text-align: left">
      Apply before {{dateValid}}.
    </div>

    <div style="align-items: self-start; justify-content: left; align-content: first">
      <badge v-for="(skill,index) in requiredSkills" :key="index" type="primary" style="margin: 1%">
        {{skill}}
      </badge>
    </div>
    <div style="margin-top: 2%; text-align: left">
      Posted {{calculateDays()}} days ago
    </div>
  </div>

</div>
</template>

<script>
import {Badge} from "../../components";
export default {
  name: "JobOffer",
  props: {companyName: String,
  datePosted: String,
  dateValid: String,
  description: String,
  jobID: String,
  position: String,
  publisherID: String,
  requiredSkills: Array()},
  data(){
    return{

    }
  },
  components: {
    Badge
  },
  mounted: function () {
  },
  methods:{
 calculateDays(){
 var today = new Date()
   var thatDay = new Date(this.datePosted)
   var days = today.getDate()-thatDay.getDate()
   return days
 }
  }
}
</script>

<style scoped>

</style>