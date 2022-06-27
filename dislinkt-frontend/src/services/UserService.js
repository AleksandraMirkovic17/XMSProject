import axios from 'axios';

const USER_API_BASE_URL = 'http://localhost:4200/user';

class UserService{

    getUsers(){
       return axios.get(USER_API_BASE_URL);
    }

    registerUser(user){
        return axios.post(USER_API_BASE_URL, user);
    }
}

export default new UserService();