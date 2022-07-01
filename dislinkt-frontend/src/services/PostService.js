import axios from 'axios';
import UserService from "./UserService";

const USER_API_BASE_URL = 'http://localhost:4200/';

class PostService{
    getAllPosts(){
        return axios.get(USER_API_BASE_URL+"post")
    }

    createPost(post){
        print("printing new post", post.posttext)
        return axios.post(USER_API_BASE_URL+"post", post)
    }

}

export default new PostService()