<template>
  <div>
  <div style="display: flex; flex-direction: row" >
      <div style="width: 80%; padding: 5%">
        <div style="margin-bottom: 2%">
          <h2>Generate API token</h2>
          <div style="display: flex; flex-direction: row">
            <input class="form-control" v-model="apiKey" readonly="" style="margin-top: 1.5%" type="text"/>
            <input type="button" class="btn-primary btn-round" style=" border: 0" value="Generate" v-on:click="generateAPI()"/>
          </div>
          <p v-if="apiKey!=''" style="text-align: left; font-size: 10pt" class="description">It is valid for 30 days!</p>


        </div>
        <div >
            <div>
              <h2>Edit your profile</h2>
            </div>
          <table style="width: 100%">
            <tr>
              <td style="width: 30%">Name</td>
              <td>
                <input class="form-control" type="text" name="first-name" v-model="loggedUserNew.name" required="required" @change="userInfoHasChanged()" >
              </td>
            </tr>
            <tr>
              <td style="width: 30%">Surname</td>
              <td>
                <input type="text" class="form-control" name="last-name" v-model="loggedUserNew.surname" required="required" @change="userInfoHasChanged()">
              </td>
            </tr>
            <tr>
              <td style="width: 30%">Username</td>
              <td>
                <input type="text" class="form-control" name="email" v-model="loggedUserNew.username" disabled=yes required="required" @change="userInfoHasChanged()">
              </td>
            </tr>
            <tr>
              <td>Email</td>
              <td>
                <input type="email" class="form-control" name="email" v-model="loggedUserNew.email" disabled=yes required="required" @change="userInfoHasChanged()">
              </td>
            </tr>
            <tr>
              <td>Contact phone</td>
              <td>
                <input class="form-control" type="text" name="phone" v-model="loggedUserNew.phone" required="required" @change="userInfoHasChanged()" >
              </td>
            </tr>
            <tr>
              <td>Birthday</td>
              <td>
                <input class="form-control" placeholder="dd.mm.yyyy." type="datetime-local" name="dateOfBirth" v-model="loggedUserNew.dateOfBirth" required="required" @change="userInfoHasChanged()" >
              </td>
            </tr>
            <tr>
              <td colspan="2">
                <h4 style="width: 100%" class="description">
                  Change password
                </h4>
              </td>

            </tr>
            <tr>
              <td>Old password</td>
              <td>
                <input class="form-control" type="password" name="old-password" v-model="loggedUserNew.oldPasswordGuess" @change="userInfoHasChanged()">

              </td>
            </tr>
            <tr>
              <td>New password</td>
              <td>
                <input class="form-control" type="password" name="new-password" v-model="loggedUserNew.newPassword" @change="userInfoHasChanged()" >

              </td>
            </tr>
            <tr>
              <td>Repeat new password</td>
              <td>
                <input class="form-control" type="password" name="confirm-new-password" v-model="loggedUserNew.newPasswordConfirmation" @change="userInfoHasChanged()" >
              </td>
            </tr>
            <tr>
