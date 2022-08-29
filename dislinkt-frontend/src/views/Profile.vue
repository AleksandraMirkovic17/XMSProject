<template>
<div>
  <div>
    <div class="page-header clear-filter" filter-color="orange">
      <parallax
          class="page-header-image"
          style="background-image:url('../assets/mountain.jpg')"
      >
      </parallax>
      <div class="container">
        <div class="photo-container" style="background-color: #b7b7b7; align-content: center; vertical-align: center">
          <h1 style="position: relative; align-content: center; margin-top: 30%;" v-if="user && user.name && user.surname">
            {{user.name.charAt(0).toUpperCase()}}{{user.surname.charAt(0).toUpperCase()}}
          </h1>
        </div>
        <h3 class="title">{{user.name}} {{user.surname}}</h3>
        <p class="category" v-if="user.username">@{{user.username.toLowerCase()}}</p>
        <div class="content">
          <div class="social-description">
            <h2>26</h2>
            <p>Connections</p>
          </div>
          <div class="social-description">
            <h2>{{usersPosts.length}}</h2>
            <p>Posts</p>
          </div>
          <div class="social-description">
            <h2>26</h2>
            <p>Job offers</p>
          </div>
        </div>
      </div>
    </div>
    <div class="section" style="display: flex; flex-direction: row">
      <nav v-if="user.username==loggedUserDetails.username">
        <div class="d-flex flex-column flex-shrink-0 p-3 bg-light" style="width: 280px;" >
          <ul id="sidebar" class="nav nav-pills flex-column mb-auto">
            <li v-on:click="DisplayFeed">
              <a href="#" class="nav-link link-dark" aria-current="page">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-house-door-fill" viewBox="0 0 16 16">
                  <path d="M6.5 14.5v-3.505c0-.245.25-.495.5-.495h2c.25 0 .5.25.5.5v3.5a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-7a.5.5 0 0 0-.146-.354L13 5.793V2.5a.5.5 0 0 0-.5-.5h-1a.5.5 0 0 0-.5.5v1.293L8.354 1.146a.5.5 0 0 0-.708 0l-6 6A.5.5 0 0 0 1.5 7.5v7a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5z"/>
                </svg>
                Feed
              </a>
            </li>
            <li v-on:click="DisplayProfile">
              <a href="#anchor1" class="nav-link link-dark" aria-current="page">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-postcard-fill" viewBox="0 0 16 16">
                  <path d="M11 8h2V6h-2v2Z"/>
                  <path d="M0 4a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V4Zm8.5.5a.5.5 0 0 0-1 0v7a.5.5 0 0 0 1 0v-7ZM2 5.5a.5.5 0 0 0 .5.5H6a.5.5 0 0 0 0-1H2.5a.5.5 0 0 0-.5.5ZM2.5 7a.5.5 0 0 0 0 1H6a.5.5 0 0 0 0-1H2.5ZM2 9.5a.5.5 0 0 0 .5.5H6a.5.5 0 0 0 0-1H2.5a.5.5 0 0 0-.5.5Zm8-4v3a.5.5 0 0 0 .5.5h3a.5.5 0 0 0 .5-.5v-3a.5.5 0 0 0-.5-.5h-3a.5.5 0 0 0-.5.5Z"/>
                </svg>
                My posts
              </a>
            </li>
            <li v-on:click="DisplayChat">
              <a href="#" class="nav-link link-dark" aria-current="page">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-chat" viewBox="0 0 16 16">
                <path d="M2.678 11.894a1 1 0 0 1 .287.801 10.97 10.97 0 0 1-.398 2c1.395-.323 2.247-.697 2.634-.893a1 1 0 0 1 .71-.074A8.06 8.06 0 0 0 8 14c3.996 0 7-2.807 7-6 0-3.192-3.004-6-7-6S1 4.808 1 8c0 1.468.617 2.83 1.678 3.894zm-.493 3.905a21.682 21.682 0 0 1-.713.129c-.2.032-.352-.176-.273-.362a9.68 9.68 0 0 0 .244-.637l.003-.01c.248-.72.45-1.548.524-2.319C.743 11.37 0 9.76 0 8c0-3.866 3.582-7 8-7s8 3.134 8 7-3.582 7-8 7a9.06 9.06 0 0 1-2.347-.306c-.52.263-1.639.742-3.468 1.105z"/>
              </svg>
                Chat
              </a>
            </li>
              <li  v-on:click="DisplayConnections">
                <a href="#" class="nav-link link-dark">
                  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-globe2" viewBox="0 0 16 16">
                    <path d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8zm7.5-6.923c-.67.204-1.335.82-1.887 1.855-.143.268-.276.56-.395.872.705.157 1.472.257 2.282.287V1.077zM4.249 3.539c.142-.384.304-.744.481-1.078a6.7 6.7 0 0 1 .597-.933A7.01 7.01 0 0 0 3.051 3.05c.362.184.763.349 1.198.49zM3.509 7.5c.036-1.07.188-2.087.436-3.008a9.124 9.124 0 0 1-1.565-.667A6.964 6.964 0 0 0 1.018 7.5h2.49zm1.4-2.741a12.344 12.344 0 0 0-.4 2.741H7.5V5.091c-.91-.03-1.783-.145-2.591-.332zM8.5 5.09V7.5h2.99a12.342 12.342 0 0 0-.399-2.741c-.808.187-1.681.301-2.591.332zM4.51 8.5c.035.987.176 1.914.399 2.741A13.612 13.612 0 0 1 7.5 10.91V8.5H4.51zm3.99 0v2.409c.91.03 1.783.145 2.591.332.223-.827.364-1.754.4-2.741H8.5zm-3.282 3.696c.12.312.252.604.395.872.552 1.035 1.218 1.65 1.887 1.855V11.91c-.81.03-1.577.13-2.282.287zm.11 2.276a6.696 6.696 0 0 1-.598-.933 8.853 8.853 0 0 1-.481-1.079 8.38 8.38 0 0 0-1.198.49 7.01 7.01 0 0 0 2.276 1.522zm-1.383-2.964A13.36 13.36 0 0 1 3.508 8.5h-2.49a6.963 6.963 0 0 0 1.362 3.675c.47-.258.995-.482 1.565-.667zm6.728 2.964a7.009 7.009 0 0 0 2.275-1.521 8.376 8.376 0 0 0-1.197-.49 8.853 8.853 0 0 1-.481 1.078 6.688 6.688 0 0 1-.597.933zM8.5 11.909v3.014c.67-.204 1.335-.82 1.887-1.855.143-.268.276-.56.395-.872A12.63 12.63 0 0 0 8.5 11.91zm3.555-.401c.57.185 1.095.409 1.565.667A6.963 6.963 0 0 0 14.982 8.5h-2.49a13.36 13.36 0 0 1-.437 3.008zM14.982 7.5a6.963 6.963 0 0 0-1.362-3.675c-.47.258-.995.482-1.565.667.248.92.4 1.938.437 3.008h2.49zM11.27 2.461c.177.334.339.694.482 1.078a8.368 8.368 0 0 0 1.196-.49 7.01 7.01 0 0 0-2.275-1.52c.218.283.418.597.597.932zm-.488 1.343a7.765 7.765 0 0 0-.395-.872C9.835 1.897 9.17 1.282 8.5 1.077V4.09c.81-.03 1.577-.13 2.282-.287z"/>
                  </svg>
                  Connections
                </a>
              </li>
            <li  v-if="display=='connections'" v-on:click="DisplayConnectionsFriends">
              <a href="#" class="nav-link link-dark">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-people" viewBox="0 0 16 16">
                  <path d="M15 14s1 0 1-1-1-4-5-4-5 3-5 4 1 1 1 1h8zm-7.978-1A.261.261 0 0 1 7 12.996c.001-.264.167-1.03.76-1.72C8.312 10.629 9.282 10 11 10c1.717 0 2.687.63 3.24 1.276.593.69.758 1.457.76 1.72l-.008.002a.274.274 0 0 1-.014.002H7.022zM11 7a2 2 0 1 0 0-4 2 2 0 0 0 0 4zm3-2a3 3 0 1 1-6 0 3 3 0 0 1 6 0zM6.936 9.28a5.88 5.88 0 0 0-1.23-.247A7.35 7.35 0 0 0 5 9c-4 0-5 3-5 4 0 .667.333 1 1 1h4.216A2.238 2.238 0 0 1 5 13c0-1.01.377-2.042 1.09-2.904.243-.294.526-.569.846-.816zM4.92 10A5.493 5.493 0 0 0 4 13H1c0-.26.164-1.03.76-1.724.545-.636 1.492-1.256 3.16-1.275zM1.5 5.5a3 3 0 1 1 6 0 3 3 0 0 1-6 0zm3-2a2 2 0 1 0 0 4 2 2 0 0 0 0-4z"/>
                </svg>
                Friends
              </a>
            </li>
            <li  v-if="display=='connections'" v-on:click="DisplayConnectionsRequests">
              <a href="#" class="nav-link link-dark">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-person-plus-fill" viewBox="0 0 16 16">
                  <path d="M1 14s-1 0-1-1 1-4 6-4 6 3 6 4-1 1-1 1H1zm5-6a3 3 0 1 0 0-6 3 3 0 0 0 0 6z"/>
                  <path fill-rule="evenodd" d="M13.5 5a.5.5 0 0 1 .5.5V7h1.5a.5.5 0 0 1 0 1H14v1.5a.5.5 0 0 1-1 0V8h-1.5a.5.5 0 0 1 0-1H13V5.5a.5.5 0 0 1 .5-.5z"/>
                </svg>
                Requests
              </a>
            </li>
            <li  v-if="display=='connections'" v-on:click="DisplayConnectionsBlockeds">
              <a href="#" class="nav-link link-dark">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-person-plus-fill" viewBox="0 0 16 16">
                  <path d="M1 14s-1 0-1-1 1-4 6-4 6 3 6 4-1 1-1 1H1zm5-6a3 3 0 1 0 0-6 3 3 0 0 0 0 6z"/>
                  <path fill-rule="evenodd" d="M13.5 5a.5.5 0 0 1 .5.5V7h1.5a.5.5 0 0 1 0 1H14v1.5a.5.5 0 0 1-1 0V8h-1.5a.5.5 0 0 1 0-1H13V5.5a.5.5 0 0 1 .5-.5z"/>
                </svg>
                Blockeds
              </a>
            </li>



            <li v-on:click="DisplayJobOffers">
              <a href="#" class="nav-link link-dark">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-briefcase-fill" viewBox="0 0 16 16">
                  <path d="M6.5 1A1.5 1.5 0 0 0 5 2.5V3H1.5A1.5 1.5 0 0 0 0 4.5v1.384l7.614 2.03a1.5 1.5 0 0 0 .772 0L16 5.884V4.5A1.5 1.5 0 0 0 14.5 3H11v-.5A1.5 1.5 0 0 0 9.5 1h-3zm0 1h3a.5.5 0 0 1 .5.5V3H6v-.5a.5.5 0 0 1 .5-.5z"/>
                  <path d="M0 12.5A1.5 1.5 0 0 0 1.5 14h13a1.5 1.5 0 0 0 1.5-1.5V6.85L8.129 8.947a.5.5 0 0 1-.258 0L0 6.85v5.65z"/>
                </svg>
                Job offers
              </a>
            </li>
            <li v-on:click="DisplayMyJobOffers">
              <a href="#" class="nav-link link-dark">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-file-earmark" viewBox="0 0 16 16">
                  <path d="M14 4.5V14a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2h5.5L14 4.5zm-3 0A1.5 1.5 0 0 1 9.5 3V1H4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V4.5h-2z"/>
                </svg>
                My job offers
              </a>
            </li>
            <li  v-on:click="DisplayProfile">
              <a href="#" class="nav-link link-dark">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-person-fill" viewBox="0 0 16 16">
                  <path d="M3 14s-1 0-1-1 1-4 6-4 6 3 6 4-1 1-1 1H3zm5-6a3 3 0 1 0 0-6 3 3 0 0 0 0 6z"/>
                </svg>
                My profile
              </a>
            </li>
            <li  v-on:click="DisplaySettings">
              <a href="#" class="nav-link link-dark">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-gear-fill" viewBox="0 0 16 16">
                  <path d="M9.405 1.05c-.413-1.4-2.397-1.4-2.81 0l-.1.34a1.464 1.464 0 0 1-2.105.872l-.31-.17c-1.283-.698-2.686.705-1.987 1.987l.169.311c.446.82.023 1.841-.872 2.105l-.34.1c-1.4.413-1.4 2.397 0 2.81l.34.1a1.464 1.464 0 0 1 .872 2.105l-.17.31c-.698 1.283.705 2.686 1.987 1.987l.311-.169a1.464 1.464 0 0 1 2.105.872l.1.34c.413 1.4 2.397 1.4 2.81 0l.1-.34a1.464 1.464 0 0 1 2.105-.872l.31.17c1.283.698 2.686-.705 1.987-1.987l-.169-.311a1.464 1.464 0 0 1 .872-2.105l.34-.1c1.4-.413 1.4-2.397 0-2.81l-.34-.1a1.464 1.464 0 0 1-.872-2.105l.17-.31c.698-1.283-.705-2.686-1.987-1.987l-.311.169a1.464 1.464 0 0 1-2.105-.872l-.1-.34zM8 10.93a2.929 2.929 0 1 1 0-5.86 2.929 2.929 0 0 1 0 5.858z"/>
                </svg>
                Settings
              </a>
            </li>

          </ul>
        </div>
      </nav>
      <div  class="container" >
        <chatbox v-if="display=='chat'"></chatbox>
        <div v-if="display=='profile'">
          <div class="button-container" v-if="loggedUserDetails.username!=user.username">
            <a v-if="this.loggedUserFollows==false && connectionStatus!='ACCEPT' && connectionStatus!='PENDING' && connectionStatus!='A_BLOCK_B' && connectionStatus!='B_BLOCK_A'" href="#button" class="btn btn-primary btn-round btn-lg" v-on:click="follow">Connect</a>
            <a v-if="connectionStatus=='ACCEPT'" href="#button" class="btn btn-primary btn-round btn-lg" v-on:click="follow">✔ Accept</a>
            <a v-if="connectionStatus=='ACCEPT'" href="#button" class="btn btn-default btn-round btn-lg" v-on:click="RemoveFriendRequest">✖ Decline</a>
            <a v-if="loggedUserFollows==true" href="#button" class="btn btn-primary btn-round btn-lg" v-on:click="deleteFriend">✔ Connected</a>
            <a v-if="connectionStatus=='PENDING'" class="btn btn-primary btn-round btn-lg" v-on:click="unsendRequest">Request sent</a>
            <a v-if="connectionStatus!='A_BLOCK_B' && connectionStatus!='B_BLOCK_A'" v-on:click="blockUser" href="#button" class="btn btn-default btn-round btn-lg">Block</a>
            <a v-if="connectionStatus=='A_BLOCK_B'" v-on:click="unblockUser" href="#button" class="btn btn-default btn-round btn-lg">Unblock</a>

          </div>
          <div v-if="(loggedUserFollows || user.Public || loggedUserDetails.username==user.username)  && connectionStatus!='A_BLOCK_B' && connectionStatus!='B_BLOCK_A'">
            <h3 class="title" v-if="user.biography!=''">About me</h3>
            <h5 class="description" v-if="user.biography!=''">
              {{user.biography}}
            </h5>
            <h3 class="title">Basic information</h3>
            <div class="prinf" style="align-content: center; justify-content: center; horiz-align: center;">
              <div style="align-content: center; horiz-align: center; text-align: center; align-items: center;">

                <i slot="label"   class="now-ui-icons ui-1_calendar-60" style="width: 25pt; height:25pt;"></i>
              </div>
              <div style="text-align: center;">
                <h5 style="align-content: center; justify-content: center; margin: 0">
                  {{user.dateOfBirth}}
                </h5>
              </div>
            </div>
            <div class="prinf" style=" margin-top: 0; text-align: center; align-content: center" v-if="user.contactPhone!=''">
              <div style="align-content: center; horiz-align: center; text-align: center; align-items: center;">

                <svg xmlns="http://www.w3.org/2000/svg" width="25" height="25" style=" " fill="currentColor" class="bi bi-telephone-fill" viewBox="0 0 16 16">
                  <path fill-rule="evenodd" d="M1.885.511a1.745 1.745 0 0 1 2.61.163L6.29 2.98c.329.423.445.974.315 1.494l-.547 2.19a.678.678 0 0 0 .178.643l2.457 2.457a.678.678 0 0 0 .644.178l2.189-.547a1.745 1.745 0 0 1 1.494.315l2.306 1.794c.829.645.905 1.87.163 2.611l-1.034 1.034c-.74.74-1.846 1.065-2.877.702a18.634 18.634 0 0 1-7.01-4.42 18.634 18.634 0 0 1-4.42-7.009c-.362-1.03-.037-2.137.703-2.877L1.885.511z"/>
                </svg>
              </div>
              <div style="text-align: center;">
                <h5 style="align-content: center; justify-content: center; margin: 0">
                  {{user.contactPhone}}
                </h5>
              </div>
            </div>

            <div class="prinf" style="margin-top: 0; margin-bottom: 10%" v-if="user.email!=''">
              <div style="align-content: center; horiz-align: center; text-align: center; align-items: center;">
                <svg xmlns="http://www.w3.org/2000/svg" style=" " width="25" height="25" fill="currentColor" class="bi bi-envelope-fill" viewBox="0 0 16 16">
                  <path d="M.05 3.555A2 2 0 0 1 2 2h12a2 2 0 0 1 1.95 1.555L8 8.414.05 3.555ZM0 4.697v7.104l5.803-3.558L0 4.697ZM6.761 8.83l-6.57 4.027A2 2 0 0 0 2 14h12a2 2 0 0 0 1.808-1.144l-6.57-4.027L8 9.586l-1.239-.757Zm3.436-.586L16 11.801V4.697l-5.803 3.546Z"/>
                </svg>
              </div>
              <div style="text-align: center;">
                <h5 style="align-content: center; justify-content: center; margin: 0">
                  {{user.email}}
                </h5>
              </div>
            </div>
          </div>
          <div v-if="(loggedUserFollows || user.Public || loggedUserDetails.username==user.username)  && connectionStatus!='A_BLOCK_B' && connectionStatus!='B_BLOCK_A'" class="additional-info" style="display: flex; flex-direction: row; width: 100%">
            <div class="experience profile-panel" style="width: 55%; margin: 2%;">
              <h3>Experience</h3>
              <h4 class="description" style="text-align: left">
                Education
              </h4>
              <div v-if="user.educationExperiences && user.educationExperiences.length>0">
                <div class="education-profile" v-for="(experience,index) in user.educationExperiences" :key="index" >
                  <div style="display: flex; flex-direction: row; margin-bottom: 2%">
                    <div class="diploma" style="width: 20%; margin: 5%;">
                      <svg xmlns="http://www.w3.org/2000/svg" width="45" height="45" fill="currentColor" class="bi bi-mortarboard-fill" viewBox="0 0 16 16">
                        <path d="M8.211 2.047a.5.5 0 0 0-.422 0l-7.5 3.5a.5.5 0 0 0 .025.917l7.5 3a.5.5 0 0 0 .372 0L14 7.14V13a1 1 0 0 0-1 1v2h3v-2a1 1 0 0 0-1-1V6.739l.686-.275a.5.5 0 0 0 .025-.917l-7.5-3.5Z"/>
                        <path d="M4.176 9.032a.5.5 0 0 0-.656.327l-.5 1.7a.5.5 0 0 0 .294.605l4.5 1.8a.5.5 0 0 0 .372 0l4.5-1.8a.5.5 0 0 0 .294-.605l-.5-1.7a.5.5 0 0 0-.656-.327L8 10.466 4.176 9.032Z"/>
                      </svg>
                    </div>
                    <div class="experience information">
                      <p style="margin-bottom: 1%; font-size: 11pt; text-align: left">{{formattedEducationType(experience.type)}}</p>
                      <h5 style="margin-top: 1%; margin-bottom: 1%; text-align: left">{{experience.institutionName}}</h5>
                      <p style="margin-bottom: 1%; font-size: 11pt; text-align: left">{{experience.startDate}} - {{experience.endDate}}</p>
                    </div>

                  </div>
                </div>

              </div>
              <div v-if="user.educationExperiences && user.educationExperiences.length==0">
                <p class="nav-pills-warning">No information about education.</p>
              </div>

              <hr>
              <h4 class="description" style="text-align: left">
                Work
              </h4>
              <div v-if="user.workExperiences && user.workExperiences.length>0">
                <div  v-for="(experience,index) in user.workExperiences" :key="index">
                  <div style="display: flex; flex-direction: row; margin-bottom: 2%">
                    <div class="diploma" style="width: 20%; margin: 5%;">
                      <svg xmlns="http://www.w3.org/2000/svg" width="45" height="45" fill="currentColor" class="bi bi-briefcase-fill" viewBox="0 0 16 16">
                        <path d="M6.5 1A1.5 1.5 0 0 0 5 2.5V3H1.5A1.5 1.5 0 0 0 0 4.5v1.384l7.614 2.03a1.5 1.5 0 0 0 .772 0L16 5.884V4.5A1.5 1.5 0 0 0 14.5 3H11v-.5A1.5 1.5 0 0 0 9.5 1h-3zm0 1h3a.5.5 0 0 1 .5.5V3H6v-.5a.5.5 0 0 1 .5-.5z"/>
                        <path d="M0 12.5A1.5 1.5 0 0 0 1.5 14h13a1.5 1.5 0 0 0 1.5-1.5V6.85L8.129 8.947a.5.5 0 0 1-.258 0L0 6.85v5.65z"/>
                      </svg>
                    </div>
                    <div class="experience information">
                      <h5 style=" margin-bottom: 1%; text-align: left">{{experience.organizationName}}</h5>

                      <p style="margin-top: 1%; margin-bottom: 1%; font-size: 11pt; text-align: left">{{experience.positionName}}</p>
                      <p style="margin-bottom: 1%; font-size: 11pt; text-align: left">{{experience.startDate}} - {{experience.endDate}}</p>
                    </div>

                  </div>
                </div>

              </div>
              <div v-if="user.workExperiences.length ==0">
                <p class="nav-pills-warning">No information about job positions.</p>
              </div>
            </div>
            <div class="skills profile-panel" style="width: 35%; margin: 2%;">
              <h3 style="text-align: left">Skills({{user.skills.length}})</h3>
              <div v-if="user.skills.length>0" style="text-align: left">
                <badge type="primary"  style="margin-right: 1%" v-for="(skill,index) in user.skills" :key="index" class="primary">
                  {{skill.name}}
                </badge>

              </div>
              <hr>
              <h3 style="text-align: left">Interests({{user.interests.length}})</h3>
              <div v-if="user.skills.length>0" style="text-align: left">
                <badge type="info" style="margin-right: 1%" v-for="(interest,index) in user.interests" :key="index" >
                  {{interest.name}}
                </badge>

              </div>
            </div>
          </div>
          <div class="row" v-if="(loggedUserFollows || user.Public || loggedUserDetails.username==user.username)  && connectionStatus!='A_BLOCK_B' && connectionStatus!='B_BLOCK_A'">
            <tabs
                pills
                class="nav-align-center"
                tab-content-classes="gallery"
                tab-nav-classes="nav-pills-just-icons"
                type="primary"
            >
              <a id="anchor1"></a>

              <tab-pane selected title="Posts">

                <i slot="label" class="now-ui-icons education_paper"></i>
                <div>
                  <h3>Posts</h3>
                  <div class="create-new" style="border: 0.2pt solid #e95e38; padding: 2%; " v-if="loggedUserDetails.username == user.username">
                    <div class="head-line" style="display: flex; flex-direction: row">
                      <svg xmlns="http://www.w3.org/2000/svg" width="30" height="30" fill="white" class="bi bi-file-earmark-post-fill" viewBox="0 0 16 16">
                        <path d="M9.293 0H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V4.707A1 1 0 0 0 13.707 4L10 .293A1 1 0 0 0 9.293 0zM9.5 3.5v-2l3 3h-2a1 1 0 0 1-1-1zm-5-.5H7a.5.5 0 0 1 0 1H4.5a.5.5 0 0 1 0-1zm0 3h7a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-.5.5h-7a.5.5 0 0 1-.5-.5v-7a.5.5 0 0 1 .5-.5z"/>
                      </svg>
                      <h4 class="description" style="margin-left: 1%; margin-top: 0.2%; ">New post</h4>
                    </div>
                    <div class="post-content" style="display: flex; flex-direction: row" >
                      <div class="hello" style="margin-top: 1%">
                        <picture-input
                            ref="pictureInput"
                            width="250"
                            height="220"
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
                            <textarea class="form-control" placeholder="Write a post..." id="floatingTextarea" style="font-size: 12pt; height: 180px" v-model = "postText"></textarea>
                            <label for="floatingTextarea" style="font-size: 12pt">Post text</label>
                          </div>
                        </div>
                        <div class="post-link" style="display: flex; flex-direction: row">
                          <div class="input-field">
                            <input type="text" class="form-control" id="exampleFormControlInput1" v-model="link" placeholder="Add link">
                          </div>
                          <button type="button" class="btn btn-light" style="border-radius: 90%" v-on:click="addLink()">+</button>
                        </div>
                        <p v-for="(link, index) in links" :key="index" class="link-dark" style="text-align: left">
                          <a>{{link}}</a>
                        </p>
                      </div>
                    </div>
                    <div class="footer">
                      <button type="button" class="btn btn-primary" v-on:click="CreatePost">Post</button>
                    </div>
                  </div>
                  <Posts :feed-posts="false" :users-posts="true" :userid="userID" :username="empty"></Posts>

                </div>
              </tab-pane>

              <tab-pane title="Job offers">
                <a id="anchor2"></a>

                <i slot="label" class="now-ui-icons business_briefcase-24"></i>
                <div>
                  <h3>Job offers</h3>
                </div>
                <div class="create-new" style="border: 0.2pt solid #e95e38; padding: 2%; " v-if="loggedUserDetails.username == user.username">
                  <div class="head-line" style="display: flex; flex-direction: row">
                    <svg xmlns="http://www.w3.org/2000/svg" width="30" height="30" fill="white" class="bi bi-file-earmark-post-fill" viewBox="0 0 16 16">
                      <path d="M9.293 0H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V4.707A1 1 0 0 0 13.707 4L10 .293A1 1 0 0 0 9.293 0zM9.5 3.5v-2l3 3h-2a1 1 0 0 1-1-1zm-5-.5H7a.5.5 0 0 1 0 1H4.5a.5.5 0 0 1 0-1zm0 3h7a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-.5.5h-7a.5.5 0 0 1-.5-.5v-7a.5.5 0 0 1 .5-.5z"/>
                    </svg>
                    <h4 class="description" style="margin-left: 1%; margin-top: 0.2%; ">Publish new job offer</h4>
                  </div>
                  <div class="post-content" style="display: flex; flex-direction: row" >
                    <div class="content2" style="width: 100%">
                      <div class="post-text" >
                        <div class="form-floating" style="margin-bottom: 2%;">
                          <input id="position" type="text" class="form-control" v-model="newJobOffer.companyName" placeholder="Organization name">
                          <label for="position" style="font-size: 12pt">Organization</label>
                        </div>
                        <div class="form-floating">
                          <input id="company" type="text" class="form-control" v-model="newJobOffer.position" placeholder="Organization name">
                          <label for="company" style="font-size: 12pt">Position</label>
                        </div>
                        <div class="form-floating" style="margin-bottom: 3%">
                          <textarea class="form-control" placeholder="Job description..." id="jobDescription" style="font-size: 12pt; height: 180px" v-model = "newJobOffer.jobDescription"></textarea>
                          <label for="jobDescription" style="font-size: 12pt">Job description</label>
                        </div>
                        <div class="form-floating" style="margin-bottom: 2%">
                          <input id="jobValid" type="datetime-local" class="form-control" v-model="newJobOffer.dateValid" required="required" @change="userInfoHasChanged()">
                          <label for="jobValid" style="font-size: 12pt">Job valid</label>
                        </div>
                        <div style="display: flex; flex-direction: row; margin-bottom: 2%;">
                          <div class="form-floating" style="width: 100%">
                            <input id="newSkill" class="form-control" placeholder="New skill" type="text" name="new-skill" v-model="newSkill">
                            <label for="newSkill" style="font-size: 12pt">New skill</label>
                          </div>
                          <button class="btn-primary btn-round btn-icon"  style="border: 0; margin-left: 1%;" v-on:click="addSkillToJobOfferSkills()" v-if="newSkill">
                            <i class="now-ui-icons ui-1_simple-add"></i>
                          </button>
                        </div>
                        <div>
                          <div style="display: flex; flex-direction: row; margin-right: 2%" v-for="(skill,index) in newJobOffer.requiredSkills" :key="index">
                            <badge type="primary">
                              {{skill}}
                            </badge>
                            <badge type="default" style="border-radius: 50%; cursor: pointer;" v-on:click="removeSkillFromJobOffer(skill)">
                              ✕
                            </badge>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                  <div class="footer">
                    <button type="button" class="btn btn-primary" v-on:click="PublishJobOffer">Publish</button>
                  </div>
                </div>
                <div style="margin-bottom: 5%"></div>
                <JobOffers :users-jobs="true" :all-jobs="false" :userid="userID"></JobOffers>
              </tab-pane>

            </tabs>
          </div>
        </div>
        <ChangeProfile v-if="display=='settings'" :logged-user="loggedUserDetails"></ChangeProfile>
        <div v-if="display=='joboffers'">
          <JobOffers :users-jobs="false" :all-jobs="true" ></JobOffers>
        </div>
        <div v-if="display=='myjoboffers'">
          <JobOffers :users-jobs="true" :all-jobs="false" :userid="userID" ></JobOffers>

        </div>
        <Posts v-if="display=='feed'" :username="loggedUserDetails.username" :userid="userID" :feed-posts="true" :users-posts="false"></Posts>

        <Connections :display="this.displayConnections" v-if="display=='connections'"></Connections>




      </div>
      <div v-if="user.username==loggedUserDetails.username" class="d-flex flex-column flex-shrink-0 p-3 bg-light" style="width: 20%;">
        <Recommendation></Recommendation>
      </div>
    </div>
  </div>

  <div v-if="!(loggedUser || user.Public)">
    <div class="profile-panel">
      Log in to view {{user.name}} {{user.surname}}'s profile
    </div>
  </div>
  <div v-if="loggedUser && connectionStatus=='B_BLOCK_A'">
    <div style="width: 100%; justify-content: center; text-align: center" >
      You cannot view this user, because you are blocked by them
    </div>
  </div>
  <div v-if="loggedUser && connectionStatus=='A_BLOCK_B'">
    <div style="width: 100%; justify-content: center; text-align: center">
      You cannot view this user, because you blocked them

    </div>

  </div>

