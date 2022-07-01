<template>
<div style="display: flex; flex-direction: row" >
  <div class="profile-info">

  </div>
  <div class="posts">
    <div class="create-new">
      <div class="head-line" style="display: flex; flex-direction: row">
        <svg xmlns="http://www.w3.org/2000/svg" width="30" height="30" fill="white" class="bi bi-file-earmark-post-fill" viewBox="0 0 16 16">
          <path d="M9.293 0H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V4.707A1 1 0 0 0 13.707 4L10 .293A1 1 0 0 0 9.293 0zM9.5 3.5v-2l3 3h-2a1 1 0 0 1-1-1zm-5-.5H7a.5.5 0 0 1 0 1H4.5a.5.5 0 0 1 0-1zm0 3h7a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-.5.5h-7a.5.5 0 0 1-.5-.5v-7a.5.5 0 0 1 .5-.5z"/>
        </svg>
        <h4 style="margin-left: 1%; margin-top: 0.2%; font-weight: bolder; font-style: italic; color: white">Create new post</h4>
      </div>
      <div class="post-content" style="display: flex; flex-direction: row" >
      <div class="hello" style="margin-top: 1%">
        <picture-input
            ref="pictureInput"
            width="350"
            height="350"
            margin="16"
            accept="image/jpeg,image/png"
            size="10"
            button-class="btn"
            :custom-strings="{
        upload: '<h1>Bummer!</h1>',
        drag: 'Drag or click to select photo',
        remove: 'Remove photo'
      }"
            @change="onChange">
        </picture-input>
      </div>
        <div class="content2">
          <div class="post-text">
            <div class="form-floating">
              <textarea class="form-control" placeholder="Write a post..." id="floatingTextarea" style="font-size: 12pt; height: 250px" v-model = "postText"></textarea>
              <label for="floatingTextarea" style="font-size: 12pt">Post text</label>
            </div>

          </div>
          <div class="post-link" style="display: flex; flex-direction: row">
            <div class="input-field">
              <input type="text" class="form-control" id="exampleFormControlInput1" v-model="link" placeholder="Add link">
            </div>
            <button type="button" class="btn btn-light" style="border-radius: 90%" v-on:click="addLink()">+</button>
          </div>
        </div>

      </div>

      <div class="footer">
        <button type="button" class="btn btn-primary" v-on:click="CreatePost">Post</button>

      </div>


    </div>
    <div class="view-all-users-posts">
    <h4 style="margin-top: 1%; color: white">Older posts</h4>
    </div>

  </div>

</div>


</template>

<script>
import PictureInput from 'vue-picture-input'
import PostService from '../services/PostService'
import UserService from "../services/UserService";
export default {
  name: "Profile",
  components:{
    PictureInput
  },
  data(){
    return{
      postText: '',
      image: '',
      links: new Array(),
      link: '',
      user: '',
      usersPosts: []

    }
  },
  mounted() {
  this.user = localStorage.getItem('user')

    },

  methods:{
    onChange (image) {
      console.log('New picture selected!')
      if (image) {
        console.log('Picture loaded.')
        this.image = image
      } else {
        console.log('FileReader API not supported: use the <form>, Luke!')
      }
    },

    addLink(){
      console.log("New link", this.link)
      this.links.push(this.link)
      this.link=''

    },
    emptyFields: function(){
      this.image = '';
      this.links = new Array();
      this.postText = ''
      this.link = ''
    },
    CreatePost: function (){
      console.log("U create post fji")
      PostService.createPost({
        "user": "",
        "posttext": this.postText,
        "image": this.image,
        "links": this.links,
        "isdeleted": false
      }).then(res => {
        console.log("New post", res.data)
        this.emptyFields();
      }).catch(err =>{
            console.log(err);
            alert("It is not possible to create post!")
      }
       );
    },

  }

}
</script>

<style scoped>
.profile-info{
  width: 30%;
  height: 100%;
  background-color: #151560;
  border: #c80000 solid 2pt;
}

.posts{
  width: 70%;
  height: 100%;
  border: royalblue solid 2pt;
}

.posts .create-new{
  height: 20%;
  border: white solid 2pt;
  border-radius: 20px;
  margin: 1%;
  padding: 1%;
}

.posts .create-new .head-line{
  margin: 1%;
  padding: 1%;
  border-bottom: white solid 2pt;

}

.posts .create-new .footer{
  margin-top: 3%;
}

.post-text{

}

.content2{
  margin: 1%;
  width: 70%;
  height: 100%;
}

.post-link{
  width: 100%;
  margin-top: 1%;
}

.post-link .input-field{
  width: 100%;
  margin-right: 1%;
}


.posts .view-all-users-posts{
  height: 80%;
  border: aqua dashed 2pt;
}

</style>