<td colspan="2"><hr></td>
            </tr>
            <tr>
              <td colspan="2">
                <h4 style="width: 100%" class="description">
                  Additional information
                </h4>
              </td>
            </tr>
            <tr>
              <td colspan="2">
                <div class="form-floating" style="border-top: 0.5pt solid #e95e38; border-left: 0.5pt solid #e95e38; ">
                  <textarea style="height: 180px" id="biographyArea"  class="biography-textarea form-control" colspan="2" placeholder="Write biography" v-model="loggedUserNew.biography" @change="userInfoHasChanged()" ></textarea>
                  <label for="biographyArea" style="font-size: 12pt">Biography</label>
                </div>



              </td>
            </tr>
            <tr>
              <td colspan="2">
                <hr>
              </td>
            </tr>
            <tr>
              <td v-bind:rowspan="2" style="vertical-align: top">Skills</td>
              <td>
                <input class="form-control" placeholder="New skill" type="text" name="new-skill" v-model="newSkill">
              </td>
              <td>
                <button class="btn-primary btn-round btn-icon" style="border: 0; padding-top:7%; padding-bottom: 0; padding-right: 25%; padding-left: 25%;" v-on:click="addSkill()" v-if="newSkill">
                  <i class="now-ui-icons ui-1_simple-add"></i>
                </button>

              </td>
            </tr>
            <tr >
            <td>
              <div style="display: flex; flex-direction: row; margin-right: 2%" v-for="(skill,index) in loggedUserNew.skills" :key="index">
                <badge type="primary">
                  {{skill.name}}
                </badge>
                <badge type="default" style="border-radius: 50%; cursor: pointer;" v-on:click="removeSkill(skill.name)">
                  ✕
                </badge>
              </div>

            </td>

            </tr>
            <tr>
              <td colspan="2">
                <hr>
              </td>
            </tr>
            <tr>
              <td v-bind:rowspan="2" style="vertical-align: top">Interests</td>
              <td>
                <input class="form-control" placeholder="New interest" type="text" name="new-interest" v-model="newInterest">
              </td>
              <td>
                <button class="btn-primary btn-round btn-icon" style="border: 0; padding-top:7%; padding-bottom: 0; padding-right: 25%; padding-left: 25%;" v-on:click="addInterest()" v-if="newInterest">
                  <i class="now-ui-icons ui-1_simple-add"></i>
                </button>

              </td>
            </tr>
            <tr >
              <td>
                <div style="display: flex; flex-direction: row; margin-right: 2%" v-for="(interest,index) in loggedUserNew.interests" :key="index">
                  <badge type="primary">
                    {{interest.name}}
                  </badge>
                  <badge type="default" style="border-radius: 50%; cursor: pointer;" v-on:click="removeInterest(interest.name)">
                    ✕
                  </badge>
                </div>

              </td>

            </tr>
            <tr>
              <td colspan="2">
                <hr>
              </td>
            </tr>
            <tr>
              <td v-bind:rowspan="5" style="vertical-align: top">Education</td>
              <td>
                <select class="form-control" v-model="newEducation.type">
                  <option selected value="PRIMARY_EDUCATION">Primary</option>
                  <option value="SECONDARY_EDUCATION">Secondary</option>
                  <option value="COLLEGE_EDUCATION">College</option>
                </select>

              </td>
            </tr>
            <tr>
              <td>
                <input placeholder="Institution name" class="form-control" type="text" name="new-education-institution" v-model="newEducation.institutionName">
              </td>
            </tr>
            <tr>
              <td>
                <div style="display: flex; flex-direction: row;">
                  <input type="datetime-local" class="form-control" v-model="newEducation.startDate" required="required" @change="userInfoHasChanged()">
                  <input type="datetime-local" class="form-control" v-model="newEducation.endDate" required="required" @change="userInfoHasChanged()" >

                </div>
              </td>
            </tr>
            <tr>
              <td>
                <button class="btn-primary btn-round" style="width: 100%; border: 0pt;" v-on:click="addEducation()" v-if="newEducation.type && newEducation.institutionName && newEducation.startDate && newEducation.endDate">Add</button>

              </td>
            </tr>
            <tr>
              <td>
                <div v-for="(experience,index) in loggedUserNew.educationExperiences" :key="index" class="row d-flex mt-4" style="border-bottom: 1px solid gray">
                  <div class="col-md-6">
                    <p>{{formattedEducationType(experience.type)}}</p>
                    <h4>{{experience.institutionName}}</h4>
                  </div>
                  <div class="col-md-4">
                    <p>{{experience.startDate}} - {{experience.endDate}}</p>
                  </div>
                  <div class="col-md-2">
                    <div class="button_minus" v-on:click="removeEducation(index)" v-if="isUserLoggedIn()"></div>
                  </div>
                </div>
              </td>
            </tr>
            <tr>
              <td colspan="2">
                <hr>
              </td>
            </tr>
            <tr>
              <td v-bind:rowspan="5" style="vertical-align: top">Work</td>
              <td>
                <input type="text" class="form-control" v-model="newWork.organizationName" placeholder="Organization name">
              </td>
            </tr>
            <tr>
              <td>
                <input type="text" class="form-control"  v-model="newWork.positionName" placeholder="Position">

              </td>
            </tr>
            <tr>
              <div style="display: flex; flex-direction: row;">
                <input type="datetime-local" class="form-control" v-model="newWork.startDate" required="required" @change="userInfoHasChanged()">
                <input type="datetime-local" class="form-control" v-model="newWork.endDate" required="required" @change="userInfoHasChanged()" >
              </div>
            </tr>
            <tr>
              <td>
                <button class="btn-primary btn-round" style="width: 100%; border: 0pt;"  v-on:click="addWork()" v-if="newWork.positionName && newWork.organizationName && newWork.startDate && newWork.endDate">Add</button>

              </td>
            </tr>
            <tr>
              <td>
                <div v-for="(experience,index) in loggedUserNew.workExperiences" :key="index" class="row d-flex mt-4" style="border-bottom: 1px solid gray">
                  <div class="col-md-6">
                    <p>{{experience.positionName}}</p>
                    <h4>{{experience.organizationName}}</h4>
                  </div>
                  <div class="col-md-4">
                    <p>{{experience.startDate}} - {{experience.endDate}}</p>
                  </div>
                  <div class="col-md-2">
                    <div class="button_minus" v-on:click="removeWork(index)" v-if="isUserLoggedIn()"></div>
                  </div>
                </div>
              </td>
            </tr>
            <tr>
              <td colspan="2">
                <hr>
              </td>
            </tr>
            <tr>
              <td>Visibility</td>
              <td>
                <div class="form-check form-switch" style="width: 180px">
                  <span class="switch">
                    <input type="checkbox" class="switch" id="isPublic" v-model="loggedUserNew.public" @change="userInfoHasChanged()">
                    <label for="isPublic" v-if="loggedUserNew.public">Public</label>
                    <label for="isPublic" v-if="!loggedUserNew.public">Private</label>
                  </span>
                </div>
              </td>
            </tr>
            <tr>
              <td>Followed Profile Notifications</td>
              <td>
                <div class="form-check form-switch" style="width: 180px">
                  <span class="switch">
                    <input type="checkbox" class="switch" id="notificationsFollowedProfiles" v-model="loggedUserNew.notificationsFollowedProfiles" @change="userInfoHasChanged()">
                    <label for="notificationsFollowedProfiles" v-if="loggedUserNew.notificationsFollowedProfiles">On</label>
                    <label for="notificationsFollowedProfiles" v-if="!loggedUserNew.notificationsFollowedProfiles">Off</label>
                  </span>
                </div>
              </td>
            </tr>
            <tr>
              <td>Message Notifications</td>
              <td>
                <div class="form-check form-switch" style="width: 180px">
                  <span class="switch">
                    <input type="checkbox" class="switch" id="notificationsMessages" v-model="loggedUserNew.notificationsMessages" @change="userInfoHasChanged()">
                    <label for="notificationsMessages" v-if="loggedUserNew.notificationsMessages">On</label>
                    <label for="notificationsMessages" v-if="!loggedUserNew.notificationsMessages">Off</label>
                  </span>
                </div>
              </td>
            </tr>
            <tr>
              <td>Post Notifications</td>
              <td>
                <div class="form-check form-switch" style="width: 180px">
                  <span class="switch">
                    <input type="checkbox" class="switch" id="notificationsPosts" v-model="loggedUserNew.notificationsPosts" @change="userInfoHasChanged()">
                    <label for="notificationsPosts" v-if="loggedUserNew.notificationsPosts">On</label>
                    <label for="notificationsPosts" v-if="!loggedUserNew.notificationsPosts">Off</label>
                  </span>
                </div>
              </td>
            </tr>
            <tr>
              <td colspan="2">
                <hr>
              </td>
            </tr>
            <tr>
              <td>

              </td>
              <td>
                <div style="display: flex; flex-direction: row;">
                  <input type="button" class="btn-primary btn-round" style="width: 50%; border: 0" value="Update" v-if="isUserInfoChanged" :disabled="!isAllInputValid()" v-on:click="updateUser()"/>
                  <input type="button" class="btn-default btn-round"  style="width: 50%; border: 0" value="Reset" v-if="isUserInfoChanged"/>

                </div>
              </td>
            </tr>
          </table>




          </div>


        <div class="profile-panel">
        </div>
        <div class="profile-panel">
          <h4>Block user</h4>
          <p>When you block a member on Dislinkt, here's what will happen: You won't be able to access each other's profiles on Dislinkt. You won't be able to message each other on Dislinkt. You won't be able to see each other's shared content.</p>
          <button  type="button" class="btn btn-light" style=" right: 10%; margin-top: 2%; margin-left: 0.5%; border: 1pt black solid;" v-on:click="blockUser">Block user</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import UserService from "../../services/UserService";
