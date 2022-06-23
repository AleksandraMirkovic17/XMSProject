<template>
 <div class="main-home">
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
</template>

<script>
export default {
  name: "Home"
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


</style>