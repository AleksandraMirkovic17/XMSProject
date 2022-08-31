<template>
  <div>
    <div>
      <div class="page-header clear-filter" filter-color="red">
        <parallax
            class="page-header-image"
            style="background-image:url('img/mountain.jpg')"
        >

        </parallax>
        <div class="container">
          <div class="content-center brand">
            <img style="width: 300px" class="n-logo" src="../assets/icons8-d-67.png" alt="../assets/icons8-d-67.jpg" />
            <h1 style="font-weight: bolder; color: #242424">Dislinkt</h1>
            <h3>Widen your horizons.</h3>
            <div v-if="!loggedUser">
            <button type="button" class="btn btn-dark" style="width: 50%" onclick="window.location.href='/register'">Register</button>
              <br>
              <br>
              <p style="font-size: 12pt">
                Already have an account?
              </p>
              <button type="button" class="btn btn-dark" style="width: 50%" onclick="window.location.href='/login'">Log in</button>
            </div>
          </div>
          <h6 class="category category-absolute">
          </h6>
        </div>
      </div>
    </div>
    <!--Likes Modal -->
    <!--DislikesModal -->
    <!--CommentsModal -->
  </div>

</template>

<script>
import PostService from "../services/PostService";
import { Parallax } from '@/components';

export default {
  name: "Home",
  bodyClass: 'index-page',
  components: {
    Parallax,

  },
  data(){
    return{
      loggedUser: "",
      feedPosts: new Array(),
      selectedPost: "",
      selectedPostLikes: new Array(),
      selectedPostDislikes: new Array(),
      selectedPostComments: new Array(),
      newCommentContent: ""
    }
  },
  mounted: function(){
    this.loggedUser = localStorage.getItem('user')
    PostService.getFeedForUser(JSON.parse(this.loggedUser).username).then(response1 =>{
      this.feedPosts = response1.data.Feed
      if (this.feedPosts.length>0){
        console.log("selecting post")
        console.log("Selected", this.feedPosts[0])
       this.changeSelectedPost(this.feedPosts[0].Id)
      }
    })
    .catch( err => {
          console.log("Error while getting feed!", err)
        }
    )
  },
  methods:{
    checkUserLiked() {
      if(this.loggedUser){
      var username = JSON.parse(this.loggedUser).username;
      for (var likes of this.selectedPostLikes){
        if(username == likes.username){
          this.userLiked = true;
          return;
        }
      }
    }
      console.log("user liked false")
      this.userLiked = false;

    },
    checkUserDisliked(){
      if(this.loggedUser){
        var username = JSON.parse(this.loggedUser).username;
        for (var dislike of this.selectedPostDislikes){
          if(username == dislike.username){
            this.userDisliked = true;
            return;
          }
        }
      }
      this.userDisliked = false;
    },
    likePost(){
      PostService.reactToPost({
        "username" : JSON.parse(this.loggedUser).username,
        "postId" : this.selectedPost.Id,
        "reactionType": 2
      })
          .then( respo =>{
            this.changeSelectedPost(respo.data.post.id)
          })
    },
    dislikePost(){
      PostService.reactToPost({
        "username" : JSON.parse(this.loggedUser).username,
        "postId" : this.selectedPost.Id,
        "reactionType": 1
      })
          .then( respo =>{
            console.log(respo.data.post)
            this.changeSelectedPost(respo.data.post.id)
          })
    },
    commentOnPost(){
      console.log(this.loggedUser)
      PostService.createComment({
        "id": "",
        "content": this.newCommentContent,
        "username": JSON.parse(this.loggedUser).username
      }, this.selectedPost.Id)
          .then( response =>{
            this.changeSelectedPost(response.data.post.id)
            this.newCommentContent = ""
          })
          .catch( err => {
                alert("It is impossible to leave a comment right now!");
                console.log(err)
              }
          )
    },
    changeSelectedPost(){
    },
  }
}
</script>

<style scoped>
.main-home{
  margin-left: 20%;
  margin-top: 5%;
  margin-bottom: 0;
  width: 60%;
  height: 700px;
  horiz-align: center;
  vertical-align: center;
  background-color: #e85a4f;
  border-radius: 20px;
  display: flex;
  flex-direction: row;
}

.text-home{
  padding: 5%;
  max-width: 50%;
}

.img-home{
  width: 50%;
}

.dropdown-toggle { outline: 0; }

.btn-toggle {
  padding: .25rem.5rem;
  font-weight: 600;
  color: rgba(0, 0, 0, .65);
  background-color: transparent;
}
.btn-toggle:hover,
.btn-toggle:focus {
  color: rgba(0, 0, 0, .85);
  background-color: #d2f4ea;
}

.btn-toggle::before {
  width: 1.25em;
  line-height: 0;
  content: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' width='16' height='16' viewBox='0 0 16 16'%3e%3cpath fill='none' stroke='rgba%280,0,0,.5%29' stroke-linecap='round' stroke-linejoin='round' stroke-width='2' d='M5 14l6-6-6-6'/%3e%3c/svg%3e");
  transition: transform .35s ease;
  transform-origin: .5em 50%;
}

.btn-toggle[aria-expanded="true"] {
  color: rgba(0, 0, 0, .85);
}
.btn-toggle[aria-expanded="true"]::before {
  transform: rotate(90deg);
}

.btn-toggle-nav a {
  padding: .1875rem .5rem;
  margin-top: .125rem;
  margin-left: 1.25rem;
}
.btn-toggle-nav a:hover,
.btn-toggle-nav a:focus {
  background-color: #d2f4ea;
}

.scrollarea {
  overflow-y: auto;
}

.post-view{
  margin: 2%;
  padding: 2%;
  background-color: whitesmoke;
}

.post-view-nav{
  height: 10%;
  border-bottom: #242424 1pt solid;
}

.post-view-person-icon{
}

.post-view-username{
  margin-left: 3%;
}

.post-view-date{

}

.view-all-users-posts{
  margin-left: 25%;
  margin-right: 25%;
}



</style>