import ConnectionService from "../../services/ConnectionService";
import {Badge} from "../../components"
import {Switch} from "../../components"

export default {
  name: "ChangeProfile",
  props: {loggedUser : Object},
  data(){
    return{
      loggedUserNew: Object,
      user: Object,
      isUserInfoChanged: false,
      newEducation: {},
      newWork: {},
      newSkill: "",
      newInterest: "",
      apiKey: ""
    }
  },
  components: {Badge, [Switch.name]: Switch},
  mounted: function () {
    this.loggedUserNew = this.loggedUser
  },
  methods:{
    updateUser() {
      UserService.updateUser( {"user": {"user" : this.loggedUserNew}, "oldUser": {"user":this.loggedUser}}).then(() => {
        this.$router.go();
      });
    },
    addSkill(){
      this.loggedUserNew.skills.push({
        id: "",
        name: this.newSkill
      })
      this.isUserInfoChanged = true
    },
    removeSkill(name){
      for(let index in this.loggedUserNew.skills) {
        if(this.loggedUserNew.skills[index].name === name) {
          this.loggedUserNew.skills.splice(index,1)
        }
      }
      this.isUserInfoChanged = true
    },
    addInterest(){
      this.loggedUserNew.interests.push({
        id: "",
        name: this.newInterest
      })
      this.isUserInfoChanged = true
    },
    removeInterest(name){
      for(let index in this.loggedUserNew.interests) {
        if(this.loggedUserNew.interests[index].name === name) {
          this.loggedUserNew.interests.splice(index,1)
        }
      }
      this.isUserInfoChanged = true
    },
    addEducation(){
      this.loggedUserNew.educationExperiences.push(this.newEducation)
      this.isUserInfoChanged = true
    },
    removeEducation(index){
      this.loggedUserNew.educationExperiences.splice(index,1)
      this.isUserInfoChanged = true
    },
    addWork(){
      this.loggedUserNew.workExperiences.push(this.newWork)
      this.isUserInfoChanged = true
    },
    removeWork(index){
      this.loggedUserNew.educationExperiences.splice(index,1)
      this.isUserInfoChanged = true
    },
    formattedEducationType(educationType){
      if(educationType === "SECONDARY_EDUCATION")
        return "Secondary"
      if(educationType == "COLLEGE_EDUCATION")
        return "College"
      return "Primary"
    },
    userInfoHasChanged() {
      this.isUserInfoChanged = true
    },
    isAllInputValid() {
      if(this.loggedUserNew.firstName == "")
        return false
      if(this.loggedUserNew.lastName == "")
        return false
      if(this.loggedUserNew.phone == "")
        return false
      if((this.loggedUserNew.oldPasswordGuess && (!this.loggedUserNew.newPassword || !this.loggedUserNew.newPasswordConfirmation))
          || (this.loggedUserNew.newPassword && (!this.loggedUserNew.oldPasswordGuess || !this.loggedUserNew.newPasswordConfirmation))
          || (this.loggedUserNew.newPasswordConfirmation && (!this.loggedUserNew.oldPasswordGuess || !this.loggedUserNew.newPassword)))
        return false;
      return true
    },
    isUserLoggedIn() {
      return this.loggedUser
    },
    isSomeoneLoggedIn(){
      console.log("In checking is someone logged",this.loggedUser)
      return this.loggedUser && JSON.parse(this.loggedUser).username!=null;
    },
    unblockUser(){
      if(this.connectionStatus == 'A_BLOCK_B'){
        ConnectionService.UnblockUser(this.loggedUserDetails.id, this.user.id)
            .then( response => {
              console.log("unblocking:", response);
              this.getConnectionDetails()
            })
            .catch(err => {
                  console.log(err)
                  alert("Error unblocking user!")
                }
            )
      }

    },
    blockUser(){
      if (this.connectionStatus!='A_BLOCK_B' && this.connectionStatus!='B_BLOCK_A'){
        ConnectionService.BlockUser(this.loggedUserDetails.id, this.user.id)
            .then( response => {
              console.log("blocking:", response);
              this.getConnectionDetails()
            })
            .catch(err => {
                  console.log(err)
                  alert("Error blocking user!")
                }
            )
      }
    },
    generateAPI(){
      UserService.generateApiToken({"username":this.loggedUser.username})
      .then(response =>{
        this.apiKey = response.data.apiToken
      })
      .catch(err => {
        console.log(err)
      })

    }
  }
}
</script>

