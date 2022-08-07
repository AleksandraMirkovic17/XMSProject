<template>
  <div>
    <div v-if="!loggedUser" class="main-home">
      <div class="img-home">
        <img src="../assets/slika1.jpg" style="border-radius: 20px; width: 100%; height: 100%">
      </div>
      <div class="text-home">
        <h1>Welcome to dislinkt!</h1>
        <p style="font-size: 12pt">
          Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.
        </p>
        <button type="button" class="btn btn-dark" style="width: 80%" onclick="window.location.href='/register'">Register</button>
        <br>
        <br>
        <p style="font-size: 12pt">
          Already have an account?
        </p>
        <button type="button" class="btn btn-dark" style="width: 80%" onclick="window.location.href='/login'">Log in</button>
      </div>

    </div>
    <div v-if="loggedUser">
      <div class="view-all-users-posts">
        <h4 style="margin-top: 1%; color: white">Feed</h4>
        <div v-if="this.feedPosts.length==0">
          <h3>There is no posts yet!</h3>
        </div>
        <div v-for="(post, index) in feedPosts" :key="index">
          <div class="post-view" v-on:mouseover="changeSelectedPost(post.Id)">
            <div class="post-view-nav" style="display: flex; flex-direction: row">
              <div class="post-view-person-icon">
                <svg xmlns="http://www.w3.org/2000/svg" width="26" height="26"
                     fill="currentColor" class="bi bi-person-circle" viewBox="0 0 16 16">
                  <path d="M11 6a3 3 0 1 1-6 0 3 3 0 0 1 6 0z"/>
                  <path fill-rule="evenodd" d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8zm8-7a7 7 0 0 0-5.468 11.37C3.242 11.226 4.805 10 8 10s4.757 1.225 5.468 2.37A7 7 0 0 0 8 1z"/>
                </svg>
              </div>
              <div class="post-view-username">
                <h5>{{post.User}}</h5>
              </div>
              <div class="post-view-date" style="margin-left: 50%">
                <h5>{{post.Date}}</h5>
              </div>
            </div>
            <div class="post-view-image" style="margin: 2%">
              <img v-bind:src="post.Image" style="width: 100%">
            </div>
            <div class="post-txt" style="position:relative;">
              <h6>{{ post.PostText }}</h6>
            </div>
            <div class="post-links">
              <div v-if="!post.Links">no links</div>
              <div v-for="(link, index) in post.Links" :key="index">
                <a :href="link"><h7>{{ link}}</h7></a>
              </div>
            </div>
            <div class="post-additiona" style="display: flex; flex-direction: row; width: 100%">
              <div class="likes" style="width: 33%;" >
                <button type="button" style="width: 100%; padding: 2%" data-bs-toggle="modal" data-bs-target="#likesModal">
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
                <button style="width: 100%; padding: 2%" data-bs-toggle="modal" data-bs-target="#dislikesModal">
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
                <button style="width: 100%; padding: 2%" data-bs-toggle="modal" data-bs-target="#commentsModal" v-on:click="changeSelectedPost(post.Id)">
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
    <!--Likes Modal -->
    <div class="modal fade" id="likesModal" tabindex="-1" aria-labelledby="likesModalLabel" aria-hidden="true">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="exampleModalLabel">Likes</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body">
            <div v-if="loggedUser">
              <div v-if="userLiked" style="width: 100%">
                <button style="width: 100%; padding: 3%" v-if="userLiked" v-on:click="deleteReaction()">
                  <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="blue" class="bi bi-hand-thumbs-up-fill" viewBox="0 0 16 16">
                    <path d="M6.956 1.745C7.021.81 7.908.087 8.864.325l.261.066c.463.116.874.456 1.012.965.22.816.533 2.511.062 4.51a9.84 9.84 0 0 1 .443-.051c.713-.065 1.669-.072 2.516.21.518.173.994.681 1.2 1.273.184.532.16 1.162-.234 1.733.058.119.103.242.138.363.077.27.113.567.113.856 0 .289-.036.586-.113.856-.039.135-.09.273-.16.404.169.387.107.819-.003 1.148a3.163 3.163 0 0 1-.488.901c.054.152.076.312.076.465 0 .305-.089.625-.253.912C13.1 15.522 12.437 16 11.5 16H8c-.605 0-1.07-.081-1.466-.218a4.82 4.82 0 0 1-.97-.484l-.048-.03c-.504-.307-.999-.609-2.068-.722C2.682 14.464 2 13.846 2 13V9c0-.85.685-1.432 1.357-1.615.849-.232 1.574-.787 2.132-1.41.56-.627.914-1.28 1.039-1.639.199-.575.356-1.539.428-2.59z"/>
                  </svg>
                  <h5>Like</h5>
                </button>
              </div>
              <div v-else style="width: 100%">
                <button style="width: 100%; padding: 3%" v-if="userLiked==false" v-on:click="likePost()">
                  <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="blue" class="bi bi-hand-thumbs-up" viewBox="0 0 16 16">
                    <path d="M8.864.046C7.908-.193 7.02.53 6.956 1.466c-.072 1.051-.23 2.016-.428 2.59-.125.36-.479 1.013-1.04 1.639-.557.623-1.282 1.178-2.131 1.41C2.685 7.288 2 7.87 2 8.72v4.001c0 .845.682 1.464 1.448 1.545 1.07.114 1.564.415 2.068.723l.048.03c.272.165.578.348.97.484.397.136.861.217 1.466.217h3.5c.937 0 1.599-.477 1.934-1.064a1.86 1.86 0 0 0 .254-.912c0-.152-.023-.312-.077-.464.201-.263.38-.578.488-.901.11-.33.172-.762.004-1.149.069-.13.12-.269.159-.403.077-.27.113-.568.113-.857 0-.288-.036-.585-.113-.856a2.144 2.144 0 0 0-.138-.362 1.9 1.9 0 0 0 .234-1.734c-.206-.592-.682-1.1-1.2-1.272-.847-.282-1.803-.276-2.516-.211a9.84 9.84 0 0 0-.443.05 9.365 9.365 0 0 0-.062-4.509A1.38 1.38 0 0 0 9.125.111L8.864.046zM11.5 14.721H8c-.51 0-.863-.069-1.14-.164-.281-.097-.506-.228-.776-.393l-.04-.024c-.555-.339-1.198-.731-2.49-.868-.333-.036-.554-.29-.554-.55V8.72c0-.254.226-.543.62-.65 1.095-.3 1.977-.996 2.614-1.708.635-.71 1.064-1.475 1.238-1.978.243-.7.407-1.768.482-2.85.025-.362.36-.594.667-.518l.262.066c.16.04.258.143.288.255a8.34 8.34 0 0 1-.145 4.725.5.5 0 0 0 .595.644l.003-.001.014-.003.058-.014a8.908 8.908 0 0 1 1.036-.157c.663-.06 1.457-.054 2.11.164.175.058.45.3.57.65.107.308.087.67-.266 1.022l-.353.353.353.354c.043.043.105.141.154.315.048.167.075.37.075.581 0 .212-.027.414-.075.582-.05.174-.111.272-.154.315l-.353.353.353.354c.047.047.109.177.005.488a2.224 2.224 0 0 1-.505.805l-.353.353.353.354c.006.005.041.05.041.17a.866.866 0 0 1-.121.416c-.165.288-.503.56-1.066.56z"/>
                  </svg>
                  <h5>Like</h5>
                </button>
              </div>
            </div>
            <div v-for="(reaction, index) in selectedPostLikes"
                 :key="index">
              <div class="reaction-view" style="display: flex; flex-direction: row">
                <div class="reaction-icon" style="margin: 1%">
                  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-hand-thumbs-up-fill" viewBox="0 0 16 16">
                    <path d="M6.956 1.745C7.021.81 7.908.087 8.864.325l.261.066c.463.116.874.456 1.012.965.22.816.533 2.511.062 4.51a9.84 9.84 0 0 1 .443-.051c.713-.065 1.669-.072 2.516.21.518.173.994.681 1.2 1.273.184.532.16 1.162-.234 1.733.058.119.103.242.138.363.077.27.113.567.113.856 0 .289-.036.586-.113.856-.039.135-.09.273-.16.404.169.387.107.819-.003 1.148a3.163 3.163 0 0 1-.488.901c.054.152.076.312.076.465 0 .305-.089.625-.253.912C13.1 15.522 12.437 16 11.5 16H8c-.605 0-1.07-.081-1.466-.218a4.82 4.82 0 0 1-.97-.484l-.048-.03c-.504-.307-.999-.609-2.068-.722C2.682 14.464 2 13.846 2 13V9c0-.85.685-1.432 1.357-1.615.849-.232 1.574-.787 2.132-1.41.56-.627.914-1.28 1.039-1.639.199-.575.356-1.539.428-2.59z"/>
                  </svg>
                </div>
                <h5 style="margin: 1%">{{ reaction.username }}</h5>
              </div>


            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
          </div>
        </div>
      </div>
    </div>
    <!--DislikesModal -->
    <div class="modal fade" id="dislikesModal" tabindex="-1" aria-labelledby="dislikesModalLabel" aria-hidden="true">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="exampleModalLabel1">Dislikes</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body">
            <div v-if="loggedUser">


              <div v-if="userDisliked" style="width: 100%">
                <button style="width: 100%; padding: 3%" v-if="userDisliked" v-on:click="deleteReaction()">
                  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="blue" class="bi bi-hand-thumbs-down-fill" viewBox="0 0 16 16">
                    <path d="M6.956 14.534c.065.936.952 1.659 1.908 1.42l.261-.065a1.378 1.378 0 0 0 1.012-.965c.22-.816.533-2.512.062-4.51.136.02.285.037.443.051.713.065 1.669.071 2.516-.211.518-.173.994-.68 1.2-1.272a1.896 1.896 0 0 0-.234-1.734c.058-.118.103-.242.138-.362.077-.27.113-.568.113-.856 0-.29-.036-.586-.113-.857a2.094 2.094 0 0 0-.16-.403c.169-.387.107-.82-.003-1.149a3.162 3.162 0 0 0-.488-.9c.054-.153.076-.313.076-.465a1.86 1.86 0 0 0-.253-.912C13.1.757 12.437.28 11.5.28H8c-.605 0-1.07.08-1.466.217a4.823 4.823 0 0 0-.97.485l-.048.029c-.504.308-.999.61-2.068.723C2.682 1.815 2 2.434 2 3.279v4c0 .851.685 1.433 1.357 1.616.849.232 1.574.787 2.132 1.41.56.626.914 1.28 1.039 1.638.199.575.356 1.54.428 2.591z"/>
                  </svg>
                  <h5>Dislike</h5>
                </button>
              </div>
              <div v-else style="width: 100%">
                <button style="width: 100%; padding: 3%" v-if="userDisliked==false" v-on:click="dislikePost()">
                  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="blue" class="bi bi-hand-thumbs-down" viewBox="0 0 16 16">
                    <path d="M8.864 15.674c-.956.24-1.843-.484-1.908-1.42-.072-1.05-.23-2.015-.428-2.59-.125-.36-.479-1.012-1.04-1.638-.557-.624-1.282-1.179-2.131-1.41C2.685 8.432 2 7.85 2 7V3c0-.845.682-1.464 1.448-1.546 1.07-.113 1.564-.415 2.068-.723l.048-.029c.272-.166.578-.349.97-.484C6.931.08 7.395 0 8 0h3.5c.937 0 1.599.478 1.934 1.064.164.287.254.607.254.913 0 .152-.023.312-.077.464.201.262.38.577.488.9.11.33.172.762.004 1.15.069.13.12.268.159.403.077.27.113.567.113.856 0 .289-.036.586-.113.856-.035.12-.08.244-.138.363.394.571.418 1.2.234 1.733-.206.592-.682 1.1-1.2 1.272-.847.283-1.803.276-2.516.211a9.877 9.877 0 0 1-.443-.05 9.364 9.364 0 0 1-.062 4.51c-.138.508-.55.848-1.012.964l-.261.065zM11.5 1H8c-.51 0-.863.068-1.14.163-.281.097-.506.229-.776.393l-.04.025c-.555.338-1.198.73-2.49.868-.333.035-.554.29-.554.55V7c0 .255.226.543.62.65 1.095.3 1.977.997 2.614 1.709.635.71 1.064 1.475 1.238 1.977.243.7.407 1.768.482 2.85.025.362.36.595.667.518l.262-.065c.16-.04.258-.144.288-.255a8.34 8.34 0 0 0-.145-4.726.5.5 0 0 1 .595-.643h.003l.014.004.058.013a8.912 8.912 0 0 0 1.036.157c.663.06 1.457.054 2.11-.163.175-.059.45-.301.57-.651.107-.308.087-.67-.266-1.021L12.793 7l.353-.354c.043-.042.105-.14.154-.315.048-.167.075-.37.075-.581 0-.211-.027-.414-.075-.581-.05-.174-.111-.273-.154-.315l-.353-.354.353-.354c.047-.047.109-.176.005-.488a2.224 2.224 0 0 0-.505-.804l-.353-.354.353-.354c.006-.005.041-.05.041-.17a.866.866 0 0 0-.121-.415C12.4 1.272 12.063 1 11.5 1z"/>
                  </svg>
                  <h5>Dislike</h5>
                </button>
              </div>
            </div>
            <div v-for="(reaction, index) in selectedPostDislikes"
                 :key="index">
              <div class="reaction-view" style="display: flex; flex-direction: row">
                <div class="reaction-icon" style="margin: 1%">
                  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-hand-thumbs-down-fill" viewBox="0 0 16 16">
                    <path d="M6.956 14.534c.065.936.952 1.659 1.908 1.42l.261-.065a1.378 1.378 0 0 0 1.012-.965c.22-.816.533-2.512.062-4.51.136.02.285.037.443.051.713.065 1.669.071 2.516-.211.518-.173.994-.68 1.2-1.272a1.896 1.896 0 0 0-.234-1.734c.058-.118.103-.242.138-.362.077-.27.113-.568.113-.856 0-.29-.036-.586-.113-.857a2.094 2.094 0 0 0-.16-.403c.169-.387.107-.82-.003-1.149a3.162 3.162 0 0 0-.488-.9c.054-.153.076-.313.076-.465a1.86 1.86 0 0 0-.253-.912C13.1.757 12.437.28 11.5.28H8c-.605 0-1.07.08-1.466.217a4.823 4.823 0 0 0-.97.485l-.048.029c-.504.308-.999.61-2.068.723C2.682 1.815 2 2.434 2 3.279v4c0 .851.685 1.433 1.357 1.616.849.232 1.574.787 2.132 1.41.56.626.914 1.28 1.039 1.638.199.575.356 1.54.428 2.591z"/>
                  </svg>
                </div>
                <h5 style="margin: 1%">{{ reaction.username }}</h5>
              </div>

            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
          </div>
        </div>
      </div>
    </div>
    <!--CommentsModal -->
    <div class="modal fade" id="commentsModal" tabindex="-1" aria-labelledby="commentsModalLabel" aria-hidden="true">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="exampleModalLabel2">Comments</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body">
            <div v-for="(comment, index) in selectedPostComments"
                 :key="index">
              <div class="comment-view" style="display: flex; flex-direction: row; width: 100%; border: 1pt solid black; border-radius: 20px; margin:1%;">
                <div class="comment-person-icon" style="margin: 1%; width: 10%">
                  <svg xmlns="http://www.w3.org/2000/svg" width="25" height="25" fill="currentColor" class="bi bi-person-circle" viewBox="0 0 16 16">
                    <path d="M11 6a3 3 0 1 1-6 0 3 3 0 0 1 6 0z"/>
                    <path fill-rule="evenodd" d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8zm8-7a7 7 0 0 0-5.468 11.37C3.242 11.226 4.805 10 8 10s4.757 1.225 5.468 2.37A7 7 0 0 0 8 1z"/>
                  </svg>
                </div>
                <div class="comment-view-info" style="margin: 1%; width: 90%">
                  <div class="comment-basic-info" style="border-bottom: 1pt black solid; display: flex; flex-direction: row; width: 100%">
                    <h4>@{{ comment.username }}</h4>
                    <p style="margin-left: 40%; font-size: 9pt">{{comment.date}}</p>
                  </div>
                  <div class="comment-content" style="margin: 1%;">
                    <h5 style="font-style: italic">
                      ,,{{comment.content}}"
                    </h5>
                  </div>

                </div>


              </div>


            </div>
            <div v-if="loggedUser">
              <div class="post-text">
                <div class="form-floating">
                  <textarea class="form-control" placeholder="Write a comment..." id="floatingTextarea1" style="font-size: 12pt; height: 250px" v-model = "newCommentContent"></textarea>
                  <label for="floatingTextarea1" style="font-size: 12pt">Write a comment</label>
                </div>

              </div>

            </div>
          </div>
          <div class="modal-footer">
            <button v-if="loggedUser" type="button" class="btn btn-primary" v-on:click="commentOnPost()">Comment</button>
            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
          </div>
        </div>
      </div>
    </div>
  </div>

