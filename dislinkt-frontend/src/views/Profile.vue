<template>

<div style="display: flex; flex-direction: row" >
  <div class="profile-panel">
    <div class="register-show">
      <h2>{{user.username}}</h2>
            <div class="row d-flex mt-5">
                <div class="col-md-12">
                    <label class="input_label">
                        <input type="email" name="email" v-model="user.email" disabled=yes required="required" @change="userInfoHasChanged()">
                        <span class="keep_hovered">EMail</span>
                    </label>
                </div>
            </div>
            <div class="row d-flex mt-4">
                <div class="col-md-12">
                    <label class="input_label">
                        <input type="email" name="email" v-model="user.username" disabled=yes required="required" @change="userInfoHasChanged()">
                        <span class="keep_hovered">Username</span>
                    </label>
                </div>
            </div>
            <div class="row d-flex mt-4" v-if="isUserLoggedIn()">
                <div class="col-md-4">
                    <label class="input_label">
                        <input type="password" name="old-password" v-model="user.oldPasswordGuess" @change="userInfoHasChanged()" :disabled="!(isUserLoggedIn())">
                        <span class="keep_hovered">Old Password</span>
                    </label>
                </div>
                <div class="col-md-4">
                    <label class="input_label">
                        <input type="password" name="new-password" v-model="user.newPassword" @change="userInfoHasChanged()" :disabled="!(isUserLoggedIn())">
                        <span class="keep_hovered">New Password</span>
                    </label>
                </div>
                <div class="col-md-4">
                    <label class="input_label">
                        <input type="password" name="confirm-new-password" v-model="user.newPasswordConfirmation" @change="userInfoHasChanged()" :disabled="!(isUserLoggedIn())">
                        <span class="keep_hovered">Confirm New Password</span>
                    </label>
                </div>
            </div>
            <div class="row d-flex mt-4">
                <div class="col-md-6">
                    <label class="input_label">
                        <input type="text" name="first-name" v-model="user.name" required="required" @change="userInfoHasChanged()" :disabled="!(isUserLoggedIn())">
                        <span class="keep_hovered">First Name</span>
                    </label>
                </div>
                <div class="col-md-6">
                    <label class="input_label">
                        <input type="text" name="last-name" v-model="user.surname" required="required" @change="userInfoHasChanged()" :disabled="!(isUserLoggedIn())">
                        <span class="keep_hovered">Last Name</span>
                    </label>
                </div>
            </div>
            <div class="row d-flex mt-4">
                <div class="col-md-12">
                    <label class="input_label">
                        <input type="text" name="phone" v-model="user.phone" required="required" @change="userInfoHasChanged()" :disabled="!(isUserLoggedIn())">
                        <span class="keep_hovered">Contact Phone</span>
                    </label>
                </div>
            </div>
      <input type="button" value="Update" v-if="isUserInfoChanged" :disabled="!isAllInputValid()" v-on:click="updateUser()"/>
      <input type="button" value="Reset" v-if="isUserInfoChanged" v-on:click="loadUserData()"/>
    </div>
  </div>
  <div class="posts">

    <div class="create-new" v-if="isUserLoggedIn()">
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
      <div v-if="this.usersPosts.length==0">
        <h3>There is no posts yet!</h3>

      </div>
      <div v-for="(post, index) in usersPosts" :key="index">
        <div class="post-view">
          <div class="post-view-nav" style="display: flex; flex-direction: row">
            <div class="post-view-person-icon">
              <svg xmlns="http://www.w3.org/2000/svg" width="26" height="26"
                   fill="currentColor" class="bi bi-person-circle" viewBox="0 0 16 16">
                <path d="M11 6a3 3 0 1 1-6 0 3 3 0 0 1 6 0z"/>
                <path fill-rule="evenodd" d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8zm8-7a7 7 0 0 0-5.468 11.37C3.242 11.226 4.805 10 8 10s4.757 1.225 5.468 2.37A7 7 0 0 0 8 1z"/>
              </svg>
            </div>
            <div class="post-view-username">
              <h5>{{user.username}}</h5>
            </div>
            <div class="post-view-date" style="margin-left: 50%">
              <h5>{{post.date}}</h5>

            </div>
          </div>
          <div class="post-view-image" style="margin: 2%">
          <img v-bind:src="post.image" style="width: 100%">
          </div>
          <div class="post-txt" style="position:relative;">
            <h6>{{ post.posttext }}</h6>
          </div>
          <div class="post-links">
            <div v-if="post.links.length==0">no links</div>
            <div v-for="(link, index) in post.links" :key="index">
              <a href="link"><h7>{{ link}}</h7></a>

            </div>

          </div>
          <div class="post-additiona" style="display: flex; flex-direction: row; width: 100%">
            <div class="likes" style="width: 33%;">
              <button type="button" data-bs-popper="" style="width: 100%; padding: 2%">
                <div class="likes-icon">
                  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-hand-thumbs-up-fill" viewBox="0 0 16 16">
                    <path d="M6.956 1.745C7.021.81 7.908.087 8.864.325l.261.066c.463.116.874.456 1.012.965.22.816.533 2.511.062 4.51a9.84 9.84 0 0 1 .443-.051c.713-.065 1.669-.072 2.516.21.518.173.994.681 1.2 1.273.184.532.16 1.162-.234 1.733.058.119.103.242.138.363.077.27.113.567.113.856 0 .289-.036.586-.113.856-.039.135-.09.273-.16.404.169.387.107.819-.003 1.148a3.163 3.163 0 0 1-.488.901c.054.152.076.312.076.465 0 .305-.089.625-.253.912C13.1 15.522 12.437 16 11.5 16H8c-.605 0-1.07-.081-1.466-.218a4.82 4.82 0 0 1-.97-.484l-.048-.03c-.504-.307-.999-.609-2.068-.722C2.682 14.464 2 13.846 2 13V9c0-.85.685-1.432 1.357-1.615.849-.232 1.574-.787 2.132-1.41.56-.627.914-1.28 1.039-1.639.199-.575.356-1.539.428-2.59z"/>
                  </svg>
                  <br>
                  <h5 style="margin-top: 2%">
                    Like
                  </h5>
                </div>
              </button>
            </div>
            <div class="dislikes" style="width: 33%;">
              <button style="width: 100%; padding: 2%">
                <div class="likes-icon">
                  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-hand-thumbs-down-fill" viewBox="0 0 16 16">
                    <path d="M6.956 14.534c.065.936.952 1.659 1.908 1.42l.261-.065a1.378 1.378 0 0 0 1.012-.965c.22-.816.533-2.512.062-4.51.136.02.285.037.443.051.713.065 1.669.071 2.516-.211.518-.173.994-.68 1.2-1.272a1.896 1.896 0 0 0-.234-1.734c.058-.118.103-.242.138-.362.077-.27.113-.568.113-.856 0-.29-.036-.586-.113-.857a2.094 2.094 0 0 0-.16-.403c.169-.387.107-.82-.003-1.149a3.162 3.162 0 0 0-.488-.9c.054-.153.076-.313.076-.465a1.86 1.86 0 0 0-.253-.912C13.1.757 12.437.28 11.5.28H8c-.605 0-1.07.08-1.466.217a4.823 4.823 0 0 0-.97.485l-.048.029c-.504.308-.999.61-2.068.723C2.682 1.815 2 2.434 2 3.279v4c0 .851.685 1.433 1.357 1.616.849.232 1.574.787 2.132 1.41.56.626.914 1.28 1.039 1.638.199.575.356 1.54.428 2.591z"/>
                  </svg>
                  <br>
                  <h5 style="margin-top: 2%">
                    Dislike
                  </h5>
                </div>
              </button>

            </div>
            <div class="comments" style="width: 33%;">
              <button style="width: 100%; padding: 2%">
                <div class="likes-icon">
                  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-chat-right-dots" viewBox="0 0 16 16">
                    <path d="M2 1a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h9.586a2 2 0 0 1 1.414.586l2 2V2a1 1 0 0 0-1-1H2zm12-1a2 2 0 0 1 2 2v12.793a.5.5 0 0 1-.854.353l-2.853-2.853a1 1 0 0 0-.707-.293H2a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2h12z"/>
                    <path d="M5 6a1 1 0 1 1-2 0 1 1 0 0 1 2 0zm4 0a1 1 0 1 1-2 0 1 1 0 0 1 2 0zm4 0a1 1 0 1 1-2 0 1 1 0 0 1 2 0z"/>
                  </svg>
                  <br>
                  <h5 style="margin-top: 2%">
                    Comments
                  </h5>
                </div>
              </button>
            </div>
          </div>

        </div>

      </div>
    </div>
  </div>
