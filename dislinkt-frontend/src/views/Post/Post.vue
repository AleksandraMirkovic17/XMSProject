<template>
  <div>
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
          <h5>{{username}}</h5>
        </div>
        <div class="post-view-date" style="margin-left: 50%">
          <h5>{{date}}</h5>
        </div>
      </div>
      <div class="post-view-image" style="margin: 2%">
        <img v-bind:src="image" style="width: 100%">
      </div>
      <div class="post-txt" style="position:relative;">
        <p>{{ posttext }}</p>
      </div>
      <div class="post-links" style="margin-bottom: 3%">
        <div v-if="links && links.length==0">no links</div>
        <div v-for="(link, index) in links" :key="index">
          <a :href="link"><h6>{{ link}}</h6></a>
        </div>
      </div>
      <div class="post-additiona" style="display: flex; flex-direction: row; width: 100%">
        <div>
          <button type="primary" class="btn-default" style="border: 0pt" data-bs-toggle="modal" data-bs-target="#likesModal">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-hand-thumbs-up-fill" viewBox="0 0 16 16">
              <path d="M6.956 1.745C7.021.81 7.908.087 8.864.325l.261.066c.463.116.874.456 1.012.965.22.816.533 2.511.062 4.51a9.84 9.84 0 0 1 .443-.051c.713-.065 1.669-.072 2.516.21.518.173.994.681 1.2 1.273.184.532.16 1.162-.234 1.733.058.119.103.242.138.363.077.27.113.567.113.856 0 .289-.036.586-.113.856-.039.135-.09.273-.16.404.169.387.107.819-.003 1.148a3.163 3.163 0 0 1-.488.901c.054.152.076.312.076.465 0 .305-.089.625-.253.912C13.1 15.522 12.437 16 11.5 16H8c-.605 0-1.07-.081-1.466-.218a4.82 4.82 0 0 1-.97-.484l-.048-.03c-.504-.307-.999-.609-2.068-.722C2.682 14.464 2 13.846 2 13V9c0-.85.685-1.432 1.357-1.615.849-.232 1.574-.787 2.132-1.41.56-.627.914-1.28 1.039-1.639.199-.575.356-1.539.428-2.59z"/>
            </svg>
            <br>
            <h5 style="margin-top: 2%">
              Like
            </h5>
          </button>
        </div>
        <div>
          <button type="primary" class="btn-default" style="border: 0pt" data-bs-toggle="modal" data-bs-target="#dislikesModal">
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
        <div>
          <button type="primary" class="btn-default" style="border: 0pt" data-bs-toggle="modal" data-bs-target="#commentsModal">
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

</template>

<script>
export default {
  name: "Post",
  props:{

    username: String,
    date: String,
    posttext: String,
    image: String,
    links: Array,
    postid: String
  },
  data(){
    return{
      loggedUser: {},
    }
  },
  components:{

  },
  mounted: function () {
    this.loggedUser = localStorage.getItem('user')
  },
  methods:{


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

</style>