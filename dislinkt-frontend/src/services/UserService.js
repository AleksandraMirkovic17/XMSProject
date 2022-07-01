import axios from 'axios';

const USER_API_BASE_URL = 'http://localhost:4200/user';

class UserService{

    getUsers(){
       return axios.get(USER_API_BASE_URL);
    }

    searchUsers(param){
        return axios.get(USER_API_BASE_URL+"/"+ param)
    }

    registerUser(user){
        return axios.post(USER_API_BASE_URL, user);
    }

    login(user) {
        return axios.post('http://localhost:4200/login', user);
      }
  
      logout() {
        localStorage.removeItem('user');
      }
}

export default new UserService();