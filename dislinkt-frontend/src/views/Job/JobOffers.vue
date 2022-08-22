<template>
<div>
  <div v-if="jobsToShow.length==0">
    <h4 class="description">There is no published jobs. Come back later!</h4>
  </div>
  <div v-else style="width: 45%">
    <div style="display: flex; flex-direction: row">
      <input type="text" class="form-control" v-model="searchParams" placeholder="Search"/>
    </div>
  </div>
<div v-for="(job,index) in jobsToShow" :key="index">
  <JobOffer v-if="filter(job)" :company-name="job.companyName"
            :date-posted="job.datePosted"
  :date-valid="job.dateValid"
  :description="job.jobDescription"
  :job-i-d="job.jobID"
  :position="job.position"
  :publisher-i-d="job.publisherID"
  :required-skills="job.requiredSkills">
  </JobOffer>
</div>

</div>
</template>

<script>
import JobService from "../../services/JobService";
import JobOffer from "./JobOffer";
export default {
  name: "JobOffers",

  props: {allJobs: Boolean, usersJobs: Boolean, userid: String},
  data(){
    return{
      jobsToShow: new Array(),
      user: String,
      searchParams: ""

    }
  },
  components: {JobOffer},
  mounted: function () {
    this.user = new String(this.userid)
    if (this.usersJobs){
      JobService.GetJobsByPublisher(this.user).then(response=>{
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
    filter: function(job){
      if(job.position.includes(this.searchParams) || job.companyName.includes(this.searchParams)){
        return true
      }
      return false
    }

  }
}
</script>

<style scoped>

</style>