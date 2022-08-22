<template>
<div>
  <div v-if="jobsToShow.length==0">
    <h4 class="description">There is no posts yet. Come back later!</h4>
  </div>
<div v-for="(job,index) in jobsToShow" :key="index">
  <
</div>

</div>
</template>

<script>
import JobService from "../../services/JobService";
import JobOffer from "./JobOffer";
export default {
  name: "JobOffers",
  props: {allJobs: Boolean, usersJobs: Boolean, userID: String},
  data(){
    return{
      jobsToShow: new Array()


    }
  },
  components: {JobOffer},
  mounted: function () {
    if (this.usersJobs){
      JobService.GetJobsByPublisher(this.userID).then(response=>{
        this.jobsToShow = response.data.jobs
      }).catch(err=>{
        console.log(err)
        alert("It is impossible to load jobs by publisher!")
      })
    }
    if(this.allJobs){
      JobService.GetAllJobs().then(response=>{
        this.jobsToShow = response.data.jobs
      }).catch(err =>{
        console.log(err)
        alert("It is impossible to load all jobs!")
      })

    }

  },
  methods:{

  }
}
</script>

<style scoped>

</style>