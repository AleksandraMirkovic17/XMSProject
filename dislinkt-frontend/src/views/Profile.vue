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
      isUserInfoChanged: false
    }
  },
  components:{
    PictureInput
  },
  mounted: function() {
    this.loadUserData()
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
    loadUserData() {
        UserService.getUserByUsername(this.$route.params.id).then(response => {
            this.user = response.data.User
            this.originalUser = this.user
        })
        .catch(err => {
            console.error(err);
            if(err.response.status == 403)
                this.$router.push("/unauthorized")
        })
        this.isUserInfoChanged = false
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