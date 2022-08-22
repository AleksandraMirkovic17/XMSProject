import axios from 'axios';

const USER_API_BASE_URL = 'http://localhost:4200/';

class JobService {
    PublishPost(post){
        return axios.post(USER_API_BASE_URL+"job", post)
    }

    GetJobRecommendations(userID){
        return axios.get(USER_API_BASE_URL+"job/recommendation/"+userID)
    }

    GetJobsByPublisher(userID){
        console.log(USER_API_BASE_URL+"job/user/"+userID)
        return axios.get(USER_API_BASE_URL+"job/user/"+userID)
    }

    GetJobsBySearch(param){
        return axios.get(USER_API_BASE_URL+"job/search/"+param)
    }

    GetAllJobs(){
        return axios.get(USER_API_BASE_URL+"job")
    }

    GetJobById(jobId){
        return axios.get(USER_API_BASE_URL+"job/"+jobId)
    }

}

export default new JobService()