import axios from 'axios';

const USER_API_BASE_URL = 'http://localhost:4200/';

class JobService {
    PublishPost(post){
        return axios.post(USER_API_BASE_URL+"job", post)
    }

}

export default new JobService()