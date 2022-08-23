<template>
    <div style="width: 100%">
        <aside>
          <header>
            <input type="text" placeholder="search">
          </header>
          <ul>
            <li v-for="(user,index) in friends" :key="index" v-on:click="selectUser(user)">
                <div class="profile-icon" style="width: 20%; margin-left: 20px">
                    <div class="photo-container-small">
                        <p style="position: relative; align-content: center; margin: 13%; font-weight: bold;" v-if="user && user.name && user.surname">
                          {{user.name.charAt(0).toUpperCase()}}{{user.surname.charAt(0).toUpperCase()}}
                        </p>
                    </div>
                </div>
                <div>
                    <h2>{{user.name}} {{user.surname}}</h2>
                </div>
            </li>
          </ul>
        </aside>
        <main>
          <header v-if="selectedUser.name">
            <div>
              <h2>Chat with {{selectedUser.name}} {{selectedUser.surname}}</h2>
            </div>
          </header>
          <header v-if="!selectedUser.name">
            <div>
              <h2>No chat open</h2>
            </div>
          </header>
          <ul id="chat">
            <li :class="{ me: messageIsFromMe(message), you: !messageIsFromMe(message) }" v-for="(message,index) in relevantMessages" :key="index">
              <div class="entete">
                <h3>{{message.date}}</h3>
                <h2>Me</h2>
                <span class="status blue"></span>
              </div>
              <div class="triangle"></div>
              <div class="message">
                {{message.content}}
              </div>
            </li>
          </ul>
          <footer>
            <textarea placeholder="Type your message" v-model="messageContent"></textarea>
            <button type="button" class="btn btn-dark" style="width: 30%; height: 40px; float: right" v-on:click="sendMessage()">Send</button>
          </footer>
        </main>
    </div>
</template>

<script>
import UserService from "../../services/UserService";
import ConnectionService from "../../services/ConnectionService";
import MessageService from "../../services/MessageService";
export default {
  name: "Chatbox",
  data(){
    return{
      loggedUser: "",
      loggedUserDetails: "",
      selectedUser: {},
      friends: new Array(),
      messages: new Array(),
      messageContent: "",
    }
  },
  mounted: function () {
    this.loggedUser = localStorage.getItem('user')
    UserService.getUserByUsername(JSON.parse(this.loggedUser).username).then(response =>{
      this.loggedUserDetails = response.data.User
      ConnectionService.GetFriends(this.loggedUserDetails.id).then(response1 =>{
        var friendNodes =  response1.data.users
        for (var f of friendNodes){
          UserService.getUserById(f.userID)
          .then(response2 =>{
            console.log(response2.data)
            this.friends.push(response2.data.User)
          })
          .catch(err =>{
            console.log("User with id: ", f.userID, " does not exists in UserService, just in ConnectionService!", err)
          })
        }

      })
      .catch(err1 =>{
        alert("It is impossible to load friends!")
        console.log(err1)
      })
      MessageService.getMessagesByUser(this.loggedUserDetails.id).then(response1 => {
        this.messages = response1.data.userMessages
      })
    })
    .catch(err =>{
      alert("It s impossible to get logged user!")
      console.log(err)
    })

  },
  methods:{
    selectUser(user){
      this.selectedUser = user
    },
    messageIsFromMe(message){
      return message.fromUser == this.loggedUserDetails.id
    },
    sendMessage(){
      let messageToSend = {
        fromUser: this.loggedUserDetails.id,
        toUser: this.selectedUser.id,
        content: this.messageContent
      }
      MessageService.sendMessage(messageToSend).then( () => {
        MessageService.getMessagesByUser(this.loggedUserDetails.id).then(response1 => {
          this.messages = response1.data.userMessages
        })
      })
    }
  },
  computed: {
    relevantMessages() {
      return this.messages.filter(m => m.fromUser == this.selectedUser.id || m.toUser == this.selectedUser.id)
    }
  }
}
</script>