</div>
</template>

<script>
import PictureInput from 'vue-picture-input'
import PostService from '../services/PostService'
import UserService from '../services/UserService'
export default {
  name: "Profile",
  data(){
    return{
      postText: '',
      image: '',
      links: new Array(),
      link: '',
      user: {},
      loggedUser: {},
      originalUser: {},
      usersPosts: new Array(),
      isUserInfoChanged: false,
    }
  },
  components:{
    PictureInput
  },
  mounted: function() {
    UserService.getUserByUsername(this.$route.params.id).then(response => {
      this.user = response.data.User
      this.originalUser = this.user
      PostService.getAllPostsByUser(this.user.id).then(response1 =>{
        this.usersPosts = response1.data.posts;
        console.log("UserPosts", this.usersPosts)

      })
      .catch(err1 =>{
        console.log(err1)
        alert("User posts are unavailable!")
      })
    })
        .catch(err => {
          console.error(err);
          if(err.response.status == 403)
            this.$router.push("/unauthorized")
        })
    this.isUserInfoChanged = false
    this.loggedUser = localStorage.getItem('user')

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
        "user": this.originalUser.id,
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
    loadUserData() {
    },
    updateUser() {
        UserService.updateUser(this.user).then(() => {
            this.$router.go();
        });
    },
    userInfoHasChanged() {
        this.isUserInfoChanged = true
    },
    isAllInputValid() {
        if(this.user.firstName == "")
            return false
        if(this.user.lastName == "")
            return false
        if(this.user.contactPhone == "")
            return false
        if(this.user.address == "")
            return false
        if(this.user.city == "")
            return false
        if(this.user.country == "")
            return false
        if((this.user.oldPasswordGuess && (!this.user.newPassword || !this.user.newPasswordConfirmation))
            || (this.user.newPassword && (!this.user.oldPasswordGuess || !this.user.newPasswordConfirmation))
            || (this.user.newPasswordConfirmation && (!this.user.oldPasswordGuess || !this.user.newPassword)))
            return false;
        return true
    },
    isUserLoggedIn() {
      return this.loggedUser && JSON.parse(this.loggedUser).username == this.user.username
    }
  }

}
</script>