</div>
</template>

<script>
import PictureInput from 'vue-picture-input'
import PostService from '../services/PostService'
import UserService from '../services/UserService'
import ConnectionService from '../services/ConnectionService'
import JobService from "../services/JobService";
import $ from 'jquery'
import Recommendation from "./Connections/Recommendation";
import {Badge} from "../components";
import { Tabs, TabPane } from '@/components';
import ChangeProfile from "./Profile/ChangeProfile";
import JobOffers from "./Job/JobOffers";
import Posts from "./Post/Posts";
import Connections from "./Connections";
import Chatbox from "./Profile/Chatbox";
window.addEventListener('scroll', HandleScroll )

function HandleScroll(){
  var el =document.getElementById('sidebar');
  if ($(this).scrollTop() > 550){
    el.style.position= 'fixed'
    el.style.top='0';
    el.style.marginTop='4.5%'
  }
  if ($(this).scrollTop() < 550){
    el.style.position ='sticky'
    el.style.marginTop='0'
  }
}

export default {
  name: "Profile",
  bodyClass: 'profile-page',
  data(){
    return{
      postText: '',
      image: '',
      links: new Array(),
      link: '',
      user: {},
      loggedUser: {},
      loggedUserDetails : {},
      loggedUserFollows: false,
      connectionStatus: '',
      originalUser: {},
      usersPosts: new Array(),
      userID : String,


      selectedPost: "",
      selectedPostComments: [],
      selectedPostLikes: [],
      selectedPostDislikes: [],
      newCommentContent: '',
      userLiked: false,
      userDisliked: false,
      display: 'profile',
      newJobOffer: new Object(),
      newSkill: '',
      displayConnections: "friends"
    }
  },
  components:{
    PictureInput,
    Recommendation,
    Badge,
    Tabs,
    TabPane,
    ChangeProfile,
    JobOffers,
    Connections,
    Chatbox,
    Posts
  },
  mounted: function() {
    this.newJobOffer.requiredSkills = new Array();

    this.loggedUser = localStorage.getItem('user')

    UserService.getUserByUsername(this.$route.params.id).then(response => {
      this.user = response.data.User
      this.userID = this.user.id
      console.log("user to show", this.user)
      this.originalUser = this.user
      if (this.loggedUser)
      {
        UserService.getUserByUsername(JSON.parse(this.loggedUser).username).then(response2 => {
          this.loggedUserDetails = response2.data.User

          if(this.loggedUserDetails.id!= this.user.id)
          {
            ConnectionService.GetConnectionDetail(this.loggedUserDetails.id, this.user.id)
                .then(response3 =>{
                  //console.log(response3.data.relation)
                  this.connectionStatus = response3.data.relation
                  if(response3.data.relation == "CONNECTED"){
                    this.loggedUserFollows = true
                    //console.log("connected", this.loggedUserFollows)
                  } else{
                    this.loggedUserFollows = false
                    //console.log("not connected", this.loggedUserFollows)
                  }
                })
                .catch(err3 =>{
                  console.log("Connection details unavailable", err3)
                  alert("Connection details are unavailable!")
                })
          }

        })
            .catch(err2 =>{
              console.log(err2)
              alert("Logged user unavailable!")
            })

      }
      PostService.getAllPostsByUser(this.user.id).then(response1 =>{
        this.usersPosts = response1.data.posts;
        //console.log("UserPosts", this.usersPosts)
        if(this.usersPosts.length>0){
          this.selectedPost = this.usersPosts[0]
          this.isUserInfoChanged = false

          this.updatedSelectedPost()

        }

      })
          .catch(err1 =>{
            console.log(err1)
            //alert("User posts are unavailable!")
          })
    })
        .catch(err => {
          console.error(err);
          if(err.response.status == 403)
            this.$router.push("/unauthorized")
        })
  },
  methods:{
    DisplayFeed: function(){
      this.display='feed'
    },
     DisplayConnections: function(){
      this.display='connections'
    },
    DisplayConnectionsFriends: function(){
      this.displayConnections='friends'
    },
    DisplayConnectionsRequests: function(){
      this.displayConnections='requests'
    },
    DisplayConnectionsBlockeds: function(){
      this.displayConnections='blockeds'
    },

    DisplayChat: function(){
      this.display='chat'
    },
    DisplayJobOffers: function(){
      this.display='joboffers'
    },
    DisplayMyJobOffers: function(){
      this.display='myjoboffers'
    },
    DisplayProfile: function(){
      this.display='profile'
    },
    DisplaySettings: function(){
      this.display ='settings'
    },
    getConnectionDetails: function(){
      ConnectionService.GetConnectionDetail(this.loggedUserDetails.id, this.user.id)
          .then(response3 =>{
            //console.log(response3.data.relation)
            this.connectionStatus = response3.data.relation
            if(response3.data.relation == "CONNECTED"){
              this.loggedUserFollows = true
             // console.log("logged user follows changed to ", this.loggedUserFollows)
            } else{
              this.loggedUserFollows = false
             // console.log("logged user follows changed to ", this.loggedUserFollows)

            }
          })
          .catch(err3 =>{
            console.log("Connection details unavailable", err3)
            alert("Connection details are unavailable!")
          })
    },
    unsendRequest(){
      ConnectionService.UnsendFriendRequest(this.loggedUserDetails.id, this.user.id)
      .then( response => {
        console.log("connecting:", response);
        this.getConnectionDetails()
      })
          .catch(err => {
                console.log(err)
                alert("Error unsending request!")
              }
          )

    },
    RemoveFriendRequest(){
      ConnectionService.DeclineFriendRequest(this.loggedUserDetails.id, this.user.id)
      .then(response => {
        console.log("declining friend request: ", response);
        this.getConnectionDetails()
      })
      .catch(err => {
        console.log(err)
        alert("Error declining request!")
      })

    },
    formattedEducationType(educationType){
      if(educationType === "SECONDARY_EDUCATION")
        return "Secondary"
      if(educationType == "COLLEGE_EDUCATION")
        return "College"
      return "Primary"
    },
    deleteFriend(){
      ConnectionService.RemoveFriend(this.loggedUserDetails.id, this.user.id)
      .then(response => {
        console.log("deleting friend: ", response)
        this.getConnectionDetails()
      })
      .catch(err => {
        console.log(err)
        alert("Error deleting friend!")
      })

    },
    follow(){
      //alert(this.loggedUserDetails.username)
console.log(this.connectionStatus)
      if(this.user.Public || this.connectionStatus == 'ACCEPT'){
        ConnectionService.Connect(this.loggedUserDetails.id, this.user.id)
            .then( response => {
              console.log("connecting:", response);
              this.getConnectionDetails()
            })
            .catch(err => {
                  console.log(err)
                  alert("Error creating connection!")
                }
            )
      } else{
        ConnectionService.SendFriendRequest(this.loggedUserDetails.id, this.user.id)
            .then( response => {
              console.log("connecting sending request:", response);
              this.getConnectionDetails()
            })
            .catch(err => {
                  console.log(err)
                  alert("Error sending request!")
                }
            )
      }


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
      this.postText = '';
      this.link = '';
      this.newJobOffer.publisherId ="";
      this.newJobOffer.dateValid = "";
      this.newJobOffer.companyName="";
      this.newJobOffer.position="";
      this.newJobOffer.jobDescription="";
      this.newJobOffer.requiredSkills = new Array();
    },
    CreatePost: function (){
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
    PublishJobOffer: function(){
      console.log("U pusblish job fji")
      this.newJobOffer.datePosted = new Date();
      this.newJobOffer.jobID= "";
      this.newJobOffer.publisherId=this.loggedUserDetails.id;
      JobService.PublishPost(this.newJobOffer).then(res => {
        console.log("New job", res.data)
        this.emptyFields();
      }).catch(err =>{
            console.log(err);
            alert("It is not possible to publish job!")
          }
      );


    },
    addSkillToJobOfferSkills: function(){
      this.newJobOffer.requiredSkills.push(this.newSkill)
      this.newSkill = ""
    },
    removeSkillFromJobOffer: function (skill){
      alert("Skill to delete"+skill)
      for(let index in this.newJobOffer.requiredSkills) {
        if(this.newJobOffer.requiredSkills[index] === skill) {
          this.newJobOffer.requiredSkills.splice(index,1)
        }
      }
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

#sidebar{
  position:sticky;
  z-index:110;
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
    .profile-panel{
        background-color: rgba(255,255, 255, 1);
        color:#242424;
        position: relative;
        width:80%;
        margin:auto;
        margin-top: 15px;
        margin-bottom: 15px;
        padding: 20px;
        text-align:left;
        transition:.3s ease-in-out;
        z-index:0;
        box-shadow: 0 0 15px 9px #00000096;
        text-align:center;
    }
    
    .profile-panel input[type="button"] {
        max-width: 150px;
        width: 100%;
        margin-left: 20px;
        background: #e98074;
        color: #f9f9f9;
        border: none;
        padding: 10px;
        text-transform: uppercase;
        border-radius: 10px;
        float:right;
        cursor:pointer;
        transition: all .3s;
    }
    .profile-panel input[type="button"]:disabled {
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

    .profile-panel input[type="button"]:hover:enabled {
        transform: scale(110%);
        background: #e98074;
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
        border-bottom: 1px solid #eea098;
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
        border-bottom: 1px solid #e98074;
        box-shadow: 0 1px 0 0 #e98074;
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
        color: #e98074;
    }
    label.input_label > input:focus {
        border-bottom: 1px solid #e98074;
        box-shadow: 0 1px 0 0 #e98074;
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
        border-bottom: 1px solid #e98074;
        box-shadow: 0 1px 0 0 #e98074;
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
        color: #e98074;
    }
    label.input_label > select:valid + span.keep_hovered {
        transform: translateY(-25px) scale(0.8);
        transform-origin: 0;
        color: #e98074;
    }
    label.input_label > input:valid + span.keep_hovered {
        transform: translateY(-25px) scale(0.8);
        transform-origin: 0;
        color: #e98074;
    }
    label.input_label > select:focus {
        border-bottom: 1px solid #e98074;
        box-shadow: 0 1px 0 0 #e98074;
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
        border:1px solid #eea098;
        border-radius: 10px;
        resize: none;
        height: 400px;
        width: 100%;
        transition: all .3s;
    }

    .deletion-request-side-panel textarea:hover {
        border:1px solid #e98074;
        border-radius: 10px;
        resize: none;
        height: 400px;
        width: 100%;
        transition: all .3s;
    }

    .deletion-request-side-panel textarea:focus {
        border:1px solid #e98074;
        border-radius: 10px;
        resize: none;
        height: 400px;
        width: 100%;
        transition: all .3s;
    }

    /* ADD BUTTON */
    
    .button_plus {
      position: absolute;
      width: 35px;
      height: 35px;
      background: #fff;
      cursor: pointer;
      border: 2px solid #157c12;
      border-radius: 20px;
    }
    
    .button_plus:after {
      content: '';
      position: absolute;
      transform: translate(-50%, -50%);
      height: 4px;
      width: 50%;
      background: #157c12;
      top: 50%;
      left: 50%;
    }
    
    .button_plus:before {
      content: '';
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      background: #157c12;
      height: 50%;
      width: 4px;
    }
    
    .button_plus:hover:before,
    .button_plus:hover:after {
      background: #fff;
      transition: 0.2s;
    }
    
    .button_plus:hover {
      background-color: #157c12;
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
      content: '';
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
    .biography-textarea textarea {
        border:1px solid rgb(143, 176, 233);
        border-radius: 10px;
        resize: none;
        height: 400px;
        width: 100%;
        transition: all .3s;
    }

    .biography-textarea textarea:hover {
        border:1px solid rgba(0,95,255,1);
        border-radius: 10px;
        resize: none;
        height: 400px;
        width: 100%;
        transition: all .3s;
    }

    .biography-textarea textarea:focus {
        border:1px solid rgba(0,95,255,1);
        border-radius: 10px;
        resize: none;
        height: 400px;
        width: 100%;
        transition: all .3s;
    }

    .customactive{
      background-color: #999999;
      color: white;
      font-weight: bold;
      border-radius: 20px;
    }
</style>
