<template>
<div>
  <div class="view-all-users-posts">
    <div v-if="postsToShow && postsToShow.length==0">
      <h4 class="description">There is no posts yet. Come back later!</h4>
    </div>
    <div v-for="(post, index) in postsToShow" :key="index">
      <Post v-if="post.id"  :username="username" :date="post.date" :posttext="post.posttext" :image="post.image" :links="post.links" :postid="post.id"></Post>
      <Post v-if="post.Id"  :username="post.User" :date="post.Date" :posttext="post.PostText" :image="post.Image" :links="post.Links" :postid="post.Id"></Post>

    </div>
  </div>

</div>
</template>

<script>
import PostService from "../../services/PostService";
import Post from "./Post";
export default {
  name: "Posts",
  props:{
    feedPosts: Boolean,
    usersPosts: Boolean,
    userid: String,
    username: String
  },
  data(){
    return{
      postsToShow: new Array()
    }
  },
  components:{
    Post

  },
  mounted: function () {
    if (this.feedPosts){
      PostService.getFeedForUser(this.username).then(response1 =>{
        this.postsToShow = response1.data.Feed
      })

    }
    if (this.usersPosts){
      console.log("getting all posts by user")
      PostService.getAllPostsByUser(this.userid).then(response =>{
        this.postsToShow = response.data.posts;
      })
    }
  },
  methods:{

  }
}
</script>

<style scoped>

.view-all-users-posts{

}

</style>