import axios from 'axios';

const USER_API_BASE_URL = 'http://localhost:4200/user';

class UserService{

    getUsers(){
       return axios.get(USER_API_BASE_URL);
    }

    registerUser(user){
        return axios.post(USER_API_BASE_URL, user);
    }

    login(user) {
        return axios
          .post('http://localhost:4200/login', user)
          .then(response => {
            console.log(response.data.token)
            if (response.data.token) {
              localStorage.setItem('user', JSON.stringify(response.data));
            }
            return response.data;
          });
      }
  
      logout() {
        localStorage.removeItem('user');
      }
}

export default new UserService();