<style scoped>
/* ADD BUTTON */

.button_plus {

  width: 25px;
  height: 25px;
  background: #fff;
  cursor: pointer;
  border: 2px solid #e56830;
  border-radius: 20px;
}

.button_plus:after {
  content: '';
  color: #e5e5e5;
  transform: translate(-50%, -50%);
  height: 4px;
  width: 3px;
  padding: 5%;
  background: #e56830;
  margin-left: 45%;
  top: 50%;
  left: 50%;
}

.button_plus:before {
  content: '';



  padding-right: 10%;
  padding-left: 10%;
  transform: translate(-50%, -50%);
  background: #e56830;
  width: 10px;
  height: 2%;
}

.button_plus:hover:before,
.button_plus:hover:after {
  background: #fff;
  transition: 0.2s;
}

.button_plus:hover {
  background-color: #e56830;
  transition: 0.2s;
}

/* REMOVE BUTTON */

.button_minus {
  position: absolute;
  width: 35px;
  height: 35px;
  background: #fff;
  cursor: pointer;
  border: 2px solid #cc0000;
  border-radius: 20px;
}

.button_minus:after {
  content: '-';
  position: absolute;
  transform: translate(-50%, -50%);
  height: 4px;
  width: 50%;
  background: #cc0000;
  top: 50%;
  left: 50%;
}