</template>

<script>
import PostService from "../services/PostService";
export default {
  name: "Home",
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
          alert("Error while getting feed!")

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
    changeSelectedPost(postID){
      PostService.getPostByIdWithUsername(postID)
      .then(resp =>{
        this.selectedPost = resp.data
        console.log("changing selected", this.selectedPost)
        PostService.getLikes(this.selectedPost.Id)
            .then(response =>{
              this.selectedPostLikes = response.data.owner;
              console.log("Selected posts likes: ", this.selectedPostLikes)
              this.checkUserLiked()
            })
            .catch(err =>{
              alert("It is impossible to get likes!")
              console.log(err)
            })
        PostService.getDislikes(this.selectedPost.Id)
            .then(response =>{
              this.selectedPostDislikes = response.data.owner
              this.checkUserDisliked()
            })
            .catch(err =>{
              alert("It is impossible to get dislikes!")
              console.log(err)
            })
        PostService.getComments(this.selectedPost.Id)
            .then(response =>{
              this.selectedPostComments = response.data.comment
              console.log("comments", this.selectedPostComments)
            })
            .catch(err =>{
              alert("It is impossible to get comments!")
              console.log(err)
            })
      })
      .catch(err =>{
        console.log("Error while getting post ", err)
        alert("It is impossible to get post by id with username!")
      })

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