<style scoped>
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

.profile-info{
  width: 30%;
  height: 100%;
  background-color: #151560;
  border: #c80000 solid 2pt;
}

.posts{
  width: 70%;
  height: 100%;
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
}

    .login-reg-panel{
        position: relative;
        top: 50%;
        transform: translateY(50%);
        text-align:center;
        width:50%;
        right:0;left:0;
        margin:auto;
        height:400px;
        background: rgb(9,53,121);
        background: linear-gradient(90deg, rgba(9,53,121,1) 0%, rgba(9,51,121,1) 35%, rgba(0,95,255,1) 100%);
    }
    .profile-panel{
        background-color: rgba(255,255, 255, 1);
        position: relative;
        top: 20%;
        right:0;left:0;
        width:40%;
        margin:auto;
        margin-top: 20px;
        text-align:center;
        transition:.3s ease-in-out;
        z-index:0;
        box-shadow: 0 0 15px 9px #00000096;
    }
    .login-reg-panel input[type="radio"]{
        position:relative;
        display:none;
    }
    .login-reg-panel{
        color:#B8B8B8;
    }
    .login-reg-panel #label-login, 
    .login-reg-panel #label-register{
        border:1px solid white;
        padding:5px 5px;
        width:150px;
        display:block;
        text-align:center;
        border-radius:10px;
        cursor:pointer;
        font-weight: 600;
        font-size: 18px;
    }
    .login-info-box{
        width:30%;
        padding:0 50px;
        top:20%;
        left:0;
        position:absolute;
        text-align:left;
    }
    .register-info-box{
        width:30%;
        padding:0 50px;
        top:20%;
        right:0;
        position:absolute;
        text-align:left;
        color: whitesmoke
    }
    .right-log{right:50px !important;}

    .register-show{
        z-index: 1;
        transition:0.3s ease-in-out;
        color:#242424;
        text-align:left;
        padding:50px;
    }
    
    .register-show input[type="button"] {
        max-width: 150px;
        width: 100%;
        margin-left: 20px;
        background: rgba(0,95,255,1);
        color: #f9f9f9;
        border: none;
        padding: 10px;
        text-transform: uppercase;
        border-radius: 10px;
        float:right;
        cursor:pointer;
        transition: all .3s;
    }

    .register-show input[type="button"]:disabled {
        transform: scale(110%);
        background: rgb(162, 196, 255);
        color: #f9f9f9;
        border: none;
        padding: 10px;
        text-transform: uppercase;
        border-radius: 10px;
        float:right;
        cursor:auto;
        transition: all .3s;
    }

    .register-show input[type="button"]:hover:enabled {
        transform: scale(110%);
        background: rgba(0,95,255,1);
        color: #f9f9f9;
        border: none;
        padding: 10px;
        text-transform: uppercase;
        border-radius: 10px;
        float:right;
        cursor:pointer;
        transition: all .3s;
    }

    /* INPUT FIELDS */
    .wrapper {
        max-width: 560px;
        margin: 100px auto;
    }
    label.input_label {
        position: relative;
        display: block;
        margin: 0 0 40px 0;
    }
    label.input_label > input {
        position: relative;
        background-color: transparent;
        border: none;
        border-bottom: 1px solid #9e9e9e;
        border-radius: 0;
        outline: none;
        height: 45px;
        width: 100%;
        font-size: 16px;
        padding: 0;
        box-shadow: none;
        box-sizing: content-box;
        transition: all .3s;
    }
    label.input_label:hover > input {
        position: relative;
        background-color: transparent;
        border: none;
        border-bottom: 1px solid rgb(143, 176, 233);
        border-radius: 0;
        outline: none;
        height: 45px;
        width: 100%;
        font-size: 16px;
        padding: 0;
        box-shadow: none;
        box-sizing: content-box;
        transition: all .3s;
    }
    label.input_label > input:valid {
        border-bottom: 1px solid rgba(0,95,255,1);
        box-shadow: 0 1px 0 0 rgba(0,95,255,1);
    }
    label.input_label > span {
        color: #9e9e9e;
        position: absolute;
        top: -10px;
        left: 0;
        font-size: 16px;
        cursor: text;
        transition: .2s ease-out;
    }
    label.input_label > input:focus + span {
        transform: translateY(-25px) scale(0.8);
        transform-origin: 0;
        color: rgba(0,95,255,1);
    }
    label.input_label > input:focus {
        border-bottom: 1px solid rgba(0,95,255,1);
        box-shadow: 0 1px 0 0 rgba(0,95,255,1);
    }   
    label.input_label > select {
        position: relative;
        background-color: transparent;
        border: none;
        border-bottom: 1px solid #9e9e9e;
        border-radius: 0;
        outline: none;
        height: 45px;
        width: 100%;
        font-size: 16px;
        margin: 0 0 0 0;
        padding: 0;
        box-shadow: none;
        box-sizing: content-box;
        transition: all .3s;
    }
    label.input_label > select:valid {
        border-bottom: 1px solid rgba(0,95,255,1);
        box-shadow: 0 1px 0 0 rgba(0,95,255,1);
    }
    label.input_label > span {
        color: #9e9e9e;
        position: absolute;
        top: -10px;
        left: 0;
        font-size: 16px;
        cursor: text;
        transition: .2s ease-out;
    }
    label.input_label > select:focus + span {
        transform: translateY(-25px) scale(0.8);
        transform-origin: 0;
        color: rgba(0,95,255,1);
    }
    label.input_label > select:valid + span.keep_hovered {
        transform: translateY(-25px) scale(0.8);
        transform-origin: 0;
        color: rgba(0,95,255,1);
    }
    label.input_label > input:valid + span.keep_hovered {
        transform: translateY(-25px) scale(0.8);
        transform-origin: 0;
        color: rgba(0,95,255,1);
    }
    label.input_label > select:focus {
        border-bottom: 1px solid rgba(0,95,255,1);
        box-shadow: 0 1px 0 0 rgba(0,95,255,1);
    }

    .deletion-request-side-panel{
        background-color: rgba(255,255, 255, 1);
        position:absolute;
        width: 40%;
        right:calc(75% + 200px);
        z-index:0;
        box-shadow: 0 0 15px 9px #00000096;
        color:#242424;
        text-align:left;
        padding:50px;
    }
    .deletion-request-side-panel textarea {
        border:1px solid rgb(143, 176, 233);
        border-radius: 10px;
        resize: none;
        height: 400px;
        width: 100%;
        transition: all .3s;
    }

    .deletion-request-side-panel textarea:hover {
        border:1px solid rgba(0,95,255,1);
        border-radius: 10px;
        resize: none;
        height: 400px;
        width: 100%;
        transition: all .3s;
    }

    .deletion-request-side-panel textarea:focus {
        border:1px solid rgba(0,95,255,1);
        border-radius: 10px;
        resize: none;
        height: 400px;
        width: 100%;
        transition: all .3s;
    }

    /* GENDER SELECTION */  
    input[type="radio"] {
        display: none;
    }

    input[type="radio"] + label {
        z-index: 10;
        margin: 0 10px 10px 0;
        position: relative;
        color: #ced4da;
        text-shadow: 0 1px 0 rgba(255, 255, 255, 0.1);
        font-weight: bold;
        background-color: #ffffff;
        border: 2px solid #ced4da;
        cursor: pointer;
        transition: all 200ms ease;
    }

    input[type="radio"].male_option:hover + label {
        color: white;
        background-color: rgb(143, 176, 233);
        border: 2px solid white;
        transition: all .3s;
    }

    input[type="radio"].male_option:checked + label {
        color: white;
        background-color: rgba(0,95,255,1);
        border: 2px solid white;
        transition: all .3s;
    }

    input[type="radio"].female_option:hover + label {
        color: white;
        background-color: #ffd1d9;
        border: 2px solid white;
        transition: all .3s;
    }

    input[type="radio"].female_option:checked + label {
        color: white;
        background-color: #FB9AAC;
        border: 2px solid white;
        transition: all .3s;
    }

    input[type="radio"] + label {
        padding: 5px 40%;
        border-radius: 10px;
    }

    .user-profile-red-button {
    max-width: 150px;
    width: 100%;
    background: rgb(228, 40, 40);
    color: #f9f9f9;
    border: none;
    padding: 10px;
    border-radius: 10px;
    text-transform: uppercase;
    float:right;
    cursor:pointer;
}

</style>