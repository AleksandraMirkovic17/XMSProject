import axios from 'axios';

const USER_API_BASE_URL = 'http://localhost:4200/';

class PostService{
    getAllPosts(){
        return axios.get(USER_API_BASE_URL+"post")
    }

    getAllPostsByUser(userId){
        return axios.get(USER_API_BASE_URL+"post/user/"+userId)
    }

    createPost(post){
        print("printing new post", post.posttext)
        return axios.post(USER_API_BASE_URL+"post", post)
    }

    getLikes(postId){
        return axios.get(USER_API_BASE_URL+"post/"+postId+"/likes")
    }

    getDislikes(postId){
        return axios.get(USER_API_BASE_URL+"post/"+postId+"/dislikes")
    }

    getComments(postId){
        return axios.get(USER_API_BASE_URL+"post/"+postId+"/comment")
    }

    reactToPost(reaction){
        return axios.post(USER_API_BASE_URL+"post/react", reaction)
    }

    createComment(newComment, postID){
        return axios.post(USER_API_BASE_URL+"post/"+postID+"/comment", newComment)
    }

}

export default new PostService()