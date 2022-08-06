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

    </div>
  </div>

</template>

<script>
export default {
  name: "Home",
  data(){
    return{
      loggedUser: "",
    }
  },
  mounted: function(){
    this.loggedUser = localStorage.getItem('user')
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



</style>