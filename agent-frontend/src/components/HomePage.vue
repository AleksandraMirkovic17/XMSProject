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
      <a type="button" class = "btn btn-primary btn-lg"  href="/login">Login
      </a>
<div class="modal fade" id="login" tabindex="-1" role="dialog" aria-labelledby="login" aria-hidden="true">
  <div class="modal-dialog" role="document">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title" id="loginn">Modal title</h5>
        <button type="button" class="close" data-dismiss="modal" aria-label="Close">
          <span aria-hidden="true">&times;</span>
        </button>
      </div>
      <div class="modal-body">
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
        <button type="button" class="btn btn-primary">Save changes</button>
      </div>
    </div>
  </div>
</div>
      </div>
      <div>
      <h1>Dont have an account?</h1>
      <p class="lead my-4"></p>
      <a type="button" class = "btn btn-primary btn-lg"  href="/clientRegistration">Register
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
          
    localStorage.removeItem('token');
    this.user = this.$store.state.userType
    if(this.user == 'Client')
      window.location.href = "/clientHomePage";
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