<style scoped>
    #container{
      width:750px;
      height:800px;
      background:#eff3f7;
      margin:0 auto;
      font-size:0;
      border-radius:5px;
      overflow:hidden;
    }
    aside{
      width:260px;
      height:800px;
      background-color:#3b3e49;
      display:inline-block;
      font-size:15px;
      vertical-align:top;
    }
    main{
      width:490px;
      height:800px;
      display:inline-block;
      font-size:15px;
      vertical-align:top;
    }

    aside header{
      padding:30px 20px;
    }
    aside input{
      width:100%;
      height:50px;
      line-height:50px;
      padding:0 50px 0 20px;
      background-color:#5e616a;
      border:none;
      border-radius:3px;
      color:#fff;
      background-image:url(https://s3-us-west-2.amazonaws.com/s.cdpn.io/1940306/ico_search.png);
      background-repeat:no-repeat;
      background-position:170px;
      background-size:40px;
    }
    aside input::placeholder{
      color:#fff;
    }
    aside ul{
      padding-left:0;
      margin:0;
      list-style-type:none;
      overflow-y:scroll;
      height:690px;
    }
    aside li{
      padding:10px 0;
    }
    aside li:hover{
      background-color:#5e616a;
    }
    h2,h3{
      margin:0;
    }
    aside li img{
      border-radius:50%;
      margin-left:20px;
      margin-right:8px;
    }
    aside li div{
      display:inline-block;
      vertical-align:top;
      margin-top:12px;
    }
    aside li h2{
      font-size:14px;
      color:#fff;
      font-weight:normal;
      margin-bottom:5px;
    }
    aside li h3{
      font-size:12px;
      color:#7e818a;
      font-weight:normal;
    }

    .status{
      width:8px;
      height:8px;
      border-radius:50%;
      display:inline-block;
      margin-right:7px;
    }
    .green{
      background-color:#58b666;
    }
    .orange{
      background-color:#ff725d;
    }
    .blue{
      background-color:#6fbced;
      margin-right:0;
      margin-left:7px;
    }

    main header{
      height:110px;
      padding:30px 20px 30px 40px;
    }
    main header > *{
      display:inline-block;
      vertical-align:top;
    }
    main header img:first-child{
      border-radius:50%;
    }
    main header img:last-child{
      width:24px;
      margin-top:8px;
    }
    main header div{
      margin-left:10px;
      margin-right:145px;
    }
    main header h2{
      font-size:16px;
      margin-bottom:5px;
    }
    main header h3{
      font-size:14px;
      font-weight:normal;
      color:#7e818a;
    }

    #chat{
      padding-left:0;
      margin:0;
      list-style-type:none;
      overflow-y:scroll;
      height:535px;
      border-top:2px solid #fff;
      border-bottom:2px solid #fff;
    }
    #chat li{
      padding:10px 30px;
    }
    #chat h2,#chat h3{
      display:inline-block;
      font-size:13px;
      font-weight:normal;
    }
    #chat h3{
      color:#bbb;
    }
    #chat .entete{
      margin-bottom:5px;
    }
    #chat .message{
      padding:20px;
      color:#fff;
      line-height:25px;
      max-width:90%;
      display:inline-block;
      text-align:left;
      border-radius:5px;
    }
    #chat .me{
      text-align:right;
    }
    #chat .you .message{
      background-color:#58b666;
    }
    #chat .me .message{
      background-color:#6fbced;
    }
    #chat .triangle{
      width: 0;
      height: 0;
      border-style: solid;
      border-width: 0 10px 10px 10px;
    }
    #chat .you .triangle{
        border-color: transparent transparent #58b666 transparent;
        margin-left:15px;
    }
    #chat .me .triangle{
        border-color: transparent transparent #6fbced transparent;
        margin-left:375px;
    }

    main footer{
      height:155px;
      padding:20px 30px 10px 20px;
    }
    main footer textarea{
      resize:none;
      border:none;
      display:block;
      width:100%;
      height:80px;
      border-radius:3px;
      padding:20px;
      font-size:13px;
      margin-bottom:13px;
    }
    main footer textarea::placeholder{
      color:#ddd;
    }
    main footer img{
      height:30px;
      cursor:pointer;
    }
    main footer a{
      text-decoration:none;
      text-transform:uppercase;
      font-weight:bold;
      color:#6fbced;
      vertical-align:top;
      margin-left:333px;
      margin-top:5px;
      display:inline-block;
    }
</style>