.button_minus:hover:before,
.button_minus:hover:after {
  background: #fff;
  transition: 0.2s;
}

.button_minus:hover {
  background-color: #cc0000;
  transition: 0.2s;
}

/* SWITCH */

.switch {
    font-size: 1rem;
    position: relative;
}
.switch input {
    position: absolute;
    height: 1px;
    width: 1px;
    background: none;
    border: 0;
    clip: rect(0 0 0 0);
    clip-path: inset(50%);
    overflow: hidden;
    padding: 0;
}
.switch input + label {
    position: relative;
    min-width: calc(calc(2.375rem * .8) * 2);
    border-radius: calc(2.375rem * .8);
    height: calc(2.375rem * .8);
    line-height: calc(2.375rem * .8);
    display: inline-block;
    cursor: pointer;
    outline: none;
    user-select: none;
    vertical-align: middle;
    text-indent: calc(calc(calc(2.375rem * .8) * 2) + .5rem);
}
.switch input + label::before, .switch input + label::after {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    width: calc(calc(2.375rem * .8) * 2);
    bottom: 0;
    display: block;
}
.switch input + label::before {
    right: 0;
    background-color: #dee2e6;
    border-radius: calc(2.375rem * .8);
    transition: 0.2s all;
}
.switch input + label::after {
    top: 2px;
    left: 2px;
    width: calc(calc(2.375rem * .8) - calc(2px * 2));
    height: calc(calc(2.375rem * .8) - calc(2px * 2));
    border-radius: 50%;
    background-color: white;
    transition: 0.2s all;
}
.switch input:checked + label::before {
    background-color: #e56830;
}
.switch input:checked + label::after {
    margin-left: calc(2.375rem * .8);
}
.switch input:focus + label::before {
    outline: none;
    box-shadow: 0 0 0 0.2rem #f34d0098;
}
.switch input:disabled + label {
    color: #868e96;
    cursor: not-allowed;
}
.switch input:disabled + label::before {
    background-color: #e9ecef;
}
.switch.switch-sm {
    font-size: 0.875rem;
}
.switch.switch-sm input + label {
    min-width: calc(calc(1.9375rem * .8) * 2);
    height: calc(1.9375rem * .8);
    line-height: calc(1.9375rem * .8);
    text-indent: calc(calc(calc(1.9375rem * .8) * 2) + .5rem);
}
.switch.switch-sm input + label::before {
    width: calc(calc(1.9375rem * .8) * 2);
}
.switch.switch-sm input + label::after {
    width: calc(calc(1.9375rem * .8) - calc(2px * 2));
    height: calc(calc(1.9375rem * .8) - calc(2px * 2));
}
.switch.switch-sm input:checked + label::after {
    margin-left: calc(1.9375rem * .8);
}
.switch.switch-lg {
    font-size: 1.25rem;
}
.switch.switch-lg input + label {
    min-width: calc(calc(3rem * .8) * 2);
    height: calc(3rem * .8);
    line-height: calc(3rem * .8);
    text-indent: calc(calc(calc(3rem * .8) * 2) + .5rem);
}
.switch.switch-lg input + label::before {
    width: calc(calc(3rem * .8) * 2);
}
.switch.switch-lg input + label::after {
    width: calc(calc(3rem * .8) - calc(2px * 2));
    height: calc(calc(3rem * .8) - calc(2px * 2));
}
.switch.switch-lg input:checked + label::after {
    margin-left: calc(3rem * .8);
}
.switch + .switch {
    margin-left: 1rem;
}

</style>