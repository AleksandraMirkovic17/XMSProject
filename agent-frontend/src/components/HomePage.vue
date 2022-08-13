<template>
<section class="bg-dark text-light p-5 text-center text-sm-start">
<br>
<br>
<br>

  <div class = "container pt-100">
    <div class = "d-sm-flex align-items-center justify-content-between">
      <div>
      <h1>Log in to your account</h1>
      <p class="lead my-4"> </p>
      <a type="button" class = "btn btn-primary btn-lg"  href="/userLogin">Login
      </a>

      </div>
      <div>
      <h1>Dont have an account?</h1>
      <p class="lead my-4"></p>
      <a type="button" class = "btn btn-primary btn-lg"  href="/userRegistration">Register
      </a>
      </div>
      <div>
        <img class="img-fluid w-30 s-sm-block" src = "">
      </div>
    </div>
  </div>
</section>

</template>

<script>


export default {
  name: 'HomePage',

  data: function(){
    return{
      user: null
    }
  },

  mounted(){
          
  },


  methods:{

  }
}
</script>
<style>

#home{   
  margin: 0 auto;
  text-align: center;
}

#home button{
	padding:10px;
	border:none;
	background-color:#2ECC71;
	color:#fff;
	font-weight:600;
	border-radius:20px;
	width:600px;
  height:50px;

}

#home button:hover{
	background: #333;
  }

</style>